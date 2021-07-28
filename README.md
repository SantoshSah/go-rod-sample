1) add .env file
2) add below content to .env file
PRODUCT_TYPE_URL=https://nepallife.com.np/en/products

# Test commands
$ go test -timeout 60s -run ^TestGetProducts$ github.com/SantoshSah/go-rod-sample

$ go test -timeout 200s -run ^TestGetProductsForType$ github.com/SantoshSah/go-rod-sample