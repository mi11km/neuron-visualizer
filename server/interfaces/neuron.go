package interfaces

import (
	"context"
	"fmt"
	"os"
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

type NeuronServiceServer struct {
	neuronSimulationPath string
	swcLineRegex         *regexp.Regexp
}

func NewNeuronServiceServer(neuronSimulationDir string) (*NeuronServiceServer, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	swcLineRegex, err := regexp.Compile(`^(\d+\s+){2}(-?\d+\.\d+\s+){4}-?\d+\s*$`)
	if err != nil {
		return nil, err
	}
	return &NeuronServiceServer{
		neuronSimulationPath: path.Join(currentDir, neuronSimulationDir),
		swcLineRegex:         swcLineRegex,
	}, nil
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
	// swcファイルのパスは以下のようになっている
	// <binary_execute_dir>/simulations/<neuron_name>/<neuron_name>.swc
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
