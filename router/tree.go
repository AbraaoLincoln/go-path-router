package router

import (
	"net/http"
	"strings"
)

type Tree struct {
	Root *Node
}

func (t *Tree) Insert(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	splitPath := strings.Split(path, "/")
	currentNode := t.Root

	for index, value := range splitPath {
		if isPathVariable(value) {
			updateNodePathVariable(value, handler, currentNode)
			continue
		}

		if isEndpoint(index, len(splitPath)) && hasTheSameValueAsTheCurrentNode(value, currentNode.value) {
			updateNodeHandler(handler, currentNode)
			break
		}

		if currentNode.hasChild(value) {
			currentNode = currentNode.getChild(value)
			continue
		}

		newNode := NewNodeWith(value, handler)
		currentNode.addChild(value, newNode)
		currentNode = newNode
	}
}

func isPathVariable(value string) bool {
	return strings.HasPrefix(value, ":")
}

func updateNodePathVariable(value string, handler func(w http.ResponseWriter, r *http.Request), node *Node) {

	if node.pathVariable != nil {
		node.pathVariable.handle = handler
		return
	}

	newNode := NewNodeWith(value, handler)
	node.pathVariable = newNode
}

func isEndpoint(currentIndex int, qtyOfPoints int) bool {
	return currentIndex == (qtyOfPoints - 1)
}

func updateNodeHandler(handler func(w http.ResponseWriter, r *http.Request), node *Node) {
	node.handle = handler
}

func hasTheSameValueAsTheCurrentNode(value string, currentNodeValue string) bool {
	return value == currentNodeValue
}
