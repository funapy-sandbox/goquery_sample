package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	doc, err := goquery.NewDocument("http://www.gsi.go.jp/KOKUJYOHO/CENTER/kendata/kumamoto_heso.htm")
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("table").Each(func(_ int, s *goquery.Selection) {
		lat := s.Find("TR").Eq(1).Find("TD").Eq(2).Text()
		lon := s.Find("TR").Eq(2).Find("TD").Eq(2).Text()
		fmt.Println(lat)
		fmt.Println(lon)
	})
}
