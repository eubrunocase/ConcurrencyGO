package main

import (
	"C2-Golang/internal/fetcher"
	"C2-Golang/internal/processor"
	"C2-Golang/internal/models"
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	priceChannel := make(chan models.PriceDetails)
	done := make(chan bool)

	go fetcher.FectchPrices(priceChannel)
	go processor.ShowPricesAVG(priceChannel, done)

	<- done

	fmt.Printf("Tempo total: %s\n", time.Since(start))
}
