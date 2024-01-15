package ecs

import "reflect"

type ComponentData interface{}

type Component interface {
	Name() string
	Update()
	Data() ComponentData
	IsA(interface{}) bool
	Type() reflect.Type
}

func CheckComponent[T Component](a interface{}) bool {
	if _, ok := a.(T); ok {
		return true
	}
	return false
}
