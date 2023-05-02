package node

import (
	"fmt"
	"log"
)

func ExampleAttr() {
	node, err := ParseHTML(`<div data-foo="value">foo!</div>`)
	if err != nil {
		log.Fatal(err)
	}
	if nodes := node.FindAll(0, nil, Attr("data-foo", "value")); len(nodes) != 1 {
		log.Fatalf("expected nodes %d; got %d", 1, len(nodes))
	} else {
		fmt.Println(nodes[0].Readable())
	}
	node, err = ParseHTML(`<input name="email"/>`)
	if err != nil {
		log.Fatal(err)
	}
	if nodes := node.FindAll(0, nil, Attr("name", "email")); len(nodes) != 1 {
		log.Fatalf("expected nodes %d; got %d", 1, len(nodes))
	} else {
		fmt.Println(nodes[0].Readable())
	}
	// Output:
	// <div data-foo="value">foo!</div>
	// <input name="email"/>
}

func ExampleClass() {
	node, err := ParseHTML(`<p class="body strikeout"></p>`)
	if err != nil {
		log.Fatal(err)
	}
	if nodes := node.FindAll(0, nil, Class("body strikeout")); len(nodes) != 1 {
		log.Fatalf("expected nodes %d; got %d", 1, len(nodes))
	} else {
		fmt.Println(nodes[0].Readable())
	}
	if nodes := node.FindAll(0, nil, Class("strikeout body")); len(nodes) != 1 {
		log.Fatalf("expected nodes %d; got %d", 1, len(nodes))
	} else {
		fmt.Println(nodes[0].Readable())
	}
	if nodes := node.FindAll(0, nil, ClassStrict("body strikeout")); len(nodes) != 1 {
		log.Fatalf("expected nodes %d; got %d", 1, len(nodes))
	} else {
		fmt.Println(nodes[0].Readable())
	}
	if nodes := node.FindAll(0, nil, ClassStrict("strikeout body")); len(nodes) != 0 {
		log.Fatalf("expected nodes %d; got %d", 0, len(nodes))
	} else {
		fmt.Println(nodes)
	}
	// Output:
	// <p class="body strikeout"></p>
	// <p class="body strikeout"></p>
	// <p class="body strikeout"></p>
	// []
}

func ExampleHtmlNode() {
	node, err := ParseHTML("<a><b>text1</b><c>text2</c></a>")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(node.Find(0, B).NextSibling().Readable())
	fmt.Println(node.Find(0, Tag("c")).PrevSibling().Readable())
	fmt.Println(node.Find(0, B).PrevSibling())
	fmt.Println(node.Find(0, Tag("c")).NextSibling())
	fmt.Println(node.Find(0, B).String().String())
	fmt.Println(node.Find(0, B).String().NextSibling())
	// Output:
	// <c>text2</c>
	// <b>text1</b>
	// <nil>
	// <nil>
	// text1
	// <nil>
}
