go_gilt_api
==================

[![GoDoc](https://godoc.org/github.com/infomofo/go_gilt_api?status.png)](https://godoc.org/github.com/infomofo/go_gilt_api)

go_gilt_api is a simple, transparent Go package for accessing the Gilt API. 

Successful API queries return native Go structs that can be used immediately, with no need for type assertions.


Examples
-------------

### Authentication

You will need a gilt api key which you can get at [api.gilt.com](http://api.gilt.com) and instantiate your client as:

````go
api := go_gilt_api.NewGiltApi("your-api-key")
````

### Endpoints

go_gilt_api will implement the endpoints defined in the [api.gilt.com resources](https://dev.gilt.com/documentation/resources.html) For clarity, in most 
cases, the function name is simply the name of the HTTP method and the endpoint 
(e.g., the endpoint `GET /sales/active` is provided by the function `GetSalesActive`).

### Example

````go
	activeSales, err := api.GetSalesActive()
	if err != nil {
		fmt.Errorf("GetSearch yielded error %s", err.Error())
		panic(err)
	}
	fmt.Println(activeSales)
````