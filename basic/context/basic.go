package main

import (
	"context"
	"fmt"
	"time"
)

// 上下文基础示例

func main() {
	parentCtx := context.Background()
	ctx, cancel := context.WithTimeout(parentCtx, 10000*time.Millisecond)
	defer cancel()

	func(ctx context.Context) {
		for {
			select {
			case <-time.After(1 * time.Second):
				fmt.Println("overslept")
			case <-ctx.Done():
				fmt.Println(ctx.Err())
				return
			}
		}
	}(ctx)
}
