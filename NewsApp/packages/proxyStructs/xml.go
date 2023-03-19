package proxyStructs

import (
	"encoding/xml"
	"fmt"
	"time"
)

type Feed struct{
	XMLName xml.Name		`xml:"rss"`
	Chanel Channel			`xml:"channel"`

}

type Channel struct {
	List []Article `xml:"item"`
}


type Article struct{
	ID int
	Title string		`xml:"title"`
	Description string	`xml:"description"`
	PubDate timeInt64	`xml:"pubDate"`
	Link string			`xml:"link"`
}


type timeInt64 int64

func (t *timeInt64) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error{
	var s string
	err := d.DecodeElement(&s, &start)
	if err != nil{
		return err
	}

	var tm time.Time
	// s = strings.ReplaceAll(s, ",", "")
	tm, err = time.Parse("Mon, 2 Jan 2006 15:04:05 -0700", s)
	if err != nil{
		tm, err = time.Parse("Mon, 2 Jan 2006 15:04:05 GMT", s)
		if err != nil{
			fmt.Println(s)
			return err
		}
	}

	*t = timeInt64(tm.Unix())
	return nil
}

/*tm, err = time.Parse("Mon 2 Jan 2006 15:04:05 -0700", s)
if err != nil{
tm, err = time.Parse("Mon 2 Jan 2006 15:04:05 GMT", s)
if err != nil{
fmt.Println(s)
return err
}
}*/