package main

import (
	"bufio"
	"fmt"
	"log"
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
		oldVal := n.Value.(StringVal).word
		newVal := value.(StringVal).word
		oldVal = append(oldVal, newVal...)
		k := n.Value.(StringVal).key
		n.Value = StringVal{k, oldVal}

		// oldVal := n.Value.(ProcVal).proc
		// newVal := value.(ProcVal).proc
		// oldVal = append(oldVal, newVal...)
		// k := n.Value.(ProcVal).key
		// n.Value = ProcVal{k, oldVal}
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
		// l := len(n.Value.(ProcVal).proc)
		// if l > 1 {
		// 	k := n.Value.(ProcVal).key
		// 	proc := n.Value.(ProcVal).proc[:l-1]
		// 	n.Value = ProcVal{k, proc}
		// 	return n
		// }
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

type StringVal struct {
	key  string
	word []string
}

func (i StringVal) Comp(val Val) int8 {
	v := val.(StringVal)
	if i.key > v.key {
		return 1
	} else if i.key < v.key {
		return -1
	} else {
		return 0
	}
}

func main() {

	// word := "decide"

	// l := len(word)
	// // // ml := l - 1
	// // // cl := ml

	// // !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	// for ml := l - 1; ml > 0; ml-- {
	// 	for i := 0; i <= l-ml; i++ {
	// 		key := word[i : ml+i]
	// 		fmt.Println(key)
	// 	}
	// }
	// return

	// // !!!
	// for ml := l - 2; ml > 0; ml-- {
	// 	// for i := 1; i <= l-ml; i++ {
	// 	i := 1
	// 	key := word[ml:l]
	// 	fmt.Println(ml, i, key)
	// 	// }
	// }
	// return

	// in := bufio.NewReader(os.Stdin)

	f, err := os.Open("/Users/brisk/vscode/contest/lab09/tests/01")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	in := bufio.NewReader(f)

	// ts := time.Now()

	n := 0 // dict size
	fmt.Fscan(in, &n)

	trees := make([]*Btree, 10)
	for i := 0; i < 9; i++ {
		trees[i] = NewTree()
	}

	for i := 0; i < n; i++ {
		word := ""
		fmt.Fscan(in, &word)
		// word = "decide"
		l := len(word)
		trees[l-1].Insert(StringVal{word, []string{word}})
		// for ml := l - 2; ml > 0; ml-- {
		// 	key := word[ml:l]
		// 	fmt.Println(ml, key)
		// 	tn := len(key) - 1
		// 	trees[tn].Insert(StringVal{key, []string{word}})
		// }
		for ml := l - 1; ml > 0; ml-- {
			for i := 0; i <= l-ml; i++ {
				key := word[i : ml+i]
				// fmt.Println(key)
				tn := len(key) - 1
				trees[tn].Insert(StringVal{key, []string{word}})
			}
		}
		// break
	}
	// fmt.Println("\n\n\n")

	q := 0 // req count
	fmt.Fscan(in, &q)
	for i := 0; i < q; i++ {
		req := ""
		fmt.Fscan(in, &req)

		// req = "code"

		br := false
		l := len(req)

		for t := 0; t < 9; t++ {

			if trees[t] != nil {
				if trees[t].len == 0 {
					continue
				}
			}

			// fmt.Println(t, trees[t].len)
			node := trees[t].root.get(StringVal{req, nil}) // l-1
			if node != nil {
				word := node.Value.(StringVal).word
				br = false
				for _, v := range word {
					if v == req {
						continue
					}
					// fmt.Println(req, v, "!!!!")
					fmt.Println(v)
					br = true
					break
				}
				if br {
					break
				}
			}
		}
		if br {
			continue
		}

		for ml := l - 1; ml > 0; ml-- {
			for i := 0; i <= l-ml; i++ {
				key := req[i : ml+i]
				// fmt.Println(key)
				// tn := len(key) - 1

				for t := 0; t < 9; t++ {

					if trees[t].root == nil {
						continue
					}
					node := trees[t].root.get(StringVal{key, nil})
					if node != nil {
						word := node.Value.(StringVal).word
						minLen := 10 //len(req)
						foundWord := ""
						for _, v := range word {
							if v == req {
								continue
							}
							if len(v) <= minLen {
								minLen = len(v)
								foundWord = v
							}
						}
						// for _, v := range word {
						// 	if v == req {
						// 		continue
						// 	}
						// 	fmt.Println(req, v, "???")
						// 	br = true
						// 	break
						// }
						if foundWord != "" {
							// fmt.Println(req, foundWord, "++++++")
							fmt.Println(foundWord)
							br = true
							break
						}
					}
				}
				if br {
					break
				}
			}
			if br {
				break
			}
		}
		// if !br {
		// 	fmt.Println(req, "not found")
		// }
		// break
	}

	// te := time.Since(ts)
	// fmt.Println(te)
}
