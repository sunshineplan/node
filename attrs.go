package node

var _ Attributes = attributes{}

// Attributes is an interface that describes a node's attributes with
// methods for getting and iterating over key-value pairs.
type Attributes interface {
	// Range calls the provided function for each key-value pair in the Attributes
	// iteration stops if the function returns false for any pair.
	Range(func(key, value string) bool)

	// Get returns the value associated with the specified key and
	// a boolean indicating whether the key exists in the Attributes.
	Get(key string) (value string, exists bool)
}

// attributes is a private struct that implements the Attributes interface.
type attributes map[string]string

// Range calls the provided function for each key-value pair in the attributes
// iteration stops if the function returns false for any pair.
func (attrs attributes) Range(f func(key, value string) bool) {
	for k, v := range attrs {
		if !f(k, v) {
			break
		}
	}
}

// Get returns the value associated with the specified key and
// a boolean indicating whether the key exists in the attributes.
func (attrs attributes) Get(key string) (value string, exists bool) {
	v, ok := attrs[key]
	return v, ok
}
