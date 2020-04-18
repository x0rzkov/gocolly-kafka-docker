package main

import (
	"github.com/gocolly/colly"
	"log"
	"strconv"
	"strings"
)

// voiture-occasion-renault-kadjar-1-5-blue-dci-115ch-black--FVO16533323.htm

func handleStandalone (url string) {
	var a Announce

	domain := "www.autoreflex.com"

	c := colly.NewCollector(
		colly.AllowedDomains(domain),
		colly.MaxDepth(1),
	)

	// Announce title
	c.OnHTML("div.header", func(e *colly.HTMLElement) {
		title := e.ChildText("h1")
		//log.Println(title)
		a.Title = title
	})

	// Price
	c.OnHTML("div.prix-container", func(e *colly.HTMLElement) {
		price := e.ChildText("div.prix")
		price = strings.TrimSuffix(price, "€")
		price = strings.Replace(price, " ", "", -1)
		n, _ := strconv.Atoi(price)
		//log.Println(n)
		a.Price = n
	})

	// Seller and phone
	c.OnHTML("div.contact", func(e *colly.HTMLElement) {
		/*garageName := e.ChildText("a[title]")
		log.Println(garageName)*/

		phone := e.ChildAttr("a[data-seller-telephone]","data-seller-telephone")
		//log.Println(phone)
		a.Phone = phone
	})

	// Seller name and address
	c.OnHTML("div.small-12.large-5.columns.mg-bottom", func (e *colly.HTMLElement) {
		garageName := e.ChildText("h3")
		//log.Println(garageName)
		a.GarageName = garageName

		garageAddress := e.ChildText("p")
		// Removes <br> tag
		garageAddress = strings.Replace(garageAddress, "\n", "", -1)
		// Removes the 56 spaces before the postal code (no idea they've put 56 blanks here)
		garageAddress = strings.Replace(garageAddress, "                                                        ", " ", -1)
		//log.Println(garageAddress)
		a.GarageAddress = garageAddress
	})

	c.OnHTML("div.specs", func(e *colly.HTMLElement) {
		//km := e.ChildText("li")
		attr := e.ChildTexts("li")

		a.CarMileage = attr[0] // KM
		a.CarFirstRegistration = attr[1]
		a.AnnouncePostalCode = attr[2]
		a.CarEngine = attr[3]
		a.Transmission = attr[4]
		a.FiscalPower = attr[5]

		/*for i, j := range attr {
			log.Print(i, " ", j, " | ")
		}*/
	})

	if err := c.Visit("http://" + domain + "/" + url); err != nil {
		log.Printf("Error when visiting %s", "http://" + domain + "/" + url)
	}

	log.Println(a)
	//log.Println(url)
}
