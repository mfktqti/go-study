package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			_, file, line, _ := runtime.Caller(3)
			fmt.Printf("string(debug.Stack()): %v\n", string(debug.Stack()))
			fmt.Printf("file(%d):%s, err:%+v\n", line, file, err)
		}
	}()
	add()
}

func add() int {
	d := 0
	fmt.Printf("%d", 33/d)
	return 1 + 1
}
