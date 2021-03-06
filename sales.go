package go_gilt_api

const (
	salesUrl     = "sales"
	activeJson   = "/active.json"
	upcomingJson = "/upcoming.json"
)

type salesResponse struct {
	Sales []SaleListObject `json:"sales"`
}

// Returns a list of active sales on gilt.com
// for more info see: https://dev.gilt.com/documentation/resources.html#toc_163
func (a GiltApi) GetSalesActive() (sales []SaleListObject, err error) {
	response_ch := make(chan response)
	activeSales := new(salesResponse)
	a.queryQueue <- query{baseUrl + salesUrl + activeJson, &activeSales, response_ch}
	return activeSales.Sales, (<-response_ch).err
}

// Returns the list of active sales for a particular store
func (a GiltApi) GetSalesActiveInStore(store Store) (sales []SaleListObject, err error) {
	response_ch := make(chan response)
	activeSales := new(salesResponse)
	a.queryQueue <- query{baseUrl + salesUrl + "/" + string(store) + activeJson, &activeSales, response_ch}
	return activeSales.Sales, (<-response_ch).err
}

// Returns a list of upcoming sales sales on gilt.com
// For more info see: https://dev.gilt.com/documentation/resources.html#toc_164
func (a GiltApi) GetSalesUpcoming() (sales []SaleListObject, err error) {
	response_ch := make(chan response)
	upcomingSales := new(salesResponse)
	a.queryQueue <- query{baseUrl + salesUrl + upcomingJson, &upcomingSales, response_ch}
	return upcomingSales.Sales, (<-response_ch).err
}

// Returns the list of upcoming sales for a particular store
func (a GiltApi) GetSalesUpcomingInStore(store Store) (sales []SaleListObject, err error) {
	response_ch := make(chan response)
	upcomingSales := new(salesResponse)
	a.queryQueue <- query{baseUrl + salesUrl + "/" + string(store) + upcomingJson, &upcomingSales, response_ch}
	return upcomingSales.Sales, (<-response_ch).err
}

// Retrieves detailed sale information for a given sale key and store
// for more info see: https://dev.gilt.com/documentation/resources.html#toc_165
func (a GiltApi) GetSaleDetail(store Store, saleKey string) (saleDetails SaleDetailObject, err error) {
	response_ch := make(chan response)
	a.queryQueue <- query{baseUrl + salesUrl + "/" + string(store) + "/" + saleKey + detailJson, &saleDetails, response_ch}
	return saleDetails, (<-response_ch).err
}

// Retrieves detailed sale details for a given sale list object
func (a GiltApi) GetSaleDetailFromListObject(saleList SaleListObject) (saleDetails SaleDetailObject, err error) {
	return a.GetSaleDetail(saleList.Store, saleList.SaleKey)
}
