package server

import (
	"fmt"
	"net/http"
	"strings"
)

type NodeHandler interface {
	Find(method string, path []string) (*Node, []string)
	Register(method string, path []string, up SingUp)
	GetPath() string
}

type Node struct {
	Method string
	Path   string
	Up     SingUp
	Nodes  []*Node
}

func (n *Node) Find(method string, path []string) (*Node, []string) {
	length := len(path)
	if length < 1 {
		if n.Method != method {
			return nil, []string{}
		}
		return n, []string{}
	}
	for _, n := range n.Nodes {
		if n.GetPath() == path[0] {
			s, x := n.Find(method, path[1:])
			if s != nil {
				return s, x
			}
		}
	}
	return n, path
}

func (n *Node) Register(method string, path []string, up SingUp) {
	node := Node{
		Method: method,
		Path:   path[0],
		Up:     nil,
		Nodes:  []*Node{},
	}
	if len(path[1:]) > 0 {
		node.Register(method, path[1:], up)
	} else {
		node.Up = up
	}

	n.Nodes = append(n.Nodes, &node)
}

func (n *Node) GetAll() {
	fmt.Println(n.Nodes)
}

func (n Node) GetPath() string {
	return n.Path
}

func (n Node) GetSingUp() SingUp {
	return n.Up
}

type HandlerNode struct {
	Node *Node
}

var _ Handler = &HandlerNode{}
var _ NodeHandler = &Node{}

func (h HandlerNode) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.Trim(r.URL.Path, "/")
	split := strings.Split(path, "/")
	method := r.Method
	if m := h.Find(method, split); m != nil {
		context := NewContext(w, r)
		m(context)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("未匹配到方法"))
	}
}

func (h *HandlerNode) GetAll() {
	fmt.Println(h.Node.Nodes[0].Nodes[0].Path)
}

func (h *HandlerNode) Route(method, path string, fn SingUp) {
	path = strings.Trim(path, "/")
	split := strings.Split(path, "/")
	find, i := h.Node.Find(method, split)
	if len(i) > 0 {
		find.Register(method, i, fn)
	}
}

func (h *HandlerNode) Find(method string, paths []string) SingUp {
	find, i := h.Node.Find(method, paths)
	if len(i) > 0 {
		return nil
	}
	if find == nil {
		return nil
	}
	return find.Up
}

func NewNodeHandler() Handler {
	return &HandlerNode{
		Node: &Node{},
	}
}
