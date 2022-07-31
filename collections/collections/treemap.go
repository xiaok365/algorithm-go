package collections

import "fmt"

const (
	COLOR_RED   = false
	COLOR_BLACK = true
)

type RBNode struct {
	Key                interface{}
	Value              interface{}
	Color              bool
	Lson, Rson, Parent *RBNode
}

type TreeMap struct {
	Root    *RBNode
	Size    int
	Compare func(interface{}, interface{}) int
}

func NewTreeMap(cmp func(interface{}, interface{}) int) *TreeMap {
	return &TreeMap{
		Compare: cmp,
	}
}

func LeftOf(p *RBNode) *RBNode {
	return p.Lson
}
func RightOf(p *RBNode) *RBNode {
	return p.Rson
}
func ParentOf(p *RBNode) *RBNode {
	return p.Parent
}

func IsRed(p *RBNode) bool {
	return p != nil && p.Color == COLOR_RED
}

func IsBlack(p *RBNode) bool {
	return !IsRed(p)
}

func (t *TreeMap) IsEmpty() bool {
	return t.Root == nil
}

func (t *TreeMap) Find(key interface{}) *RBNode {
	p := t.Root
	for p != nil {
		if p.Key == key {
			break
		}
		if t.Compare(p.Key, key) < 0 {
			p = p.Rson
		} else {
			p = p.Lson
		}
	}
	return p
}

func (t *TreeMap) FindNextMin(key interface{}) *RBNode {

	tail := t.Root
	var last *RBNode

	for tail != nil {
		last = tail
		if t.Compare(key, tail.Key) <= 0 {
			tail = tail.Lson
		} else {
			tail = tail.Rson
		}
	}

	for last != nil && t.Compare(last.Key, key) >= 0 {
		last = last.Parent
	}

	if last != nil {
		return last
	}

	return t.FindMax(t.Root)
}

func (t *TreeMap) FindNextMax(key interface{}) *RBNode {

	tail := t.Root
	var last *RBNode

	for tail != nil {
		last = tail
		if t.Compare(key, tail.Key) >= 0 {
			tail = tail.Rson
		} else {
			tail = tail.Lson
		}
	}

	for last != nil && t.Compare(last.Key, key) <= 0 {
		last = last.Parent
	}

	if last != nil {
		return last
	}

	return t.FindMin(t.Root)
}

func (t *TreeMap) FindMin(p *RBNode) *RBNode {
	if p == nil {
		return nil
	}
	for p.Lson != nil {
		p = p.Lson
	}
	return p
}

func (t *TreeMap) FindMax(p *RBNode) *RBNode {
	if p == nil {
		return nil
	}
	for p.Rson != nil {
		p = p.Rson
	}
	return p
}

func (t *TreeMap) LeftRotate(x *RBNode) *RBNode {
	y := x.Rson

	x.Rson = y.Lson
	if y.Lson != nil {
		y.Lson.Parent = x
	}

	y.Parent = x.Parent

	if x.Parent == nil {
		t.Root = y
	} else if x == LeftOf(ParentOf(x)) {
		x.Parent.Lson = y
	} else {
		x.Parent.Rson = y
	}

	y.Lson = x
	x.Parent = y

	return y
}

func (t *TreeMap) RightRotate(y *RBNode) *RBNode {
	x := y.Lson

	y.Lson = x.Rson
	if x.Rson != nil {
		x.Rson.Parent = y
	}

	x.Parent = y.Parent

	if y.Parent == nil {
		t.Root = x
	} else if y == LeftOf(ParentOf(y)) {
		y.Parent.Lson = x
	} else {
		y.Parent.Rson = x
	}

	x.Rson = y
	y.Parent = x

	return x
}

func (t *TreeMap) Insert(key, value interface{}) {

	p := &RBNode{Key: key, Value: value, Color: COLOR_RED}

	exist := t.Find(key)
	if exist != nil {
		return
	}

	t.Size++

	tail := t.Root
	var last *RBNode

	for tail != nil {
		last = tail
		if t.Compare(p.Key, tail.Key) < 0 {
			tail = tail.Lson
		} else {
			tail = tail.Rson
		}
	}
	p.Parent = last
	if last == nil {
		t.Root = p
	} else if t.Compare(p.Key, last.Key) < 0 {
		last.Lson = p
	} else {
		last.Rson = p
	}

	t.InsertFixup(p)
}

func (t *TreeMap) InsertFixup(p *RBNode) {

	var parent, grandParent, uncle *RBNode

	parent = ParentOf(p)

	for parent != nil && IsRed(parent) {
		grandParent = ParentOf(parent)
		if parent == LeftOf(grandParent) {
			uncle = RightOf(grandParent)
			if IsRed(uncle) {
				parent.Color = COLOR_BLACK
				uncle.Color = COLOR_BLACK
				grandParent.Color = COLOR_RED
				p = grandParent
			} else {
				if p == RightOf(parent) {
					t.LeftRotate(parent)
					temp := parent
					parent = p
					p = temp
				}
				parent.Color = COLOR_BLACK
				grandParent.Color = COLOR_RED
				t.RightRotate(grandParent)
			}
		} else {
			uncle = LeftOf(grandParent)
			if IsRed(uncle) {
				parent.Color = COLOR_BLACK
				uncle.Color = COLOR_BLACK
				grandParent.Color = COLOR_RED
				p = grandParent
			} else {
				if p == LeftOf(parent) {
					t.RightRotate(parent)
					temp := parent
					parent = p
					p = temp
				}
				parent.Color = COLOR_BLACK
				grandParent.Color = COLOR_RED
				t.LeftRotate(grandParent)
			}
		}
		parent = ParentOf(p)
	}
	t.Root.Color = COLOR_BLACK
}

func (t *TreeMap) Transplant(u, v *RBNode) {
	if u.Parent == nil {
		t.Root = v
	} else if u == LeftOf(ParentOf(u)) {
		u.Parent.Lson = v
	} else {
		u.Parent.Rson = v
	}

	if v != nil {
		v.Parent = u.Parent
	}
}

func (t *TreeMap) Remove(key interface{}) {
	if t.IsEmpty() {
		return
	}

	p := t.Find(key)
	if p == nil {
		fmt.Println("remove failed, key=", key)
		return
	}

	t.Size--

	if p.Lson != nil && p.Rson != nil {
		replacement := t.FindMin(p.Rson)
		p.Value = replacement.Value
		p.Key = replacement.Key
		p = replacement
	}

	var successor *RBNode

	if p.Lson != nil {
		successor = p.Lson
	} else {
		successor = p.Rson
	}

	if successor != nil {
		t.Transplant(p, successor)
		if p.Color == COLOR_BLACK {
			t.RemoveFixup(successor)
		}
	} else {
		if p.Color == COLOR_BLACK {
			t.RemoveFixup(p)
		}
		t.Transplant(p, nil)
	}
}

func (t *TreeMap) RemoveFixup(p *RBNode) {
	var brother, parent *RBNode

	for p != t.Root && IsBlack(p) {

		parent = ParentOf(p)

		if p == LeftOf(parent) {
			brother = parent.Rson
			if IsRed(brother) {
				brother.Color = COLOR_BLACK
				parent.Color = COLOR_RED
				t.LeftRotate(parent)
				brother = RightOf(parent)
			}
			if IsBlack(brother.Lson) && IsBlack(brother.Rson) {
				brother.Color = COLOR_RED
				p = parent
			} else {
				if IsBlack(brother.Rson) {
					brother.Color = COLOR_RED
					brother.Lson.Color = COLOR_BLACK
					t.RightRotate(brother)
					brother = RightOf(parent)
				}
				brother.Rson.Color = COLOR_BLACK
				brother.Color = parent.Color
				parent.Color = COLOR_BLACK
				t.LeftRotate(parent)
				break
			}
		} else {
			brother = parent.Lson
			if IsRed(brother) {
				brother.Color = COLOR_BLACK
				parent.Color = COLOR_RED
				t.RightRotate(parent)
				brother = LeftOf(parent)
			}
			if IsBlack(brother.Lson) && IsBlack(brother.Rson) {
				brother.Color = COLOR_RED
				p = parent
			} else {
				if IsBlack(brother.Lson) {
					brother.Color = COLOR_RED
					brother.Rson.Color = COLOR_BLACK
					t.LeftRotate(brother)
					brother = LeftOf(parent)
				}
				brother.Lson.Color = COLOR_BLACK
				brother.Color = parent.Color
				parent.Color = COLOR_BLACK
				t.RightRotate(parent)
				break
			}
		}
	}

	p.Color = COLOR_BLACK
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func (t *TreeMap) GetHeight(p *RBNode) int {
	if p == nil {
		return 0
	}
	return max(t.GetHeight(p.Lson)+1, t.GetHeight(p.Rson)+1)
}

func (t *TreeMap) Height() int {
	return t.GetHeight(t.Root)
}

func (t *TreeMap) Print(p *RBNode) {
	if p == nil {
		return
	}

	fmt.Println(fmt.Sprintf("key=%d, color=%t", p.Key, p.Color))
	t.Print(p.Lson)
	t.Print(p.Rson)
}

func (t *TreeMap) PrintTree() {
	t.Print(t.Root)
}
