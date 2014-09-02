package go_gilt_api

const (
	productsUrl    = "products/"
	categoriesJson = "categories.json"
	sizesJson = "sizes.json"
)

// Get product details for a given product id
// see https://dev.gilt.com/documentation/resources.html#toc_167
func (a GiltApi) GetProductDetails(productId int) (productDetail ProductDetail, err error) {
	return a.GetProductDetailsFromUrl(baseUrl + productsUrl + string(productId) + detailJson)
}

// Get product details for a given product url
// see https://dev.gilt.com/documentation/resources.html#toc_167
func (a GiltApi) GetProductDetailsFromUrl(url string) (productDetail ProductDetail, err error) {
	response_ch := make(chan response)
	a.queryQueue <- query{url, &productDetail, response_ch}
	return productDetail, (<-response_ch).err
}

type categoriesResponse struct {
	Categories []string `json:"categories"`
}

// Get a list of all distinct product categories
// more info at https://dev.gilt.com/documentation/resources.html#toc_168
func (a GiltApi) GetProductCategories() (categories []string, err error) {
	response_ch := make(chan response)
	categoriesResponseJson := new(categoriesResponse)
	a.queryQueue <- query{baseUrl + productsUrl + categoriesJson, &categoriesResponseJson, response_ch}
	return categoriesResponseJson.Categories, (<-response_ch).err
}

// Get a list of all distinct size keys
// more info at https://dev.gilt.com/documentation/resources.html#toc_169
func (a GiltApi) GetProductSizes() (sizes Sizes, err error) {
	response_ch := make(chan response)
	a.queryQueue <- query{baseUrl + productsUrl + sizesJson, &sizes, response_ch}
	return sizes, (<-response_ch).err
}
