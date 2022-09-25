package main

import (
	"bufio"
	"fmt"
	"os"
)

type Proc struct {
	Power     uint64
	BusyUntil uint64
}

type Task struct {
	StartTime uint64
	WorkTime  uint64
}

type Node struct {
	Value uint64
	Data  Proc
	Left  *Node
	Right *Node
}

func (n *Node) Insert(value uint64, data Proc) error {
	if value == n.Value {
		value++
		data.BusyUntil = value
	}
	switch {
	case value == n.Value:
		return nil
	case value < n.Value:
		if n.Left == nil {
			n.Left = &Node{Value: value, Data: data}
			return nil
		}
		return n.Left.Insert(value, data)
	case value > n.Value:
		if n.Right == nil {
			n.Right = &Node{Value: value, Data: data}
			return nil
		}
		return n.Right.Insert(value, data)
	}
	return nil
}

func (n *Node) Find(s uint64) (Proc, bool) {
	if n == nil {
		return Proc{}, false
	}
	switch {
	case s == n.Value:
		return n.Data, true
	case s < n.Value:
		return n.Left.Find(s)
	default:
		return n.Right.Find(s)
	}
}

func (n *Node) FindMin() Proc {
	if n.Left == nil {
		return n.Data
	}
	return n.Left.FindMin()
}

func (n *Node) findMax(parent *Node) (*Node, *Node) {
	if n == nil {
		return nil, parent
	}
	if n.Right == nil {
		return n, parent
	}
	return n.Right.findMax(n)
}

func (n *Node) replaceNode(parent, replacement *Node) error {
	if n == parent.Left {
		parent.Left = replacement
		return nil
	}
	parent.Right = replacement
	return nil
}

func (n *Node) Delete(s uint64, parent *Node) error {
	switch {
	case s < n.Value:
		return n.Left.Delete(s, n)
	case s > n.Value:
		return n.Right.Delete(s, n)
	default:
		if n.Left == nil && n.Right == nil {
			n.replaceNode(parent, nil)
			return nil
		}
		if n.Left == nil {
			n.replaceNode(parent, n.Right)
			return nil
		}
		if n.Right == nil {
			n.replaceNode(parent, n.Left)
			return nil
		}
		replacement, replParent := n.Left.findMax(n)
		n.Value = replacement.Value
		n.Data = replacement.Data
		return replacement.Delete(replacement.Value, replParent)
	}
}

type Tree struct {
	Root *Node
}

func (t *Tree) Insert(value uint64, data Proc) error {
	if t.Root == nil {
		t.Root = &Node{Value: value, Data: data}
		return nil
	}
	return t.Root.Insert(value, data)
}

func (t *Tree) Find(s uint64) (Proc, bool) {
	if t.Root == nil {
		return Proc{}, false
	}
	return t.Root.Find(s)
}

func (t *Tree) FindMin() Proc {
	min := t.Root.FindMin()
	return min
}

func (t *Tree) Delete(s uint64) error {
	fakeParent := &Node{Right: t.Root}
	err := t.Root.Delete(s, fakeParent)
	if err != nil {
		return err
	}
	if fakeParent.Right == nil {
		t.Root = nil
	}
	t.Root = fakeParent.Right
	return nil
}

func (t *Tree) Traverse(n *Node, f func(*Node)) {
	if n == nil {
		return
	}
	t.Traverse(n.Left, f)
	f(n)
	t.Traverse(n.Right, f)
}

func main() {
	in := bufio.NewReader(os.Stdin)

	n := 0 // procs
	m := 0 // tasks
	fmt.Fscan(in, &n, &m)

	freeProcs := &Tree{}
	for i := 0; i < n; i++ {
		proc := Proc{}
		fmt.Fscan(in, &proc.Power)
		freeProcs.Insert(proc.Power, proc)
	}
	busyProcs := &Tree{}

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

		for busyProcs.Root != nil {
			// proc := busyProcs.Root.FindMin()
			proc := busyProcs.FindMin()

			if proc.BusyUntil <= curTime {
				busyProcs.Delete(proc.BusyUntil)
				proc.BusyUntil = 0
				freeProcs.Insert(proc.Power, proc)
			} else {
				break
			}
		}
		if freeProcs.Root == nil {
			continue
		}

		// proc := freeProcs.Root.FindMin()
		proc := freeProcs.FindMin()
		proc.BusyUntil = curTime + curTask.WorkTime
		busyProcs.Insert(proc.BusyUntil, proc)
		freeProcs.Delete(proc.Power)

		totalPower += curTask.WorkTime * proc.Power
	}

	fmt.Println(totalPower)
}
