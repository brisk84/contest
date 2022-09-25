package main

import (
	"bufio"
	"fmt"
	"os"
)

type Btree struct {
	root   *Node
	values []Val
	len    int
}

type Val interface {
	Comp(val Val) int8
}

type Node struct {
	Value       Val
	left, right *Node
	height      int8
}

func NewTree() *Btree { return new(Btree).Init() }

func (t *Btree) Init() *Btree {
	t.root = nil
	t.values = nil
	t.len = 0
	return t
}

func (t *Btree) Empty() bool {
	return t.root == nil
}

func (t *Btree) NotEmpty() bool {
	return t.root != nil
}

func (t *Btree) balance() int8 {
	if t.root != nil {
		return balance(t.root)
	}
	return 0
}

func (t *Btree) Insert(value Val) *Btree {
	added := false
	t.root = insert(t.root, value, &added)
	if added {
		t.len++
	}
	t.values = nil
	return t
}

func insert(n *Node, value Val, added *bool) *Node {
	if n == nil {
		*added = true
		return (&Node{Value: value}).Init()
	}
	c := value.Comp(n.Value)
	if c > 0 {
		n.right = insert(n.right, value, added)
	} else if c < 0 {
		n.left = insert(n.left, value, added)
	} else {
		oldVal := n.Value.(ProcVal).proc
		newVal := value.(ProcVal).proc
		oldVal = append(oldVal, newVal...)
		k := n.Value.(ProcVal).key
		n.Value = ProcVal{k, oldVal}
		*added = false
		return n
	}

	n.height = n.maxHeight() + 1
	c = balance(n)

	if c > 1 {
		c = value.Comp(n.left.Value)
		if c < 0 {
			return n.rotateRight()
		} else if c > 0 {
			n.left = n.left.rotateLeft()
			return n.rotateRight()
		}
	} else if c < -1 {
		c = value.Comp(n.right.Value)
		if c > 0 {
			return n.rotateLeft()
		} else if c < 0 {
			n.right = n.right.rotateRight()
			return n.rotateLeft()
		}
	}
	return n
}

func (t *Btree) Len() int {
	return t.len
}

func (t *Btree) Delete(value Val) *Btree {
	deleted := false
	t.root = deleteNode(t.root, value, &deleted)
	if deleted {
		t.len--
	}
	t.values = nil
	return t
}

func deleteNode(n *Node, value Val, deleted *bool) *Node {
	if n == nil {
		return n
	}

	c := value.Comp(n.Value)

	if c < 0 {
		n.left = deleteNode(n.left, value, deleted)
	} else if c > 0 {
		n.right = deleteNode(n.right, value, deleted)
	} else {
		l := len(n.Value.(ProcVal).proc)
		if l > 1 {
			k := n.Value.(ProcVal).key
			proc := n.Value.(ProcVal).proc[:l-1]
			n.Value = ProcVal{k, proc}
			return n
		}
		if n.left == nil {
			t := n.right
			n.Init()
			return t
		} else if n.right == nil {
			t := n.left
			n.Init()
			return t
		}
		t := n.right.min()
		n.Value = t.Value
		n.right = deleteNode(n.right, t.Value, deleted)
		*deleted = true
	}

	//re-balance
	if n == nil {
		return n
	}
	n.height = n.maxHeight() + 1
	bal := balance(n)
	if bal > 1 {
		if balance(n.left) >= 0 {
			return n.rotateRight()
		}
		n.left = n.left.rotateLeft()
		return n.rotateRight()
	} else if bal < -1 {
		if balance(n.right) <= 0 {
			return n.rotateLeft()
		}
		n.right = n.right.rotateRight()
		return n.rotateLeft()
	}

	return n
}

func (n *Node) Init() *Node {
	n.height = 1
	n.left = nil
	n.right = nil
	return n
}

func height(n *Node) int8 {
	if n != nil {
		return n.height
	}
	return 0
}

func balance(n *Node) int8 {
	if n == nil {
		return 0
	}
	return height(n.left) - height(n.right)
}

func (n *Node) get(val Val) *Node {
	var node *Node
	c := val.Comp(n.Value)
	if c < 0 {
		if n.left != nil {
			node = n.left.get(val)
		}
	} else if c > 0 {
		if n.right != nil {
			node = n.right.get(val)
		}
	} else {
		node = n
	}
	return node
}

func (n *Node) rotateRight() *Node {
	l := n.left
	// Rotation
	l.right, n.left = n, l.right

	// update heights
	n.height = n.maxHeight() + 1
	l.height = l.maxHeight() + 1

	return l
}

func (n *Node) rotateLeft() *Node {
	r := n.right
	// Rotation
	r.left, n.right = n, r.left

	// update heights
	n.height = n.maxHeight() + 1
	r.height = r.maxHeight() + 1

	return r
}

func (n *Node) min() *Node {
	current := n
	for current.left != nil {
		current = current.left
	}
	return current
}

func (n *Node) maxHeight() int8 {
	rh := height(n.right)
	lh := height(n.left)
	if rh > lh {
		return rh
	}
	return lh
}

type ProcVal struct {
	key  uint64
	proc []Proc
}

func (p ProcVal) Comp(val Val) int8 {
	v := val.(ProcVal)
	if p.key > v.key {
		return 1
	} else if p.key < v.key {
		return -1
	} else {
		return 0
	}
}

type Proc struct {
	Power     uint64
	BusyUntil uint64
}

type Task struct {
	StartTime uint64
	WorkTime  uint64
}

func main() {

	// f, err := os.Open("/Users/brisk/vscode/contest/lab08/tests/07")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer f.Close()
	// in := bufio.NewReader(f)

	in := bufio.NewReader(os.Stdin)

	n := 0 // procs
	m := 0 // tasks
	fmt.Fscan(in, &n, &m)

	freeProcs := NewTree()
	busyProcs := NewTree()

	for i := 0; i < n; i++ {
		proc := Proc{}
		fmt.Fscan(in, &proc.Power)
		freeProcs.Insert(ProcVal{proc.Power, []Proc{proc}})
	}

	tasks := make([]Task, m)
	for i := 0; i < m; i++ {
		var startTime uint64
		var workTime uint64
		fmt.Fscan(in, &startTime, &workTime)
		newTask := Task{
			StartTime: startTime,
			WorkTime:  workTime,
		}
		tasks[i] = newTask
	}

	var totalPower uint64
	var curTime uint64

	for _, curTask := range tasks {
		curTime = curTask.StartTime

		for busyProcs.root != nil {
			proc := busyProcs.root.min().Value.(ProcVal).proc
			br := false
			for _, p := range proc {
				if p.BusyUntil <= curTime {
					busyProcs.Delete(ProcVal{p.BusyUntil, nil})
					p.BusyUntil = 0
					freeProcs.Insert(ProcVal{p.Power, []Proc{p}})
				} else {
					br = true
					break
				}
			}
			if br {
				break
			}
		}
		if freeProcs.root == nil {
			continue
		}

		proc := freeProcs.root.min().Value.(ProcVal).proc[0]
		busyUntil := curTime + curTask.WorkTime
		proc.BusyUntil = busyUntil
		busyProcs.Insert(ProcVal{proc.BusyUntil, []Proc{proc}})
		freeProcs.Delete(ProcVal{proc.Power, nil})
		totalPower += curTask.WorkTime * proc.Power
	}
	fmt.Println(totalPower)
}
