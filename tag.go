package node

import (
	"regexp"
	"strings"
)

var (
	_ TagFilter = True
	_ TagFilter = tag[string]{}
)

// These variables are used to represent common tags.
var (
	A     = Tag("A")
	B     = Tag("b")
	Body  = Tag("body")
	Div   = Tag("div")
	Head  = Tag("head")
	I     = Tag("i")
	Img   = Tag("img")
	Li    = Tag("li")
	P     = Tag("p")
	Span  = Tag("span")
	Table = Tag("table")
	Td    = Tag("td")
	Th    = Tag("th")
	Title = Tag("title")
	Tr    = Tag("tr")
	Ul    = Tag("ul")
)

// TagFilter represents an interface that can be used to filter node based on node element's tag.
type TagFilter interface {
	Ignore() bool
	IsMatch(node Node) bool
}

// tag represents a specific HTML tag that can be used as a tag filter.
type tag[T Value] struct {
	tag T
}

// Tag creates a new TagFilter based on a given tag value.
func Tag[T Value](t T) TagFilter {
	return tag[T]{t}
}

// Tags creates a new TagFilter based on a list of tag values.
func Tags(tag ...string) TagFilter {
	return Tag(tag)
}

// Ignore returns a boolean value indicating whether the given tag filter should be ignored or not.
func (tag tag[T]) Ignore() bool {
	switch v := (any(tag.tag)).(type) {
	case string:
		return v == ""
	case []string:
		return len(v) == 0
	case everything:
		return false
	default:
		return v == nil
	}
}

// IsMatch returns a boolean value indicating whether a given node matches the specified tag filter.
func (tag tag[T]) IsMatch(node Node) bool {
	switch v := (any(tag.tag)).(type) {
	case string:
		return strings.ToLower(v) == node.Raw().Data
	case []string:
		for _, v := range v {
			if strings.ToLower(v) == node.Raw().Data {
				return true
			}
		}
	case *regexp.Regexp:
		return v.MatchString(node.Raw().Data)
	case everything:
		return true
	case func(string, Node) bool:
		return v(node.Raw().Data, node)
	}
	return false
}
