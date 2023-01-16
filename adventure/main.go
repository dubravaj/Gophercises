package main

import (
	cyoa "dubravaj/adventure/cyoa/adventure"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {

	port := flag.Int("port", 3000, "port to start CYOA app")
	filename := flag.String("file", "gopher.json", "JSON file with definition of choose your own adventure story")
	flag.Parse()

	file, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	story, err := cyoa.JSONStory(file)
	if err != nil {
		panic(err)
	}

	tmpl := template.Must(template.New("").Parse(storyTpl))
	h := cyoa.NewHandler(story, cyoa.WithTemplate(tmpl), cyoa.WithPathFn(customPathFn))
	mux := http.NewServeMux()
	mux.Handle("/story/", h)
	fmt.Printf("Starting the server on port: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), mux))

}

func customPathFn(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)
	fmt.Println(path)

	if path == "/story" || path == "/story/" {
		path = "/story/intro"
	}

	return path[len("/story/"):]
}

var storyTpl = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Choose Your Own Adventure</title>
</head>
<body>
    <h1>{{.Title}}</h1>
    {{range .Paragraphs}}
    <p>{{.}}</p>
    {{end}}
    <ul>
    {{range .Options}}
    <li><a href="/story/{{.Chapter}}">{{.Text}}</a></li>
    {{end}}
    </ul>
</body>
</html>
`
