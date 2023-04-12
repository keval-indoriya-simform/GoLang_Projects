package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//ctx := context.TODO()
	ctx := context.Background()
	doSomething(ctx)
}

func doSomething(ctx context.Context) {
	ctx, cancleCtx := context.WithCancel(ctx)

	printCh := make(chan int)
	go doAnother(ctx, printCh)

	for num := 1; num <= 3; num++ {
		printCh <- num
	}

	cancleCtx()

	time.Sleep(100 * time.Millisecond)
	fmt.Println("doSomething -> finished")
}

func doAnother(ctx context.Context, printCh <-chan int) {
	for {
		select {
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				fmt.Printf("doAnother err: %s\n", err)
			}
			fmt.Println("doAnother Finished")
			return
		case num := <-printCh:
			fmt.Printf("doAnother : %d\n", num)
		}
	}
}
