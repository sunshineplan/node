package node

import "regexp"

var (
	_ Filter       = text[string]{}
	_ StringFilter = text[string]{}
	_ StringFilter = True
)

// StringFilter interface extends the Filter interface and defines
// a method for checking if the filter represents an string filter.
type StringFilter interface {
	Filter
	IsString() bool
}

type text[T Value] struct {
	string T
}

// String returns a StringFilter with the specified value.
func String[T Value](t T) StringFilter {
	return text[T]{t}
}

// Text is an alias of String.
func Text[T Value](t T) StringFilter {
	return String(t)
}

// IsAttribute returns false, indicating that the filter does not represent an attribute filter.
func (text[T]) IsAttribute() bool {
	return false
}

// IsString returns true, indicating that the filter represents an string filter.
func (text[T]) IsString() bool {
	return true
}

// IsMatch returns true if the string filter matches the given node.
func (text text[T]) IsMatch(node Node) bool {
	textNode := node.String()
	if textNode == nil {
		return false
	}
	switch v := (any(text.string)).(type) {
	case string:
		return textNode.String() == v
	case []string:
		for _, v := range v {
			if textNode.String() == v {
				return true
			}
		}
	case *regexp.Regexp:
		return v.MatchString(textNode.String())
	case everything:
		return textNode.String() != ""
	case func(string, Node) bool:
		return v(textNode.String(), node)
	}
	return false
}
