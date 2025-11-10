package leet_list

// 141_142 环形链表
// https://leetcode.cn/problems/linked-list-cycle/?envType=study-plan-v2&envId=top-100-liked

// 142 环形链表2
// https://leetcode.cn/problems/linked-list-cycle-ii/description/?envType=study-plan-v2&envId=top-100-liked

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

////使用双指针
//func detectCycle(head *ListNode) *ListNode {
//	if head == nil || head.Next == nil {
//		return nil
//	}
//	slow, fast := head, head
//	flag := false
//	for fast != nil && fast.Next != nil {
//		slow = slow.Next
//		fast = fast.Next.Next
//		if slow == fast {
//			flag = true
//			break
//		}
//	}
//	if flag {
//		fast = head
//		for slow != fast {
//			fast = fast.Next
//			slow = slow.Next
//		}
//		return slow
//	}
//	return nil
//}

// 使用map来存储
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	cur := head
	seen := make(map[*ListNode]any)
	for cur != nil {
		if _, ok := seen[cur]; ok {
			return cur
		}
		seen[cur] = nil
		cur = cur.Next
	}
	return nil
}

func main() {

}
