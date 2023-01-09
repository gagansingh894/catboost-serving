package main

import (
	"context"
	"fmt"
	pb "github.com/gagansingh894/catboost-serving/internal/pkg/pb/cbserving"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"math/rand"
	"strconv"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	defer conn.Close()
	cbmClient := pb.NewDeploymentServiceClient(conn)

	// prepare request
	var times int64

	for i := 0; i < 500; i++ {
		req := createPredictionRequest(rand.Intn(500)+1, 125, "model_v4")
		start := time.Now()
		_, err = cbmClient.GetPredictions(context.Background(), req)
		if err != nil {
			log.Fatalf("failed to call CBM Service: %v", err)
		}
		elapsed := time.Since(start)
		fmt.Println("time taken: ", elapsed.Milliseconds())
		times += elapsed.Milliseconds()
	}
	fmt.Println("average time taken:", times/500)

}

func createPredictionRequest(numRecords, numFeatures int, modelName string) *pb.GetPredictionsRequest {
	inputData := make([]*pb.GetPredictionsRequest_InputData, numRecords)

	for i, _ := range inputData {
		data := make(map[string]float32)
		for j := 0; j < numFeatures; j++ {
			featureName := fmt.Sprintf("feature_%s", strconv.Itoa(j))
			data[featureName] = rand.Float32() * float32(rand.Intn(20))
		}
		inputData[i] = &pb.GetPredictionsRequest_InputData{Input: data}
	}

	return &pb.GetPredictionsRequest{
		ModelName: modelName,
		ModelTask: pb.ModelTask_MODEL_TASK_REGRESSION,
		InputData: inputData,
	}
}
