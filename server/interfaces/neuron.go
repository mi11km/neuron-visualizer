package interfaces

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"os"
	"os/exec"
	"path"
	"regexp"
	"strconv"
	"strings"

	"github.com/mi11km/neuron-visualizer/server/openapi"
)

func NewNeuronVisualizerServer(neuronSimulationDir string) (*NeuronVisualizerServer, error) {
	rootPath, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	swcLineRegex, err := regexp.Compile(`^(\d+\s+){2}(-?\d+\.\d+\s+){4}-?\d+\s*$`)
	if err != nil {
		return nil, err
	}
	return &NeuronVisualizerServer{
		rootPath:             rootPath,
		neuronSimulationPath: path.Join(rootPath, neuronSimulationDir),
		swcLineRegex:         swcLineRegex,
		compartmentTypeMap: map[int64]openapi.NeuronCompartmentTypeName{
			1: openapi.SOMA,
			2: openapi.AXON,
			3: openapi.BASALDENDRITE,
			4: openapi.APICALDENDRITE,
		},
	}, nil
}

type NeuronVisualizerServer struct {
	rootPath             string
	neuronSimulationPath string
	swcLineRegex         *regexp.Regexp
	compartmentTypeMap   map[int64]openapi.NeuronCompartmentTypeName
}

func (n *NeuronVisualizerServer) HealthCheck(w http.ResponseWriter, r *http.Request) {
	if mediaType := r.Header.Get("Accept"); mediaType != "application/json" {
		http.Error(w, "Unsupported media type", http.StatusUnsupportedMediaType)
		return
	}

	res, err := json.Marshal(
		&openapi.HealthCheckResponse{
			Status:  openapi.OK,
			Message: "Server is running",
		},
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err = w.Write(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (n *NeuronVisualizerServer) GetNeurons(w http.ResponseWriter, r *http.Request) {
	if mediaType := r.Header.Get("Accept"); mediaType != "application/json" {
		http.Error(w, "Unsupported media type", http.StatusUnsupportedMediaType)
		return
	}

	entries, err := os.ReadDir(n.neuronSimulationPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	neurons := make([]openapi.Neuron, 0, len(entries))
	for _, entry := range entries {
		// フォルダ名がそのままニューロン名になる
		if entry.IsDir() && entry.Name() != "" {
			neurons = append(
				neurons, openapi.Neuron{
					Name: entry.Name(),
				},
			)
		}
	}
	res, err := json.Marshal(openapi.GetNeuronsResponse{Neurons: neurons})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err = w.Write(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (n *NeuronVisualizerServer) GetNeuronCompartments(
	w http.ResponseWriter, r *http.Request, neuronName openapi.NeuronName,
) {
	if mediaType := r.Header.Get("Accept"); mediaType != "application/json" {
		http.Error(w, "Unsupported media type", http.StatusUnsupportedMediaType)
		return
	}

	// swcファイル(ニューロンのコンパートメント情報)を読み込む
	// swcファイルのパスはデフォルトで以下のようになっている
	// <root>/simulations/<neuron_name>/<neuron_name>.swc
	swcFilePath := path.Join(n.neuronSimulationPath, neuronName, fmt.Sprintf("%s.swc", neuronName))
	swcFile, err := os.ReadFile(swcFilePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	validateSwcLine := func(line string) []string {
		// swcファイルの各行は以下のような形式になっている
		// <compartment_id> <compartment_type> <x> <y> <z> <radius> <parent_compartment_id>
		if !n.swcLineRegex.MatchString(line) {
			return nil
		}
		words := strings.Fields(line)
		if len(words) != 7 {
			return nil
		}
		return words
	}

	// swcファイルの各行をパースしてニューロンのコンパートメント情報を取得する
	lines := strings.Split(string(swcFile), "\n")
	neuronCompartments := make([]openapi.NeuronCompartment, 0, len(lines))

	// somaの位置を 0, 0, 0 にするために、取得する
	var somaPositionX, somaPositionY, somaPositionZ float64
	for _, line := range lines {
		var err error
		words := validateSwcLine(line)
		if words == nil {
			continue
		}

		compartmentTypeInt, err := strconv.ParseInt(words[1], 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// somaのみを対象とする
		if compartmentTypeInt != 1 {
			continue
		}

		somaPositionX, err = strconv.ParseFloat(words[2], 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		somaPositionY, err = strconv.ParseFloat(words[3], 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		somaPositionZ, err = strconv.ParseFloat(words[4], 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		break
	}

	for _, line := range lines {
		words := validateSwcLine(line)
		if words == nil {
			continue
		}

		// 変換処理
		id, err := strconv.ParseInt(words[0], 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		compartmentTypeInt, err := strconv.ParseInt(words[1], 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		positionX, err := strconv.ParseFloat(words[2], 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		positionY, err := strconv.ParseFloat(words[3], 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		positionZ, err := strconv.ParseFloat(words[4], 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		radius, err := strconv.ParseFloat(words[5], 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		parentId, err := strconv.ParseInt(words[6], 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		neuronCompartments = append(
			neuronCompartments, openapi.NeuronCompartment{
				Id:        id,
				ParentId:  parentId,
				PositionX: positionX - somaPositionX,
				PositionY: positionY - somaPositionY,
				PositionZ: positionZ - somaPositionZ,
				Radius:    radius,
				Type: openapi.NeuronCompartmentType{
					Id:   compartmentTypeInt,
					Name: n.compartmentTypeMap[compartmentTypeInt],
				},
			},
		)
	}

	res, err := json.Marshal(openapi.GetNeuronCompartmentsResponse{Compartments: neuronCompartments})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err = w.Write(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (n *NeuronVisualizerServer) GetNeuronMembranePotentials(
	w http.ResponseWriter, r *http.Request, neuronName openapi.NeuronName,
) {
	// SSEのみをサポートする
	if mediaType := r.Header.Get("Accept"); mediaType != "text/event-stream" {
		http.Error(w, "Unsupported media type", http.StatusUnsupportedMediaType)
		return
	}

	// 一時的にシミュレーションの実行ディレクトリに移動する
	// <root>/simulations/<neuron_name>/
	simulationRootPath := path.Join(n.neuronSimulationPath, neuronName)
	if err := os.Chdir(simulationRootPath); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer func() {
		if err := os.Chdir(n.rootPath); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}()
	// シミュレーションの実行ファイルのパスはデフォルトで以下のようになっている
	// <root>/simulations/<neuron_name>/main
	simulationBinPath := path.Join(simulationRootPath, "main")
	if _, err := os.Stat(simulationBinPath); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cmd := exec.CommandContext(r.Context(), simulationBinPath)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := cmd.Start(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}
	cw := httputil.NewChunkedWriter(w)
	defer func() {
		if err := cw.Close(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}()

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		// シミュレーションの出力は以下のような形式になっている
		// <time_step>,<membrane_potential_1>,<membrane_potential_2>,...
		words := strings.Split(scanner.Text(), ",")
		timeStep, err := strconv.ParseFloat(words[0], 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		membranePotentials := make([]float64, 0, len(words)-1)
		for _, potentialStr := range words[1:] {
			membranePotential, err := strconv.ParseFloat(potentialStr, 64)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			membranePotentials = append(membranePotentials, membranePotential)
		}

		res, err := json.Marshal(
			openapi.GetNeuronCompartmentsMembranePotentialResponse{
				TimeStep:           timeStep,
				MembranePotentials: membranePotentials,
			},
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if _, err = cw.Write(res); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		flusher.Flush()
	}
	if err := scanner.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	stderrBytes, err := io.ReadAll(stderr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := cmd.Wait(); err != nil {
		http.Error(w, string(stderrBytes), http.StatusInternalServerError)
		return
	}
}
