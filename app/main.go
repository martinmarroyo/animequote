/*
animequote fetches a random anime quote from your favorite anime title, character, or
simply any random anime. It uses the AnimeChan API (https://animechan.xyz/) to generate
these quotes.

Usage:

	animequote [flags]

The flags are:

	-title
		Enter an anime title to get a random quote from that anime.
	-character
		Enter an anime character name to get a random quote from them.

When no flag is provided, the program will generate and return a random quote from a random anime.

Examples:

	Get a quote by title:
		animequote -title "One Piece"

	Get a quote by character name:
		animequote -character Guts
*/
package main

import (
	"flag"
	"fmt"
)

func main() {
	searchType, query := parseArgs()
	animeQuote := getAnimeQuote(searchType, query)
	fmt.Println(animeQuote)
}

func parseArgs() (string, string) {
	title := flag.String("title", "foo", "Enter an anime title to get a random quote from that anime.")
	character := flag.String("character", "bar", "Enter an anime character name to get a random quote from them.")
	flag.Parse()
	query := "random"
	searchType := "random"
	if *title != "foo" || (*title != "foo" && *character != "bar") {
		query = *title
		searchType = "title"
	} else {
		query = *character
		searchType = "character"
	}
	return searchType, query
}
