package object

type Environment struct {
	store map[string]Object
    outer *Environment
}

func NewEnv() *Environment {
	s := make(map[string]Object)
    return &Environment{store: s, outer: nil}
}

func NewEnclosedEnvironment(outer *Environment) *Environment {
    env := NewEnv()
    env.outer = outer

    return env
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
    if !ok && e.outer != nil {
        obj, ok = e.outer.Get(name)
    }
	return obj, ok
}

func (e *Environment) Insert(name string, val Object) Object {
	e.store[name] = val
	return val
}
