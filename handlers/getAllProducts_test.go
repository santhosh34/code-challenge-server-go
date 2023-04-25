package handlers

import (
	"testing"
)

func TestGetAllProducts(t *testing.T) {
	arrayOfProducts := getAllProducts()
	got := len(arrayOfProducts)
	want := 6

	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, arrayOfProducts)
	}

	t.Log(arrayOfProducts)
}
