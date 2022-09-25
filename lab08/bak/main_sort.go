package main

import (
	"errors"
	"fmt"
	"math/rand"
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
	val   Proc
	left  *TreeNode
	right *TreeNode
}

func (t *TreeNode) Insert(value Proc) error {
	if t == nil {
		return errors.New("Tree is nil")
	}
	if t.val == value {
		return errors.New("This node value already exists")
	}
	if t.val.Power > value.Power {
		if t.left == nil {
			t.left = &TreeNode{val: value}
			return nil
		}
		return t.left.Insert(value)
	}
	if t.val.Power < value.Power {
		if t.right == nil {
			t.right = &TreeNode{val: value}
			return nil
		}
		return t.right.Insert(value)
	}
	return nil
}

func (t *TreeNode) Find(value uint64) (TreeNode, bool) {
	if t == nil {
		return TreeNode{}, false
	}
	switch {
	case value == t.val.Power:
		return *t, true
	case value < t.val.Power:
		return t.left.Find(value)
	default:
		return t.right.Find(value)
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

func main() {

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
	var busyProcs *TreeNode

	freeProcs := &TreeNode{val: Proc{1000000, 0}}

	tasks := make([]Task, m)
	tasks[0] = Task{
		StartTime: 1000000000,
		WorkTime:  1000000,
	}

	// tasks := make(map[int]int, m)
	// tasks[1000000000] = 1000000
	// tasks[2] = 5
	// tasks[3] = 7
	// tasks[4] = 10
	// tasks[5] = 5
	// tasks[6] = 100
	// tasks[9] = 2

	// totalPower := 0
	var totalPower uint64
	// minTime := 1000000000
	// maxTime := 1000000000

	var curTime uint64

	// for curTime := minTime; curTime <= maxTime; curTime++ {

	for _, curTask := range tasks {
		curTime = curTask.StartTime
		proc := freeProcs.FindMin()
		proc.BusyUntil = curTime + curTask.WorkTime
		freeProcs.Delete(proc)
		busyProcs.Insert(proc)
		totalPower += curTask.WorkTime * proc.Power
	}

	fmt.Println(totalPower)
}

func qsort(a []Proc) []Proc {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	// Pick a pivot
	pivotIndex := rand.Uint64() % uint64(len(a))

	// Move the pivot to the right
	a[pivotIndex], a[right] = a[right], a[pivotIndex]

	// Pile elements smaller than the pivot on the left
	for i := range a {
		if a[i].Power < a[right].Power {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}

	// Place the pivot after the last smaller element
	a[left], a[right] = a[right], a[left]

	// Go down the rabbit hole
	qsort(a[:left])
	qsort(a[left+1:])

	return a
}
