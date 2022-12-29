package urlshort

import (
	"fmt"
	"net/http"

	"gopkg.in/yaml.v3"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if redirectUrl, ok := pathsToUrls[r.URL.String()]; ok {
			fmt.Println("Match found: ", redirectUrl)
			http.Redirect(w, r, redirectUrl, http.StatusMovedPermanently)
			return
		}
		fallback.ServeHTTP(w, r)

	}
}

type yamlPair struct {
	Path string
	Url  string
}

type yamlPairs struct {
	YamlData []yamlPair
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//   - path: /some-path
//     url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {

	parsedYaml, err := parseYaml(yml)

	if err != nil {
		panic("Error")
	}

	yamlMap := buildMap(parsedYaml)

	return MapHandler(yamlMap, fallback), err
}

func parseYaml(yml []byte) (yamlPairs, error) {
	var pairs yamlPairs

	err := yaml.Unmarshal(yml, &pairs.YamlData)

	return pairs, err
}

func buildMap(parsedYaml yamlPairs) map[string]string {

	yamlData := make(map[string]string, len(parsedYaml.YamlData))
	fmt.Println(yamlData)

	for _, data := range parsedYaml.YamlData {
		yamlData[data.Path] = data.Url
	}

	return yamlData
}
