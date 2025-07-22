// package impl
// @author: chenzhewei.97
// @create date: 2025/7/22
package impl

import "context"

type TestProcessor struct {
	output int64
	num1   int64
	num2   int64
}

func NewTestProcessor(ctx context.Context, num1 int64, num2 int64) *TestProcessor {
	return &TestProcessor{
		num1:   num1,
		num2:   num2,
		output: 0,
	}
}

func (p *TestProcessor) Run(ctx context.Context) {
	p.output = p.num1 + p.num2
}

func (p *TestProcessor) GetResult(ctx context.Context) int64 {
	return p.output
}
