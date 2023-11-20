package node

import (
	"reflect"
	"regexp"
	"strings"
	"testing"
)

func TestFindAll(t *testing.T) {
	if nodes := soup.FindAll(0, Title); len(nodes) != 1 {
		t.Errorf("expected nodes %d; got %d", 1, len(nodes))
	} else if html := nodes[0].Readable(); html != "<title>The Dormouse's story</title>" {
		t.Errorf("expected html %q; got %q", "<title>The Dormouse's story</title>", html)
	}
	if nodes := soup.FindAll(0, P, Class("title")); len(nodes) != 1 {
		t.Errorf("expected nodes %d; got %d", 1, len(nodes))
	} else if html := nodes[0].Readable(); html != `<p class="title"><b>The Dormouse's story</b></p>` {
		t.Errorf("expected html %q; got %q", `<p class="title"><b>The Dormouse's story</b></p>`, html)
	}
	if nodes := soup.FindAll(0, A); len(nodes) != 3 {
		t.Errorf("expected nodes %d; got %d", 3, len(nodes))
	} else {
		expected := []string{elsie, lacie, tillie}
		for i, node := range nodes {
			if html := node.Readable(); html != expected[i] {
				t.Errorf("expected html #%d %q; got %q", i, expected[i], html)
			}
		}
	}
	if nodes := soup.FindAll(0, nil, Id("link2")); len(nodes) != 1 {
		t.Errorf("expected nodes %d; got %d", 1, len(nodes))
	} else if html := nodes[0].Readable(); html != lacie {
		t.Errorf("expected html %q; got %q", lacie, html)
	}
	if nodes := soup.FindAll(0, nil, String(regexp.MustCompile("sisters"))); len(nodes) != 1 {
		t.Errorf("expected nodes %d; got %d", 1, len(nodes))
	} else if text := nodes[0].Readable(); text != "Once upon a time there were three little sisters; and their names were\n" {
		t.Errorf("expected text %q; got %q", "Once upon a time there were three little sisters; and their names were\n", text)
	}
	if nodes := soup.FindAll(0, nil, Attr("href", regexp.MustCompile("elsie"))); len(nodes) != 1 {
		t.Errorf("expected nodes %d; got %d", 1, len(nodes))
	} else if html := nodes[0].Readable(); html != elsie {
		t.Errorf("expected html %q; got %q", elsie, html)
	}
	if nodes := soup.FindAll(0, nil, Id(True)); len(nodes) != 3 {
		t.Errorf("expected nodes %d; got %d", 3, len(nodes))
	} else {
		expected := []string{elsie, lacie, tillie}
		for i, node := range nodes {
			if html := node.Readable(); html != expected[i] {
				t.Errorf("expected html #%d %q; got %q", i, expected[i], html)
			}
		}
	}
	if nodes := soup.FindAll(0, nil, Attr("href", regexp.MustCompile("elsie")), Id("link1")); len(nodes) != 1 {
		t.Errorf("expected nodes %d; got %d", 1, len(nodes))
	} else if html := nodes[0].Readable(); html != elsie {
		t.Errorf("expected html %q; got %q", elsie, html)
	}
}

func TestFindN(t *testing.T) {
	if nodes := soup.FindN(0, 2, A); len(nodes) != 2 {
		t.Errorf("expected nodes %d; got %d", 2, len(nodes))
	} else {
		expected := []string{elsie, lacie}
		for i, node := range nodes {
			if html := node.Readable(); html != expected[i] {
				t.Errorf("expected html #%d %q; got %q", i, expected[i], html)
			}
		}
	}
}

func TestFind(t *testing.T) {
	if nodes := soup.FindN(0, 1, Title); len(nodes) != 1 {
		t.Errorf("expected nodes %d; got %d", 1, len(nodes))
	} else if html := nodes[0].Readable(); html != "<title>The Dormouse's story</title>" {
		t.Errorf("expected html %q; got %q", "<title>The Dormouse's story</title>", html)
	}
	if html := soup.Find(0, Title).Readable(); html != "<title>The Dormouse's story</title>" {
		t.Errorf("expected html %q; got %q", "<title>The Dormouse's story</title>", html)
	}
	if node := soup.Find(0, Tag("nosuchtag")); node != nil {
		t.Errorf("expected node nil; got %q", node.Readable())
	}
	if html := soup.Find(0, Head).Find(0, Title).Readable(); html != "<title>The Dormouse's story</title>" {
		t.Errorf("expected html %q; got %q", "<title>The Dormouse's story</title>", html)
	}
}

func TestFindMethod(t *testing.T) {
	aString := soup.FindString(0, String("Lacie"))
	if text := aString.String(); text != "Lacie" {
		t.Errorf("expected string %q; got %q", "Lacie", text)
	}
	if nodes := aString.FindAll(Parent, A); len(nodes) != 1 {
		t.Errorf("expected nodes %d; got %d", 1, len(nodes))
	} else if html := nodes[0].Readable(); html != lacie {
		t.Errorf("expected html %q; got %q", lacie, html)
	}
	if html := aString.Find(Parent, P).Readable(); !strings.HasPrefix(html, `<p class="story">Once upon a time there were`) {
		t.Errorf("expected html %q; got %q", `<p class="story">Once upon a time there were`, html)
	}
	if nodes := aString.FindAll(Parent, P, Class("title")); len(nodes) != 0 {
		t.Errorf("expected nodes %d; got %d", 0, len(nodes))
	}
	firstLink := soup.Find(0, A)
	if nodes := firstLink.FindAll(NextSibling, A); len(nodes) != 2 {
		t.Errorf("expected nodes %d; got %d", 2, len(nodes))
	} else {
		expected := []string{lacie, tillie}
		for i, node := range nodes {
			if html := node.Readable(); html != expected[i] {
				t.Errorf("expected html #%d %q; got %q", i, expected[i], html)
			}
		}
	}
	if html := soup.Find(0, P, Class("story")).Find(NextSibling, P).Readable(); html != `<p class="story">...</p>` {
		t.Errorf("expected html %q; got %q", `<p class="story">...</p>`, html)
	}
	lastLink := soup.Find(0, A, Id("link3"))
	if nodes := lastLink.FindAll(PrevSibling, A); len(nodes) != 2 {
		t.Errorf("expected nodes %d; got %d", 2, len(nodes))
	} else {
		expected := []string{lacie, elsie}
		for i, node := range nodes {
			if html := node.Readable(); html != expected[i] {
				t.Errorf("expected html #%d %q; got %q", i, expected[i], html)
			}
		}
	}
	if html := soup.Find(0, P, Class("story")).Find(PrevSibling, P).Readable(); html != `<p class="title"><b>The Dormouse's story</b></p>` {
		t.Errorf("expected html %q; got %q", `<p class="title"><b>The Dormouse's story</b></p>`, html)
	}
	if strings := firstLink.FindAllString(Next, True); len(strings) != 9 {
		t.Errorf("expected nodes %d; got %d", 9, len(strings))
	} else {
		expected := []string{"Elsie", ",\n", "Lacie", " and\n", "Tillie",
			";\nand they lived at the bottom of a well.", "\n\n", "...", "\n"}
		for i, node := range strings {
			if html := node.String(); html != expected[i] {
				t.Errorf("expected html #%d %q; got %q", i, expected[i], html)
			}
		}
	}
	if html := firstLink.Find(Next, P).Readable(); html != `<p class="story">...</p>` {
		t.Errorf("expected html %q; got %q", `<p class="story">...</p>`, html)
	}
	if nodes := firstLink.FindAll(Previous, P); len(nodes) != 2 {
		t.Errorf("expected nodes %d; got %d", 2, len(nodes))
	} else {
		expected := []string{
			`<p class="story">Once upon a time there were three little sisters;`,
			`<p class="title"><b>The Dormouse's story</b></p>`,
		}
		for i, node := range nodes {
			if html := node.Readable(); !strings.HasPrefix(html, expected[i]) {
				t.Errorf("expected html #%d %q; got %q", i, expected[i], html)
			}
		}
	}
	if html := firstLink.Find(Previous, Title).Readable(); html != `<title>The Dormouse's story</title>` {
		t.Errorf("expected html %q; got %q", `<title>The Dormouse's story</title>`, html)
	}
	if nodes := soup.Find(0, Tag("html")).FindAll(NoRecursive, Title); len(nodes) != 0 {
		t.Errorf("expected nodes %d; got %d", 0, len(nodes))
	}
}

func TestSelectAll(t *testing.T) {
	if nodes := soup.SelectAll("title"); len(nodes) != 1 {
		t.Errorf("expected nodes %d; got %d", 1, len(nodes))
	} else if html := nodes[0].Readable(); html != "<title>The Dormouse's story</title>" {
		t.Errorf("expected html %q; got %q", "<title>The Dormouse's story</title>", html)
	}
	if nodes := soup.SelectAll("p:nth-of-type(3)"); len(nodes) != 1 {
		t.Errorf("expected nodes %d; got %d", 1, len(nodes))
	} else if html := nodes[0].Readable(); html != `<p class="story">...</p>` {
		t.Errorf("expected html %q; got %q", `<p class="story">...</p>`, html)
	}
	if nodes := soup.SelectAll("body a"); len(nodes) != 3 {
		t.Errorf("expected nodes %d; got %d", 3, len(nodes))
	} else {
		expected := []string{elsie, lacie, tillie}
		for i, node := range nodes {
			if html := node.Readable(); html != expected[i] {
				t.Errorf("expected html #%d %q; got %q", i, expected[i], html)
			}
		}
	}
	if nodes := soup.SelectAll("html head title"); len(nodes) != 1 {
		t.Errorf("expected nodes %d; got %d", 1, len(nodes))
	} else if html := nodes[0].Readable(); html != "<title>The Dormouse's story</title>" {
		t.Errorf("expected html %q; got %q", "<title>The Dormouse's story</title>", html)
	}
	if nodes := soup.SelectAll("head > title"); len(nodes) != 1 {
		t.Errorf("expected nodes %d; got %d", 1, len(nodes))
	} else if html := nodes[0].Readable(); html != "<title>The Dormouse's story</title>" {
		t.Errorf("expected html %q; got %q", "<title>The Dormouse's story</title>", html)
	}
	if nodes := soup.SelectAll("p > a"); len(nodes) != 3 {
		t.Errorf("expected nodes %d; got %d", 3, len(nodes))
	} else {
		expected := []string{elsie, lacie, tillie}
		for i, node := range nodes {
			if html := node.Readable(); html != expected[i] {
				t.Errorf("expected html #%d %q; got %q", i, expected[i], html)
			}
		}
	}
	if nodes := soup.SelectAll("p > a:nth-of-type(2)"); len(nodes) != 1 {
		t.Errorf("expected nodes %d; got %d", 1, len(nodes))
	} else if html := nodes[0].Readable(); html != lacie {
		t.Errorf("expected html %q; got %q", lacie, html)
	}
	if nodes := soup.SelectAll("p > #link1"); len(nodes) != 1 {
		t.Errorf("expected nodes %d; got %d", 1, len(nodes))
	} else if html := nodes[0].Readable(); html != elsie {
		t.Errorf("expected html %q; got %q", elsie, html)
	}
	if nodes := soup.SelectAll("body > a"); len(nodes) != 0 {
		t.Errorf("expected nodes %d; got %d", 0, len(nodes))
	}
	if nodes := soup.SelectAll(".sister"); len(nodes) != 3 {
		t.Errorf("expected nodes %d; got %d", 3, len(nodes))
	} else {
		expected := []string{elsie, lacie, tillie}
		for i, node := range nodes {
			if html := node.Readable(); html != expected[i] {
				t.Errorf("expected html #%d %q; got %q", i, expected[i], html)
			}
		}
	}
	if nodes := soup.SelectAll("[class~=sister]"); len(nodes) != 3 {
		t.Errorf("expected nodes %d; got %d", 3, len(nodes))
	} else {
		expected := []string{elsie, lacie, tillie}
		for i, node := range nodes {
			if html := node.Readable(); html != expected[i] {
				t.Errorf("expected html #%d %q; got %q", i, expected[i], html)
			}
		}
	}
	if nodes := soup.SelectAll("#link1"); len(nodes) != 1 {
		t.Errorf("expected nodes %d; got %d", 1, len(nodes))
	} else if html := nodes[0].Readable(); html != elsie {
		t.Errorf("expected html %q; got %q", elsie, html)
	}
	if nodes := soup.SelectAll("a#link2"); len(nodes) != 1 {
		t.Errorf("expected nodes %d; got %d", 1, len(nodes))
	} else if html := nodes[0].Readable(); html != lacie {
		t.Errorf("expected html %q; got %q", lacie, html)
	}
	if nodes := soup.SelectAll("#link1,#link2"); len(nodes) != 2 {
		t.Errorf("expected nodes %d; got %d", 2, len(nodes))
	} else {
		expected := []string{elsie, lacie}
		for i, node := range nodes {
			if html := node.Readable(); html != expected[i] {
				t.Errorf("expected html #%d %q; got %q", i, expected[i], html)
			}
		}
	}
	if nodes := soup.SelectAll("a[href]"); len(nodes) != 3 {
		t.Errorf("expected nodes %d; got %d", 3, len(nodes))
	} else {
		expected := []string{elsie, lacie, tillie}
		for i, node := range nodes {
			if html := node.Readable(); html != expected[i] {
				t.Errorf("expected html #%d %q; got %q", i, expected[i], html)
			}
		}
	}
	if nodes := soup.SelectAll(`a[href="http://example.com/elsie"]`); len(nodes) != 1 {
		t.Errorf("expected nodes %d; got %d", 1, len(nodes))
	} else if html := nodes[0].Readable(); html != elsie {
		t.Errorf("expected html %q; got %q", elsie, html)
	}
	if nodes := soup.SelectAll(`a[href^="http://example.com/"]`); len(nodes) != 3 {
		t.Errorf("expected nodes %d; got %d", 3, len(nodes))
	} else {
		expected := []string{elsie, lacie, tillie}
		for i, node := range nodes {
			if html := node.Readable(); html != expected[i] {
				t.Errorf("expected html #%d %q; got %q", i, expected[i], html)
			}
		}
	}
	if nodes := soup.SelectAll(`a[href$="tillie"]`); len(nodes) != 1 {
		t.Errorf("expected nodes %d; got %d", 1, len(nodes))
	} else if html := nodes[0].Readable(); html != tillie {
		t.Errorf("expected html %q; got %q", tillie, html)
	}
	if nodes := soup.SelectAll(`a[href*=".com/el"]`); len(nodes) != 1 {
		t.Errorf("expected nodes %d; got %d", 1, len(nodes))
	} else if html := nodes[0].Readable(); html != elsie {
		t.Errorf("expected html %q; got %q", elsie, html)
	}
}

func TestSelect(t *testing.T) {
	if html := soup.Select(".sister").Readable(); html != elsie {
		t.Errorf("expected html %q; got %q", elsie, html)
	}
}

func TestXPath(t *testing.T) {
	if nodes := soup.XPath("//title[1]"); len(nodes) != 1 {
		t.Errorf("expected nodes %d; got %d", 1, len(nodes))
	} else if html := nodes[0].Readable(); html != "<title>The Dormouse's story</title>" {
		t.Errorf("expected html %q; got %q", "<title>The Dormouse's story</title>", html)
	}
	if node := soup.XPath("//nosuchtag"); len(node) != 0 {
		t.Errorf("expected node nil; got %q", node[0].Readable())
	}
	if nodes := soup.XPath("//head/title"); len(nodes) != 1 {
		t.Errorf("expected nodes %d; got %d", 1, len(nodes))
	} else if html := nodes[0].Readable(); html != "<title>The Dormouse's story</title>" {
		t.Errorf("expected html %q; got %q", "<title>The Dormouse's story</title>", html)
	}
}

func TestEvaluate(t *testing.T) {
	if _, err := soup.Evaluate("$test"); err == nil {
		t.Errorf("expected error; got nil")
	}
	if res, err := soup.Evaluate("count(//a)"); err != nil {
		t.Error(err)
	} else if v, ok := res.(float64); !ok {
		t.Errorf("expect type float64; got %s", reflect.TypeOf(res))
	} else if v != 3 {
		t.Errorf("expected count 3; got %g", v)
	}
	if res, err := soup.Evaluate(`local-name(//a)`); err != nil {
		t.Error(err)
	} else if v, ok := res.(string); !ok {
		t.Errorf("expect type string; got %s", reflect.TypeOf(res))
	} else if v != "a" {
		t.Errorf("expected %q; got %q", "a", v)
	}
	if res, err := soup.Evaluate(`boolean(//table)`); err != nil {
		t.Error(err)
	} else if v, ok := res.(bool); !ok {
		t.Errorf("expect type bool; got %s", reflect.TypeOf(res))
	} else if v {
		t.Error("expected false; got true")
	}
}
