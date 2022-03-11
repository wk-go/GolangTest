package main

import (
	"context"
	"log"
	"time"
)

// 上下文基础示例

func main() {
	parentCtx := context.Background()

	//with timeout
	ctx, cancel := context.WithTimeout(parentCtx, 2*time.Second)
	defer cancel()
	func(ctx context.Context) {
		for {
			select {
			case <-time.After(300 * time.Millisecond):
				log.Println("Doing something...")
			case <-ctx.Done():
				log.Println(ctx.Err())
				return
			}
		}
	}(ctx)

	// with Value
	withValueFunc := func(ctx context.Context) {
		val := ctx.Value("name")
		name := val.(string)
		log.Println("The name is", name)
	}
	ctxWithValue := context.WithValue(parentCtx, "name", "zhang san")
	withValueFunc(ctxWithValue)

	// with Cancel
	withCancelFunc := func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				log.Println(ctx.Err())
				return
			default:
				log.Println("Working...")
				time.Sleep(200 * time.Millisecond)
			}
		}
	}
	ctxWithCancel, cancel2 := context.WithCancel(parentCtx)
	go withCancelFunc(ctxWithCancel)
	for i := 0; i < 2; i++ {
		time.Sleep(time.Second)
	}
	cancel2()

	log.Println("Done.")
}
