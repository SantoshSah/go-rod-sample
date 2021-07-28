package main

import (
	"fmt"
	"log"
	"os"

	"github.com/SantoshSah/go-rod-sample/types"
	"github.com/go-rod/rod"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	products := getProducts(os.Getenv("PRODUCT_TYPE_URL"))

	// Iterate over final slice to print list of insurance products
	for _, product := range products {
		// fmt.Printf("Type:%s, Name: %s, URL: %s, Desc: %s \n", product.Type, product.Name, product.URL, product.Description)
		fmt.Printf("Type:%s, Name: %s \n", product.Type, product.Name)
	}
}

func getProducts(productTypeURL string) (products []types.InsuranceProduct) {
	// Connect to product type page
	productTypePage := rod.New().MustConnect().MustPage(productTypeURL)

	// Though connection to page will be closed implicitly by rod, let's close it explicitly
	// to maintain good programming practice to release resources
	defer productTypePage.Close()

	// Find container div that contains list of product types
	productTypeDiv := productTypePage.MustElement("#main > div.container > div")

	// Find list of prouct type
	productTypesDiv := productTypeDiv.MustElements("div.category-list-item")

	// Make buffered chanel of lengh equal to number of product type
	productChan := make(chan []types.InsuranceProduct, len(productTypesDiv))

	// Make slice to store all products from all channels
	products = make([]types.InsuranceProduct, 0)

	// Iterate over product types to find products of particular product type
	for _, productTypeDiv := range productTypesDiv {
		// Find product type and product link
		productTypeLink := productTypeDiv.MustElement("a")
		productType := fmt.Sprintf("%s", productTypeLink.MustText())
		productLink := fmt.Sprintf("%s", productTypeLink.MustProperty("href"))

		// Execute go routines to get list of products of particular product type
		go getProductsForType(productLink, productType, productChan)

		// Get slices of products
		result := <-productChan

		// Append products from channel to main slice to accumulate products from all channels
		products = append(products, result...)
	}

	return
}

// getProductsForType method connect to product type page, extract list of product of that type
// and send slices of products to main go routine
func getProductsForType(productLink, productType string, productChan chan<- []types.InsuranceProduct) {
	// Connect to product page
	productPage := rod.New().MustConnect().MustPage(productLink)

	// Defer product page close
	defer productPage.Close()

	// Find product div
	productDiv := productPage.MustElement("#main > div.page.Products > div.container > div")

	// Find list of products
	productDivs := productDiv.MustElements("div.align-items-stretch")

	// Make slice of products type to hold all products of that type
	insurances := make([]types.InsuranceProduct, 0)

	// Iterate over list of products to extract product name, URL and description if available
	for _, productDiv := range productDivs {
		// Find product link
		productLink := productDiv.MustElement("div > div.card-body.d-flex.flex-column.small > div > a:nth-child(1)")

		// Find product name and URL
		productName := fmt.Sprintf("%s", productLink.MustText())
		productURL := fmt.Sprintf("%s", productLink.MustProperty("href"))
		productDesc := ""

		// Description may not be available for all product. So find description if available
		descriptionDiv := productLink.MustNext()
		descriptionP, err := descriptionDiv.Element("p")
		if err == nil {
			productDesc = fmt.Sprintf("%s", descriptionP.MustText())
		}

		// Append product to main slice
		insurances = append(insurances, types.InsuranceProduct{
			Type:        productType,
			Name:        productName,
			URL:         productURL,
			Description: productDesc,
		})
	}

	// Send slice of products to main go routine over channel
	productChan <- insurances
}
