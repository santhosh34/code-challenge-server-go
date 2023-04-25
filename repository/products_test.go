package repository

import (
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
)

func TestPopulateTemperature(t *testing.T) {

	t.Log(ProductTypes)
	product1 := ProductTypes[0]
	returnedProduct, err := product1.PopulateTemperature(product1.Id)
	t.Log(returnedProduct)
	t.Log(err)

}

func serverMock() *httptest.Server {
	handler := http.NewServeMux()
	handler.HandleFunc("/products", productsMock)

	srv := httptest.NewServer(handler)

	return srv
}

func productsMock(w http.ResponseWriter, r *http.Request) {
	testProduct := new(Product)
	testProduct.Id = "1"
	testProduct.Temperature = 6
	_, _ = w.Write([]byte("mock server response"))

}

func TestPopulateTempRangeStatusForHighTemp(t *testing.T) {
	testProduct := new(Product)
	testProduct.Id = "1"
	testProduct.MinTemperature = 5
	testProduct.MaxTemperature = 10
	testProduct.Temperature = 15
	dummyReturnProduct, err := testProduct.populateTempRangeStatus()
	expectedHigh := regexp.MustCompile("too high")
	if !expectedHigh.MatchString(dummyReturnProduct.TempRangeStatus) || err != nil {
		t.Fatalf(`Test failed as expected: "%v"  is different from actual: "%v". Error: %v `, dummyReturnProduct.TempRangeStatus, expectedHigh, err)
	}
}

func TestPopulateTempRangeStatusForLowTemp(t *testing.T) {
	testProduct := new(Product)
	testProduct.Id = "1"
	testProduct.MinTemperature = 5
	testProduct.MaxTemperature = 10
	testProduct.Temperature = 2
	dummyReturnProduct, err := testProduct.populateTempRangeStatus()
	expectedLow := regexp.MustCompile("too low")
	if !expectedLow.MatchString(dummyReturnProduct.TempRangeStatus) || err != nil {
		t.Fatalf(`Test failed as expected: "%v"  is different from actual: "%v". Error: %v `, dummyReturnProduct.TempRangeStatus, expectedLow, err)
	}
}

func TestPopulateTempRangeStatusForWithInRangeScenario(t *testing.T) {

	testProduct := new(Product)
	testProduct.Id = "1"
	testProduct.MinTemperature = 5
	testProduct.MaxTemperature = 10
	testProduct.Temperature = 6
	dummyReturnProduct, err := testProduct.populateTempRangeStatus()
	expectedWithInNormal := regexp.MustCompile("all good")
	if !expectedWithInNormal.MatchString(dummyReturnProduct.TempRangeStatus) || err != nil {
		t.Fatalf(`Test failed as expected: "%v"  is different from actual: "%v". Error: %v `, dummyReturnProduct.TempRangeStatus, expectedWithInNormal, err)
	}
}
