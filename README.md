# Description
This is a sample project for scraping Insurance Product Type and list of insurance products of a particular product type from website https://nepallife.com.np/en/home

Website: https://nepallife.com.np/en/home
Product Page: https://nepallife.com.np/en/products

# Setup
1) add .env file
2) add below content to .env file
PRODUCT_TYPE_URL=https://nepallife.com.np/en/products

# Command to run
$ go run .

# Test commands
$ go test -timeout 60s -run ^TestGetProducts$ github.com/SantoshSah/go-rod-sample

$ go test -timeout 60s -run ^TestGetProductsForType$ github.com/SantoshSah/go-rod-sample