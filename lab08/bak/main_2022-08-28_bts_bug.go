package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
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
	key     uint64
	data    []Proc
	left    *Node
	right   *Node
	parent  *Node
	balance int
}

func (n *Node) Find(key uint64) *Node {
	if key == n.key {
		return n
	}

	if n.left != nil && key < n.key {
		return n.left.Find(key)
	}

	if n.right != nil && key > n.key {
		return n.right.Find(key)
	}
	return nil
}

func (n *Node) FindMin() *Node {
	if n.left != nil {
		return n.left.FindMin()
	}
	return n
}

func (t *Node) findParentNode(key uint64) *Node {
	node := t
	for {
		if node.key == key {
			break
		}
		nodeAddr := node.selectNextNode(key)
		n := *nodeAddr
		if n == nil {
			break
		}
		node = n
	}
	return node
}

func (n *Node) Insert(key uint64, data Proc) *Node {
	newNode := &Node{
		key: key,
	}
	newNode.data = append(newNode.data, data)
	lastNode := n.findParentNode(key)
	if lastNode.key == key {
		lastNode.data = append(lastNode.data, data)
		return n
		// return lastNode
	}
	newNode.parent = lastNode

	nextNode := lastNode.selectNextNode(key)
	*nextNode = newNode
	newNode.recalculateBalance(1)
	return newNode.balanceTree()
}

func (n *Node) selectNextNode(key uint64) **Node {
	if key < n.key {
		return &n.left
	}
	return &n.right
}

// Функция ребалансировки относительно текущего узла, принимает 1 на добавление узла -1 на удаление
func (n *Node) recalculateBalance(direction int) {
	currentNode := n

	for currentNode.parent != nil {
		balance := 1 * direction
		if currentNode.key < currentNode.parent.key {
			balance = -1 * direction
		}
		currentNode.parent.balance += balance
		if direction == 1 && currentNode.parent.balance == 0 ||
			direction == -1 &&
				currentNode.parent.left != nil &&
				currentNode.parent.right != nil &&
				(currentNode.parent.balance-balance >= 0 && currentNode.parent.balance > 0 || currentNode.parent.balance-balance < 0 && currentNode.parent.balance < 0) {
			break
		}

		currentNode = currentNode.parent
	}
}

func (n *Node) Remove(key uint64) *Node {
	if key < n.key {
		return n.left.Remove(key)
	}
	if key > n.key {
		return n.right.Remove(key)
	}

	if len(n.data) > 1 {
		n.data = n.data[:(len(n.data) - 1)]
		return n.findRootNod()
	}

	nextNode := n.parent.selectNextNode(key)
	// Первый случай, у узла нет потомков
	if n.left == nil && n.right == nil {
		*nextNode = nil
		n.parent.balance = 0
		n.parent.recalculateBalance(-1)
		return n.parent.balanceTree()
	}

	// Второй случай, у зла 1 потомок
	if n.left == nil {
		*nextNode = n.right
		n.right.parent = n.parent
		n.right.recalculateBalance(-1)
		return n.right.balanceTree()
	}

	if n.right == nil {
		*nextNode = n.left
		n.left.parent = n.parent
		n.left.recalculateBalance(-1)
		return n.left.balanceTree()
	}

	// 3 случай, у удаляемого узла 2 потомка
	maxLeftNode := n.right
	for {
		if maxLeftNode.left == nil {
			break
		}
		maxLeftNode = n.right.left
	}

	maxLeftNode.left = n.left
	n.left.parent = maxLeftNode

	n.right.parent = n.parent
	n.right.balance = n.balance - 1
	// Перерасчитываем баланс для родителя, если правое поддерево было максимальным по длине
	if n.balance > 0 && n.right.balance <= 0 {
		n.right.recalculateBalance(-1)
	}
	*nextNode = n.right
	return n.right.balanceTree()
}

func (n *Node) String() string {
	if n == nil {
		return ""
	}
	text := ""
	parentNode := n
	for parentNode != nil {
		parentNode = parentNode.parent
		text += "-"
	}
	text += " " + strconv.Itoa(int(n.key)) + " " + strconv.Itoa(n.balance) + "\n"

	text += n.left.String()
	text += n.right.String()

	return text
}

func (n *Node) balanceTree() *Node {
	if n.right == nil {
		return n
	}
	currentNode := n
	for {
		if currentNode.balance == 2 {
			if currentNode.right.balance < 0 {
				currentNode.right = currentNode.right.rotateRight()
			}
			currentNode = currentNode.rotateLeft()
			break
		}
		if currentNode.balance == -2 {
			if currentNode.left.balance > 0 {
				currentNode.left = currentNode.left.rotateLeft()
			}
			currentNode = currentNode.rotateRight()
			break
		}

		if currentNode.parent == nil {
			break
		}

		currentNode = currentNode.parent
	}

	return currentNode.findRootNod()
}

func (n *Node) rotateLeft() *Node {
	if n.parent == nil {
		return n
	}
	p := n.right
	p.parent = n.parent
	if n.parent != nil {
		parent := n.parent.selectNextNode(n.key)
		*parent = p
	}
	n.right = p.left
	if p.left != nil {
		p.left.parent = n
	}

	p.left = n
	n.parent = p

	n.balance = 0
	p.balance -= 1
	p.recalculateBalance(-1)
	return p
}

func (n *Node) rotateRight() *Node {
	if n.parent == nil {
		return n
	}
	q := n.left
	q.parent = n.parent
	parent := n.parent.selectNextNode(n.key)
	*parent = q
	n.left = q.right
	if q.right != nil {
		q.right.parent = n
	}

	q.right = n
	n.parent = q

	n.balance = 0
	q.balance += 1
	q.recalculateBalance(+1)
	return q
}

func (n *Node) findRootNod() *Node {
	rootNode := n

	for rootNode.parent != nil {
		rootNode = rootNode.parent
	}

	return rootNode
}

type Tree struct {
	root *Node
}

func (t *Tree) Insert(key uint64, data Proc) {
	if t.root == nil {
		node := &Node{
			key:     key,
			balance: 0,
		}
		node.data = append(node.data, data)
		t.root = node
		return
	}
	t.root = t.root.Insert(key, data)
}

func (t *Tree) Remove(key uint64) {
	if t.root == nil {
		return
	}
	if t.root.left == nil && t.root.right == nil {
		t.root = nil
		return
	}
	fakeParent := &Node{right: t.root}
	t.root.parent = fakeParent
	fakeParent.Remove(key)
	t.root.parent = nil
	t.root = fakeParent.right
	// t.roo
	// err := t.Root.Delete(s, fakeParent)
	// if err != nil {
	// 	return err
	// }
	// if fakeParent.Right == nil {
	// 	t.Root = nil
	// }
	// t.Root = fakeParent.Right
	// !!! orig
	// t.root = t.root.Remove(key)
}

func (t *Tree) FindMin() []Proc {
	if t.root == nil {
		return nil
	}
	return t.root.FindMin().data
}

func main() {

	// p1 := Proc{1, 1}
	// p2 := Proc{2, 2}
	// p3 := Proc{3, 3}

	// t := Tree{}
	// t.Insert(9, p1)
	// t.Insert(6, p2)
	// t.Insert(3, p3)
	// // t.Insert(4, p3)

	// // t.Remove(5)
	// // t.Remove(2)
	// // // t.Remove(6)

	// // t.Insert(2, p2)

	// // x := 0
	// // _ = x
	// return

	ts := time.Now()

	f, err := os.Open("/Users/brisk/vscode/contest/lab08/tests/05")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	in := bufio.NewReader(f)

	// in := bufio.NewReader(os.Stdin)

	n := 0 // procs
	m := 0 // tasks
	fmt.Fscan(in, &n, &m)

	freeProcs := Tree{}
	busyProcs := Tree{}

	for i := 0; i < n; i++ {
		proc := Proc{}
		fmt.Fscan(in, &proc.Power)
		freeProcs.Insert(proc.Power, proc)
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

	// cntr := 0
	for _, curTask := range tasks {
		// if cntr%1000 == 0 {
		// if freeProcs.Root != nil {
		// 	freeProcs.Root = freeProcs.Root.balanceTree()
		// }
		// if busyProcs.Root != nil {
		// 	busyProcs.Root = busyProcs.Root.balanceTree()
		// }
		// }
		curTime = curTask.StartTime

		for busyProcs.root != nil {
			proc := busyProcs.FindMin()
			br := false
			for _, p := range proc {
				if p.BusyUntil <= curTime {
					// busyProcs.Delete(p.BusyUntil, p)
					busyProcs.Remove(p.BusyUntil)
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
		if freeProcs.root == nil {
			continue
		}

		n := freeProcs.FindMin()
		proc := n[0]

		busyUntil := curTime + curTask.WorkTime
		proc.BusyUntil = busyUntil
		busyProcs.Insert(proc.BusyUntil, proc)
		// freeProcs.Delete(proc.Power, proc)
		freeProcs.Remove(proc.Power)

		totalPower += curTask.WorkTime * proc.Power
		// fmt.Println(curTime, proc.Power, curTask.StartTime, totalPower)
	}

	fmt.Println(totalPower)

	te := time.Since(ts)
	fmt.Println(te)
}
