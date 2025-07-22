// package Go_Impl_Countdownlatch
// @author: chenzhewei.97
// @create date: 2025/7/22
package main

import (
	"context"
	"fmt"
	"github.com/Ryan1997-tongji/Go-Impl-Countdownlatch/concurrent"
	"github.com/Ryan1997-tongji/Go-Impl-Countdownlatch/impl"
)

func main() {
	ctx := context.Background()
	addProcessors := []concurrent.Runnable{impl.NewTestProcessor(ctx, 1, 2), impl.NewTestProcessor(ctx, 3, 4), impl.NewTestProcessor(ctx, 5, 6)}
	simpleRunner := concurrent.NewSimpleRunner(addProcessors)
	_ = simpleRunner.JoinGoroutines(ctx)
	var finalResult int64
	for _, p := range addProcessors {
		output := p.(*impl.TestProcessor).GetResult(ctx)
		finalResult += output
	}
	fmt.Println("Final Result:", finalResult)
}
