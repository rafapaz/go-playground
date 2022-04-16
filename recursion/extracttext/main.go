package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	var text string

	resp, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Printf("Não pegou %s\n", os.Args[1])
		os.Exit(1)
	}

	defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)
	//fmt.Printf("%s\n", body)
	doc, err := html.Parse(resp.Body)

	text = extract(text, doc)
	fmt.Printf("%+s\n", text)
}

func extract(text string, n *html.Node) string {

	if n.Type == html.TextNode {
		text += n.Data
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		text = extract(text, c)
	}

	return text
}
