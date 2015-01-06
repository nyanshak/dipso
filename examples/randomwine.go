package main

import (
	"os"
	"log"
	"github.com/nyanshak/dipso"
	"flag"
	"errors"
	"math/rand"
	"time"
	"strconv"
	"fmt"
)

var (
	apiKey = os.Getenv("WINE_API_KEY")
	shipState = flag.String("state", "", "state where wine will be shipped")
	totalCost = flag.Float64("cost", 0, "total cost (minus shipping")
	rating = flag.Int("rating", 0, "Optional: minimum professional rating")
	randomGenerator = rand.New(rand.NewSource(time.Now().UnixNano()))
	NoWineFoundError = errors.New("No Wine Found")
)

func init() {
	if apiKey == "" {
		log.Fatalln("API key for Wine.com not found")
	}
	flag.Parse()
	if *shipState == "" || *totalCost <= float64(0) || *rating < 0 {
		flag.Usage()
	}
}

func getRandomWine(api dipso.WineApi, shipState string, cost float64, rating int) (dipso.Wine, error) {
	redWineFilter := "categories(490+124)"
	searchStr := "state=" + shipState + "&instock=true&filter=" + redWineFilter

	if cost > 0 {
		costStr := fmt.Sprintf("%.2f", cost)
		searchStr += "+price(0|" + costStr + ")"
	}

	if rating > 0 {
		ratingStr := fmt.Sprintf("%d", rating)
		searchStr += "+rating(" + ratingStr + "|100)"
	}

	wines, err := api.Search(searchStr + "&size=0&offset=0")

	var blankWine dipso.Wine
	if err != nil {
		return blankWine, err
	}

	if wines.Total == 0 {
		return blankWine, NoWineFoundError
	}

	wineId := int64(randomGenerator.Intn(wines.Total))
	wines, err = api.Search(searchStr + "&size=1&offset=" + strconv.FormatInt(wineId, 10))

	if err != nil {
		return blankWine, err
	}

	return wines.List[0], nil
}

func main() {
	api := dipso.NewWineApi(apiKey)

	wines := []dipso.Wine{}

	remainingMoney := *totalCost
	for {
		wine, err := getRandomWine(*api, "TX", remainingMoney, *rating)
		if err != nil {
			if err != NoWineFoundError {
				log.Fatalln(err)
			} else {
				break
			}
		}

		if wine.PriceRetail > remainingMoney {
			continue
		} else {
			remainingMoney -= wine.PriceRetail
		}
		wines = append(wines, wine)
	}

	log.Printf("%d wine(s) for $%.2f (minus shipping)", len(wines), *totalCost - remainingMoney)
	for _, v := range wines {
		log.Println(v.Retail.Url)
	}
}
