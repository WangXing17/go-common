package leet_list

import "math"

// 148 排序链表
// https://leetcode.cn/problems/sort-list/?envType=study-plan-v2&envId=top-100-liked

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummyNode := &ListNode{
		Val:  math.MinInt,
		Next: head,
	}
	cur := head
	for cur != nil {
		dCur := dummyNode
		for dCur != nil {
			if dCur.Val > cur.Val {

			}
			dCur = dCur.Next
		}
		cur = cur.Next
	}

	return dummyNode.Next
}

func main() {

}
