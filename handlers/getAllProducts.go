package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"server/repository"
	"strconv"
	"time"
)

func GetAllProductsWithTempStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(getAllProducts())

}

func getAllProducts() []repository.Product {
	products := repository.ProductTypes

	start := time.Now()
	defer func() {
		fmt.Println("Execution Time: ", time.Since(start).Seconds())
	}()

	ch := make(chan repository.Product)

	for index, specProduct := range products {
		go func(index int, specProduct repository.Product) {
			product, err := specProduct.PopulateTemperature(strconv.Itoa(index + 1))
			if err != nil {
				log.Println("error getting product", err)
			}
			ch <- product
		}(index, specProduct)
	}

	var productsnew []repository.Product
	for i := 0; i < len(products); i++ {
		prod := <-ch
		productsnew = append(productsnew, prod)
		fmt.Println(productsnew[i])
	}
	return productsnew
}
