// Server app that servers requests on a certain port and responds to requests
// with calculated response
//
package main

import (
	"fmt"
	"log"
	"net"

	"github.com/goakshit/boromir/config"
	calcPB "github.com/goakshit/boromir/internal/proto/calc"
	calcSVC "github.com/goakshit/boromir/svc/calc"

	"google.golang.org/grpc"
)

func main() {

	// Initiate the config
	config := config.New()

	// Initialise grpc server
	grpcServer := grpc.NewServer()

	// Intialise the calc service and register calc service with grpc server
	calcSVC := calcSVC.NewCalcService()
	calcPB.RegisterCalcServiceServer(grpcServer, calcSVC)

	conn, err := net.Listen("tcp", fmt.Sprintf(":%d", config.ServerPort))
	if err != nil {
		panic(err)
	}

	log.Println("Serving on Port:", config.ServerPort)
	err = grpcServer.Serve(conn)
	if err != nil {
		panic(err)
	}
}
