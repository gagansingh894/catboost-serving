package main

import (
	"context"
	"flag"
	"fmt"
	pb "github.com/gagansingh894/catboost-serving/internal/pkg/pb/cbserving"
	"log"
	"math"
	"math/rand"
	"strconv"
	"time"

	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	isParallelPtr := flag.Bool("parallel", false, " Determine whether to parallelize request ")
	numRecordsPtr := flag.Int("records", 300, " Number of records in a single request for which predictions are done")
	numIterPtr := flag.Int("iter", 500, " Number of iterations for benchmarking")
	flag.Parse()

	fmt.Println("creating client")
	conn, err := grpc.Dial("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	defer conn.Close()
	cbmClient := pb.NewDeploymentServiceClient(conn)

	var (
		records   int
		totalTime int64
	)

	numRecords := *numRecordsPtr
	numIter := *numIterPtr

	if numRecords < 4 {
		log.Fatalf("failed to run benchmark. Minimum number of records required is 4")
	}

	fmt.Println("sending requests...")
	if *isParallelPtr {
		fmt.Println("benchmarking using parallel mode")
	} else {
		fmt.Println("benchmarking using single mode")
	}

	for i := 0; i < numIter; i++ {

		in := createPredictionRequestInputData(rand.Intn(numRecords)+4, 125)
		req := createPredictionRequest(in, "model_v4")

		start := time.Now()

		if *isParallelPtr {
			makeParallelRequests(cbmClient, req)
		} else {
			makeSingleRequests(cbmClient, req)
		}

		elapsed := time.Since(start)

		records += *numRecordsPtr
		totalTime += elapsed.Milliseconds()
	}

	fmt.Println("average time taken:", totalTime/int64(numIter))
	fmt.Println("average records:", records/numIter)
}

func dividePredictionRequest(in *pb.GetPredictionsRequest, n int) []*pb.GetPredictionsRequest {
	s := int(math.Ceil(float64(len(in.InputData) / n)))
	out := make([]*pb.GetPredictionsRequest, n)

	t := len(in.InputData)
	for i := 0; i < n; i++ {
		var inputDatas []*pb.GetPredictionsRequest_InputData
		for j := 0; j < s; j++ {
			inputDatas = append(inputDatas, in.InputData[j])
		}

		out[i] = &pb.GetPredictionsRequest{
			ModelName: in.ModelName,
			ModelTask: in.ModelTask,
			InputData: inputDatas,
		}
		t -= s
		if t < s {
			s = t
		}
	}
	return out
}

func createPredictionRequestInputData(numRecords, numFeatures int) []*pb.GetPredictionsRequest_InputData {
	inputData := make([]*pb.GetPredictionsRequest_InputData, numRecords)

	for i := range inputData {
		data := make(map[string]float32)
		for j := 0; j < numFeatures; j++ {
			featureName := fmt.Sprintf("feature_%s", strconv.Itoa(j))
			data[featureName] = rand.Float32() * float32(rand.Intn(20))
		}
		inputData[i] = &pb.GetPredictionsRequest_InputData{Input: data}
	}

	return inputData
}

func createPredictionRequest(in []*pb.GetPredictionsRequest_InputData, modelName string) *pb.GetPredictionsRequest {
	return &pb.GetPredictionsRequest{
		ModelName: modelName,
		ModelTask: pb.ModelTask_MODEL_TASK_REGRESSION,
		InputData: in,
	}
}

func makeSingleRequests(c pb.DeploymentServiceClient, r *pb.GetPredictionsRequest) {
	_, err := c.GetPredictions(context.Background(), r)
	if err != nil {
		log.Fatalf("failed to call CBM Service: %v", err)
	}
}

func makeParallelRequests(c pb.DeploymentServiceClient, r *pb.GetPredictionsRequest) {
	numPartitions := 4
	reqs := dividePredictionRequest(r, numPartitions)
	out := make([]*pb.GetPredictionsResponse, numPartitions)
	g, ctx := errgroup.WithContext(context.Background())

	for j := 0; j < numPartitions; j++ {
		partitionNum := j
		g.Go(func() error {
			pred, err := c.GetPredictions(ctx, reqs[partitionNum])
			if err != nil {
				log.Fatalf("failed to call CBM Service: %v", err)
			}
			out[partitionNum] = pred
			return nil
		})
	}
	err := g.Wait()
	if err != nil {
		log.Fatalf("failed to get predictions: %s", err)
	}

	_ = combinePredictions(out)
}

func combinePredictions(in []*pb.GetPredictionsResponse) []float64 {
	var out []float64
	for _, predResponse := range in {
		for _, pred := range predResponse.Predictions {
			out = append(out, pred)
		}
	}
	return out
}
