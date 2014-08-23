package go_gilt_api

const (
	SalesUrl = "sales"
)

type Sales struct {
	Sales []Sale `json:"sales"`
}

func (a GiltApi) GetSalesActive() (activeSales Sales, err error) {
	response_ch := make(chan response)
	a.queryQueue <- query{BaseUrl + SalesUrl + "/active.json", &activeSales, response_ch}
	return activeSales, (<-response_ch).err
}
