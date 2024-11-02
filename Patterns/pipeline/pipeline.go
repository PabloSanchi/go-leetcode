package pipeline

import "fmt"

// Primitive is a generic type def
type Primitive interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64 | string | bool
}

// Pipe is a generic transformation with independent input and output types
type Pipe[IN Primitive, OUT Primitive] interface {
	Process(input IN) (OUT, error)
}

// Executable is a function that takes any input and returns any output (wrapper)
type Executable func(input any) (any, error)

// Pipeline is a sequence of Pipes to be executed in order
type Pipeline struct {
	pipes []Executable
}

// NewPipeline creates a new Pipeline
func NewPipeline() *Pipeline {
	return &Pipeline{pipes: []Executable{}}
}

// Add allows adding any Pipe, adapting each Pipe’s input-output types to the Pipeline
func Add[IN Primitive, OUT Primitive](p *Pipeline, pipe Pipe[IN, OUT]) {
	executable := func(input any) (any, error) {
		typedInput, ok := input.(IN)
		if !ok {
			return nil, fmt.Errorf("type mismatch: expected %T but got %T", *new(IN), input)
		}
		return pipe.Process(typedInput)
	}
	p.pipes = append(p.pipes, executable)
}

// Process executes each pipe in sequence, passing output to the next as input
func (p *Pipeline) Process(input any) (any, error) {
	var output = input
	var err error

	for _, executable := range p.pipes {
		output, err = executable(output)
		if err != nil {
			return nil, err
		}
	}
	return output, nil
}
