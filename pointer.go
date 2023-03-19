package objects

import "reflect"

func memaddr(ptr any) uintptr {
	value := reflect.ValueOf(ptr)
	return value.Pointer()
}
