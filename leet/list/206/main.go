package leet_list

// 206 反转链表
// https://leetcode.cn/problems/reverse-linked-list/?envType=study-plan-v2&envId=top-100-liked

type ListNode struct {
	Val  int
	Next *ListNode
}

////定义3个指针，交替操作
//func reverseList(head *ListNode) *ListNode {
//	if head == nil {
//		return head
//	}
//	var pre *ListNode = nil
//	var cur *ListNode = head
//	for cur != nil {
//		var next = cur.Next
//		cur.Next = pre
//		pre = cur
//		cur = next
//	}
//	return pre
//}

// 使用递归
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var newHead *ListNode = reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
