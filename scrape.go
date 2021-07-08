package main

import (
	"encoding/json"
	"os"

	// "fmt"

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

	// vegFoodsSlice := []map[string][]byte{}

	c.OnHTML("a.dj-thumb-link", func(e *colly.HTMLElement) {

		title := e.Attr("title")
		href := e.Attr("href")
		img := e.ChildAttr("img", "src")

		vegFoods := VegFoods{
			Title: title,
			Img:   img,
			Href:  href,
		}

		// fmt.Println(vegFoods)
		vegFoodJson, _ := json.Marshal(vegFoods)
		// vegFoodsSlice = append(vegFoodsSlice, vegFoodJson)

		// vegFoodsStr = append(vegFoodJson, vegFoodsStr)

		// fmt.Println(vegFoodJson)

		// if err := os.WriteFile("file.txt", []byte(vegFoodJson), 0666); err != nil {
		// 	log.Fatal(err)
		// }

		f, err := os.OpenFile("file.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			panic(err)
		}

		defer f.Close()

		vegFoodsStr := string(vegFoodJson)

		if _, err = f.WriteString(vegFoodsStr); err != nil {
			panic(err)
		}
	})

	// if err := os.WriteFile("file.txt", []byte(vegFoodsSlice), 0666); err != nil {
	// 	log.Fatal(err)
	// }

	c.Visit("https://recipeforvegans.com/")

}
