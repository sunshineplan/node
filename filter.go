package node

import (
	"regexp"
	"strings"
)

var _ Filter = attribute[string]{}

// Filter is an interface that describes a filter that can be used to select nodes.
type Filter interface {
	// IsAttribute returns true if the filter represents an attribute filter.
	IsAttribute() bool

	// IsMatch returns true if the filter matches the given node.
	IsMatch(node Node) bool
}

// Value is an interface that represents a value that can be used as a filter.
type Value interface {
	// Value can be one of the following types:
	// - string: a simple string value
	// - []string: a slice of strings
	// - *regexp.Regexp: a regular expression
	// - everything: a special value that matches any node
	// - func(string, Node) bool: a function that takes a string and a node and returns true or false
	string | []string | *regexp.Regexp | everything | func(string, Node) bool
}

// True is a special value that matches any node.
var True everything

type everything struct{}

func (everything) Ignore() bool      { return false }
func (everything) IsAttribute() bool { return true }
func (everything) IsString() bool    { return true }
func (everything) IsMatch(Node) bool { return true }

// attribute is a struct that represents an attribute filter.
type attribute[T Value] struct {
	name  string
	value T
}

// Attr returns a new attribute filter with the specified name and value.
func Attr[T Value](name string, value T) Filter {
	return attribute[T]{strings.ToLower(name), value}
}

// Id returns a new attribute filter for the "id" attribute with the specified value.
func Id[T Value](id T) Filter {
	return Attr("id", id)
}

// IsAttribute returns true, indicating that the filter represents an attribute filter.
func (attribute[T]) IsAttribute() bool {
	return true
}

// IsMatch returns true if the attribute filter matches the given node
func (attribute attribute[T]) IsMatch(node Node) bool {
	switch v := (any(attribute.value)).(type) {
	case string:
		// If the attribute name is "class", use the class filter to match the node's class attribute.
		if attribute.name == "class" {
			return class[T]{attribute.value}.IsMatch(node)
		} else if value, ok := getAttribute(node, attribute.name); !ok {
			return false
		} else {
			return value == v
		}
	case []string:
		if attribute.name == "class" {
			return class[T]{attribute.value}.IsMatch(node)
		} else if value, ok := getAttribute(node, attribute.name); !ok {
			return false
		} else {
			for _, v := range v {
				if value == v {
					return true
				}
			}
		}
	case *regexp.Regexp:
		if value, ok := getAttribute(node, attribute.name); !ok {
			return false
		} else {
			return v.MatchString(value)
		}
	case everything:
		_, ok := getAttribute(node, attribute.name)
		return ok
	case func(string, Node) bool:
		if value, ok := getAttribute(node, attribute.name); !ok {
			return false
		} else {
			return v(value, node)
		}
	}
	return false
}

// getAttribute returns the value of the specified attribute of the given node.
// It returns the attribute value and true if the attribute exists, empty string and false otherwise.
func getAttribute(node HtmlNode, name string) (string, bool) {
	if attr := node.Attrs(); attr == nil {
		return "", false
	} else {
		attr, ok := attr.Get(name)
		return attr, ok
	}
}

// isAttributeFilter checks if a list of filters only contains attribute filters.
// It returns true if all filters are attribute filters, false otherwise.
func isAttributeFilter(filters []Filter) bool {
	for _, i := range filters {
		if !i.IsAttribute() {
			return false
		}
	}
	return true
}
