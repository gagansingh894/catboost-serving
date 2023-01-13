package main

import (
	"fmt"
	v1 "github.com/gagansingh894/catboost-serving/go/api/v1"
	"github.com/gagansingh894/catboost-serving/go/internal/modelmanager"
	pb "github.com/gagansingh894/catboost-serving/go/internal/pkg/pb/cbserving"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	fmt.Println("Welcome to CATBOOST Serving!")
	modelManager := modelmanager.NewModelManager("../artefacts/")
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
