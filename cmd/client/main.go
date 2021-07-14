// Client app that takes command line arguments and based on method passed,
// makes gRPC calls to the respective stub with the arguments.
// Example:
//  ./client -method add -a 1 -b 2
//  3
package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/goakshit/boromir/config"
	calcPB "github.com/goakshit/boromir/internal/proto/calc"
	"google.golang.org/grpc"
)

func main() {

	var (
		calcResponse *calcPB.CalcResponse
		err          error
	)

	// Initialise the config
	config := config.New()

	// Initialise the command line flags
	method := flag.String("method", "add", "calculation method")
	a := flag.Float64("a", 0, "First value")
	b := flag.Float64("b", 0, "Second value")

	// Parse the flags
	flag.Parse()

	conn, e := grpc.Dial(fmt.Sprintf("%s:%d", config.ServerHost, config.ServerPort), grpc.WithInsecure())
	if e != nil {
		panic(e)
	}
	defer conn.Close()

	// Initialise the calc service client
	client := calcPB.NewCalcServiceClient(conn)

	// Based on method, call the respective gRPC method
	switch *method {
	case "add":
		calcResponse, err = client.Add(context.Background(), &calcPB.CalcRequest{A: *a, B: *b})
	case "sub":
		calcResponse, err = client.Sub(context.Background(), &calcPB.CalcRequest{A: *a, B: *b})
	case "mul":
		calcResponse, err = client.Mul(context.Background(), &calcPB.CalcRequest{A: *a, B: *b})
	case "div":
		calcResponse, err = client.Div(context.Background(), &calcPB.CalcRequest{A: *a, B: *b})
	default:
		log.Println("Unknown method:", *method)
		return
	}

	if err != nil {
		log.Println(err.Error())
		return
	}
	fmt.Println(calcResponse.GetResult())
}
