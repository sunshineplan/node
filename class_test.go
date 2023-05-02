package node

import (
	"regexp"
	"testing"
)

func TestClass(t *testing.T) {
	if nodes := soup.FindAll(0, A, Class("sister")); len(nodes) != 3 {
		t.Errorf("expected nodes %d; got %d", 3, len(nodes))
	} else {
		expected := []string{
			`<a href="http://example.com/elsie" class="sister" id="link1">Elsie</a>`,
			`<a href="http://example.com/lacie" class="sister" id="link2">Lacie</a>`,
			`<a href="http://example.com/tillie" class="sister" id="link3">Tillie</a>`,
		}
		for i, node := range nodes {
			if html := node.Readable(); html != expected[i] {
				t.Errorf("expected html #%d %q; got %q", i, expected[i], html)
			}
		}
	}
	if nodes := soup.FindAll(0, nil, Class(regexp.MustCompile("itl"))); len(nodes) != 1 {
		t.Errorf("expected nodes %d; got %d", 1, len(nodes))
	} else {
		if html := nodes[0].Readable(); html != `<p class="title"><b>The Dormouse's story</b></p>` {
			t.Errorf("expected html %q; got %q", `<p class="title"><b>The Dormouse's story</b></p>`, html)
		}
	}
	if nodes := soup.FindAll(0, A, Class(func(class string, node Node) bool {
		return node.HasAttr("class") && len(class) == 6
	})); len(nodes) != 3 {
		t.Errorf("expected nodes %d; got %d", 3, len(nodes))
	} else {
		expected := []string{
			`<a href="http://example.com/elsie" class="sister" id="link1">Elsie</a>`,
			`<a href="http://example.com/lacie" class="sister" id="link2">Lacie</a>`,
			`<a href="http://example.com/tillie" class="sister" id="link3">Tillie</a>`,
		}
		for i, node := range nodes {
			if html := node.Readable(); html != expected[i] {
				t.Errorf("expected html #%d %q; got %q", i, expected[i], html)
			}
		}
	}
	if nodes := soup.FindAll(0, A, Attr("class", "sister")); len(nodes) != 3 {
		t.Errorf("expected nodes %d; got %d", 3, len(nodes))
	} else {
		expected := []string{
			`<a href="http://example.com/elsie" class="sister" id="link1">Elsie</a>`,
			`<a href="http://example.com/lacie" class="sister" id="link2">Lacie</a>`,
			`<a href="http://example.com/tillie" class="sister" id="link3">Tillie</a>`,
		}
		for i, node := range nodes {
			if html := node.Readable(); html != expected[i] {
				t.Errorf("expected html #%d %q; got %q", i, expected[i], html)
			}
		}
	}
}
