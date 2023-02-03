/*
 * @Notice: edit notice here
 * @Author: zhulei
 * @Date: 2022-12-12 17:54:32
 * @LastEditors: zhulei
 * @LastEditTime: 2023-02-03 13:33:14
 */
package main

import (
	"fmt"
	"sync"
	"time"
)

type elemNode struct {
	IdleTimeout time.Duration
	// Conn *net.conn
	Count int
	Prev  *elemNode
	Next  *elemNode
}

type linkList struct {
	Head  *elemNode
	Tail  *elemNode
	Len   int
	Mutex *sync.Mutex
}

func (pl *linkList) Append(timeout time.Duration, count int) {
	node := &elemNode{
		IdleTimeout: timeout,
		Count:       count,
	}
	pl.Mutex.Lock()
	defer pl.Mutex.Unlock()
	if pl.Head == nil {
		pl.Head = node
		pl.Tail = node

	} else {
		// 头节点修改，如果是循环链表
		currentHeadNode := pl.Head
		currentHeadNode.Prev = node
		// 尾节点修改
		currentTailNode := pl.Tail
		currentTailNode.Next = node

		// 新添加节点修改
		node.Prev = currentTailNode
		node.Next = currentHeadNode
		// 链表修改tail指针
		pl.Tail = node
	}
	pl.Len++
}

// 遍历链表所有节点
func (pl *linkList) foreach() {

	p := pl.Head
	for {
		if p == pl.Tail {
			fmt.Println("last: ", p.Count)
			break
		}
		fmt.Println(p.Count)
		p = p.Next
	}
}

// 从尾部出链表队列
func (pl *linkList) Pop() *elemNode {
	pl.Mutex.Lock()
	defer pl.Mutex.Unlock()
	if pl.Len <= 0 {
		return nil
	}
	currentTailNode := pl.Tail

	// 头节点修改，如果是循环链表
	currentHeadNode := pl.Head
	currentHeadNode.Prev = currentTailNode.Prev
	// 新的尾节点修改
	newCurrentTailNode := currentTailNode.Prev
	newCurrentTailNode.Next = currentHeadNode

	// 链表长度减1
	pl.Len--
	return currentTailNode
}

func main() {
	// JSRUN引擎2.0，支持多达30种语言在线运行，全仿真在线交互输入输出。
	// fmt.Println("Hello world!   -  go.jsrun.net ")
	IdleTimeout := 0 * time.Second
	if IdleTimeout <= 0 {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
	count := 0

	pl := &linkList{}

	for {
		if count > 100 {
			break
		}
		pl.Append(600*time.Second, count)
		count++
	}
	fmt.Println(pl.Len)

	pl.foreach()

}
