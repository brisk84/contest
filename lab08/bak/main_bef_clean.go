package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Btree represents an AVL tree
type Btree struct {
	root   *Node
	values []Val
	len    int
}

// Val interface to define the compare method used to insert and find values
type Val interface {
	Comp(val Val) int8
}

// Node represents a node in the tree with a value, left and right children, and a height/balance of the node.
type Node struct {
	Value       Val
	left, right *Node
	height      int8
}

// New returns a new btree
func New() *Btree { return new(Btree).Init() }

// Init initializes all values/clears the tree and returns the tree pointer
func (t *Btree) Init() *Btree {
	t.root = nil
	t.values = nil
	t.len = 0
	return t
}

// String returns a string representation of the tree values
func (t *Btree) String() string {
	return fmt.Sprint(t.Values())
}

// Empty returns true if the tree is empty
func (t *Btree) Empty() bool {
	return t.root == nil
}

// NotEmpty returns true if the tree is not empty
func (t *Btree) NotEmpty() bool {
	return t.root != nil
}

func (t *Btree) balance() int8 {
	if t.root != nil {
		return balance(t.root)
	}
	return 0
}

// Insert inserts a new value into the tree and returns the tree pointer
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
		// n.Value = value
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

// InsertAll inserts all the values into the tree and returns the tree pointer
func (t *Btree) InsertAll(values []Val) *Btree {
	for _, v := range values {
		t.Insert(v)
	}
	return t
}

// Contains returns true if the tree contains the specified value
func (t *Btree) Contains(value Val) bool {
	return t.Get(value) != nil
}

// ContainsAny returns true if the tree contains any of the values
func (t *Btree) ContainsAny(values []Val) bool {
	for _, v := range values {
		if t.Contains(v) {
			return true
		}
	}
	return false
}

// ContainsAll returns true if the tree contains all of the values
func (t *Btree) ContainsAll(values []Val) bool {
	for _, v := range values {
		if !t.Contains(v) {
			return false
		}
	}
	return true
}

// Get returns the node value associated with the search value
func (t *Btree) Get(value Val) Val {
	var node *Node
	if t.root != nil {
		node = t.root.get(value)
	}
	if node != nil {
		return node.Value
	}
	return nil
}

// Len return the number of nodes in the tree
func (t *Btree) Len() int {
	return t.len
}

// Head returns the first value in the tree
func (t *Btree) Head() Val {
	if t.root == nil {
		return nil
	}
	var beginning = t.root
	for beginning.left != nil {
		beginning = beginning.left
	}
	if beginning == nil {
		for beginning.right != nil {
			beginning = beginning.right
		}
	}
	if beginning != nil {
		return beginning.Value
	}
	return nil
}

// Tail returns the last value in the tree
func (t *Btree) Tail() Val {
	if t.root == nil {
		return nil
	}
	var beginning = t.root
	for beginning.right != nil {
		beginning = beginning.right
	}
	if beginning == nil {
		for beginning.left != nil {
			beginning = beginning.left
		}
	}
	if beginning != nil {
		return beginning.Value
	}
	return nil
}

// Values returns a slice of all the values in tree in order
func (t *Btree) Values() []Val {
	if t.values == nil {
		t.values = make([]Val, t.len)
		t.Ascend(func(n *Node, i int) bool {
			t.values[i] = n.Value
			return true
		})
	}
	return t.values
}

// Delete deletes the node from the tree associated with the search value
func (t *Btree) Delete(value Val) *Btree {
	deleted := false
	t.root = deleteNode(t.root, value, &deleted)
	if deleted {
		t.len--
	}
	t.values = nil
	return t
}

// DeleteAll deletes the nodes from the tree associated with the search values
func (t *Btree) DeleteAll(values []Val) *Btree {
	for _, v := range values {
		t.Delete(v)
	}
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
			// *deleted = false
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

// Pop deletes the last node from the tree and returns its value
func (t *Btree) Pop() Val {
	value := t.Tail()
	if value != nil {
		t.Delete(value)
	}
	return value
}

// Pull deletes the first node from the tree and returns its value
func (t *Btree) Pull() Val {
	value := t.Head()
	if value != nil {
		t.Delete(value)
	}
	return value
}

// NodeIterator expresses the iterator function used for traversals
type NodeIterator func(n *Node, i int) bool

// Ascend performs an ascending order traversal of the tree calling the iterator function on each node
// the iterator will continue as long as the NodeIterator returns true
func (t *Btree) Ascend(iterator NodeIterator) {
	var i int
	if t.root != nil {
		t.root.iterate(iterator, &i, true)
	}
}

// Descend performs a descending order traversal of the tree using the iterator
// the iterator will continue as long as the NodeIterator returns true
func (t *Btree) Descend(iterator NodeIterator) {
	var i int
	if t.root != nil {
		t.root.rIterate(iterator, &i, true)
	}
}

// Debug prints out useful debug information about the tree for debugging purposes
func (t *Btree) Debug() {
	fmt.Println("----------------------------------------------------------------------------------------------")
	if t.Empty() {
		fmt.Println("tree is empty")
	} else {
		fmt.Println(t.Len(), "elements")
	}

	t.Ascend(func(n *Node, i int) bool {
		if t.root.Value == n.Value {
			fmt.Print("ROOT ** ")
		}
		n.Debug()
		return true
	})
	fmt.Println("----------------------------------------------------------------------------------------------")
}

// Init initializes the values of the node or clears the node and returns the node pointer
func (n *Node) Init() *Node {
	n.height = 1
	n.left = nil
	n.right = nil
	return n
}

// String returns a string representing the node
func (n *Node) String() string {
	return fmt.Sprint(n.Value)
}

// Debug prints out useful debug information about the tree node for debugging purposes
func (n *Node) Debug() {
	var children string
	if n.left == nil && n.right == nil {
		children = "no children |"
	} else if n.left != nil && n.right != nil {
		children = fmt.Sprint("left child:", n.left.String(), " right child:", n.right.String())
	} else if n.right != nil {
		children = fmt.Sprint("right child:", n.right.String())
	} else {
		children = fmt.Sprint("left child:", n.left.String())
	}

	fmt.Println(n.String(), "|", "height", n.height, "|", "balance", balance(n), "|", children)
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

func (n *Node) iterate(iterator NodeIterator, i *int, cont bool) {
	if n != nil && cont {
		n.left.iterate(iterator, i, cont)
		cont = iterator(n, *i)
		*i++
		n.right.iterate(iterator, i, cont)
	}
}

func (n *Node) rIterate(iterator NodeIterator, i *int, cont bool) {
	if n != nil && cont {
		n.right.iterate(iterator, i, cont)
		cont = iterator(n, *i)
		*i++
		n.left.iterate(iterator, i, cont)
	}
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

	// t := New()

	// t.Insert(ProcVal{1, Proc{1, 1}})
	// t.Insert(ProcVal{2, Proc{2, 2}})
	// t.Insert(ProcVal{3, Proc{3, 3}})

	// fmt.Println(t.root.min().Value.(ProcVal).key)
	// return

	// ts := time.Now()

	f, err := os.Open("/Users/brisk/vscode/contest/lab08/tests/07")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	in := bufio.NewReader(f)

	// in := bufio.NewReader(os.Stdin)

	fw, err := os.Create("/Users/brisk/vscode/contest/lab08/test07_new.txt")
	defer fw.Close()
	out := bufio.NewWriter(fw)

	n := 0 // procs
	m := 0 // tasks
	fmt.Fscan(in, &n, &m)

	freeProcs := New()
	busyProcs := New()

	for i := 0; i < n; i++ {
		proc := Proc{}
		fmt.Fscan(in, &proc.Power)
		// freeProcs.Insert(proc.Power, proc)
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
			// proc := busyProcs.FindMin()
			proc := busyProcs.root.min().Value.(ProcVal).proc
			br := false
			for _, p := range proc {
				if p.BusyUntil <= curTime {
					// busyProcs.Delete(p.BusyUntil, p)
					// busyProcs.Remove(p.BusyUntil)
					busyProcs.Delete(ProcVal{p.BusyUntil, nil})
					p.BusyUntil = 0
					// freeProcs.Insert(p.Power, p)
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

		n := freeProcs.root.min().Value.(ProcVal).proc
		proc := n[0]

		busyUntil := curTime + curTask.WorkTime
		proc.BusyUntil = busyUntil
		busyProcs.Insert(ProcVal{proc.BusyUntil, []Proc{proc}})

		freeProcs.Delete(ProcVal{proc.Power, nil})

		totalPower += curTask.WorkTime * proc.Power
		if totalPower == 151090802988738 {
			x := 0
			_ = x
		}
		fmt.Fprintln(out, totalPower)
	}

	fmt.Println(totalPower)

	// te := time.Since(ts)
	// fmt.Println(te)
}
