package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	fmt.Printf("This is the cance function: %#v\n", cancel)
	testCtx(ctx)
}

func testCtx(ctx context.Context) {
	fmt.Printf("This is the context: %#v\n", ctx)
}
