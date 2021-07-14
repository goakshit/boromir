package calc

import (
	"context"
	"fmt"
	"testing"

	"github.com/goakshit/boromir/internal/proto/calc"
)

func TestAdd(t *testing.T) {

	// Prepare the test data
	data := []struct {
		a   float64
		b   float64
		res float64
		err error
	}{
		{1, 2, 3, nil},
		{-2, 2, 0, nil},
		{2, -2, 0, nil},
		{-2, -2, -4, nil},
		{-2.5, -2, -4.5, nil},
	}

	calcSVC := NewCalcService()

	for _, d := range data {
		res, err := calcSVC.Add(context.Background(), &calc.CalcRequest{
			A: d.a,
			B: d.b,
		})
		if err != nil && d.err.Error() != err.Error() {
			t.Errorf("Add(%v, %v) returned error: %v", d.a, d.b, err)
		}
		if res != nil && res.Result != d.res {
			t.Errorf("Add(%v, %v) returned %v, want %v", d.a, d.b, res.Result, d.res)
		}
	}
}

func TestSub(t *testing.T) {

	// Prepare the test data
	data := []struct {
		a   float64
		b   float64
		res float64
		err error
	}{
		{1, 2, -1, nil},
		{-2, 2, -4, nil},
		{2, -2, 4, nil},
		{-2, -2, 0, nil},
		{3, 2, 1, nil},
		{10.75, 2.5, 8.25, nil},
	}

	calcSVC := NewCalcService()
	for _, d := range data {
		res, err := calcSVC.Sub(context.Background(), &calc.CalcRequest{
			A: d.a,
			B: d.b,
		})
		if err != nil && d.err.Error() != err.Error() {
			t.Errorf("Sub(%v, %v) returned error: %v", d.a, d.b, err)
		}
		if res != nil && res.Result != d.res {
			t.Errorf("Sub(%v, %v) returned %v, want %v", d.a, d.b, res.Result, d.res)
		}
	}
}

func TestMul(t *testing.T) {

	// Prepare the test data
	data := []struct {
		a   float64
		b   float64
		res float64
		err error
	}{
		{1, 2, 2, nil},
		{-2, 2, -4, nil},
		{2, -2, -4, nil},
		{-2, -2, 4, nil},
		{0, 1, 0, nil},
		{2.5, 7, 17.5, nil},
	}

	calcSVC := NewCalcService()
	for _, d := range data {
		res, err := calcSVC.Mul(context.Background(), &calc.CalcRequest{
			A: d.a,
			B: d.b,
		})
		if err != nil && d.err.Error() != err.Error() {
			t.Errorf("Mul(%v, %v) returned error: %v", d.a, d.b, err)
		}
		if res != nil && res.Result != d.res {
			t.Errorf("Mul(%v, %v) returned %v, want %v", d.a, d.b, res.Result, d.res)
		}

	}
}

func TestDiv(t *testing.T) {

	// Prepare the test data
	data := []struct {
		a   float64
		b   float64
		res float64
		err error
	}{
		{1, 2, 0.5, nil},
		{-2, 2, -1, nil},
		{2, -2, -1, nil},
		{-2, -2, 1, nil},
		{10, 0, 0, fmt.Errorf("Divide by zero error")},
		{14.5, 7, 2.0714285714285716, nil},
	}

	calcSVC := NewCalcService()
	for _, d := range data {
		res, err := calcSVC.Div(context.Background(), &calc.CalcRequest{
			A: d.a,
			B: d.b,
		})
		if err != nil && d.err.Error() != err.Error() {
			t.Errorf("Div(%v, %v) returned error: %v", d.a, d.b, err)
		}
		if res != nil && res.Result != d.res {
			t.Errorf("Div(%v, %v) returned %v, want %v", d.a, d.b, res.Result, d.res)
		}
	}
}
