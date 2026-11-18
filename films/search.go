package films

import (
        "encoding/xml"
        "fmt"
        "net/http"
)


func SearchFilms(nick string) ([]Item, error) {
	resp, err := http.Get("https://letterboxd.com/" + nick + "/rss/")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var dados Rss
	err = xml.NewDecoder(resp.Body).Decode(&dados)
	if err != nil {
		return nil, err
	}

	items := dados.Channel.Items

	if len(items) == 0 {
		return nil, fmt.Errorf("nenhum item encontrado para %s", nick)
	}
	return items, nil
}


