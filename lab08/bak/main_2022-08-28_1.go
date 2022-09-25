package main

import (
	"fmt"
	"strconv"
)

type Node struct {
	key     int
	left    *Node
	right   *Node
	parent  *Node
	balance int
}

func (n *Node) Find(key int) *Node {
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

func (t *Node) findParentNode(key int) *Node {
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

func (n *Node) Insert(key int) *Node {
	newNode := &Node{
		key: key,
	}
	lastNode := n.findParentNode(key)
	if lastNode.key == key {
		return lastNode
	}
	newNode.parent = lastNode

	nextNode := lastNode.selectNextNode(key)
	*nextNode = newNode
	newNode.recalculateBalance(1)
	return newNode.balanceTree()
}

func (n *Node) selectNextNode(key int) **Node {
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

func (n *Node) Remove(key int) *Node {
	if key < n.key {
		return n.left.Remove(key)
	}
	if key > n.key {
		return n.right.Remove(key)
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
	text += " " + strconv.Itoa(n.key) + " " + strconv.Itoa(n.balance) + "\n"

	text += n.left.String()
	text += n.right.String()

	return text
}

func (n *Node) balanceTree() *Node {
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

func main() {
	node := &Node{
		key:     1,
		balance: 0,
	}
	n := node.Insert(2)
	n = node.Insert(3)
	n = node.Insert(4)

	// n := node.Insert(10)
	// fmt.Println(n)
	// n = node.Insert(30)
	// fmt.Println(n)
	// n = node.Insert(35)
	// fmt.Println(n)
	// n = n.Insert(40)
	// n = n.Insert(42)
	// n = n.Insert(43)
	// fmt.Println(n)
	// n = n.Remove(42)
	// n = n.Remove(40)

	fmt.Println("---")
	fmt.Println(n)

}
