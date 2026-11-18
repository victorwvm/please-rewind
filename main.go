package main

import (
    "fmt"
    "github.com/victorwvm/please-rewind/films"
)

func main() {
    
    films, err := films.SearchFilms("victorwv")
    if err != nil {
        fmt.Printf("Erro ao buscar filmes: %v\n", err)
        return
    }

    fmt.Println("Filmes encontrados:")
    for _, film := range films {
        fmt.Printf("- %s\n", film.Title)
    }
}