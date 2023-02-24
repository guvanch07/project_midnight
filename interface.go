package main

import . "io"

type calculator interface {
	add(float64, float64) float64
	subtract(float64, float64) float64
	multi(float64, float64) float64
	div(float64, float64) float64
}

type adder interface {
	add(float64, float64) float64
}

type adderMulter interface {
	add(float64, float64) float64
	multi(float64, float64) float64
}

type myCalc struct{}

func (m myCalc) add(a, b float64) float64 {
	return a + b
}

func (m myCalc) subtract(a, b float32) float64 {
	return float64(a - b)
}

func (m myCalc) multi(a, b float64) float64 {
	return float64(a * b)
}

func (m myCalc) div(a, b float64) float64 {

	return a / b
}

type ReadWriter interface {
	Reader // это io.Reader
	Writer // это io.Writer
}
