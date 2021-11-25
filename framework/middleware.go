package framework

type middleware struct {
	part     string
	children []*middleware
	handlers []handlerFunc
}

func newMiddleware() *middleware {
	return &middleware{}
}

func (m *middleware) insert(parts []string, depth int, middlewares []handlerFunc) {
	if len(parts) == depth {
		m.handlers = append(m.handlers, middlewares...)
		return
	}

	part := parts[depth]
	child := m.matchChild(part)
	if child == nil {
		child = &middleware{part: part}
		m.children = append(m.children, child)
	}

	child.insert(parts, depth+1, middlewares)
}

func (m *middleware) find(parts []string, depth int) []handlerFunc {
	handlers := m.handlers
	if len(parts) == depth {
		return handlers
	}

	part := parts[depth]
	child := m.matchChild(part)
	if child == nil {
		return handlers
	}

	return append(handlers, child.find(parts, depth+1)...)
}

func (m *middleware) matchChild(part string) *middleware {
	for _, child := range m.children {
		if child.part == part {
			return child
		}
	}

	return nil
}
