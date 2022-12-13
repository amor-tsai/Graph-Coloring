package assignments

import (
	"fmt"
)

type Node struct {
	data  int   //data
	index int   //first index=1
	next  *Node //next Node
}

type NodeList struct {
	head   *Node //the first node of NodeList
	tail   *Node //the last node of NodeList
	length int   // the length of NodeList
}

/**
whether list is empty ,O(1)
@return bool
*/
func (list *NodeList) IsEmpty() bool {
	return (list.length == 0)
}

/**
@description add data into the end of list,cost O(1)
*/
func (list *NodeList) Append(data int) {
	node := &Node{data: data}
	if list.IsEmpty() {
		list.head = node
		list.head.index = 1
		list.length = 1
		list.tail = list.head
	} else {
		tail := list.tail //get the last node of list
		tail.next = node
		node.index = tail.index + 1
		list.tail = node
		list.length++
	}
}

/**
@description get the length of list
@return bool
*/
func (list *NodeList) GetLength() int {
	return list.length
}

/**
@description find the index of dumplicate number,if exists, return the index of it, or else return 0, cost O(n)
*/
func (list *NodeList) FindIndexOfDuplicateNumber(number int) int {
	if list.length == 0 {
		return 0
	}
	current := list.head
	for {
		if current.data == number {
			return current.index
		}
		if current != list.tail {
			current = current.next
		} else {
			break
		}
	}
	return 0
}

/**
@description display all elements in the list
*/
func (list *NodeList) DisplayAllElements() {
	current := list.head
	fmt.Println("start to display elements of list:")
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
	fmt.Println("it's over")
}

/**
@description delete the index element from list
@return nil
*/
func (list *NodeList) Delete(index int) { //
	panic("this function is not ready to implement")
}
