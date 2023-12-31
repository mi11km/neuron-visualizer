package interfaces

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"os/exec"
	"path"
	"regexp"
	"strconv"
	"strings"

	neuronv1 "github.com/mi11km/neuron-visualizer/server/proto/neuron/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

var _ neuronv1.NeuronServiceServer = (*NeuronServiceServer)(nil)

func NewNeuronServiceServer(neuronSimulationDir string) (*NeuronServiceServer, error) {
	rootPath, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	swcLineRegex, err := regexp.Compile(`^(\d+\s+){2}(-?\d+\.\d+\s+){4}-?\d+\s*$`)
	if err != nil {
		return nil, err
	}
	return &NeuronServiceServer{
		rootPath:             rootPath,
		neuronSimulationPath: path.Join(rootPath, neuronSimulationDir),
		swcLineRegex:         swcLineRegex,
	}, nil
}

type NeuronServiceServer struct {
	rootPath             string
	neuronSimulationPath string
	swcLineRegex         *regexp.Regexp
}

func (n *NeuronServiceServer) GetMembranePotentials(request *neuronv1.GetMembranePotentialsRequest, server neuronv1.NeuronService_GetMembranePotentialsServer) error {
	// 一時的にシミュレーションの実行ディレクトリに移動する
	// <root>/simulations/<neuron_name>/
	simulationRootPath := path.Join(n.neuronSimulationPath, request.NeuronName)
	if err := os.Chdir(simulationRootPath); err != nil {
		return fmt.Errorf("os.Chdir: %w", err)
	}
	defer func() {
		if err := os.Chdir(n.rootPath); err != nil {
			slog.Error("os.Chdir: ", slog.Any("error", err))
		}
	}()
	// シミュレーションの実行ファイルのパスはデフォルトで以下のようになっている
	// <root>/simulations/<neuron_name>/main
	simulationBinPath := path.Join(simulationRootPath, "main")
	if _, err := os.Stat(simulationBinPath); err != nil {
		return fmt.Errorf("os.Stat: %w", err)
	}

	cmd := exec.CommandContext(server.Context(), simulationBinPath)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("cmd.StdoutPipe: %w", err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return fmt.Errorf("cmd.StderrPipe: %w", err)
	}
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("cmd.Start: %w", err)
	}

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		// シミュレーションの出力は以下のような形式になっている
		// <time_step>,<membrane_potential_1>,<membrane_potential_2>,...
		words := strings.Split(scanner.Text(), ",")
		timeStep, err := strconv.ParseFloat(words[0], 32)
		if err != nil {
			return fmt.Errorf("strconv.ParseFloat: %w", err)
		}
		membranePotentials := make([]float32, 0, len(words)-1)
		for _, potentialStr := range words[1:] {
			membranePotential, err := strconv.ParseFloat(potentialStr, 32)
			if err != nil {
				return fmt.Errorf("strconv.ParseFloat: %w", err)
			}
			membranePotentials = append(membranePotentials, float32(membranePotential))
		}

		if err = server.Send(&neuronv1.GetMembranePotentialsResponse{
			TimeStep:           float32(timeStep),
			MembranePotentials: membranePotentials,
		}); err != nil {
			return fmt.Errorf("server.Send: %w", err)
		}
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("scanner.Err: %w", err)
	}

	stderrBytes, err := io.ReadAll(stderr)
	if err != nil {
		return fmt.Errorf("io.ReadAll: %w", err)
	}
	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("cmd.Wait: %w, %s", err, string(stderrBytes))
	}

	return nil
}

func (n *NeuronServiceServer) ListNeurons(ctx context.Context, empty *emptypb.Empty) (*neuronv1.ListNeuronsResponse, error) {
	entries, err := os.ReadDir(n.neuronSimulationPath)
	if err != nil {
		return nil, err
	}
	neuronNames := make([]string, 0, len(entries))
	for _, entry := range entries {
		// フォルダ名がそのままニューロン名になる
		if entry.IsDir() && entry.Name() != "" {
			neuronNames = append(neuronNames, entry.Name())
		}
	}
	return &neuronv1.ListNeuronsResponse{NeuronNames: neuronNames}, nil
}

func (n *NeuronServiceServer) GetNeuronShape(ctx context.Context, request *neuronv1.GetNeuronShapeRequest) (*neuronv1.GetNeuronShapeResponse, error) {
	// swcファイル(ニューロンのコンパートメント情報)を読み込む
	// swcファイルのパスはデフォルトで以下のようになっている
	// <root>/simulations/<neuron_name>/<neuron_name>.swc
	swcFilePath := path.Join(n.neuronSimulationPath, request.NeuronName, fmt.Sprintf("%s.swc", request.NeuronName))
	swcFile, err := os.ReadFile(swcFilePath)
	if err != nil {
		return nil, err
	}

	// swcファイルの各行をパースしてニューロンのコンパートメント情報を取得する
	lines := strings.Split(string(swcFile), "\n")
	neuronCompartments := make([]*neuronv1.NeuronCompartment, 0, len(lines))
	for _, line := range lines {
		// swcファイルの各行は以下のような形式になっている
		// <compartment_id> <compartment_type> <x> <y> <z> <radius> <parent_compartment_id>
		if !n.swcLineRegex.MatchString(line) {
			continue
		}
		words := strings.Fields(line)
		if len(words) != 7 {
			continue
		}

		// 変換処理
		id, err := strconv.ParseInt(words[0], 10, 64)
		if err != nil {
			return nil, err
		}

		compartmentTypeInt, err := strconv.ParseFloat(words[1], 32)
		if err != nil {
			return nil, err
		}
		_, ok := neuronv1.NeuronCompartmentType_name[int32(compartmentTypeInt)]
		if !ok {
			return nil, status.Errorf(codes.Unavailable, "non-exist compartment type: %v", compartmentTypeInt)
		}

		positionX, err := strconv.ParseFloat(words[2], 32)
		if err != nil {
			return nil, err
		}
		positionY, err := strconv.ParseFloat(words[3], 32)
		if err != nil {
			return nil, err
		}
		positionZ, err := strconv.ParseFloat(words[4], 32)
		if err != nil {
			return nil, err
		}
		radius, err := strconv.ParseFloat(words[5], 32)
		if err != nil {
			return nil, err
		}

		parentId, err := strconv.ParseInt(words[6], 10, 64)
		if err != nil {
			return nil, err
		}

		neuronCompartments = append(neuronCompartments, &neuronv1.NeuronCompartment{
			Id:        id,
			Type:      neuronv1.NeuronCompartmentType(int32(compartmentTypeInt)),
			PositionX: float32(positionX),
			PositionY: float32(positionY),
			PositionZ: float32(positionZ),
			Radius:    float32(radius),
			ParentId:  parentId,
		})
	}
	return &neuronv1.GetNeuronShapeResponse{
		NeuronCompartments: neuronCompartments,
	}, nil
}
