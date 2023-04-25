package repository

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Product struct {
	Id              string `json:"id"`
	Name            string `json:"name"`
	Temperature     int    `json:"temperature"`
	MinTemperature  int    `json:"minTemperature"`
	MaxTemperature  int    `json:"maxTemperature"`
	TempRangeStatus string `json:"tempRangeStatus"`
}

var ProductTypes []Product

func init() {

	ProductTypes = []Product{
		{
			Id:              "1",
			Name:            "Pilsner",
			Temperature:     0,
			MinTemperature:  4,
			MaxTemperature:  6,
			TempRangeStatus: "normal",
		},
		{
			Id:              "2",
			Name:            "IPA",
			Temperature:     0,
			MinTemperature:  5,
			MaxTemperature:  6,
			TempRangeStatus: "normal",
		},
		{
			Id:              "3",
			Name:            "Lager",
			Temperature:     0,
			MinTemperature:  4,
			MaxTemperature:  7,
			TempRangeStatus: "normal",
		},
		{
			Id:              "4",
			Name:            "Stout",
			Temperature:     0,
			MinTemperature:  6,
			MaxTemperature:  8,
			TempRangeStatus: "normal",
		},
		{
			Id:              "5",
			Name:            "Witbier",
			Temperature:     0,
			MinTemperature:  3,
			MaxTemperature:  5,
			TempRangeStatus: "normal",
		},
		{
			Id:              "6",
			Name:            "Pale Ale",
			Temperature:     0,
			MinTemperature:  4,
			MaxTemperature:  6,
			TempRangeStatus: "normal",
		},
	}
}

// Calls external aws lambda ervice for each typoe of Beer Product
// This can be further optmized by removing param id.
func (currentProduct *Product) PopulateTemperature(id string) (currentProductRef Product, err error) {

	response, err := http.Get("https://hasydbj5c4gpa2oozfpjpc677a0hxuob.lambda-url.ap-southeast-2.on.aws/sensor/" + id)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	if err := json.Unmarshal(body, &currentProduct); err != nil {
		panic(err)
	}
	currentProduct.populateTempRangeStatus()
	return *currentProduct, nil
}

func (currentProduct *Product) populateTempRangeStatus() (currentProductRef Product, err error) {
	if (currentProduct).Temperature < (currentProduct).MinTemperature {
		(currentProduct).TempRangeStatus = "too low"
	} else if (currentProduct).Temperature > (currentProduct).MaxTemperature {
		(currentProduct).TempRangeStatus = "too high"
	} else {
		(currentProduct).TempRangeStatus = "all good"
	}
	return *currentProduct, nil
}
