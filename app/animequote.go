// animequote implements the necessary functionality to call the Animechan (https://animechan.xyz/) API
// to retrieve random quotes from numerous anime series.

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// This is the structure of the quote result from the API
type AnimeQuote struct {
	Title     string `json:"anime"`
	Character string `json:"character"`
	Quote     string `json:"quote"`
}

// getAnimeQuote parses the given searchType and query, then
// fetches the requested quote from the API, returning a string
// that includes the character name, title, and quote
func getAnimeQuote(searchType, query string) string {
	response, err := http.Get(createUrl(searchType, query))
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var payload AnimeQuote
	json.Unmarshal(body, &payload)
	return payload.Speak()
}

// createUrl generates the url that we will use to query the API
// based on the searchType. Currently, it only checks for two searchTypes,
// `title` and `character`. These are passed in via the command line. If neither
// are given, then we default to a randomly generated quote.
func createUrl(searchType, query string) string {
	url := "https://animechan.xyz/api/random/"
	if searchType == "title" && query != "foo" {
		url += "anime?title="
		urlTitle := strings.ToLower(query)
		urlTitle = strings.Replace(query, " ", "%20", -1)
		url += urlTitle
	} else if searchType == "character" && query != "bar" {
		url += "character?name="
		urlTitle := strings.ToLower(query)
		urlTitle = strings.Replace(urlTitle, " ", "%20", -1)
		url += urlTitle
	}
	return url
}

// Speak generates a quote string from the given AnimeQuote result from the API.
func (aq AnimeQuote) Speak() string {
	quote := "Sorry, no results were found for the given title. Please try a different anime."
	if len(aq.Title) > 0 {
		quote = fmt.Sprintf("%s, from the anime %s, says:\n\"%s\"", aq.Character, aq.Title, aq.Quote)
	}
	return quote
}
