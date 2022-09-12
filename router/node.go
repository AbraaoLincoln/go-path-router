package router

import "net/http"

type Node struct {
	value        string
	children     map[string]*Node
	handle       func(w http.ResponseWriter, r *http.Request)
	pathVariable *Node
}

func NewNode() Node {
	return Node{children: make(map[string]*Node)}
}

func NewNodeWith(value string, handler func(w http.ResponseWriter, r *http.Request)) *Node {
	return &Node{
		value:    value,
		children: make(map[string]*Node),
		handle:   handler,
	}
}

func (n *Node) addChild(name string, newChild *Node) {
	n.children[name] = newChild
}

func (n *Node) hasChild(childrenName string) bool {
	return n.children[childrenName] != nil
}

func (n *Node) getChild(childrenName string) *Node {
	return n.children[childrenName]
}

func (n *Node) hasPathVariable() bool {
	return n.pathVariable != nil
}
