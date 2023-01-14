package main

import (
	"flag"
	"fmt"
	v1 "github.com/gagansingh894/catboost-serving/api/v1"
	"github.com/gagansingh894/catboost-serving/internal/modelmanager"
	pb "github.com/gagansingh894/catboost-serving/internal/pkg/pb/cbserving"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	// parse cmd arguments
	artefactsPathPtr := flag.String("artefacts-path", "artefacts/", " path to directory containing models")
	portPtr := flag.String("port", "9090", " port number for gRPC server")
	flag.Parse()

	fmt.Println("Welcome to CATBOOST Serving!")
	modelManager := modelmanager.NewModelManager(*artefactsPathPtr)
	err := modelManager.Initialize()
	if err != nil {
		log.Fatalf("failed to initialize model manager: %s", err)
	}

	// configure grpc server
	cbDeploymentService := v1.NewCBDeploymentService(modelManager)
	cbDeploymentServer := grpc.NewServer()
	pb.RegisterDeploymentServiceServer(cbDeploymentServer, cbDeploymentService)
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", *portPtr))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("server listening at %v", lis.Addr())

	if err := cbDeploymentServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
