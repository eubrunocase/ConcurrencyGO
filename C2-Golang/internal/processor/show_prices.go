package processor

import "fmt"

func ShowPricesAVG(priceChannel chan float64) {
	var totalPrice float64
	countPrinces := 0.0
	for price := range priceChannel {
		totalPrice += price
		countPrinces++
		avgPrice := totalPrice / countPrinces
		fmt.Printf("Preço recebido: R$ %.2f\n | Preço médio até agora: R$ %.2f\n", price, avgPrice)
	}
}
