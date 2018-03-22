package Util

type Node struct {
	Value      interface{}
	next, prev *Node
}
type List struct {
	root   Node // 头节点
	length int  // list长度
}
// 返回List的指针
func New() *List {
	l := &List{}// 获取List{}的地址
	l.length = 0// list初始长度为0
	l.root.next = &l.root
	l.root.prev = &l.root
	return l
}
func (l *List) IsEmpty() bool {
	return l.root.next == &l.root
}
func (l *List) Length() int {
	return l.length
}

func (l *List) PushFront(elements ...interface{}) {
	for _, element := range elements {
		n := &Node{Value: element} // 注释一
		n.next = l.root.next // 新节点的next是root节点的next
		n.prev = &l.root // 新节点的prev存储的是root的地址
		l.root.next.prev = n // 原来root节点的next的prev是新节点
		l.root.next = n // 头插法 root 之后始终是新节点
		l.length++ // list 长度加1
	}
}
func (l *List) PushBack(elements ...interface{}) {
	for _, element := range elements {
		n := &Node{Value: element}
		n.next = &l.root // since n is the last element, its next should be the head
		n.prev = l.root.prev // n's prev should be the tail
		l.root.prev.next = n // tail's next should be n
		l.root.prev = n // head's prev should be n
		l.length++
	}
}
func (l *List) Find(element interface{}) int {
	index := 0
	p := l.root.next
	for p != &l.root && p.Value != element {
		p = p.next
		index++
	}
	// p不是root
	if p != &l.root {
		return index
	}

	return -1
}
func (l *List) remove(n *Node) {
	n.prev.next = n.next
	n.next.prev = n.prev
	n.next = nil
	n.prev = nil
	l.length--
}
func (l *List) Lpop() interface{} {
	if l.length == 0 {
		// null的表现形式nil
		return nil
	}
	n := l.root.next
	l.remove(n)
	return n.Value
}
func (l *List) normalIndex(index int) int {
	if index > l.length - 1 {
		index = l.length - 1
	}

	if index < -l.length {
		index = 0
	}
	// 将给定的index与length做取余处理
	index = (l.length + index) % l.length
	return index
}
func (l *List) Range(start, end int) []interface{} {
	// 获取正常的start和end
	start = l.normalIndex(start)
	end = l.normalIndex(end)
	// 声明一个interface类型的数组
	res := []interface{}{}
	// 如果上下界不符合逻辑，返回空res
	if start > end {
		return res
	}

	sNode := l.index(start)
	eNode := l.index(end)
	// 起始点和重点遍历
	for n := sNode; n != eNode; {
		// res的append方式
		res = append(res, n.Value)
		n = n.next
	}
	res = append(res, eNode.Value)
	return res
}