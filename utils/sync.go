package utils

import (
	"fmt"
	"runtime/debug"
)

func Recovery() {
	err := recover()
	if err == nil {
		return
	}
	fmt.Printf("panic occur, err: %v, Stack: %v", err, string(debug.Stack()))
}

func SafeGo(fn func()) {
	go func() {
		defer Recovery()
		fn()
	}()
}
