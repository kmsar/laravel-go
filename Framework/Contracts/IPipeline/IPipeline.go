package IPipeline

type Pipe func(passable interface{}) interface{}

type Pipeline interface {

	// Send Set the object being sent through the pipeline.
	Send(passable interface{}) Pipeline

	// Through Set the array of pipes.
	Through(pipes ...interface{}) Pipeline

	// Then Run the pipeline with a final destination callback.
	Then(destination interface{}) interface{}
}
