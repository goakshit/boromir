// Package calc provides methods for basic arithmetic operations such as
// adding, subtracting, multiplying and dividing.
package calc

import (
	"context"
	"fmt"
	"log"

	"github.com/goakshit/boromir/internal/proto/calc"
)

// Implements CalcServiceServer interface generated through proto/calc/calc.proto
type Service struct{}

// NewCalcService returns CalcServiceServer interface
func NewCalcService() calc.CalcServiceServer {
	return &Service{}
}

// Add is used for arithmetic addition operation
// Returns result and error if any
func (s *Service) Add(ctx context.Context, req *calc.CalcRequest) (*calc.CalcResponse, error) {
	log.Printf("Function name: Add; Params: a=%f, b=%f", req.GetA(), req.GetB())
	a := req.GetA()
	b := req.GetB()
	return &calc.CalcResponse{
		Result: a + b,
	}, nil
}

// Sub is used for arithmetic subtraction operation
// Returns result and error if any
func (s *Service) Sub(ctx context.Context, req *calc.CalcRequest) (*calc.CalcResponse, error) {
	log.Printf("Function name: Sub; Params: a=%f, b=%f", req.GetA(), req.GetB())
	a := req.GetA()
	b := req.GetB()
	return &calc.CalcResponse{
		Result: a - b,
	}, nil
}

// Mul is used for arithmetic multiplication operation
// Returns result and error if any
func (s *Service) Mul(ctx context.Context, req *calc.CalcRequest) (*calc.CalcResponse, error) {
	log.Printf("Function name: Mul; Params: a=%f, b=%f", req.GetA(), req.GetB())
	a := req.GetA()
	b := req.GetB()
	return &calc.CalcResponse{
		Result: a * b,
	}, nil
}

// Div is used for arithmetic division operation
// Returns result and error if any
func (s *Service) Div(ctx context.Context, req *calc.CalcRequest) (*calc.CalcResponse, error) {
	log.Printf("Function name: Div; Params: a=%f, b=%f", req.GetA(), req.GetB())
	a := req.GetA()
	b := req.GetB()

	if b == 0 {
		log.Printf("Divide by zero error for params: a=%f, b=%f", a, b)
		return nil, fmt.Errorf("Divide by zero error")
	}

	return &calc.CalcResponse{
		Result: a / b,
	}, nil
}
