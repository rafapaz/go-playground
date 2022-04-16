package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	var count map[string]int
	count = make(map[string]int)

	resp, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Printf("Não pegou %s\n", os.Args[1])
		os.Exit(1)
	}

	defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)
	//fmt.Printf("%s\n", body)
	doc, err := html.Parse(resp.Body)

	count = countnodes(count, doc)
	fmt.Printf("%+v\n", count)
}

func countnodes(count map[string]int, n *html.Node) map[string]int {

	if n.Type == html.ElementNode {
		count[n.Data]++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		count = countnodes(count, c)
	}

	return count
}
