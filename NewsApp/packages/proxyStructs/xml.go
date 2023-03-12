package proxyStructs

import "encoding/xml"

type Feed struct{
	XMLName xml.Name		`xml:"rss"`
	Chanel Channel			`xml:"channel"`

}

type Channel struct {
	List []Article `xml:"item"`
}


type Article struct{
	Title string		`xml:"title"`
	Description string	`xml:"description"`
	PubDate string		`xml:"pubDate"`
	Link string			`xml:"link"`
}

