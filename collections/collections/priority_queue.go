package collections

import (
	"fmt"
	"time"
)

const (
	MAX_HEAP = true
	MIN_HEAP = false
)

type Node struct {
	value interface{}
	t     int64
}

func NewNode(value interface{}) Node {

	return Node{
		value: value,
		t:     time.Now().UnixNano(),
	}
}

type PriorityQueue struct {
	data     []Node
	capacity int
	size     int
	heapType bool
	cmp      func(interface{}, interface{}) int
}

func NewPriorityQueue(capacity int, t bool, cmp func(interface{}, interface{}) int) *PriorityQueue {
	return &PriorityQueue{
		data:     make([]Node, capacity),
		capacity: capacity,
		size:     0,
		heapType: t,
		cmp:      cmp,
	}
}

func (p *PriorityQueue) IsEmpty() bool {
	return p.size <= 0
}

func (p *PriorityQueue) Push(t interface{}) {
	node := NewNode(t)
	fmt.Println(node.t)
	p.data[p.size] = node
	p.size++
	p.HeapUp(p.size - 1)
}

func (p *PriorityQueue) Pop() interface{} {

	if p.IsEmpty() {
		return nil
	}
	ret := p.data[0]
	p.size--
	p.data[0] = p.data[p.size]
	p.HeapDown(0, p.size-1)

	return ret
}

func (p *PriorityQueue) IsMax() bool {
	return p.heapType == MAX_HEAP
}

func (p *PriorityQueue) IsMin() bool {
	return p.heapType == MIN_HEAP
}

func (p *PriorityQueue) HeapUp(k int) {
	var parent, son int
	x := p.data[k]
	son = k
	parent = (son - 1) / 2

	for son > 0 {
		if p.IsMax() && p.Compare(p.data[parent], x) >= 0 || p.IsMin() && p.Compare(p.data[parent], x) <= 0 {
			break
		}

		p.data[son] = p.data[parent]
		son = parent
		parent = (son - 1) / 2
	}
	p.data[son] = x
}

func (p *PriorityQueue) HeapDown(k, n int) {
	var parent, son int
	x := p.data[k]
	parent = k
	son = 2*k + 1

	for son <= n {

		if son+1 <= n &&
			(p.IsMax() && p.Compare(p.data[son+1], p.data[son]) > 0 ||
				p.IsMin() && p.Compare(p.data[son+1], p.data[son]) < 0) {
			son++
		}

		if p.IsMax() && p.Compare(p.data[son], x) <= 0 || p.IsMin() && p.Compare(p.data[son], x) >= 0 {
			break
		}
		p.data[parent] = p.data[son]
		parent = son
		son = 2*parent + 1
	}
	p.data[parent] = x
}

func (p *PriorityQueue) Compare(a, b Node) int {
	if p.cmp(a.value, b.value) == 0 {

		if p.heapType == MAX_HEAP {
			return int(b.t - a.t)
		} else {
			return int(a.t - b.t)
		}
	}

	return p.cmp(a.value, b.value)
}
