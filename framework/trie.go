package framework

import "errors"

type node struct {
	pattern    string
	part       string
	children   []*node
	isWildcard bool
}

func (n *node) insert(pattern string, parts []string, depth int) {
	if len(parts) == depth {
		n.pattern = pattern
		return
	}

	part := parts[depth]
	child := n.matchChild(part)
	if child == nil {
		child = &node{part: part, isWildcard: part[0] == ':'}
		n.children = append(n.children, child)
	}

	n.insert(pattern, parts, depth+1)
}

func (n *node) find(parts []string, depth int) (string, error) {
	if len(parts) == depth {
		return n.pattern, nil
	}

	part := parts[depth]
	child := n.matchChild(part)

	if child == nil {
		return "", errors.New("not found")
	}

	return child.find(parts, depth+1)
}

func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWildcard {
			return child
		}
	}

	return nil
}

func (n *node) matchChildren(part string) []*node {
	return nil
}
