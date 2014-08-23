package go_gilt_api

//Package go_gilt_api provides structs and functions for accessing version 1
//of the Gilt API.
//
//Successful API queries return native Go structs that can be used immediately,
//with no need for type assertions.
//
//Authentication
//
//Obtain an api key on api.gilt.com and instantiate a client with:
//
//  api := go_gilt_api.NewGiltApi("your-api-key")
//
//
//Queries
//
//Executing queries on an authenticated GiltApi struct is simple.
//
//  activeSales, err := api.GetSalesActive()
//  for _ , sale := range searchResult {
//      fmt.Print(sale.Name)
//  }
//
//Endpoints
//
//go_gilt_api implements some of the endpoints defined in the Gilt API documentation: https://dev.gilt.com/documentation/resources.html.
//For clarity, in most cases, the function name is simply the name of the HTTP method and the endpoint (e.g., the endpoint `GET /sales/active` is provided by the function `GetSalesActive`).
//
//More detailed information about the behavior of each particular endpoint can be found at the official Gilt API documentation.

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	BaseUrl = "https://api.gilt.com/v1/"
	Women   = "women"
	Men     = "men"
	Kids    = "kids"
	Home    = "home"
)

type GiltApi struct {
	apiKey     string
	queryQueue chan query
	HttpClient *http.Client
}

type query struct {
	url         string
	data        interface{}
	response_ch chan response
}

type response struct {
	data interface{}
	err  error
}

//NewGiltApi takes an application-specific api token returns a GiltApi struct for that application.
//The GiltApi struct can be used for accessing any of the endpoints available.
func NewGiltApi(apiKey string) *GiltApi {
	//TODO figure out how much to buffer this channel
	//A non-buffered channel will cause blocking when multiple queries are made at the same time
	queue := make(chan query)
	c := &GiltApi{apiKey, queue, http.DefaultClient}
	go c.processQueries()
	return c
}

func (c GiltApi) addApiKey(urlStr string) string {
	return urlStr + "?apikey=" + c.apiKey
}

// apiGet issues a GET request to the Gilt API and decodes the response JSON to data.
func (c GiltApi) apiGet(urlStr string, data interface{}) error {
	resp, err := http.Get(c.addApiKey(urlStr))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return decodeResponse(resp, data)
}

// decodeResponse decodes the JSON response from the Gilt API.
func decodeResponse(resp *http.Response, data interface{}) error {
	if resp.StatusCode != 200 {
		return newApiError(resp)
	}
	jsonDebug := json.NewDecoder(resp.Body)
	return jsonDebug.Decode(data)
}

func NewApiError(resp *http.Response) *ApiError {
	body, _ := ioutil.ReadAll(resp.Body)

	return &ApiError{
		StatusCode: resp.StatusCode,
		Header:     resp.Header,
		Body:       string(body),
		URL:        resp.Request.URL,
	}
}

func (c *GiltApi) processQueries() {
	for q := range c.queryQueue {
		url := q.url
		data := q.data //This is where the actual response will be written

		response_ch := q.response_ch

		err := c.apiGet(url, data)

		response_ch <- response{data, err}
	}
}

// Close query queue
func (c *GiltApi) Close() {
	close(c.queryQueue)
}
