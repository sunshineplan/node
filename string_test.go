package node

import (
	"regexp"
	"testing"
)

func TestString(t *testing.T) {
	if nodes := soup.FindAllString(0, String("Elsie")); len(nodes) != 1 {
		t.Errorf("expected nodes %d; got %d", 1, len(nodes))
	} else if text := nodes[0].String(); text != "Elsie" {
		t.Errorf("expected string %q; got %q", "Elsie", text)
	}
	if nodes := soup.FindAllString(0, String([]string{"Tillie", "Elsie", "Lacie"})); len(nodes) != 3 {
		t.Errorf("expected nodes %d; got %d", 3, len(nodes))
	} else {
		expected := []string{"Elsie", "Lacie", "Tillie"}
		for i, node := range nodes {
			if text := node.String(); text != expected[i] {
				t.Errorf("expected string #%d %q; got %q", i, expected[i], text)
			}
		}
	}
	if nodes := soup.FindAllString(0, String(regexp.MustCompile("Dormouse"))); len(nodes) != 2 {
		t.Errorf("expected nodes %d; got %d", 2, len(nodes))
	} else {
		for i, node := range nodes {
			if text := node.String(); text != "The Dormouse's story" {
				t.Errorf("expected string #%d %q; got %q", i, "The Dormouse's story", text)
			}
		}
	}
	if nodes := soup.FindAllString(0, String(func(s string, node Node) bool {
		if parent := node.Parent(); parent != nil && parent.String() != nil {
			return node.Raw() == parent.String().Raw()
		}
		return false
	})); len(nodes) != 6 {
		t.Errorf("expected nodes %d; got %d", 6, len(nodes))
	} else {
		expected := []string{"The Dormouse's story", "The Dormouse's story", "Elsie", "Lacie", "Tillie", "..."}
		for i, node := range nodes {
			if text := node.String(); text != expected[i] {
				t.Errorf("expected string #%d %q; got %q", i, expected[i], text)
			}
		}
	}
	if nodes := soup.FindAllString(0, Text("Elsie")); len(nodes) != 1 {
		t.Errorf("expected nodes %d; got %d", 1, len(nodes))
	} else if text := nodes[0].String(); text != "Elsie" {
		t.Errorf("expected string %q; got %q", "Elsie", text)
	}
}
