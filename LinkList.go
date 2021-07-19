package sensitive

// 定义链表结构
type LinkList struct {
	head  *listNode
	tail  *listNode
	count int64
}

// 定义链表结点
type listNode struct {
	Value interface{}
	Next  *listNode
}

// 定义入队函数
func (list *LinkList) Push(v interface{}) {
	node := &listNode{
		Value: v,
	}

	if list.head == nil {
		list.head = node
	} else {
		list.tail.Next = node
	}

	list.tail = node
	list.count++
}

// 定义出队函数
func (list *LinkList) Pop() interface{} {
	if list.Empty() {
		return nil
	}

	n := list.head
	list.head = n.Next
	list.count--

	return n.Value
}

// 判断是否队列是否为空
func (list *LinkList) Empty() bool {
	return list.count == 0
}
