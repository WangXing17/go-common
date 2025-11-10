package leet_list

import "fmt"

// 138 随机链表的复制
// https://leetcode.cn/problems/copy-list-with-random-pointer/description/?envType=study-plan-v2&envId=top-100-liked

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	//在每个结点后面复制一个新的结点
	cur := head
	for cur != nil {
		newNode := &Node{
			Val:  cur.Val,
			Next: cur.Next,
		}
		cur.Next = newNode
		cur = cur.Next.Next
	}
	//为新的结点设置random指针
	cur = head
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}
	//连接所有新的结点
	dummyNode := &Node{
		Val:  -1,
		Next: head,
	}
	curDummy := dummyNode
	cur = head
	for cur != nil {
		curDummy.Next = cur.Next
		curDummy = curDummy.Next
		cur.Next = curDummy.Next
		cur = cur.Next
	}
	return dummyNode.Next
}

func main() {
	node3 := &Node{
		Val:    3,
		Next:   nil,
		Random: nil,
	}
	node2 := &Node{
		Val:    2,
		Next:   node3,
		Random: nil,
	}
	node1 := &Node{
		Val:    1,
		Next:   node2,
		Random: node3,
	}
	list := copyRandomList(node1)
	fmt.Println(list)
}
