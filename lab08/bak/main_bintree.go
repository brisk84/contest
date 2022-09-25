package main

import (
	"errors"
	"fmt"
	"log"
)

type Proc struct {
	Power     uint64
	BusyUntil uint64
}

type Task struct {
	StartTime uint64
	WorkTime  uint64
}

type TreeNode struct {
	val    Proc
	parent *TreeNode
	left   *TreeNode
	right  *TreeNode
}

func (t *TreeNode) Insert(value Proc) error {
	// if t == nil {
	// 	return errors.New("Tree is nil")
	// }
	// if t.val == value {
	// 	return errors.New("This node value already exists")
	// }
	if t.val.Power > value.Power {
		if t.left == nil {
			t.left = &TreeNode{val: value, parent: t}
			return nil
		}
		return t.left.Insert(value)
	}
	if t.val.Power < value.Power {
		if t.right == nil {
			t.right = &TreeNode{val: value, parent: t}
			return nil
		}
		return t.right.Insert(value)
	}
	return nil
}

func (t *TreeNode) FindPower(value uint64) (TreeNode, bool) {
	if t == nil {
		return TreeNode{}, false
	}
	switch {
	case value == t.val.Power:
		return *t, true
	case value < t.val.Power:
		return t.left.FindPower(value)
	default:
		return t.right.FindPower(value)
	}
}

func (t *TreeNode) FindTime(value uint64) (TreeNode, bool) {
	if t == nil {
		return TreeNode{}, false
	}
	switch {
	case value == t.val.BusyUntil:
		return *t, true
	case value < t.val.BusyUntil:
		return t.left.FindTime(value)
	default:
		return t.right.FindTime(value)
	}
}

func (t *TreeNode) Delete(value Proc) {
	t.remove(value)
}

func (t *TreeNode) remove(value Proc) *TreeNode {

	if t == nil {
		return nil
	}

	if value.Power < t.val.Power {
		t.left = t.left.remove(value)
		return t
	}
	if value.Power > t.val.Power {
		t.right = t.right.remove(value)
		return t
	}

	if t.left == nil && t.right == nil {
		t = nil
		return nil
	}

	if t.left == nil {
		t = t.right
		return t
	}
	if t.right == nil {
		t = t.left
		return t
	}

	smallestValOnRight := t.right
	for {
		if smallestValOnRight != nil && smallestValOnRight.left != nil {
			smallestValOnRight = smallestValOnRight.left
		} else {
			break
		}
	}

	t.val = smallestValOnRight.val
	t.right = t.right.remove(t.val)
	return t
}

func (t *TreeNode) FindMax() Proc {
	if t.right == nil {
		return t.val
	}
	return t.right.FindMax()
}

func (t *TreeNode) FindMin() Proc {
	if t.left == nil {
		return t.val
	}
	return t.left.FindMin()
}

func (t *TreeNode) PrintInorder() {

	if t == nil {

		return
	}

	t.left.PrintInorder()
	fmt.Print(t.val)
	t.right.PrintInorder()
}

type Node struct {
	Value string
	Data  string
	Left  *Node
	Right *Node
}

func (n *Node) Insert(value, data string) error {

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

func (n *Node) Find(s string) (string, bool) {

	if n == nil {
		return "", false
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

func (n *Node) Delete(s string, parent *Node) error {
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

func (t *Tree) Insert(value, data string) error {
	if t.Root == nil {
		t.Root = &Node{Value: value, Data: data}
		return nil
	}
	return t.Root.Insert(value, data)
}

func (t *Tree) Find(s string) (string, bool) {
	if t.Root == nil {
		return "", false
	}
	return t.Root.Find(s)
}

func (t *Tree) Delete(s string) error {

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
	values := []string{"d", "b", "c", "e", "a"}
	data := []string{"delta", "bravo", "charlie", "echo", "alpha"}
	tree := &Tree{}
	for i := 0; i < len(values); i++ {
		err := tree.Insert(values[i], data[i])
		if err != nil {
			log.Fatal("Error inserting value '", values[i], "': ", err)
		}
	}
	fmt.Print("Sorted values: | ")
	tree.Traverse(tree.Root, func(n *Node) { fmt.Print(n.Value, ": ", n.Data, " | ") })
	fmt.Println()

	s := "d"
	fmt.Print("Find node '", s, "': ")
	d, found := tree.Find(s)
	if !found {
		log.Fatal("Cannot find '" + s + "'")
	}
	fmt.Println("Found " + s + ": '" + d + "'")

	err := tree.Delete(s)
	if err != nil {
		log.Fatal("Error deleting "+s+": ", err)
	}
	fmt.Print("After deleting '" + s + "': ")
	tree.Traverse(tree.Root, func(n *Node) { fmt.Print(n.Value, ": ", n.Data, " | ") })
	fmt.Println()

	fmt.Println("Single-node tree")
	tree = &Tree{}

	tree.Insert("a", "alpha")
	fmt.Println("After insert:")
	tree.Traverse(tree.Root, func(n *Node) { fmt.Print(n.Value, ": ", n.Data, " | ") })
	fmt.Println()

	tree.Delete("a")
	fmt.Println("After delete:")
	tree.Traverse(tree.Root, func(n *Node) { fmt.Print(n.Value, ": ", n.Data, " | ") })
	fmt.Println()

	return

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

	procs := &TreeNode{val: Proc{3, 0}}
	procs.Insert(Proc{2, 0})
	procs.Insert(Proc{6, 0})
	procs.Insert(Proc{4, 0})

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

		if proc, found := procs.FindTime(curTime); found {
			proc.val.BusyUntil = 0
			// busyProcs.Delete(proc.val)
			// freeProcs.Insert(proc.val)
		}

		proc := procs.FindMin()
		proc.BusyUntil = curTime + curTask.WorkTime
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

type TreeNode2 struct {
	val   Proc
	left  *TreeNode2
	right *TreeNode2
}

func (t *TreeNode2) Insert(value Proc) error {
	// if t == nil {
	// 	return errors.New("Tree is nil")
	// }
	// if t.val == value {
	// 	return errors.New("This node value already exists")
	// }
	if t.val.BusyUntil > value.BusyUntil {
		if t.left == nil {
			t.left = &TreeNode2{val: value}
			return nil
		}
		return t.left.Insert(value)
	}
	if t.val.BusyUntil < value.BusyUntil {
		if t.right == nil {
			t.right = &TreeNode2{val: value}
			return nil
		}
		return t.right.Insert(value)
	}
	return nil
}

func (t *TreeNode2) Find(value uint64) (TreeNode2, bool) {
	if t == nil {
		return TreeNode2{}, false
	}
	switch {
	case value == t.val.BusyUntil:
		return *t, true
	case value < t.val.BusyUntil:
		return t.left.Find(value)
	default:
		return t.right.Find(value)
	}
}

func (t *TreeNode2) Delete(value Proc) {
	t.remove(value)
}

func (t *TreeNode2) remove(value Proc) *TreeNode2 {

	if t == nil {
		return nil
	}

	if value.BusyUntil < t.val.BusyUntil {
		t.left = t.left.remove(value)
		return t
	}
	if value.BusyUntil > t.val.BusyUntil {
		t.right = t.right.remove(value)
		return t
	}

	if t.left == nil && t.right == nil {
		t = nil
		return nil
	}

	if t.left == nil {
		t = t.right
		return t
	}
	if t.right == nil {
		t = t.left
		return t
	}

	smallestValOnRight := t.right
	for {
		if smallestValOnRight != nil && smallestValOnRight.left != nil {
			smallestValOnRight = smallestValOnRight.left
		} else {
			break
		}
	}

	t.val = smallestValOnRight.val
	t.right = t.right.remove(t.val)
	return t
}

func (t *TreeNode2) FindMax() Proc {
	if t.right == nil {
		return t.val
	}
	return t.right.FindMax()
}

func (t *TreeNode2) FindMin() Proc {
	if t.left == nil {
		return t.val
	}
	return t.left.FindMin()
}
