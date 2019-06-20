package tqueue

import "fmt"

//链表
type LinkedList struct {
	data     int
	nextList *LinkedList
}

var result []*LinkedList

func LinkedListFunc() {
	l32 := LinkedList{data: 32, nextList: nil}
	l26 := LinkedList{data: 26, nextList: &l32}
	l18 := LinkedList{data: 18, nextList: &l26}
	l10 := LinkedList{data: 10, nextList: &l18}
	l9 := LinkedList{data: 9, nextList: &l10}
	l8 := LinkedList{data: 8, nextList: &l9}
	l5 := LinkedList{data: 5, nextList: &l8}
	l3 := LinkedList{data: 3, nextList: &l5}
	l2 := LinkedList{data: 2, nextList: &l3}

	result = append(result, &l32)
	result = append(result, &l26)
	result = append(result, &l18)
	result = append(result, &l10)
	result = append(result, &l9)
	result = append(result, &l8)
	result = append(result, &l5)
	result = append(result, &l3)
	result = append(result, &l2)

	l6 := LinkedList{data: 6, nextList: nil}

	l2.AddLinkedList(&l6)
	l2.PrintList()
}

func (l *LinkedList) AddLinkedList(new *LinkedList) {
	if l.data <= new.data {
		(l.nextList).AddLinkedList(new)
	} else {
		for i := range result {
			if result[i].nextList == l {
				result[i].nextList = new
			}
			new.nextList = l
		}
	}
}

func (l LinkedList) PrintList() {
	fmt.Println(l.data)
	if l.nextList != nil {
		(l.nextList).PrintList()
	}
}
