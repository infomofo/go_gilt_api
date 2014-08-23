package go_gilt_api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	BaseUrl = "https://api.gilt.com/v1/"
	Women = "women"
	Men = "men"
	Kids = "kids"
	Home = "home"
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

const DEFAULT_DELAY = 0 * time.Second
const DEFAULT_CAPACITY = 5

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

// apiGet issues a GET request to the Twitter API and decodes the response JSON to data.
func (c GiltApi) apiGet(urlStr string, data interface{}) error {
	resp, err := http.Get(c.addApiKey(urlStr))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return decodeResponse(resp, data)
}

// decodeResponse decodes the JSON response from the Twitter API.
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
