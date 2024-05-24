package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
	"github.com/uzzalhcse/crawlme/common"
	"log"
	"runtime"
	"strings"
)

var Browser *rod.Browser

func main() {
	// Initialize the browser
	openBrowser()
	defer Browser.Close()
	emptyCount := 0
	for i := 0; i < common.ATTEMPTS; i++ {
		dom, page, err := crawlSite(common.URL)
		if err != nil {
			log.Printf("Error in crawlSite on attempt %d: %v", i+1, err)
			continue
		}
		defer page.Close()

		price := common.GetSellingPrice(dom)
		fmt.Printf("(%d) Price: %v\n", i+1, price)
		if price == "" {
			emptyCount++
		}
	}

	fmt.Printf("Number of times selling price was empty: %d\n", emptyCount)
}

func crawlSite(url string) (*goquery.Document, *rod.Page, error) {
	rodPage := Browser.MustPage()
	err := rodPage.SetUserAgent(&proto.NetworkSetUserAgentOverride{
		UserAgent: common.UserAgent,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("error setting user agent: %w", err)
	}

	rodPage.MustNavigate(url).MustWaitLoad()

	pageData, err := GetDom(rodPage)
	if err != nil {
		return nil, nil, err
	}
	return pageData, rodPage, nil
}

func openBrowser() {
	lnchr := launcher.New().
		Headless(true).
		NoSandbox(true).Devtools(true)
	lnchr.Set("disable-features", "Translate")

	browserCmd := "/usr/bin/google-chrome"
	if runtime.GOOS == "darwin" {
		browserCmd = "/Applications/Google Chrome.app/Contents/MacOS/Google Chrome"
	}
	controlUrl := lnchr.Bin(browserCmd).MustLaunch()
	Browser = rod.New().ControlURL(controlUrl).MustConnect()
}

func GetDom(rodPage *rod.Page) (*goquery.Document, error) {
	html, err := rodPage.HTML()
	if err != nil {
		return nil, err
	}

	pageData, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, err
	}

	return pageData, nil
}
