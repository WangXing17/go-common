package data_struct

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	var stack *ItemStack = NewStack()
	for i := 0; i < 10; i++ {
		stack.Push(i)
	}
	for !stack.Empty() {
		var v int = stack.Pop().(int)
		fmt.Println(v)
	}
}
