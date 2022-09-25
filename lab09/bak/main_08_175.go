package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
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
	Value       StringVal
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

func (t *Btree) Insert(value StringVal) *Btree {
	added := false
	t.root = insert(t.root, value, &added)
	if added {
		t.len++
	}
	t.values = nil
	return t
}

func insert(n *Node, value StringVal, added *bool) *Node {
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
		n.Value.words = append(n.Value.words, value.words[0])

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

func (n *Node) get(val StringVal) *Node {
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
	key   string
	words []string
	// words map[int]string
	count int
}

func (i StringVal) Comp(val StringVal) int8 {
	if i.key > val.key {
		return 1
	} else if i.key < val.key {
		return -1
	} else {
		return 0
	}
}

func main() {
	ts := time.Now()

	f, err := os.Open("/Users/brisk/vscode/contest/lab09/tests/16")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	in := bufio.NewReader(f)

	// in := bufio.NewReader(os.Stdin)
	n := 0 // dict size
	fmt.Fscan(in, &n)

	tree := NewTree()
	tEnds := NewTree()

	for i := 0; i < n; i++ {
		word := ""
		fmt.Fscan(in, &word)
		l := len(word)
		// m[0] = word
		tree.Insert(StringVal{word, []string{word}, 0})
		for ml := l; ml > 0; ml-- {
			for i := 0; i <= l-ml; i++ {
				key := word[i : ml+i]
				tree.Insert(StringVal{key, []string{word}, 0})
			}
		}
		tEnds.Insert(StringVal{word, []string{word}, 0})
		for ml := l - 1; ml > 0; ml-- {
			key := word[ml:l]
			tEnds.Insert(StringVal{key, []string{word}, 0})
		}
	}

	s := ""
	q := 0 // req count
	c := 0
	fmt.Fscan(in, &q)
	// sl := make([]string, q+1)
	for i := 0; i < q; i++ {
		if i%100 == 0 {
			fmt.Print(s)
			s = ""
		}
		req := ""
		fmt.Fscan(in, &req)
		br := false

		node := tEnds.root.get(StringVal{req, nil, 0})
		if node != nil {
			word := node.Value.words
			br = false
			for _, v := range word {
				if v == req {
					continue
				}
				// fmt.Println(v)
				// s += fmt.Sprintln(v)
				s += v + "\n"
				// sl[c] = v
				br = true
				break
			}
		}
		if br {
			continue
		}

		l := len(req)

		maxEqual := 0
		foundWord := ""
		foundSuff := ""

		for ml := l; ml > 0; ml-- {
			key := req[ml:l]

			node := tEnds.root.get(StringVal{key, nil, 0})
			if node != nil {
				word := node.Value.words
				br = false
				for _, v := range word {
					if v == req {
						continue
					}
					if maxEqual < len(key) {
						maxEqual = len(key)
						foundWord = v
						foundSuff = key
					}
				}
			}
		}
		if foundWord != "" {
			// fmt.Println(foundWord)
			// s += fmt.Sprintln(foundWord)
			s += foundWord + "\n"
			// sl[c] = foundWord

			continue
		}
		_ = foundSuff

		for ml := l - 1; ml > 0; ml-- {
			for i := 0; i <= l-ml; i++ {
				key := req[i : ml+i]

				node := tEnds.root.get(StringVal{key, nil, 0})
				if node != nil {
					word := node.Value.words
					br = false
					for _, v := range word {
						if v == req {
							continue
						}
						// fmt.Println(v)
						// s += fmt.Sprintln(v)
						s += v + "\n"
						// sl[c] = v

						br = true
						break
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
		if br {
			continue
		}

		node = tree.root.get(StringVal{req, nil, 0})
		if node != nil {
			word := node.Value.words
			br = false
			for _, v := range word {
				if v == req {
					continue
				}
				// fmt.Println(v)
				// s += fmt.Sprintln(v)
				s += v + "\n"
				// sl[c] = v

				br = true
				break
			}
		}
		if br {
			continue
		}

		for ml := l - 1; ml > 0; ml-- {
			for i := 0; i <= l-ml; i++ {
				key := req[i : ml+i]

				node := tree.root.get(StringVal{key, nil, 0})
				if node != nil {
					word := node.Value.words
					br = false
					for _, v := range word {
						if v == req {
							continue
						}
						// fmt.Println(v)
						// s += fmt.Sprintln(v)
						s += v + "\n"
						// sl[q] = v

						br = true
						break
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
		c++
	}
	fmt.Print(s)
	// for _, v := range sl {
	// 	fmt.Println(v)
	// }
	te := time.Since(ts)
	fmt.Println(te)
}
