package leet_list

import "fmt"

// 160 相交链表
// https://leetcode.cn/problems/intersection-of-two-linked-lists/?envType=study-plan-v2&envId=top-100-liked

type ListNode struct {
	Val  int
	Next *ListNode
}

// 定义两个指针，pA->headA, pB->headB，pA走完A后从B走，pB走完B后从A走
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pA, pB := headA, headB
	for pA != pB {
		if pA == nil {
			pA = headB
		} else {
			pA = pA.Next
		}
		if pB == nil {
			pB = headA
		} else {
			pB = pB.Next
		}
	}
	return pA
}

func main() {
	headA := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	node := getIntersectionNode(headA, nil)
	fmt.Println(*node)
}
