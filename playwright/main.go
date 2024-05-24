package main

import (
	"fmt"
	"github.com/playwright-community/playwright-go"
	"github.com/uzzalhcse/crawlme/common"
	"log"
)

func main() {
	pw, err := playwright.Run()
	handleError(err)
	defer pw.Stop()

	// Launch browser with headless mode option
	browser, err := launchBrowser(pw, true)
	handleError(err)
	defer browser.Close()

	page, err := createPage(browser, common.UserAgent)
	handleError(err)
	defer page.Close()

	checkPriceMultipleTimes(page, common.URL, common.ATTEMPTS)
}

func launchBrowser(pw *playwright.Playwright, headless bool) (playwright.Browser, error) {
	options := playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(headless),
	}
	return pw.Chromium.Launch(options)
}

func createPage(browser playwright.Browser, userAgent string) (playwright.Page, error) {
	return browser.NewPage(playwright.BrowserNewPageOptions{
		UserAgent: playwright.String(userAgent),
	})
}

func navigateToURL(page playwright.Page, url string) error {
	_, err := page.Goto(url, playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateNetworkidle,
	})
	return err
}

func getPrice(page playwright.Page, locator string) (string, error) {
	return page.Locator(locator).TextContent()
}

func handleError(err error) {
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}

func checkPriceMultipleTimes(page playwright.Page, url string, attempts int) {
	emptyCount := 0

	for i := 0; i < attempts; i++ {
		err := navigateToURL(page, url)
		handleError(err)

		price, err := getPrice(page, ".af-price.price")
		handleError(err)
		fmt.Printf("(%d) Price: %v\n", i+1, price)
		if price == "" {
			emptyCount++
		}
	}

	fmt.Printf("Number of times selling price was empty: %d\n", emptyCount)
}
