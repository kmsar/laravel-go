package Support

type Renderable interface {
	// Render Get the evaluated contents of the object.
	Render() string
}
