package main

import (
	"bufio"
	"fmt"
	"os"
)

type Proc struct {
	Power     uint64
	BusyUntil uint64
	Mark      string
}

type Task struct {
	StartTime uint64
	WorkTime  uint64
}

type Node struct {
	Value uint64
	Data  []Proc
	Left  *Node
	Right *Node
}

func (n *Node) Insert(value uint64, data Proc) error {
	// if value == n.Value {
	// 	value--
	// 	data.BusyUntil = value
	// }
	switch {
	case value == n.Value:
		// fmt.Println("?????????")
		n.Data = append(n.Data, data)
		return nil
	case value < n.Value:
		if n.Left == nil {
			node := Node{Value: value}
			node.Data = append(node.Data, data)
			n.Left = &node
			return nil
		}
		return n.Left.Insert(value, data)
	case value > n.Value:
		if n.Right == nil {
			node := Node{Value: value}
			node.Data = append(node.Data, data)
			n.Right = &node
			// n.Right = &Node{Value: value, Data: data}
			return nil
		}
		return n.Right.Insert(value, data)
	}
	return nil
}

func (n *Node) Find(s uint64) ([]Proc, bool) {
	if n == nil {
		return []Proc{}, false
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

func (n *Node) FindMin() []Proc {
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

func (n *Node) Delete(s uint64, parent *Node, data Proc) error {
	// if s < n.Value && n.Left == nil {
	// 	s++
	// }
	switch {
	case s < n.Value:
		return n.Left.Delete(s, n, data)
	case s > n.Value:
		return n.Right.Delete(s, n, data)
	default:
		if len(n.Data) > 1 {
			// fmt.Println("!!!!!!!!!!!!!", len(n.Data))
			newData := []Proc{}
			for _, v := range n.Data {
				if v.Power != data.Power {
					newData = append(newData, v)
				}
			}
			n.Data = newData
			return nil
		}
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
		return replacement.Delete(replacement.Value, replParent, data)
	}
}

type Tree struct {
	Root *Node
}

func (t *Tree) Insert(value uint64, data Proc) error {
	if t.Root == nil {
		node := Node{Value: value}
		node.Data = append(node.Data, data)
		t.Root = &node
		return nil
	}
	return t.Root.Insert(value, data)
}

func (t *Tree) Find(s uint64) ([]Proc, bool) {
	if t.Root == nil {
		return []Proc{}, false
	}
	return t.Root.Find(s)
}

func (t *Tree) FindMin() []Proc {
	min := t.Root.FindMin()
	// if min.Power != t.Min {
	// 	log.Fatal(min.Power, t.Min)
	// }
	return min
}

func (t *Tree) Delete(s uint64, data Proc) error {
	fakeParent := &Node{Right: t.Root}
	err := t.Root.Delete(s, fakeParent, data)
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

	// x := []int{1, 2, 3, 4, 5}
	// x2 := append(x[0:3], x[4:]...)
	// fmt.Println(x2)
	// return

	// p1 := Proc{1, 1, ""}
	// p2 := Proc{2, 2, ""}
	// p3 := Proc{3, 3, ""}
	// p4 := Proc{4, 4, ""}
	// p5 := Proc{5, 5, ""}

	// t := Tree{}
	// t.Insert(2, p2)
	// t.Insert(2, p5)
	// t.Insert(3, p3)
	// t.Insert(1, p1)
	// t.Insert(4, p4)
	// t.Insert(5, p5)
	// t.Traverse(t.Root, func(n *Node) { fmt.Print(n.Value, ": ", n.Data, " | ") })
	// return

	// pp := t.Root.FindMin()
	// fmt.Println(pp)
	// t.Delete(5)
	// pp = t.Root.FindMin()
	// fmt.Println(pp)
	// t.Delete(1)
	// pp = t.Root.FindMin()
	// fmt.Println(pp)
	// t.Delete(3)
	// // t.Delete(5)
	// pp = t.Root.FindMin()
	// fmt.Println(pp)

	// // pp = t.Root.FindMin()
	// // fmt.Println(pp)

	// t.Traverse(t.Root, func(n *Node) { fmt.Print(n.Value, ": ", n.Data, " | ") })

	// return

	// f, err := os.Open("/Users/brisk/vscode/contest/lab08/tests/11")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer f.Close()
	// in := bufio.NewReader(f)

	// fw, err := os.Create("/Users/brisk/vscode/contest/lab08/log2.txt")
	// defer fw.Close()
	// out := bufio.NewWriter(fw)

	in := bufio.NewReader(os.Stdin)

	n := 0 // procs
	m := 0 // tasks
	fmt.Fscan(in, &n, &m)

	freeProcs := &Tree{}
	for i := 0; i < n; i++ {
		proc := Proc{}
		fmt.Fscan(in, &proc.Power)
		// if _, found := freeProcs.Find(proc.Power); found {
		// 	xx := 0
		// 	_ = xx
		// }
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

		// if totalPower == 151181876613406 {
		// 	xx := 0
		// 	_ = xx
		// 	pp, found := freeProcs.Find(315220)
		// 	pp, found = busyProcs.Find(483228)

		// 	_ = pp
		// 	_ = found

		// }

		for busyProcs.Root != nil {
			// proc := busyProcs.Root.FindMin()
			proc := busyProcs.FindMin()

			// if proc.BusyUntil == 483228 {
			// 	xx := 0
			// 	_ = xx
			// }
			br := false
			for _, p := range proc {
				if p.BusyUntil <= curTime {
					busyProcs.Delete(p.BusyUntil, p)
					p.BusyUntil = 0
					freeProcs.Insert(p.Power, p)
				} else {
					br = true
					break
				}
			}
			if br {
				break
			}
		}
		if freeProcs.Root == nil {
			continue
		}

		// proc := freeProcs.Root.FindMin()
		proc := freeProcs.FindMin()[0]

		// if proc.Power == 315220 {
		// 	xx := 0
		// 	proc.Mark = "!!!"
		// 	_ = xx
		// }
		busyUntil := curTime + curTask.WorkTime //+ uint64(rand.Intn(2-1)+1)
		// if _, found := busyProcs.Find(busyUntil); found {
		// 	// busyUntil--
		// 	fmt.Println("!!!")
		// }
		// proc.BusyUntil = curTime + curTask.WorkTime
		proc.BusyUntil = busyUntil
		busyProcs.Insert(proc.BusyUntil, proc)
		freeProcs.Delete(proc.Power, proc)

		totalPower += curTask.WorkTime * proc.Power
		// fmt.Fprintln(out, totalPower)
	}

	fmt.Println(totalPower)
}
