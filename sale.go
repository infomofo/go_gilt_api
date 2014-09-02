package go_gilt_api

import "time"

// From https://dev.gilt.com/documentation/data_structures.html#image_url_objects
// Sale and product methods can both contain structures that expose image URLs. Images are provided at multiple
// resolutions and a particular resolution can also have more than one image associated with it. An image_url object
// consists of a set of key/value pairs with the key being the image resolution WIDTHxHEIGHT and the value being an
// array of image objects.
type ImageUrl struct {
	Url    string `json:"url"`
	Width  int    `'json:"width"`
	Height int    `'json:"height"`
}

// a sale list object provides basic information about a sale.
// more information is available at https://dev.gilt.com/documentation/data_structures.html#sale_list_object
type SaleListObject struct {
	Name        string                `json:"name"`
	Sale        string                `json:"sale"`
	SaleKey     string                `json:"sale_key"`
	Store       Store                 `json:"store"`
	Description string                `json:"description"`
	SaleUrl     string                `json:"sale_url"`
	Begins      time.Time             `json:"begins"`
	ImageUrls   map[string][]ImageUrl `json:"image_urls"`
}

// Information on a product
// more information is available at https://dev.gilt.com/documentation/data_structures.html#product_detail_object
type ProductDetail struct {
}

// a sale detail object extends the sale list object and provides additional fields (ends, products)
// more information is available at https://dev.gilt.com/documentation/data_structures.html#sale_detail_object
type SaleDetailObject struct {
	*SaleListObject
	Ends     time.Time `json:"ends"`
	Products []string  `json:"products"`
}
