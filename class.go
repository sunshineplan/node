package node

import "strings"

var (
	_ Filter = class[string]{}
	_ Filter = classStrict("")
)

// class is a struct that represents a class filter.
type class[T Value] struct {
	class T
}

// Class returns a new class filter with the specified value.
// This filter is an attribute filter.
func Class[T Value](v T) Filter {
	return class[T]{v}
}

// IsAttribute returns true, indicating that the filter represents an attribute filter.
func (class[T]) IsAttribute() bool {
	return true
}

// IsMatch returns true if the class filter matches the given node.
func (cls class[T]) IsMatch(node Node) bool {
	switch v := (any(cls.class)).(type) {
	case string:
		nodeClass, ok := getAttribute(node, "class")
		if !ok {
			return false
		}
		classA, classB := strings.Fields(nodeClass), strings.Fields(v)
		// Iterate over each class name in the given string and
		// check if it is present in the node's class attribute.
		for _, i := range classB {
			var b bool
			for _, ii := range classA {
				if i == ii {
					b = true
					break
				}
			}
			// If any class name is not present in the node's class attribute, return false.
			if !b {
				return false
			}
		}
		// If all class names are present in the node's class attribute, return true.
		return true
	case []string:
		// If T is a slice of strings, check if the node's class attribute matches any of the given strings.
		if _, ok := getAttribute(node, "class"); !ok {
			return false
		}
		for _, v := range v {
			if (class[string]{v}).IsMatch(node) {
				return true
			}
		}
		return false
	default:
		// If T is not a string or a slice of strings, check if the node's "class" attribute
		// matches the value of T using the attribute filter.
		return attribute[T]{"class", cls.class}.IsMatch(node)
	}
}

// classStrict is a struct that represents a strict class filter.
type classStrict string

// ClassStrict returns a new strict class filter with the specified string.
// This filter is an attribute filter.
func ClassStrict(cls string) Filter {
	return classStrict(cls)
}

// IsAttribute returns true, indicating that the filter represents an attribute filter.
func (classStrict) IsAttribute() bool {
	return true
}

// IsMatch returns true if the strict class filter matches the given node.
func (classStrict classStrict) IsMatch(node Node) bool {
	nodeClass, ok := getAttribute(node, "class")
	if !ok {
		return false
	}
	classA, classB := strings.Fields(nodeClass), strings.Fields(string(classStrict))
	// If the two class names are not exactly equal, return false.
	return strings.Join(classA, "|||") == strings.Join(classB, "|||")
}
