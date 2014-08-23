package go_gilt_api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const ()

type ApiError struct {
	StatusCode int
	Header     http.Header
	Body       string
	Decoded    GiltErrorResponse
	URL        *url.URL
}

func newApiError(resp *http.Response) *ApiError {
	// TODO don't ignore this error
	// TODO don't use ReadAll
	p, _ := ioutil.ReadAll(resp.Body)

	var GiltErrorResp GiltErrorResponse
	_ = json.Unmarshal(p, &GiltErrorResp)
	return &ApiError{
		StatusCode: resp.StatusCode,
		Header:     resp.Header,
		Body:       string(p),
		Decoded:    GiltErrorResp,
		URL:        resp.Request.URL,
	}
}

// ApiError supports the error interface
func (aerr ApiError) Error() string {
	return fmt.Sprintf("Get %s returned status %d, %s", aerr.URL, aerr.StatusCode, aerr.Body)
}

// Check to see if an error is a Rate Limiting error. If so, find the next available window in the header.
// Use like so:
//
//    if aerr, ok := err.(*ApiError); ok {
//  	  if isRateLimitError, nextWindow := aerr.RateLimitCheck(); isRateLimitError {
//       	<-time.After(nextWindow.Sub(time.Now()))
//  	  }
//    }
//

//GiltErrorResponse has an array of Twitter error messages
//It satisfies the "error" interface
//For the most part, Twitter seems to return only a single error message
//Currently, we assume that this always contains exactly one error message
type GiltErrorResponse struct {
	Errors []GiltError `json:"errors"`
}

func (tr GiltErrorResponse) First() error {
	return tr.Errors[0]
}

func (tr GiltErrorResponse) Error() string {
	return tr.Errors[0].Message
}

//GiltError represents a single Twitter error messages/code pair
type GiltError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (te GiltError) Error() string {
	return te.Message
}
