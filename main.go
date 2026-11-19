package main

import (
    "fmt"
    "github.com/victorwvm/please-rewind/films"
    "strings"
)

func main() {
    
    resp, err := films.SearchFilms("morguitaas")
    if err != nil {
        fmt.Printf("error fetching films: %v\n", err)
        return
    }

    fmt.Println("films found:")
    for _, film := range resp {
        _, afterSrc, _ := strings.Cut(film.Description, `"src="`)
        linkPoster, _, _ := strings.Cut(depoisSrc, `"`) //Guarda o link do poster pq o src finaliza com `"`, isso pega tudo q vem depois de src e antes de "
        
        if strings.Contains(film.Guid, "letterboxd-review-") {
            _, rest, _ := strings.Cut(film.Description, "</p> <p>")
            review, _, _ := strings.Cut(depoisPoster, "</p>")
            review = strings.TrimSpace(review)

            fmt.Printf("- %s (%s) - Rating: %s - Published: %s - Review: %s\n", film.FilmTitle, film.FilmYear, film.MemberRating, film.PubDate, review)
        }
        fmt.Printf("- %s (%s) - Rating: %s - Published: %s - Poster: %s\n", film.FilmTitle, film.FilmYear, film.MemberRating, film.PubDate, poster)
    }
}