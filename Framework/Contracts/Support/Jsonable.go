package Support

type Jsonable interface {
	//Convert the object to its JSON representation.
	ToJson() string
}
