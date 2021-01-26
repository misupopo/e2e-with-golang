package driver

import (
	"github.com/sclevine/agouti"
	"log"
)

func CreateNewBrowser(width int, height int, setImplicitWait int) (*agouti.WebDriver, *agouti.Page) {
	driver := agouti.ChromeDriver()

	if err := driver.Start(); err != nil {
		log.Fatal("Failed to start driver:", err)
	}

	page, err := driver.NewPage()
	if err != nil {
		log.Fatal("Failed to open instance:", err)
	}

	window, err := page.Session().GetWindow()
	window.SetSize(width, height)

	page.Session().SetImplicitWait(setImplicitWait)

	return driver, page
}
