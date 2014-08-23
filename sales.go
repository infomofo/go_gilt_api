package go_gilt_api

const (
	SalesUrl = "sales"
)

type SalesResponse struct {
	Sales []Sale `json:"sales"`
}

// Returns a list of active sales on gilt.com
func (a GiltApi) GetSalesActive() (sales []Sale, err error) {
	response_ch := make(chan response)
	activeSales := new(SalesResponse)
	a.queryQueue <- query{BaseUrl + SalesUrl + "/active.json", &activeSales, response_ch}
	return activeSales.Sales, (<-response_ch).err
}
