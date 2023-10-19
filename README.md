# `animequote`

`animequote` fetches a random anime quote from your favorite anime title, character, or
simply any random anime. It uses the [AnimeChan API](https://animechan.xyz/) to generate
these quotes.

Usage:
```bash
animequote [flags]
```
The flags are:

- `-title`: Enter an anime title to get a random quote from that anime.
- `-character`: Enter an anime character name to get a random quote from them.

When no flag is provided, the program will generate and return a random quote from a random anime.

Examples:

Get a quote by title:
```bash
animequote -title "One Piece"
```

Get a quote by character name:
```bash
animequote -character Guts
```