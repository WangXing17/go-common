package leet_list

// 19 删除链表的倒数第N个结点
// https://leetcode.cn/problems/remove-nth-node-from-end-of-list/description/?envType=study-plan-v2&envId=top-100-liked

type ListNode struct {
	Val  int
	Next *ListNode
}

// 定义快慢指针，快指针先走n步，然后快慢指针每次一起走1步，当快指针为nil时，慢指针正好指向倒数第n个结点
// 因为要删除倒数第n个结点，所以要找到它的前驱，所以慢指针从dummyNode开始
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyNode := &ListNode{
		Val:  -1,
		Next: head,
	}
	slow, fast := dummyNode, head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummyNode.Next
}
