package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	// "encoding/json"

	"github.com/gocolly/colly"
)

type VegFoods struct {
	Title string `json:"title"`
	Img   string `json:"img"`
	Href  string `json:"href"`
}

func main() {
	c := colly.NewCollector(
		// Restrict crawling to specific domains
		colly.AllowedDomains("recipeforvegans.com"),
	)

	// var vegFoodsStr []byte

	c.OnHTML("a.dj-thumb-link", func(e *colly.HTMLElement) {
		title := e.Attr("title")

		href := e.Attr("href")

		img := e.ChildAttr("img", "src")

		fmt.Println(title, img)

		vegFoods := VegFoods{
			Title: title,
			Img:   img,
			Href:  href,
		}

		vegFoodJson, _ := json.Marshal(vegFoods)

		// vegFoodsStr = append(vegFoodJson, vegFoodsStr)

		if err := os.WriteFile("file.txt", []byte(vegFoodJson), 0666); err != nil {
			log.Fatal(err)
		}
	})

	// if err := os.WriteFile("file.txt", []byte(vegFoodsStr), 0666); err != nil {
	// 	log.Fatal(err)
	// }

	c.Visit("https://recipeforvegans.com/")

}
