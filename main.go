package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func WishlistScrape() {

	doc, err := goquery.NewDocument("https://www.amazon.co.uk/gp/registry/wishlist/CY2FJBBFFULT")
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("[id^=itemInfo]").Each(func(i int, s *goquery.Selection) {
		name, _ := doc.Find("[id^=itemName]").Eq(i).Attr("title")
		price := strings.TrimSpace(doc.Find("[id^=itemPrice]").Eq(i).Text())
		fmt.Printf("%d Name: %s, Price: %s\n", i, name, price)
	})
}

func main() {
	WishlistScrape()
}
