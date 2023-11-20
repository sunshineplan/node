package node

import "testing"

var (
	soup, _ = ParseHTML(`<html><head><title>The Dormouse's story</title></head>
<body>
<p class="title"><b>The Dormouse's story</b></p>

<p class="story">Once upon a time there were three little sisters; and their names were
<a href="http://example.com/elsie" class="sister" id="link1">Elsie</a>,
<a href="http://example.com/lacie" class="sister" id="link2">Lacie</a> and
<a href="http://example.com/tillie" class="sister" id="link3">Tillie</a>;
and they lived at the bottom of a well.</p>

<p class="story">...</p>
`)

	elsie  = `<a href="http://example.com/elsie" class="sister" id="link1">Elsie</a>`
	lacie  = `<a href="http://example.com/lacie" class="sister" id="link2">Lacie</a>`
	tillie = `<a href="http://example.com/tillie" class="sister" id="link3">Tillie</a>`
)

func TestSoup(t *testing.T) {
	title := soup.Find(0, Title)
	if html := title.Readable(); html != "<title>The Dormouse's story</title>" {
		t.Errorf("expected html %q; got %q", "<title>The Dormouse's story</title>", html)
	}
	if name := title.Data(); name != "title" {
		t.Errorf("expected name %q; got %q", "title", name)
	}
	if s := title.String().String(); s != "The Dormouse's story" {
		t.Errorf("expected string %q; got %q", "The Dormouse's story", s)
	}
	if parentName := title.Parent().Data(); parentName != "head" {
		t.Errorf("expected parent name %q; got %q", "head", parentName)
	}
	p := soup.Find(0, P)
	if html := p.Readable(); html != `<p class="title"><b>The Dormouse's story</b></p>` {
		t.Errorf("expected html %q; got %q", `<p class="title"><b>The Dormouse's story</b></p>`, html)
	}
	if class, _ := p.Attrs().Get("class"); class != "title" {
		t.Errorf("expected class %q; got %q", "title", class)
	}
	if a := soup.Find(0, A).Readable(); a != elsie {
		t.Errorf("expected a %q; got %q", elsie, a)
	}
	if a := soup.FindAll(0, A); len(a) != 3 {
		t.Errorf("expected a %d; got %d", 3, len(a))
	} else {
		links := []string{"http://example.com/elsie", "http://example.com/lacie", "http://example.com/tillie"}
		for i, a := range a {
			if link, _ := a.Attrs().Get("href"); link != links[i] {
				t.Errorf("expected link #%d %q; got %q", i, links[i], link)
			}
		}
	}
	if a := soup.Find(0, nil, Id("link3")).Readable(); a != tillie {
		t.Errorf("expected a %q; got %q", tillie, a)
	}
	s := `The Dormouse's story

The Dormouse's story

Once upon a time there were three little sisters; and their names were
Elsie,
Lacie and
Tillie;
and they lived at the bottom of a well.

...
`
	if text := soup.Find(0, nil).GetText(); text != s {
		t.Errorf("expected text %s; got %q", s, text)
	}
}

func TestGoingDown(t *testing.T) {
	head := soup.Find(0, Head)
	if html := head.Readable(); html != "<head><title>The Dormouse's story</title></head>" {
		t.Errorf("expected html %q; got %q", "<head><title>The Dormouse's story</title></head>", html)
	}
	if html := soup.Find(0, Body).Find(0, B).Readable(); html != "<b>The Dormouse's story</b>" {
		t.Errorf("expected html %q; got %q", "<b>The Dormouse's story</b>", html)
	}
	if children := head.Children(); len(children) != 1 {
		t.Errorf("expected children %d; got %d", 1, len(children))
	} else if html := children[0].Readable(); html != "<title>The Dormouse's story</title>" {
		t.Errorf("expected html %q; got %q", "<title>The Dormouse's story</title>", html)
	} else if children := children[0].Children(); len(children) != 1 && children[0].Data() != "The Dormouse's story" {
		t.Errorf("expected children %d; got %d", 1, len(children))
	}
	if children := soup.Children(); len(children) != 1 {
		t.Errorf("expected children %d; got %d", 1, len(children))
	} else if name := children[0].Data(); name != "html" {
		t.Errorf("expected name %q; got %q", "html", name)
	}
	if l := len(head.Descendants()); l != 2 {
		t.Errorf("expected descendants %d; got %d", 2, l)
	}
	if l := len(soup.Descendants()); l != 26 {
		t.Errorf("expected descendants %d; got %d", 26, l)
	}
	if s := soup.Find(0, Title).String().String(); s != head.String().String() {
		t.Errorf("expected text %q; got %q", "The Dormouse's story", s)
	}
	if node := soup.Find(0, Tag("html")).String(); node != nil {
		t.Errorf("expected nil; got %q", node.String())
	}
	if strings := soup.Strings(); len(strings) != 15 {
		t.Errorf("expected strings %d; got %d", 15, len(strings))
	} else {
		expected := []string{
			"The Dormouse's story",
			"\n",
			"\n",
			"The Dormouse's story",
			"\n\n",
			"Once upon a time there were three little sisters; and their names were\n",
			"Elsie",
			",\n",
			"Lacie",
			" and\n",
			"Tillie",
			";\nand they lived at the bottom of a well.",
			"\n\n",
			"...",
			"\n",
		}
		for i, s := range strings {
			if s := s.String(); s != expected[i] {
				t.Errorf("expected string #%d %q; got %q", i, expected[i], s)
			}
		}
	}
	if strings := soup.StrippedStrings(); len(strings) != 10 {
		t.Errorf("expected strings %d; got %d", 10, len(strings))
	} else {
		expected := []string{
			"The Dormouse's story",
			"The Dormouse's story",
			"Once upon a time there were three little sisters; and their names were",
			"Elsie",
			",",
			"Lacie",
			"and",
			"Tillie",
			";\nand they lived at the bottom of a well.",
			"...",
		}
		for i, s := range strings {
			if s != expected[i] {
				t.Errorf("expected string #%d %q; got %q", i, expected[i], s)
			}
		}
	}
}

func TestGoingUp(t *testing.T) {
	title := soup.Find(0, Title)
	if html := title.Parent().Readable(); html != "<head><title>The Dormouse's story</title></head>" {
		t.Errorf("expected html %q; got %q", "<head><title>The Dormouse's story</title></head>", html)
	}
	if html := title.String().Parent().Readable(); html != "<title>The Dormouse's story</title>" {
		t.Errorf("expected html %q; got %q", "<title>The Dormouse's story</title>", html)
	}
	if parent := soup.Parent(); parent != nil {
		t.Errorf("expected nil; got %q", parent.GetText())
	}
	if parents := soup.Find(0, A).Parents(); len(parents) != 4 {
		t.Errorf("expected parents %d; got %d", 4, len(parents))
	} else {
		names := []string{"p", "body", "html", ""}
		for i, parent := range parents {
			if name := parent.Data(); name != names[i] {
				t.Errorf("expected name #%d %q; got %q", i, names[i], name)
			}
		}
	}
}

func TestGoingSideways(t *testing.T) {
	if node := soup.Find(0, A).NextSibling(); node.Readable() != ",\n" {
		t.Errorf("expected string %q; got %q", ",\n ", node.GetText())
	} else if html := node.NextSibling().Readable(); html != lacie {
		t.Errorf("expected html %q; got %q", lacie, html)
	}
	if nextSiblings := soup.Find(0, A).NextSiblings(); len(nextSiblings) != 5 {
		t.Errorf("expected next_siblings %d; got %d", 5, len(nextSiblings))
	} else {
		expected := []string{
			",\n",
			lacie,
			" and\n",
			tillie,
			";\nand they lived at the bottom of a well.",
		}
		for i, nextSibling := range nextSiblings {
			if html := nextSibling.Readable(); html != expected[i] {
				t.Errorf("expected html #%d %q; got %q", i, expected[i], html)
			}
		}
	}
	if prevSiblings := soup.Find(0, nil, Id("link3")).PrevSiblings(); len(prevSiblings) != 5 {
		t.Errorf("expected previous_siblings %d; got %d", 5, len(prevSiblings))
	} else {
		expected := []string{
			" and\n",
			lacie,
			",\n",
			elsie,
			"Once upon a time there were three little sisters; and their names were\n",
		}
		for i, nextSibling := range prevSiblings {
			if html := nextSibling.Readable(); html != expected[i] {
				t.Errorf("expected html #%d %q; got %q", i, expected[i], html)
			}
		}
	}
}

func TestGoingBackAndForth(t *testing.T) {
	a := soup.Find(0, A, Id("link3"))
	if html := a.Readable(); html != tillie {
		t.Errorf("expected html %q; got %q", tillie, html)
	}
	if html := a.NextSibling().Readable(); html != ";\nand they lived at the bottom of a well." {
		t.Errorf("expected html %q; got %q", ";\nand they lived at the bottom of a well.", html)
	}
	if html := a.NextNode().Readable(); html != "Tillie" {
		t.Errorf("expected html %q; got %q", "Tillie", html)
	}
	if html := a.PrevNode().Readable(); html != " and\n" {
		t.Errorf("expected html %q; got %q", " and\n", html)
	}
	if html := a.PrevNode().NextNode().Readable(); html != tillie {
		t.Errorf("expected html %q; got %q", tillie, html)
	}
	if nextNodes := a.NextNodes(); len(nextNodes) != 6 {
		t.Errorf("expected next_elements %d; got %d", 6, len(nextNodes))
	} else {
		expected := []string{
			"Tillie",
			";\nand they lived at the bottom of a well.",
			"\n\n",
			`<p class="story">...</p>`,
			"...",
			"\n",
		}
		for i, nextNode := range nextNodes {
			if html := nextNode.Readable(); html != expected[i] {
				t.Errorf("expected html #%d %q; got %q", i, expected[i], html)
			}
		}
	}
}
