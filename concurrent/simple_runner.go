// package concurrent
// @author: chenzhewei.97
// @create date: 2025/7/22
package concurrent

import (
	"context"
	"github.com/Ryan1997-tongji/Go-Impl-Countdownlatch/utils"
	"sync"
)

type Runnable interface {
	Run(ctx context.Context)
}

type SimpleRunner struct {
	waitGroup  sync.WaitGroup
	processors []Runnable
}

func NewSimpleRunner(processors []Runnable) *SimpleRunner {
	return &SimpleRunner{
		waitGroup:  sync.WaitGroup{},
		processors: processors,
	}
}

func (s *SimpleRunner) JoinGoroutines(ctx context.Context) error {

	if len(s.processors) <= 0 {
		return nil
	}

	s.waitGroup.Add(len(s.processors))
	for _, _item := range s.processors {
		item := _item
		utils.SafeGo(func() {
			defer s.waitGroup.Done()
			item.Run(ctx)
		})
	}

	s.waitGroup.Wait()
	return nil

}
