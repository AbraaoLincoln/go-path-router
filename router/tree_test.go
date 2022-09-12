package router

import (
	"fmt"
	"net/http"
	"testing"
)

func TestIsPathVariable(t *testing.T) {
	pathVariable1 := ":id"
	pathVariable2 := "id"

	if !isPathVariable(pathVariable1) {
		t.Errorf("Expecting %v to be a path variable", pathVariable1)
	}

	if isPathVariable(pathVariable2) {
		t.Errorf("Expecting %v not to be a path variable", pathVariable2)
	}
}

func TestInsert(t *testing.T) {
	path := "/users/123/posts"
	handle := func(w http.ResponseWriter, r *http.Request) { fmt.Println("handler1") }
	root := NewNode()
	tree := Tree{Root: &root}

	tree.Insert(path, handle)

	if root.children[""] == nil {
		t.Error("Expecting not to be nil")
	}

	if root.children[""].children["users"] == nil {
		t.Error("Expecting /users not to be nil")
	}

	if root.children[""].children["users"].children["123"] == nil {
		t.Error("Expecting /users/123 not to be nil")
	}

	if root.children[""].children["users"].children["123"].children["posts"] == nil {
		t.Error("Expecting /users/123/posts not to be nil")
	}

	path = "/users/profile"
	handle = func(w http.ResponseWriter, r *http.Request) { fmt.Println("handler2") }

	tree.Insert(path, handle)

	if root.children[""] == nil {
		t.Error("Expecting not to be nil")
	}

	if root.children[""].children["users"] == nil {
		t.Error("Expecting /users not to be nil")
	}

	if root.children[""].children["users"].children["profile"] == nil {
		t.Error("Expecting /users/profile not to be nil")
	}
}

func TestInsertWithPathVariable(t *testing.T) {
	path := "/users/:id/posts"
	handle := func(w http.ResponseWriter, r *http.Request) { fmt.Println("handler1") }
	root := NewNode()
	tree := Tree{Root: &root}

	tree.Insert(path, handle)

	path = "/users/new/posts"
	handle = func(w http.ResponseWriter, r *http.Request) { fmt.Println("handler2") }

	tree.Insert(path, handle)

	if root.children[""] == nil {
		t.Error("Expecting not to be nil")
	}

	if root.children[""].children["users"] == nil {
		t.Error("Expecting /users not to be nil")
	}

	if root.children[""].children["users"].pathVariable == nil {
		t.Error("Expecting /users/:id path variable not to be nil")
	}

	if root.children[""].children["users"].children["posts"] == nil {
		t.Error("Expecting /users/123/posts not to be nil")
	}
	//==================================================================
	if root.children[""] == nil {
		t.Error("Expecting not to be nil")
	}

	if root.children[""].children["users"] == nil {
		t.Error("Expecting /users not to be nil")
	}

	if root.children[""].children["users"].children["new"] == nil {
		t.Error("Expecting /users/new/posts not to be nil")
	}

	if root.children[""].children["users"].children["new"].children["posts"] == nil {
		t.Error("Expecting /users/new/posts not to be nil")
	}
}
