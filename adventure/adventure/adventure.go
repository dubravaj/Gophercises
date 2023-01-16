package cyoa

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"
)

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

var defaultHandlerTemplate = `
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
    <li><a href="/{{.Chapter}}">{{.Text}}</a></li>
    {{end}}
    </ul>
</body>
</html>
`

type handler struct {
	s      Story
	t      *template.Template
	pathFn func(r *http.Request) string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Parse(defaultHandlerTemplate))
}

type HandlerOption func(h *handler)

func WithTemplate(t *template.Template) HandlerOption {
	return func(h *handler) {
		h.t = t
	}
}

func WithPathFn(fn func(r *http.Request) string) HandlerOption {
	return func(h *handler) {
		h.pathFn = fn
	}
}

func NewHandler(s Story, options ...HandlerOption) http.Handler {
	h := handler{s, tpl, defaultPathFn}
	// apply functional options to default handler
	for _, opt := range options {
		opt(&h)
	}
	return h
}

func defaultPathFn(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)
	fmt.Println(path)

	if path == "" || path == "/" {
		// defaultly render intro chapter
		path = "/intro"
	}

	return path[1:]
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	path := h.pathFn(r)

	// create template if key is in the map
	if chapter, ok := h.s[path]; ok {
		err := h.t.Execute(w, chapter)
		if err != nil {
			log.Printf("%v", err)
			http.Error(w, "Error occured", http.StatusInternalServerError)
		}
		return
	}

	http.Error(w, "Chapter not found.", http.StatusNotFound)
}

// Create Story structure from JSON definition
func JSONStory(r io.Reader) (Story, error) {

	var story Story
	decoder := json.NewDecoder(r)
	if err := decoder.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}
