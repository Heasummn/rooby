package vm

import (
	"math"
	"strconv"
)

var (
	integerClass *RInteger
)

// RInteger is integer class
type RInteger struct {
	*BaseClass
}

// IntegerObject represents integer instances
type IntegerObject struct {
	Class *RInteger
	Value int
}

func (i *IntegerObject) objectType() objectType {
	return integerObj
}

func (i *IntegerObject) Inspect() string {
	return strconv.Itoa(i.Value)
}

func (i *IntegerObject) returnClass() Class {
	return i.Class
}

func (i *IntegerObject) equal(e *IntegerObject) bool {
	return i.Value == e.Value
}

func initilaizeInteger(value int) *IntegerObject {
	return &IntegerObject{Value: value, Class: integerClass}
}

var builtinIntegerMethods = []*BuiltInMethod{
	{
		Fn: func(receiver Object) builtinMethodBody {
			return func(vm *VM, args []Object, blockFrame *callFrame) Object {

				leftValue := receiver.(*IntegerObject).Value
				right, ok := args[0].(*IntegerObject)

				if !ok {
					return wrongTypeError(integerClass)
				}

				rightValue := right.Value
				return &IntegerObject{Value: leftValue + rightValue, Class: integerClass}
			}
		},
		Name: "+",
	},
	{
		Fn: func(receiver Object) builtinMethodBody {
			return func(vm *VM, args []Object, blockFrame *callFrame) Object {

				leftValue := receiver.(*IntegerObject).Value
				right, ok := args[0].(*IntegerObject)

				if !ok {
					return wrongTypeError(integerClass)
				}

				rightValue := right.Value
				return &IntegerObject{Value: leftValue - rightValue, Class: integerClass}
			}
		},
		Name: "-",
	},
	{
		Fn: func(receiver Object) builtinMethodBody {
			return func(vm *VM, args []Object, blockFrame *callFrame) Object {

				leftValue := receiver.(*IntegerObject).Value
				right, ok := args[0].(*IntegerObject)

				if !ok {
					return wrongTypeError(integerClass)
				}

				rightValue := right.Value
				return &IntegerObject{Value: leftValue * rightValue, Class: integerClass}
			}
		},
		Name: "*",
	},
	{
		Fn: func(receiver Object) builtinMethodBody {
			return func(vm *VM, args []Object, blockFrame *callFrame) Object {

				leftValue := receiver.(*IntegerObject).Value
				right, ok := args[0].(*IntegerObject)

				if !ok {
					return wrongTypeError(integerClass)
				}

				rightValue := right.Value
				result := math.Pow(float64(leftValue), float64(rightValue))
				return &IntegerObject{Value: int(result), Class: integerClass}
			}
		},
		Name: "**",
	},
	{
		Fn: func(receiver Object) builtinMethodBody {
			return func(vm *VM, args []Object, blockFrame *callFrame) Object {

				leftValue := receiver.(*IntegerObject).Value
				right, ok := args[0].(*IntegerObject)

				if !ok {
					return wrongTypeError(integerClass)
				}

				rightValue := right.Value
				return &IntegerObject{Value: leftValue / rightValue, Class: integerClass}
			}
		},
		Name: "/",
	},
	{
		Fn: func(receiver Object) builtinMethodBody {
			return func(vm *VM, args []Object, blockFrame *callFrame) Object {

				leftValue := receiver.(*IntegerObject).Value
				right, ok := args[0].(*IntegerObject)

				if !ok {
					return wrongTypeError(integerClass)
				}

				rightValue := right.Value

				if leftValue > rightValue {
					return TRUE
				}

				return FALSE
			}
		},
		Name: ">",
	},
	{
		Fn: func(receiver Object) builtinMethodBody {
			return func(vm *VM, args []Object, blockFrame *callFrame) Object {

				leftValue := receiver.(*IntegerObject).Value
				right, ok := args[0].(*IntegerObject)

				if !ok {
					return wrongTypeError(integerClass)
				}

				rightValue := right.Value

				if leftValue >= rightValue {
					return TRUE
				}

				return FALSE
			}
		},
		Name: ">=",
	},
	{
		Fn: func(receiver Object) builtinMethodBody {
			return func(vm *VM, args []Object, blockFrame *callFrame) Object {

				leftValue := receiver.(*IntegerObject).Value
				right, ok := args[0].(*IntegerObject)

				if !ok {
					return wrongTypeError(integerClass)
				}

				rightValue := right.Value

				if leftValue < rightValue {
					return TRUE
				}

				return FALSE
			}
		},
		Name: "<",
	},
	{
		Fn: func(receiver Object) builtinMethodBody {
			return func(vm *VM, args []Object, blockFrame *callFrame) Object {

				leftValue := receiver.(*IntegerObject).Value
				right, ok := args[0].(*IntegerObject)

				if !ok {
					return wrongTypeError(integerClass)
				}

				rightValue := right.Value

				if leftValue <= rightValue {
					return TRUE
				}

				return FALSE
			}
		},
		Name: "<=",
	},
	{
		Fn: func(receiver Object) builtinMethodBody {
			return func(vm *VM, args []Object, blockFrame *callFrame) Object {

				leftValue := receiver.(*IntegerObject).Value
				right, ok := args[0].(*IntegerObject)

				if !ok {
					return wrongTypeError(integerClass)
				}

				rightValue := right.Value

				if leftValue < rightValue {
					return initilaizeInteger(-1)
				}
				if leftValue > rightValue {
					return initilaizeInteger(1)
				}

				return initilaizeInteger(0)
			}
		},
		Name: "<=>",
	},
	{
		Fn: func(receiver Object) builtinMethodBody {
			return func(vm *VM, args []Object, blockFrame *callFrame) Object {

				leftValue := receiver.(*IntegerObject).Value
				right, ok := args[0].(*IntegerObject)

				if !ok {
					return wrongTypeError(integerClass)
				}

				rightValue := right.Value

				if leftValue == rightValue {
					return TRUE
				}

				return FALSE
			}
		},
		Name: "==",
	},
	{
		Fn: func(receiver Object) builtinMethodBody {
			return func(vm *VM, args []Object, blockFrame *callFrame) Object {

				leftValue := receiver.(*IntegerObject).Value
				right, ok := args[0].(*IntegerObject)

				if !ok {
					return wrongTypeError(integerClass)
				}

				rightValue := right.Value

				if leftValue != rightValue {
					return TRUE
				}

				return FALSE
			}
		},
		Name: "!=",
	},
	{
		Fn: func(receiver Object) builtinMethodBody {
			return func(vm *VM, args []Object, blockFrame *callFrame) Object {

				int := receiver.(*IntegerObject)
				int.Value++
				return int
			}
		},
		Name: "++",
	},
	{
		Fn: func(receiver Object) builtinMethodBody {
			return func(vm *VM, args []Object, blockFrame *callFrame) Object {

				int := receiver.(*IntegerObject)
				int.Value--
				return int
			}
		},
		Name: "--",
	},
	{
		Fn: func(receiver Object) builtinMethodBody {
			return func(vm *VM, args []Object, blockFrame *callFrame) Object {

				int := receiver.(*IntegerObject)

				return initializeString(strconv.Itoa(int.Value))
			}
		},
		Name: "to_s",
	},
	{
		Fn: func(receiver Object) builtinMethodBody {
			return func(vm *VM, args []Object, blockFrame *callFrame) Object {
				return receiver
			}
		},
		Name: "to_i",
	},
	{
		Fn: func(receiver Object) builtinMethodBody {
			return func(vm *VM, args []Object, blockFrame *callFrame) Object {

				i := receiver.(*IntegerObject)
				even := i.Value%2 == 0

				if even {
					return TRUE
				}

				return FALSE
			}
		},
		Name: "even",
	},
	{
		Fn: func(receiver Object) builtinMethodBody {
			return func(vm *VM, args []Object, blockFrame *callFrame) Object {

				i := receiver.(*IntegerObject)
				odd := i.Value%2 != 0
				if odd {
					return TRUE
				}

				return FALSE
			}
		},
		Name: "odd",
	},
	{
		Fn: func(receiver Object) builtinMethodBody {
			return func(vm *VM, args []Object, blockFrame *callFrame) Object {
				i := receiver.(*IntegerObject)
				return initilaizeInteger(i.Value + 1)
			}
		},
		Name: "next",
	},
	{
		Fn: func(receiver Object) builtinMethodBody {
			return func(vm *VM, args []Object, blockFrame *callFrame) Object {
				i := receiver.(*IntegerObject)
				return initilaizeInteger(i.Value - 1)
			}
		},
		Name: "pred",
	},
	{
		Fn: func(receiver Object) builtinMethodBody {
			return func(vm *VM, args []Object, blockFrame *callFrame) Object {
				n := receiver.(*IntegerObject)

				if n.Value < 0 {
					return newError("Expect paramentr to be greater 0. got=%d", n.Value)
				}

				if blockFrame == nil {
					return newError("Can't yield without a block")
				}

				for i := 0; i < n.Value; i++ {
					builtInMethodYield(vm, blockFrame, initilaizeInteger(i))
				}

				return n
			}
		},
		Name: "times",
	},
}

func initInteger() {
	methods := newEnvironment()

	for _, m := range builtinIntegerMethods {
		methods.set(m.Name, m)
	}

	bc := &BaseClass{Name: "Integer", Methods: methods, ClassMethods: newEnvironment(), Class: classClass, pseudoSuperClass: objectClass, superClass: objectClass}
	ic := &RInteger{BaseClass: bc}
	integerClass = ic
}
