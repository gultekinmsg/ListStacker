package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/quote", quoteHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
func quoteHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/quote" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if r.Method == http.MethodPost {

		var message ExpectedQuotes
		err := json.NewDecoder(r.Body).Decode(&message)
		if err != nil {
			log.Fatal(err)
		}

		answer := stackQuotes(message)
		output, err := json.Marshal(answer)
		if err != nil {
			log.Fatal(err)
		}

		w.Header().Set("content-type", "application/json")
		_, err = w.Write(output)
		if err != nil {
			log.Fatal(err)
		}

	} else {
		http.Error(w, "501 not implemented.", http.StatusNotImplemented)
		return
	}
}
func stackQuotes(body ExpectedQuotes) []GivenQuotes {
	var answer []GivenQuotes
	var authors []string
	for _, value := range body {
		if value.AUTHOR != "XX?.1" {
			var quotes []string
			var author string
			author = value.AUTHOR
			authors = append(authors, author)
			quotes = append(quotes, value.QUOTE)
			for index, value1 := range body {
				if value.AUTHOR == value1.AUTHOR && value.QUOTE != value1.QUOTE {
					quotes = append(quotes, value1.QUOTE)
					body[index].AUTHOR = "XX?.1"
				}
			}
			addVal := GivenQuotes{
				AUTHOR: author,
				QUOTES: quotes,
			}
			answer = append(answer, addVal)
		}

	}
	return answer
}
