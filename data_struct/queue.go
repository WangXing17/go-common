package data_struct

// ItemQueue 是一个通用队列，可以存储任意类型的数据
type ItemQueue struct {
	items []any
}

// NewQueue 创建一个新的队列
func NewQueue() *ItemQueue {
	return &ItemQueue{
		items: make([]any, 0),
	}
}

// Offer 将元素入队
func (q *ItemQueue) Offer(item any) {
	q.items = append(q.items, item)
}

// Pop 元素出队列
func (q *ItemQueue) Pop() any {
	if len(q.items) == 0 {
		return nil
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

// Empty 检查队列是否为空
func (q *ItemQueue) Empty() bool {
	return len(q.items) == 0
}

// Size 返回队列中元素的数量
func (q *ItemQueue) Size() int {
	return len(q.items)
}

// Clear 清空队列中的所有元素
func (q *ItemQueue) Clear() {
	q.items = q.items[:0]
}
