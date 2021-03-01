package utils

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Rss struct {
	Channel []Channel `xml:"channel"`
}

type Channel struct {
	Title         string   `xml:"title"`
	Link          string   `xml:"link"`
	Language      string   `xml:"language"`
	Item          []Item   `xml:"item"`
}

type Item struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"title"`
	Link        string   `xml:"link"`
	PubDate     string   `xml:"pubDate"`
	Description string   `xml:"description"`
	Source      string   `xml:"source"`
}

func GetNewsXML(search string, lang string) ([]byte, *Error) {
	resp, err := http.Get("https://news.google.com/rss/search?q=" + url.QueryEscape(search) + "&hl=" + url.QueryEscape(lang))

	if err != nil {
		return nil, NewError(500, "Internal Server Error")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, NewError(500, "Internal Server Error")
	}
	data, _ := ioutil.ReadAll(resp.Body)
	return data, nil
}
