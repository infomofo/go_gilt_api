package go_gilt_api

const (
	productsUrl = "products/"
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
