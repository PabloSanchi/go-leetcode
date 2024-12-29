package pipeline

import (
	"fmt"
	"testing"
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
	return fmt.Sprintf("Length = %d", input), nil
}

type StepThree struct{}

func NewStepThree() *StepThree {
	return &StepThree{}
}

func (s *StepThree) Process(input int) (int, error) {
	return input * 2, nil
}

func TestCreatePipeline(t *testing.T) {
	p := NewPipeline()
	if p == nil {
		t.Fail()
	}
}

func TestExecuteSingleStepPipeline(t *testing.T) {
	p := NewPipeline()
	step := NewStepOne()
	Add(p, step)

	result, err := p.Process("Hello, World!")
	if err != nil {
		t.Fail()
	}
	if result != 13 {
		t.Fail()
	}
}

func TestExecuteMultiStepPipeline(t *testing.T) {
	p := NewPipeline()
	step1 := NewStepOne()
	step2 := NewStepTwo()
	Add(p, step1)
	Add(p, step2)

	result, err := p.Process("Hello, World!")
	if err != nil {
		t.Fail()
	}
	if result != "Length = 13" {
		t.Fail()
	}
}

func TestErrorWhenPipeTypeMismatch(t *testing.T) {
	p := NewPipeline()
	step1 := NewStepOne()
	step2 := NewStepTwo()
	step3 := NewStepThree()
	Add(p, step1)
	Add(p, step2)
	Add(p, step3)

	_, err := p.Process("Hello, World!")
	if err == nil {
		t.Fail()
	}
}
