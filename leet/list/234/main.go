package leet_list

// 234回文链表
// https://leetcode.cn/problems/palindrome-linked-list/description/?envType=study-plan-v2&envId=top-100-liked

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	//1.利用快慢指针找到中点
	var fast *ListNode = head
	var slow *ListNode = head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	//2.反转链表后半部分
	var lastPart *ListNode = reverseList(slow)
	//3.前后进行比对
	for lastPart != nil {
		if lastPart.Val != head.Val {
			return false
		}
		lastPart = lastPart.Next
		head = head.Next
	}
	return true
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode = nil
	var cur *ListNode = head
	for cur != nil {
		var next *ListNode = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func main() {

}
