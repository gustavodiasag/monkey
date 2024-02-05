package object

type Environment struct {
	store map[string]Object
}

func NewEnv() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s}
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	return obj, ok
}

func (e *Environment) Insert(name string, val Object) Object {
	e.store[name] = val
	return val
}
