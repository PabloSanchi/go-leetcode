package main

import (
	"fmt"
	"pipeline"
)

type StepOne struct{}

func NewStepOne() *StepOne {
	return &StepOne{}
}

func (s *StepOne) Process(input string) (int, error) {
	size := len(input)
	return size, nil
}

type StepTwo struct{}

func NewStepTwo() *StepTwo {
	return &StepTwo{}
}

func (s *StepTwo) Process(input int) (string, error) {
	return fmt.Sprintf("The length of the string is %d", input), nil
}

func main() {
	step1 := NewStepOne()
	step2 := NewStepTwo()

	p := pipeline.NewPipeline()
	pipeline.Add(p, step1)
	pipeline.Add(p, step2)

	result, err := p.Process("Hello, World!")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
