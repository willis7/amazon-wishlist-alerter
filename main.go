package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Wish represents an item in a wishlist
type Wish struct {
	Name  string
	Price float64
}

func WishlistScrape() {

	doc, err := goquery.NewDocument("https://www.amazon.co.uk/gp/registry/wishlist/CY2FJBBFFULT")
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("[id^=itemInfo]").Each(func(i int, s *goquery.Selection) {

		name, _ := doc.Find("[id^=itemName]").Eq(i).Attr("title")
		price := strings.TrimSpace(doc.Find("[id^=itemPrice]").Eq(i).Text())

		t := strings.Trim(price, "£")
		fmt.Printf("%d Name: %s, Price: £%s\n", i, name, t)
		//s, _ := strconv.ParseFloat(t, 64)
		//w := &Wish{
		//	Name:  name,
		//	Price: s,
		//}
		//fmt.Print(w)
	})
}

func main() {
	WishlistScrape()
}
