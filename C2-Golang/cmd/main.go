package main

import (
	"C2-Golang/internal/fetcher"
	"fmt"
	"sync"
	"time"
	"C2-Golang/internal/processor"
)

func main() {
	start := time.Now()

	priceChannel := make(chan float64)
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		processor.ShowPricesAVG(priceChannel)
		}()

	go func() {
		defer wg.Done()
		priceChannel <- fetcher.FetchPriceFromSite1()
	}()
	go func() {
		defer wg.Done()
		priceChannel <- fetcher.FetchPriceFromSite2()
	}()
	go func() {
		defer wg.Done()
		priceChannel <- fetcher.FetchPriceFromSite3()
	}()

	wg.Wait()
	close(priceChannel)
	fmt.Printf("Tempo total: %s\n", time.Since(start))
}
