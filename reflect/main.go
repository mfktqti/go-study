package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 100

	t := reflect.TypeOf(a)
	v := reflect.ValueOf(a)
	fmt.Printf("t: %v\n", t)
	fmt.Printf("v: %v\n", v)

	fmt.Printf("v.CanAddr(): %v\n", v.CanAddr())

	n1, n2 := 2, 3
	args := []reflect.Value{reflect.ValueOf(n1), reflect.ValueOf(n2)}
	v2 := reflect.ValueOf(reflectCallFunc).Call(args)[0]
	fmt.Printf("v2.Kind(): %v\n", v2.Kind())
	fmt.Printf("v2.Interface(): %v\n", v2.Interface())
	fmt.Printf("  return: type=%v, value=[%d]\n", v2.Type(), v2.Int())
}

func reflectCallFunc(n1, n2 int) int {
	return n1 + n2
}
