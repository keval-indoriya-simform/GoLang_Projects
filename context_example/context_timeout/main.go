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
	ctx, cancleCtx := context.WithTimeout(ctx, 1500*time.Millisecond)

	printCh := make(chan int)
	go doAnother(ctx, printCh)

	for num := 1; num <= 3; num++ {
		select {
		case printCh <- num:
			time.Sleep(1 * time.Second)
		case <-ctx.Done():
			break
		}
	}

	defer cancleCtx()

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
