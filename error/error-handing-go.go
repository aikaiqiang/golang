package main

import "fmt"

type myErrorInfo struct {
	s string
}

func (e *myErrorInfo) Error() string {
	return e.s
}

func new(text string) error {
	return &myErrorInfo{text}
}

type NegativeSqrtError float64

func (f NegativeSqrtError) Error() string {
	return fmt.Sprintf("math: square root of negative number %g", float64(f))
}

type NetworkError interface {
	error
	Timeout() bool
	Tempoary() bool
}

func main() {

}
