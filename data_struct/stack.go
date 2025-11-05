package data_struct

// ItemStack 是一个通用栈，可以存储任何类型的数据
type ItemStack struct {
	items []any
}

// NewStack 创建一个新的栈
func NewStack() *ItemStack {
	return &ItemStack{
		items: make([]any, 0),
	}
}

// Push 将元素压入栈顶
func (s *ItemStack) Push(item any) {
	s.items = append(s.items, item)
}

// Pop 弹出栈顶元素并返回，如果栈为空则返回 nil
func (s *ItemStack) Pop() any {
	if len(s.items) == 0 {
		return nil
	}
	var v = s.Peek()
	s.items = s.items[:len(s.items)-1]
	return v
}

// Peek 返回栈顶元素但不移除，如果栈为空则返回 nil
func (s *ItemStack) Peek() any {
	if len(s.items) == 0 {
		return nil
	}
	return s.items[len(s.items)-1]
}

// Empty 检查栈是否为空
func (s *ItemStack) Empty() bool {
	return len(s.items) == 0
}

// Size 返回栈中元素的数量
func (s *ItemStack) Size() int {
	return len(s.items)
}

// Clear 清空栈中的所有元素
func (s *ItemStack) Clear() {
	s.items = s.items[:0]
}
