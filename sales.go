package go_gilt_api

const (
	salesUrl = "sales"
)

type salesResponse struct {
	Sales []SaleListObject `json:"sales"`
}

// Returns a list of active sales on gilt.com
func (a GiltApi) GetSalesActive() (sales []SaleListObject, err error) {
	response_ch := make(chan response)
	activeSales := new(salesResponse)
	a.queryQueue <- query{baseUrl + salesUrl + "/active.json", &activeSales, response_ch}
	return activeSales.Sales, (<-response_ch).err
}
