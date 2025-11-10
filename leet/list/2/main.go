package leet_list

// 2 两数相加
// https://leetcode.cn/problems/add-two-numbers/description/?envType=study-plan-v2&envId=top-100-liked

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyNode := &ListNode{
		Val:  -1,
		Next: nil,
	}
	n1, n2 := l1, l2
	carry := 0
	cur := dummyNode
	for n1 != nil || n2 != nil {
		v1 := 0
		if n1 != nil {
			v1 = n1.Val
			n1 = n1.Next
		}
		v2 := 0
		if n2 != nil {
			v2 = n2.Val
			n2 = n2.Next
		}
		sum := v1 + v2 + carry
		carry = sum / 10
		cur.Next = &ListNode{
			Val:  sum % 10,
			Next: nil,
		}
		cur = cur.Next
	}

	if carry > 0 {
		cur.Next = &ListNode{
			Val:  carry,
			Next: nil,
		}
	}
	return dummyNode.Next
}
