package Container

// Scoped	1
// Specifies that a new instance of the service will be created for each scope. In ASP.NET Core apps, a scope is created around each server Request.
//
// Singleton	0
// Specifies that a single instance of the service will be created.
//
// Transient	2
// Specifies that a new instance of the service will be created every time it is requested. =
// int mapping

type ServiceLifeTime uint

// string mapping
const (
	Singleton ServiceLifeTime = iota // single instance of the service will be created.
	Scoped                           // each server Request.
	Transient                        //a new instance of the service will be created every time it is requested
)

func (s ServiceLifeTime) String() string {
	switch s {
	case Singleton:
		return "Singleton"
	case Scoped:
		return "Scoped"
	case Transient:
		return "Transient"
	}
	return "unknown"
}

func FromInt(l uint) ServiceLifeTime {
	switch l {
	case 0:
		return Singleton
	case 1:
		return Scoped
	case 2:
		return Transient
	}
	return Singleton
}
