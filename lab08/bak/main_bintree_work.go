package main

import (
	"errors"
	"fmt"
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

	if n == nil {
		return errors.New("Cannot insert a value into a nil tree")
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
	if n == nil {
		return errors.New("replaceNode() not allowed on a nil node")
	}

	if n == parent.Left {
		parent.Left = replacement
		return nil
	}
	parent.Right = replacement
	return nil
}

func (n *Node) Delete(s uint64, parent *Node) error {
	if n == nil {
		return errors.New("Value to be deleted does not exist in the tree")
	}
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

func (t *Tree) Delete(s uint64) error {

	if t.Root == nil {
		return errors.New("Cannot delete from an empty tree")
	}
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

	// t := Tree{}
	// t.Insert(5, Proc{2, 2})
	// t.Insert(1, Proc{1, 1})
	// t.Insert(7, Proc{3, 3})
	// // t.Insert(8, Proc{3, 3})
	// // t.Insert(0, Proc{3, 3})
	// // t.Insert(3, Proc{3, 3})
	// // t.Insert(2, Proc{3, 3})

	// err1 := t.Delete(5)
	// if err1 != nil {
	// 	panic(err1)
	// }
	// t.Traverse(t.Root, func(n *Node) { fmt.Print(n.Value, ": ", n.Data, " | ") })

	// return

	// in := bufio.NewReader(os.Stdin)
	// n := 0 // procs
	// m := 0 // tasks
	// fmt.Fscan(in, &n, &m)
	// // // power := make([]int, n)
	// procs := make([]Proc, n)
	// for i := 0; i < n; i++ {
	// 	fmt.Fscan(in, &procs[i].Power)
	// }
	// sort.Slice(procs, func(i, j int) bool {
	// 	return procs[i].Power < procs[j].Power
	// })

	// minTime := 0
	// maxTime := 0
	// tasks := make(map[int]int, m)
	// for i := 0; i < m; i++ {
	// 	startTime := 0
	// 	workTime := 0
	// 	fmt.Fscan(in, &startTime, &workTime)
	// 	tasks[startTime] = workTime
	// if i == 0 {
	// 	minTime = starTime
	// }
	// 	if i == m-1 {
	// 		maxTime = startTime
	// 	}
	// }

	m := 1 // tasks
	// var m uint64 = 1
	// procs := []Proc{
	// 	{1000000, 0},
	// }
	// sort.Slice(procs, func(i, j int) bool {
	// 	return procs[i].Power < procs[j].Power
	// })
	// procs = qsort(procs)

	// var freeProcs *TreeNode
	// var busyProcs *TreeNode2
	// freeProcs := &TreeNode{val: Proc{1000000, 0}}
	// tasks := make([]Task, m)
	// tasks[0] = Task{
	// 	StartTime: 1000000000,
	// 	WorkTime:  1000000,
	// }

	// n = 4
	m = 7

	// procs := &TreeNode{val: Proc{3, 0}}
	// procs.Insert(Proc{2, 0})
	// procs.Insert(Proc{6, 0})
	// procs.Insert(Proc{4, 0})
	freeProcs := &Tree{}
	freeProcs.Insert(3, Proc{3, 0})
	freeProcs.Insert(2, Proc{2, 0})
	freeProcs.Insert(4, Proc{4, 0})
	freeProcs.Insert(6, Proc{6, 0})

	busyProcs := &Tree{}

	tasks := make([]Task, m)
	tasks = []Task{
		Task{1, 3},
		Task{2, 5},
		Task{3, 7},
		Task{4, 10},
		Task{5, 5},
		Task{6, 100},
		Task{9, 2},
	}

	// 4 7
	// 3 2 6 4
	// 1 3
	// 2 5
	// 3 7
	// 4 10
	// 5 5
	// 6 100
	// 9 2

	// totalPower := 0
	var totalPower uint64
	// minTime := 1000000000
	// maxTime := 1000000000

	var curTime uint64

	// for curTime := minTime; curTime <= maxTime; curTime++ {

	for _, curTask := range tasks {
		curTime = curTask.StartTime

		// if proc, found := busyProcs.Find(curTime); found {
		// 	proc.BusyUntil = 0
		// 	busyProcs.Delete(curTime)
		// 	freeProcs.Insert(proc.Power, proc)
		// }
		for busyProcs.Root != nil {
			// if busyProcs.Root.Left == nil {
			// 	break
			// }
			var proc *Proc
			if busyProcs.Root.Left != nil {
				proc = &busyProcs.Root.Left.Data
			} else {
				proc = &busyProcs.Root.Data
			}

			// proc := busyProcs.Root.Left.Data
			if proc.BusyUntil <= curTime {
				err := busyProcs.Delete(proc.BusyUntil)
				if err != nil {
					panic(err)
				}
				proc.BusyUntil = 0
				freeProcs.Insert(proc.Power, *proc)
			} else {
				break
			}
		}

		var proc *Proc

		if freeProcs.Root != nil {
			if freeProcs.Root.Left != nil {
				proc = &freeProcs.Root.Left.Data
			} else {
				proc = &freeProcs.Root.Data
			}
		} else {
			continue
		}
		// proc := freeProcs.Root.Left.Data
		proc.BusyUntil = curTime + curTask.WorkTime
		freeProcs.Delete(proc.Power)
		busyProcs.Insert(proc.BusyUntil, *proc)

		// procs.Delete(proc)
		// if busyProcs == nil {
		// 	busyProcs = &TreeNode2{val: proc}
		// } else {
		// 	busyProcs.Insert(proc)
		// }
		totalPower += curTask.WorkTime * proc.Power
	}

	fmt.Println(totalPower)
}
