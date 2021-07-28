package types

type InsuranceProduct struct {
	Type        string `json:"type"`
	Name        string `json:"name"`
	URL         string `json:"url"`
	Description string `json:"description"`
}

type GetProductsForTypeTest struct {
	ArgProductLink         string
	ArgProductType         string
	ExpectedProductsNumber int
}
