package node

import (
	"context"

	"golang.org/x/net/html"
)

// Finder represents a set of methods for finding nodes.
type Finder interface {
	// Find searches for a single node in the parse tree based on the specified find method and filters.
	Find(FindMethod, TagFilter, ...Filter) Node

	// FindN searches for up to n nodes in the parse tree based on the specified find method and filters.
	FindN(FindMethod, int, TagFilter, ...Filter) []Node

	// FindAll searches for all nodes in the parse tree based on the specified find method and filters.
	FindAll(FindMethod, TagFilter, ...Filter) []Node

	// FindString searches for a single text node in the parse tree based on the specified find method and filters.
	FindString(FindMethod, StringFilter) TextNode

	// FindStringN searches for up to n text nodes in the parse tree based on the specified find method and filters.
	FindStringN(FindMethod, int, StringFilter) []TextNode

	// FindAllString searches for all text nodes in the parse tree based on the specified find method and filters.
	FindAllString(FindMethod, StringFilter) []TextNode
}

// FindMethod represents the method used to search for nodes in the parse tree.
type FindMethod int

const (
	// Descendant represents a search for nodes that are descendants of the current node.
	Descendant FindMethod = iota

	// NoRecursive represents a search for nodes that are direct children of the current node.
	NoRecursive

	// Parent represents a search for the parent node of the current node.
	Parent

	// PrevSibling represents a search for the previous sibling node of the current node.
	PrevSibling

	// NextSibling represents a search for the next sibling node of the current node.
	NextSibling

	// Previous represents a search for the previous node in the parse tree.
	Previous

	// Next represents a search for the next node in the parse tree.
	Next
)

func findTextNode(tag TagFilter, filters []Filter, strict bool) bool {
	if strict || ((tag == nil || tag.Ignore()) && !isAttributeFilter(filters)) {
		return true
	}
	return false
}

func isMatchType(node Node, findTextNode bool) bool {
	if t := node.Type(); findTextNode {
		return t == html.TextNode
	} else {
		return t == html.ElementNode
	}
}

func (n *htmlNode) find(method FindMethod, text bool, limit int, tag TagFilter, filters ...Filter) (nodes []Node) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var f func(Node, ...Filter)
	f = func(node Node, filters ...Filter) {
		if ctx.Err() != nil || node == nil {
			return
		}
		if raw := node.Raw(); n.Raw() != raw &&
			isMatchType(node, findTextNode(tag, filters, text)) &&
			(tag == nil || tag.IsMatch(node)) {
			ok := true
			for _, i := range filters {
				if !i.IsMatch(node) {
					ok = false
					break
				}
			}
			if ok {
				nodes = append(nodes, node)
				if len(nodes) == limit {
					cancel()
				}
			}
		}
		switch method {
		case Descendant:
			for node := node.FirstChild(); node != nil; node = node.NextSibling() {
				f(node, filters...)
			}
		case NoRecursive:
			if raw := node.Raw(); n.Raw() != raw {
				f(node.FirstChild(), filters...)
			} else {
				f(node.NextSibling(), filters...)
			}
		case Parent:
			if node := node.Parent(); node != nil {
				f(node, filters...)
			}
		case PrevSibling:
			if node := node.PrevSibling(); node != nil {
				f(node, filters...)
			}
		case NextSibling:
			if node := node.NextSibling(); node != nil {
				f(node, filters...)
			}
		case Previous:
			if node := node.PrevNode(); node != nil {
				f(node, filters...)
			}
		case Next:
			if node := node.NextNode(); node != nil {
				f(node, filters...)
			}
		}
	}
	f(n.ToNode(), filters...)
	return
}

func (n *htmlNode) findOnce(method FindMethod, text bool, tag TagFilter, filters ...Filter) Node {
	nodes := n.find(method, text, 1, tag, filters...)
	if len(nodes) == 0 {
		return nil
	}
	return nodes[0]
}

func (n *htmlNode) findN(method FindMethod, text bool, limit int, tag TagFilter, filters ...Filter) []Node {
	if limit <= 0 {
		return nil
	}
	return n.find(method, text, limit, tag, filters...)
}

func (n *htmlNode) Find(method FindMethod, tag TagFilter, filters ...Filter) Node {
	return n.findOnce(method, false, tag, filters...)
}

func (n *htmlNode) FindN(method FindMethod, limit int, tag TagFilter, filters ...Filter) []Node {
	return n.findN(method, false, limit, tag, filters...)
}

func (n *htmlNode) FindAll(method FindMethod, tag TagFilter, filters ...Filter) []Node {
	return n.find(method, false, 0, tag, filters...)
}

func (n *htmlNode) FindString(method FindMethod, filter StringFilter) TextNode {
	if node := n.findOnce(method, true, nil, filter); node != nil {
		return node.ToTextNode()
	}
	return nil
}

func (n *htmlNode) FindStringN(method FindMethod, limit int, filter StringFilter) (res []TextNode) {
	for _, i := range n.findN(method, true, limit, nil, filter) {
		res = append(res, i.ToTextNode())
	}
	return
}

func (n *htmlNode) FindAllString(method FindMethod, filter StringFilter) (res []TextNode) {
	for _, i := range n.find(method, true, 0, nil, filter) {
		res = append(res, i.ToTextNode())
	}
	return
}
