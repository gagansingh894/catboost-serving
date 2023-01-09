package main

import (
	"fmt"
	v1 "github.com/gagansingh894/catboost-serving/api/v1"
	"github.com/gagansingh894/catboost-serving/internal/modelmanager"
	pb "github.com/gagansingh894/catboost-serving/internal/pkg/pb/cbserving"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"net"
)

func main() {
	fmt.Println("Welcome to CATBOOST Serving!")
	modelManager := modelmanager.NewModelManager("artefacts/")
	err := modelManager.Initialize()
	if err != nil {
		log.Fatalf("failed to initialize model manager: %s", err)
	}

	// start grpc server
	cbDeploymentService := v1.NewCBDeploymentService(modelManager)
	cbDeploymentServer := grpc.NewServer()
	pb.RegisterDeploymentServiceServer(cbDeploymentServer, cbDeploymentService)
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", "9090"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("server listening at %v", lis.Addr())

	if err := cbDeploymentServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

func generateInputs(numFeatures, numRecords int) [][]float32 {
	data := make([][]float32, numRecords) // numFeatures
	rows := make([]float32, numFeatures)  // numRecords

	for i := 0; i < numRecords; i++ {
		for j := 0; j < numFeatures; j++ {
			rows[j] = rand.Float32() * float32(rand.Intn(1000))
		}
		data[i] = rows
	}

	return data
}
