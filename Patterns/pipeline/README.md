# Pipeline Pattern in Go
The pipeline pattern is a design pattern used to process data in a sequence of stages.
Each stage is a self-contained unit of work that processes data and passes it to the next stage.
The pattern is useful when you have a large process that can be broken into steps, this pattern helps with that
but be aware that it can also introduce complexity and this implementation is suitable only for small pipelines.