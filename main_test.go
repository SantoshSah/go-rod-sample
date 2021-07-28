package main

import (
	"log"
	"os"
	"testing"

	"github.com/SantoshSah/go-rod-sample/types"
	"github.com/joho/godotenv"
)

func TestGetProducts(t *testing.T) {
	// Load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	products := getProducts(os.Getenv("PRODUCT_TYPE_URL"))

	gotProductCount := len(products)
	wantProductCount := 14 // Total number of products for all product types

	if gotProductCount != wantProductCount {
		t.Errorf("got %d, wanted %d", gotProductCount, wantProductCount)
	}
}

func TestGetProductsForType(t *testing.T) {
	getProductsForTypeTests := []types.GetProductsForTypeTest{
		{
			ArgProductLink:         "https://nepallife.com.np/en/products/endowment",
			ArgProductType:         "Endowment",
			ExpectedProductsNumber: 5,
		},
		{
			ArgProductLink:         "https://nepallife.com.np/en/products/anticipated",
			ArgProductType:         "Anticipated Insurance",
			ExpectedProductsNumber: 4,
		},
		{
			ArgProductLink:         "https://nepallife.com.np/en/products/whole-life",
			ArgProductType:         "Whole Life Insurance",
			ExpectedProductsNumber: 2,
		},
		{
			ArgProductLink:         "https://nepallife.com.np/en/products/term",
			ArgProductType:         "Term Insurance",
			ExpectedProductsNumber: 3,
		},
		{
			ArgProductLink:         "https://nepallife.com.np/en/products/micro-insurance",
			ArgProductType:         "Micro Insurance",
			ExpectedProductsNumber: 1,
		},
	}

	productChan := make(chan []types.InsuranceProduct, len(getProductsForTypeTests))

	for _, test := range getProductsForTypeTests {
		go getProductsForType(test.ArgProductLink, test.ArgProductType, productChan)
		result := <-productChan
		gotProductCount := len(result)

		if gotProductCount != test.ExpectedProductsNumber {
			t.Errorf("Product type - %s: got %d, wanted %d", test.ArgProductType, gotProductCount, test.ExpectedProductsNumber)
		}
	}
}
