package aho_corasick

import "container/list"

type Node struct {
	symbol    string
	n         int
	f         *Node
	accepting bool
	children  map[string]*Node
}

func (n *Node) getFailureFunction() []int {
	if len(n.children) == 0 {
		return nil
	}

	nodes := n.getAllNodes()

	out := make([]int, len(nodes)-1)
	for _, node := range nodes {
		if node.f != nil {
			out[node.n-1] = node.f.n
		}
	}

	return out
}

func (n *Node) getAllNodes() []*Node {
	var children []*Node
	for _, child := range n.children {
		children = append(children, child.getAllNodes()...)
	}

	return append([]*Node{n}, children...)
}

func getFailureFunction(patterns []string) []int {
	root := getTree(patterns)

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		element := queue.Front()
		queue.Remove(element)
		node := element.Value.(*Node)

		for _, child := range node.children {
			if node == root {
				child.f = root
				queue.PushBack(child)
				continue
			}

			f := node.f
			for {
				next, ok := f.children[child.symbol]
				if ok {
					f = next
					break
				}
				if f.f == nil {
					break
				}
				f = f.f
			}

			child.f = f
			queue.PushBack(child)
		}
	}

	return root.getFailureFunction()
}

func getTree(patterns []string) *Node {
	n := 0
	f0 := &Node{n: n}

	for _, pattern := range patterns {
		strLen := len(pattern)
		branch := f0

		for s := 0; s < strLen; s++ {
			symbol := string(pattern[s])

			next, ok := branch.children[symbol]
			if !ok {
				n++
				next = &Node{
					symbol: symbol,
					n:      n,
				}
				if branch.children == nil {
					branch.children = make(map[string]*Node)
				}
				branch.children[symbol] = next
			}
			branch = next

			if s == strLen-1 {
				branch.accepting = true
			}
		}
	}

	return f0
}
