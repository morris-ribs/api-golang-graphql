package main

import (
	"api-golang-graphql/musicutil"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
)

func main() {
	// Schema
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("query")
		result := graphql.Do(graphql.Params{
			Schema:        musicutil.MusicSchema,
			RequestString: query,
		})
		json.NewEncoder(w).Encode(result)
	})
	fmt.Println("Now server is running on port 8090")
	fmt.Println("Test with Get      : curl -g 'http://localhost:8090/graphql?query={artist{name}}'")
	http.ListenAndServe(":8090", nil)
}
