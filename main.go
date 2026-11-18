package main

import (
    "fmt"
    "github.com/victorwvm/please-rewind/films"
)

func main() {
    
    films, err := films.SearchFilms("morguitaas")
    if err != nil {
        fmt.Printf("error fetching films: %v\n", err)
        return
    }

    fmt.Println("films found:")
    for _, film := range films {
        fmt.Printf("- %s\n", film.Title)
    }
}