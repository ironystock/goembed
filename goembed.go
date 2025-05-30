package goembed

type Photo struct {
	oembed
	URL    string
	Width  int
	Height int
}
type Video struct {
	oembed
	HTML   string
	Width  int
	Height int
}
type Rich struct {
	oembed
	HTML   string
	Width  int
	Height int
}
type Link struct {
	oembed
}
type oembed struct {
	Type            string
	Version         float32
	Title           string
	AuthorName      string
	AuthorUrl       string
	ProviderName    string
	ProviderUrl     string
	CacheAge        string
	ThumbnailUrl    string
	ThumbnailWidth  int
	ThumbnailHeight int
}

//client

func Consume(u string, h string) {

}

// server
func Provide() {

}

//json decorator

//xml decorator [EW]
