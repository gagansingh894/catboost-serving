package v1

import (
	"context"
	"fmt"
	"github.com/gagansingh894/catboost-serving/go/internal/modelmanager"
	pb "github.com/gagansingh894/catboost-serving/go/internal/pkg/pb/cbserving"
	"strconv"
	"time"
)

type cbDeploymentService struct {
	pb.UnimplementedDeploymentServiceServer
	modelManager modelmanager.ModelManager
}

var _ pb.DeploymentServiceServer = (*cbDeploymentService)(nil)

func NewCBDeploymentService(modelManager modelmanager.ModelManager) pb.DeploymentServiceServer {
	return &cbDeploymentService{
		modelManager: modelManager,
	}
}

func (c *cbDeploymentService) GetPredictions(ctx context.Context, in *pb.GetPredictionsRequest) (*pb.GetPredictionsResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	preds, err := c.modelManager.GetPredictions(in.ModelName, parseProtoToModelInputs(in), [][]string{})
	if err != nil {
		return nil, fmt.Errorf("failed to get predictions: %w", err)
	}

	return &pb.GetPredictionsResponse{
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

func parseProtoToModelInputs(in *pb.GetPredictionsRequest) [][]float32 {
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
