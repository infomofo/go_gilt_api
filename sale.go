package go_gilt_api

import "time"

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
