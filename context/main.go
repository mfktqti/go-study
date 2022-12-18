package main

import (
	"context"
	"fmt"
	"time"
)

type MyContext struct {
	// 这里的 Context 是我 copy 出来的，所以前面不用加 context.
	context.Context
}

func main() {
	childCancel := false

	parentCtxCtx := context.Background()

	parentCtx, parentFunc := context.WithCancel(parentCtxCtx)
	mctx := MyContext{parentCtx}

	childCtx, childFun := context.WithCancel(mctx)
	ctx2, _ := context.WithTimeout(childCtx, 2*time.Second)
	ctx, _ := context.WithTimeout(ctx2, 3*time.Second)
	time.Sleep(5 * time.Second)
	fmt.Println(parentCtxCtx)
	fmt.Println(ctx2.Err())
	fmt.Println(ctx.Err())
	fmt.Println(parentCtx.Err())
	fmt.Println(mctx.Err())
	fmt.Println(childCtx.Err())
	// //cf()
	if childCancel {
		childFun()
	} else {
		parentFunc()
	}

	fmt.Println(ctx.Err())
	fmt.Println(parentCtx.Err())
	fmt.Println(mctx.Err())
	fmt.Println(childCtx.Err())

	// 防止主协程退出太快，子协程来不及打印
	time.Sleep(10 * time.Second)
}
