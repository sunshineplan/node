package node

import (
	"regexp"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestFilter(t *testing.T) {
	if nodes := soup.FindAll(0, B); len(nodes) != 1 {
		t.Errorf("expected b %d; got %d", 1, len(nodes))
	} else if html := nodes[0].Readable(); html != "<b>The Dormouse's story</b>" {
		t.Errorf("expected html %q; got %q", "<b>The Dormouse's story</b>", html)
	}
	if nodes := soup.FindAll(0, Tag(regexp.MustCompile("^b"))); len(nodes) != 2 {
		t.Errorf("expected ^b %d; got %d", 2, len(nodes))
	} else {
		expected := []string{"body", "b"}
		for i, node := range nodes {
			if name := node.Data(); name != expected[i] {
				t.Errorf("expected name #%d %q; got %q", i, expected[i], name)
			}
		}
	}
	if nodes := soup.FindAll(0, Tag(regexp.MustCompile("t"))); len(nodes) != 2 {
		t.Errorf("expected t %d; got %d", 2, len(nodes))
	} else {
		expected := []string{"html", "title"}
		for i, node := range nodes {
			if name := node.Data(); name != expected[i] {
				t.Errorf("expected name #%d %q; got %q", i, expected[i], name)
			}
		}
	}
	if nodes := soup.FindAll(0, Tags("a", "b")); len(nodes) != 4 {
		t.Errorf("expected nodes %d; got %d", 4, len(nodes))
	} else {
		expected := []string{"<b>The Dormouse's story</b>", elsie, lacie, tillie}
		for i, node := range nodes {
			if html := node.Readable(); html != expected[i] {
				t.Errorf("expected html #%d %q; got %q", i, expected[i], html)
			}
		}
	}
	if nodes := soup.FindAll(0, True); len(nodes) != 11 {
		t.Errorf("expected nodes %d; got %d", 11, len(nodes))
	} else {
		expected := []string{"html", "head", "title", "body", "p", "b", "p", "a", "a", "a", "p"}
		for i, node := range nodes {
			if name := node.Data(); name != expected[i] {
				t.Errorf("expected name #%d %q; got %q", i, expected[i], name)
			}
		}
	}
	if nodes := soup.FindAll(0, Tag(func(tag string, node Node) bool {
		return node.HasAttr("class") && !node.HasAttr("id")
	})); len(nodes) != 3 {
		t.Errorf("expected nodes %d; got %d", 3, len(nodes))
	} else {
		expected := []string{
			`<p class="title"><b>The Dormouse's story</b></p>`,
			`<p class="story">Once upon a time there were`,
			`<p class="story">...</p>`,
		}
		for i, node := range nodes {
			if html := node.Readable(); !strings.HasPrefix(html, expected[i]) {
				t.Errorf("expected html #%d %q; got %q", i, expected[i], html)
			}
		}
	}
	if nodes := soup.FindAll(0, nil, Attr("href", func(value string, node Node) bool {
		return !regexp.MustCompile("lacie").MatchString(value)
	})); len(nodes) != 2 {
		t.Errorf("expected nodes %d; got %d", 2, len(nodes))
	} else {
		expected := []string{elsie, tillie}
		for i, node := range nodes {
			if html := node.Readable(); !strings.HasPrefix(html, expected[i]) {
				t.Errorf("expected html #%d %q; got %q", i, expected[i], html)
			}
		}
	}
	if nodes := soup.FindAll(0, Tag(func(tag string, node Node) bool {
		return node.PrevNode() != nil && node.PrevNode().Type() == html.TextNode &&
			node.NextNode() != nil && node.NextNode().Type() == html.TextNode
	})); len(nodes) != 6 {
		t.Errorf("expected nodes %d; got %d", 6, len(nodes))
	} else {
		expected := []string{"body", "p", "a", "a", "a", "p"}
		for i, node := range nodes {
			if name := node.Data(); name != expected[i] {
				t.Errorf("expected name #%d %q; got %q", i, expected[i], name)
			}
		}
	}
}
