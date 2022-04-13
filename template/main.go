package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Movie struct {
	Title   string
	Year    string
	Rated   string
	Resumo  string `json:"Plot"`
	Ratings []Rating
}

type Rating struct {
	Source string
	Value  string
}

const TEMPLATE string = `
<html>
	<body>
		<form action="/" method="POST" >
			<input type="text" name="texto" id="texto">
			<button action="submit">Enviar</button>
		</form>
		<h1>Movie info</h1>
		<h2>Title: {{.Title}}</h2>
		<p>Year: {{.Year}}</p>
		<p>Rated: {{.Rated}}</p>
		<p>Resumo: {{.Resumo}}</p>
		<b>Ratings:</b>
		{{range .Ratings}}
			<h3>{{.Source}}</h3>: {{.Value}}
		{{end}}
	</body>
</html>
`

var report = template.Must(template.New("movie").
	Parse(TEMPLATE))

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:9000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	//words := []string{"The", "Batman"}
	var words []string

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	words = strings.Split(r.FormValue("texto"), " ")

	movie, err := Search(words)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	if err := report.Execute(w, movie); err != nil {
		log.Fatal(err)
	}
}

func Search(words []string) (*Movie, error) {
	q := url.QueryEscape(strings.Join(words, " "))
	resp, err := http.Get("http://www.omdbapi.com/?apikey=8b9a8b71&plot=full&t=" + q)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("Search query failed: %s", resp.Status)
	}

	var movie Movie
	if err := json.NewDecoder(resp.Body).Decode(&movie); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return &movie, nil
}
