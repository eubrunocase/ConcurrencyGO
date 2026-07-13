package main

import (
	"C2-Golang/internal/fetcher"
	"C2-Golang/internal/processor"
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()

	priceChannel := make(chan float64)
	var showWg sync.WaitGroup
	showWg.Add(1)

	go func() {
		defer showWg.Done()
		processor.ShowPricesAVG(priceChannel)
	}()

	go fetcher.FectchPrices(priceChannel)

	showWg.Wait()
	fmt.Printf("Tempo total: %s\n", time.Since(start))
}
