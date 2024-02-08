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
            if length > 0 {
                return arr.Elements[length - 1]
            }
            return NULL
         },
    },
}
