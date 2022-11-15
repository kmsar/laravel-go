package myco

type Extend struct {
	Category string
	Flags    []Flag
}

type IContext interface {
	Argument(index int) string
	Arguments() []string
	Option(key string) string
}

type Flag struct {
	Name     string
	Aliases  []string
	Usage    string
	Required bool
	Value    string
}

type ICommand interface {
	//Signature The name and signature of the console command.
	Signature() string
	//Description The console command description.
	Description() string
	//Extend The console command extend.
	Extend() Extend
	//Handle Execute the console command.
	Handle(ctx IContext) error
}
