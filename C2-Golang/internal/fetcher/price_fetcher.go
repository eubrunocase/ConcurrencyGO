package fetcher

import (
	"C2-Golang/internal/models"
	"math/rand"
	"sync"
	"time"
)

func FetchPriceFromSite1() models.PriceDetails {
	time.Sleep(1 * time.Second)
	Price := models.PriceDetails{
		StoreName: "Site 1",
		Value:     rand.Float64() * 100,
		TimeStamp: time.Now(),
	}
	return Price
}

func FetchPriceFromSite2() models.PriceDetails {
	time.Sleep(2 * time.Second)
	Price := models.PriceDetails{
		StoreName: "Site 2",
		Value:     rand.Float64() * 100,
		TimeStamp: time.Now(),
	}
	return Price
}

func FetchPriceFromSite3() models.PriceDetails {
	time.Sleep(3 * time.Second)
	Price := models.PriceDetails{
		StoreName: "Site 3",
		Value:     rand.Float64() * 100,
		TimeStamp: time.Now(),
	}
	return Price
}

func FectchPrices(priceChannel chan<- models.PriceDetails) {
	var wg sync.WaitGroup
	wg.Add(4)
	go func() {
		defer wg.Done()
		priceChannel <- FetchPriceFromSite1()
	}()
	go func() {
		defer wg.Done()
		priceChannel <- FetchPriceFromSite2()
	}()
	go func() {
		defer wg.Done()
		priceChannel <- FetchPriceFromSite3()
	}()

	go func () {
		defer wg.Done()
		FetchAndSendMultiplePrices(priceChannel)
	}()

	wg.Wait()
	close(priceChannel)
}

func FetchAndSendMultiplePrices(priceChannel chan<- models.PriceDetails) {
	time.Sleep(6 * time.Second)
	prices := []models.PriceDetails{
		{
			StoreName: "Site 4",
			Value:     rand.Float64() * 100,
			TimeStamp: time.Now(),
		},
		{
			StoreName: "Site 5",
			Value:     rand.Float64() * 100,
			TimeStamp: time.Now(),
		},
		{
			StoreName: "Site 6",
			Value:     rand.Float64() * 100,
			TimeStamp: time.Now(),
		},
	}
	for _, price := range prices {
		priceChannel <- price
	}

}