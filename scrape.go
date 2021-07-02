package main

import (
	"fmt"
	"strconv"

	"github.com/gocolly/colly"
)

type VegFood struct {
	Title        string   `json:"title"`
	Img          string   `json:"img"`
	Href         string   `json:"href"`
	Description  string   `json:"desc"`
	Ingredients  []string `json:"ingr"`
	Instructions string   `json:"instrc"`
}

func main() {
	c := colly.NewCollector(
		// Restrict crawling to specific domains
		colly.AllowedDomains("wellvegan.com"),
	)

	c.OnHTML("a.aligncenter ", func(e *colly.HTMLElement) {
		VegFoodHref, err := strconv.Atoi(e.Attr("href"))
		if err != nil {
			fmt.Println("Couldn't get href")
		}

		img := e.ChildText("img").Attr("src")
		// imgSrc := img.Attr("src")
		title := e.Attr("title")
		desc := string
		ingr := string
		instr := string

		c.Visit(VegFoodHref)

		c.OnHTML(".tasty-recipes-description-body", func(e *colly.HTMLElement) {
			desc = e.ChildText("p").Text + e.ChildText("a").Text
		})

		c.OnHTML(".tasty-recipes-ingredients-body ul", func(e *colly.HTMLElement) {
			ingr = e.ChildText("li").ChildText("span").Text + e.ChildText("li").Text
		})

		c.OnHTML(".tasty-recipes-instructions-body ol", func(e *colly.HTMLElement) {
			instr = e.ChildText("li").Text
		})

		vegFood := VegFood{
			Title:        title,
			Img:          img,
			Href:         VegFoodHref,
			Description:  desc,
			Ingredients:  ingr,
			Instructions: instr,
		}

	})

	c.Visit("https://wellvegan.com/recipes")

}
