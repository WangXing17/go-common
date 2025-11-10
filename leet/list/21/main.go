package leet_list

// 21 合并两个有序链表
// https://leetcode.cn/problems/merge-two-sorted-lists/description/?envType=study-plan-v2&envId=top-100-liked

type ListNode struct {
	Val  int
	Next *ListNode
}

//func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
//	dummyNode := &ListNode{
//		Val:  -1,
//		Next: nil,
//	}
//	cur := dummyNode
//	for list1 != nil && list2 != nil {
//		if list1.Val < list2.Val {
//			cur.Next = list1
//			list1 = list1.Next
//		} else {
//			cur.Next = list2
//			list2 = list2.Next
//		}
//		cur = cur.Next
//	}
//	if list1 != nil {
//		cur.Next = list1
//	}
//	if list2 != nil {
//		cur.Next = list2
//	}
//	return dummyNode.Next
//}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list2.Next, list1)
		return list2
	}
}
