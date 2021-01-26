package pages

import (
	"github.com/sclevine/agouti"
	"log"
)

func Reservation(page *agouti.Page)  {
	if err := page.Navigate("https://www.yahoo.co.jp/"); err != nil {
		log.Fatal("Failed to navigate:", err)
	}
}
