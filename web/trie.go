package web

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

	child.insert(pattern, parts, depth+1)
}

func (n *node) find(parts []string, depth int) string {
	if len(parts) == depth {
		return n.pattern
	}

	part := parts[depth]
	child := n.matchChild(part)
	if child == nil {
		return ""
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
