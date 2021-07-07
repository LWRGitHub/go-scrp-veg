package main

import (
	"fmt"

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
		colly.AllowedDomains("recipeforvegans.com"),
	)

	c.OnHTML("a.dj-thumb-link", func(e *colly.HTMLElement) {
		title := e.Attr("title")

		img := e.ChildAttr("img", "src")

		fmt.Println(title, img)

		// imgSrc := img.Attr("src")
		// title := e.Attr("title")
		// desc := string
		// ingr := string
		// instr := string

		// c.Visit(VegFoodHref)

		// c.OnHTML(".tasty-recipes-description-body", func(e *colly.HTMLElement) {
		// 	desc = e.ChildText("p") + e.ChildText("a")
		// })

		// c.OnHTML(".tasty-recipes-ingredients-body ul", func(e *colly.HTMLElement) {
		// 	ingr = e.Attr("li").Attr("span").Text + e.ChildText("li")
		// })

		// c.OnHTML(".tasty-recipes-instructions-body ol", func(e *colly.HTMLElement) {
		// 	instr = e.ChildText("li")
		// })

		// vegFood := VegFood{
		// 	Title:        title,
		// 	Img:          img,
		// 	Href:         VegFoodHref,
		// 	Description:  desc,
		// 	Ingredients:  ingr,
		// 	Instructions: instr,
		// }

	})

	c.Visit("https://recipeforvegans.com/")

}
