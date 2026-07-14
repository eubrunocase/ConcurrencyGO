package processor

import (
	"C2-Golang/internal/models"
	"fmt"
)

// func ShowPricesAVG(priceChannel <-chan float64, done chan<- bool) {
// 	var totalPrice float64
// 	countPrinces := 0.0
// 	for price := range priceChannel {
// 		totalPrice += price
// 		countPrinces++
// 		avgPrice := totalPrice / countPrinces
// 		fmt.Printf("Preço recebido: R$ %.2f\n | Preço médio até agora: R$ %.2f\n", price, avgPrice)
// 	}
// 	done <- true
// }


func ShowPricesAVG(priceChannel <-chan models.PriceDetails, done chan<- bool) {
	var totalPrice float64
	countPrinces := 0.0
	for priceDetails := range priceChannel {
		totalPrice += priceDetails.Value
		countPrinces++
		avgPrice := totalPrice / countPrinces
		fmt.Printf("Preço recebido: R$ %.2f | Loja: %s | Timestamp: %s\n | Preço médio até agora: R$ %.2f\n", priceDetails.Value, priceDetails.StoreName, priceDetails.TimeStamp.Format("2006-01-02 15:04:05"), avgPrice)
	}
	done <- true
}