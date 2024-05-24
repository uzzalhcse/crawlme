package common

import (
	"github.com/PuerkitoBio/goquery"
	"log"
)

// HandleError is a common function to handle errors by logging and exiting the program.
func HandleError(err error) {
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}

// UserAgent is a common variable representing a default user agent string.
const UserAgent = "AdvancedRestClient/16.0.4 (Windows NT 10.0; Win64; x64)"
const URL = "https://axel.as-1.co.jp/asone/d/4-2151-01/?cfrom=A0010100"

func GetSellingPrice(dom *goquery.Document) string {
	return dom.Find(".af-price.price").Text()
}
