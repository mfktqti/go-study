package main

import (
	"log"
	"reflect"
)

type route_interface interface {
	Init()
}
type Routers struct {
	m MenuRoute
	u UserRoute
}

func configRoute() {
	r := Routers{}
	//t := reflect.TypeOf(r)
	v := reflect.ValueOf(r)
	for i := 0; i < v.NumField(); i++ {
		fieldValue := reflect.New(v.Field(i).Type())
		methodValue := fieldValue.MethodByName("Init")
		if (methodValue == reflect.Value{}) {
			log.Fatalf("%v is not route_interface", v.Field(i).Type().Name())
		}
		methodValue.Call(nil)
		//fieldValue.Method(0).Call(nil)
	}
}
