package main

import (
	"context"
	"fmt"
)

func main() {
	//ctx := context.TODO()
	ctx := context.Background()
	ctx = context.WithValue(ctx, "PIN", "7123")
	doSomething(ctx)

}

func doSomething(ctx context.Context) {
	fmt.Printf("doSomething -> My key's value is %s\n", ctx.Value("PIN"))

	anotherCtx := context.WithValue(ctx, "PIN", "6454")
	doAnother(anotherCtx)

	fmt.Printf("doSomething -> My key's value is %s\n", ctx.Value("PIN"))
}

func doAnother(ctx context.Context) {
	fmt.Printf("doAnother -> My key's value is %s\n", ctx.Value("PIN"))
}
