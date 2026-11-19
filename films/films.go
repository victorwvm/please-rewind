package films

type Rss struct {
	Channel Channel `xml:"channel"`
}

type Channel struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Items       []Item `xml:"item"`
}

type Item struct {
	FilmTitle   string `xml:"filmTitle"`
	FilmYear    string `xml:"filmYear"`
	MemberRating string `xml:"memberRating"`
	PubDate    string `xml:"pubDate"`
	Description  string `xml:"description"`
	Link        string `xml:"link"`
	Guid        string `xml:"guid"`
}		