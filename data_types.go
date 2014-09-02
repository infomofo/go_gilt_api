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

// Additional details about a product
// more information is available at https://dev.gilt.com/documentation/data_structures.html#product_content_objects
type ProductContent struct {
	Description      string `json:"description"`
	Material         string `json:"material`
	Origin           string `json:"origin"`
	FitNotes         string `json:"fit_notes"`
	CareInstructions string `json:"care_instructions"`
}

type SkuAttribute struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// A SKU or "stock keeping unit" represents the most specific set of distinguishable attributes for a product that can be purchased
// more information is available at https://dev.gilt.com/documentation/data_structures.html#sku_objects
type Sku struct {
	Id                int             `json:"id"`
	InventoryStatus   InventoryStatus `json:"inventory_status"`
	MsrpPrice         string          `json:"msrp_price"`
	SalePrice         string          `json:"sale_price"`
	ShippingSurcharge string          `json:"shipping_surcharge"`
	Attributes        []SkuAttribute  `json:"attributes"`
}

// Information on a product
// more information is available at https://dev.gilt.com/documentation/data_structures.html#product_detail_object
type ProductDetail struct {
	Name       string                `json:"name"`
	Product    string                `json:"product"`
	Brand      string                `json:"brand"`
	Content    ProductContent        `json:"content"`
	ImageUrls  map[string][]ImageUrl `json:"image_urls"`
	Skus       []Sku                 `json:"skus"`
	Categories []string              `json:"categories"`
}

// a sale detail object extends the sale list object and provides additional fields (ends, products)
// more information is available at https://dev.gilt.com/documentation/data_structures.html#sale_detail_object
type SaleDetailObject struct {
	*SaleListObject
	Ends     time.Time `json:"ends"`
	Products []string  `json:"products"`
}
