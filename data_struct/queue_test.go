package data_struct

import (
	"fmt"
	"testing"
)

func Test_Queue(t *testing.T) {
	queue := NewQueue()
	for i := 0; i < 10; i++ {
		queue.Offer(i)
	}

	for !queue.Empty() {
		var v int = queue.Pop().(int)
		fmt.Println(v)
	}
}
