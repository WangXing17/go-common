package leet_list

// 25 k个一组反转链表
// https://leetcode.cn/problems/reverse-nodes-in-k-group/description/?envType=study-plan-v2&envId=top-100-liked

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummyNode := &ListNode{
		Val:  -1,
		Next: head,
	}
	dCur := dummyNode
	cur := head
	for cur != nil {
		//找到groupHead
		i := 1
		groupHead := cur
		for i < k && cur != nil {
			i++
			cur = cur.Next
		}
		//找到groupEnd
		groupEnd := cur
		//找到末尾还不够k个
		if groupEnd == nil {
			dCur.Next = groupHead
		} else {
			//断开groupEnd后的链
			next := groupEnd.Next
			groupEnd.Next = nil
			cur = next

			//反转group内的链表
			dCur.Next = reverseList(groupHead)
			dCur = groupHead
		}
	}
	return dummyNode.Next
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre *ListNode = nil
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
