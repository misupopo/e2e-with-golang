package main

import (
	"e2e-with-golang/driver"
	"e2e-with-golang/pages"
	"log"
	"time"
)

func main() {
	d, page := driver.CreateNewBrowser(1090, 768, 30000)

	pages.Reservation(page)

	time.Sleep(2 * time.Second)

	if err := d.Stop(); err != nil {
		log.Fatal("Failed to close pages and stop WebDriver:", err)
	}
}
