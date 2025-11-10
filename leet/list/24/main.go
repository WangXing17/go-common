package leet_list

// 24 两两交换链表中的结点
// https://leetcode.cn/problems/swap-nodes-in-pairs/description/?envType=study-plan-v2&envId=top-100-liked

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	dummyNode := &ListNode{
		Val:  -1,
		Next: head,
	}
	pre := dummyNode
	for pre.Next != nil && pre.Next.Next != nil {
		n1 := pre.Next
		n2 := pre.Next.Next

		n1.Next = n2.Next
		pre.Next = n2
		n2.Next = n1

		pre = n1
	}
	return dummyNode.Next
}
