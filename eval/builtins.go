package eval

import "monkey/object"

var builtins = map[string]*object.Builtin{
	"len": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments, expected 1")
			}

			switch arg := args[0].(type) {
            case *object.Array:
                return &object.Integer{Value: int64(len(arg.Elements))}
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			default:
				return newError("argument to `len` not supported")
			}
		},
	},
    "head": &object.Builtin{
        Fn: func(args ...object.Object) object.Object {
       	    if len(args) != 1 {
				return newError("wrong number of arguments, expected 1")
			}
            if args[0].Type() != object.ARRAY_OBJ {
                return newError("argument must be of type 'ARRAY'")
            }

            arr := args[0].(*object.Array)
            if len(arr.Elements) > 0 {
                return arr.Elements[0]
            }
            return NULL
         },
    },
    "last": &object.Builtin{
        Fn: func(args ...object.Object) object.Object {
       	    if len(args) != 1 {
				return newError("wrong number of arguments, expected 1")
			}
            if args[0].Type() != object.ARRAY_OBJ {
                return newError("argument must be of type 'ARRAY'")
            }

            arr := args[0].(*object.Array)
            length := len(arr.Elements)
            if length <= 0 {
               return NULL 
            }
            return arr.Elements[length - 1]
         },
    },
    "tail": &object.Builtin{
        Fn: func(args ...object.Object) object.Object {
       	    if len(args) != 1 {
				return newError("wrong number of arguments, expected 1")
			}
            if args[0].Type() != object.ARRAY_OBJ {
                return newError("argument must be of type 'ARRAY'")
            }

            arr := args[0].(*object.Array)
            length := len(arr.Elements)
            if length <= 0 {
               return NULL 
            }
            newElements := make([]object.Object, length - 1, length - 1)
            copy(newElements, arr.Elements[1:length])

            return &object.Array{Elements: newElements}
         },
    },
    "append": &object.Builtin{
        Fn: func(args ...object.Object) object.Object {
       	    if len(args) != 2 {
				return newError("wrong number of arguments, expected 2")
			}
            if args[0].Type() != object.ARRAY_OBJ {
                return newError("argument must be of type 'ARRAY'")
            }

            arr := args[0].(*object.Array)
            length := len(arr.Elements)

            newElements := make([]object.Object, length + 1, length + 1)
            copy(newElements, arr.Elements)
            newElements[length] = args[1]

            return &object.Array{Elements: newElements}
         },
    },
}
