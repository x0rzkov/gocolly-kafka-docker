package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"strconv"
)

func main () {


	log.Println("starting scrapper")
	c := colly.NewCollector(
			colly.AllowedDomains("www.autoreflex.com"),
			colly.MaxDepth(1),
		)

	c.OnHTML("tr[star-id]", func(e *colly.HTMLElement){
		//link := e.Attr("star-id")
		//link := e.Attr("")
		fmt.Println(e.ChildAttr("a[href]", "href"))
	})

	for i := 1; i < 5; i ++ {
		err := c.Visit("http://www.autoreflex.com/137.0.-1.-1.-1.0.999999.1900.999999.-1.99.0." + strconv.Itoa(i) + "?fulltext=&geoban=M137R99")
		if err != nil {
			log.Panic(err)
		}
	}

	log.Println("end")
}


