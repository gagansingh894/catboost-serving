package v1

import (
	"context"
	"fmt"
	"github.com/gagansingh894/catboost-serving/internal/modelmanager"
	"github.com/gagansingh894/catboost-serving/pkg/pb/cbserving"
	"strconv"
	"time"
)

type cbDeploymentService struct {
	cbserving.UnimplementedDeploymentServiceServer
	modelManager modelmanager.ModelManager
}

var _ cbserving.DeploymentServiceServer = (*cbDeploymentService)(nil)

func NewCBDeploymentService(modelManager modelmanager.ModelManager) cbserving.DeploymentServiceServer {
	return &cbDeploymentService{
		modelManager: modelManager,
	}
}

func (c *cbDeploymentService) GetPredictions(ctx context.Context, in *cbserving.GetPredictionsRequest) (*cbserving.GetPredictionsResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	preds, err := c.modelManager.GetPredictions(in.ModelName, parseProtoToModelInputs(in), [][]string{})
	if err != nil {
		return nil, fmt.Errorf("failed to get predictions: %w", err)
	}

	return &cbserving.GetPredictionsResponse{
		ModelTask:   in.ModelTask,
		ModelName:   in.ModelName,
		Predictions: parsePredictionsToProto(preds),
	}, nil
}

func parsePredictionsToProto(in []float64) map[string]float64 {
	out := make(map[string]float64)

	for i, v := range in {
		out[strconv.Itoa(i)] = v
	}

	return out
}

func parseProtoToModelInputs(in *cbserving.GetPredictionsRequest) [][]float32 {
	out := make([][]float32, len(in.InputData))

	for i, input := range in.InputData {
		var row []float32
		for _, value := range input.Input {
			row = append(row, value)
		}
		out[i] = row
	}

	return out
}
