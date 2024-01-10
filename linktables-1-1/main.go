package main

import (
	. "github.com/isdamir/gotype"
)

// 如何实现链表的逆序
// 描述： 给定一个带头结点的单链表，请将其逆序。
// 即如果单链表原来为 head->1->2->3->4->5->6->7,
// 则逆序后变为 head->7->6->5->4->3->2->1
func main() {
	head := &LNode{}
	createNode(head, 8)
	PrintNode("打印单链表：\n", head)

	//reverseNode1(head)
	reverseNode2(head)
	//node := reverseNode3(head)
	// 这个打印是必须带有头结点的，不然少了个第一个数据
	PrintNode("打印逆序后的链表：\n", head)

}

func reverseNode1(node *LNode) {
	if node == nil || node.Next == nil {
		return
	}

	// 定义前驱节点
	var pre *LNode
	// 定义当前节点
	var cur *LNode
	// 后继节点
	next := node.Next
	for next != nil {
		// 当前节点其实是下一个节点，从名字理解不是很直观
		cur = next.Next
		next.Next = pre
		pre = next
		// 进入下一循环
		next = cur
	}
	node.Next = pre
}

func reverseNode2(node *LNode) {
	if node == nil || node.Next == nil {
		return
	}

	// 定义前驱节点
	var pre *LNode
	// 定义当前节点
	var cur *LNode
	var next *LNode

	// 跟第一种方法是一样的，但是将当前节点是下一个节点，而不是cur=next.Next,对于我自己来讲这样比较好理解
	cur = node.Next
	for cur != nil {
		next = cur.Next
		cur.Next = pre

		// 指针移动，当前节点变为前序节点，后续节点变为当前节点，为下次循环迭代节点
		pre = cur
		cur = next
	}
	node.Next = pre
}

// 这种也行，就是改变了head的指针，后面需要带返回值，而且头结点是有数据的
func reverseNode3(head *LNode) *LNode {
	//  先声明两个变量
	//  前一个节点
	var preNode *LNode
	//preNode = nil
	//  后一个节点
	var nextNode *LNode
	// nextNode = nil
	for head != nil {
		//  保存头节点的下一个节点，
		nextNode = head.Next
		//  将头节点指向前一个节点
		head.Next = preNode
		//  更新前一个节点
		preNode = head
		//  更新头节点
		head = nextNode
	}
	return preNode
}

func createNode(head *LNode, num int) {
	cur := head
	for i := 1; i < num; i++ {
		cur.Next = &LNode{}
		cur.Next.Data = i
		cur = cur.Next
	}
}
