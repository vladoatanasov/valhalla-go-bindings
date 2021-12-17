package valhalla

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type Client struct {
	Endpoint   string
	httpClient *http.Client
}

func New(endpoint string) *Client {
	var t = http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = 100
	t.MaxConnsPerHost = 100
	t.MaxIdleConnsPerHost = 100
	var httpClient = &http.Client{
		Timeout:   time.Second * 10,
		Transport: t,
	}
	return &Client{
		Endpoint:   endpoint,
		httpClient: httpClient,
	}
}

func (c *Client) request(method, resource string, body io.Reader) ([]byte, error) {

	req, err := http.NewRequest(method, c.Endpoint+"/"+resource, body)
	if err != nil {
		return nil, err
	}

	// todo: hide debug information once I'm done
	// dump, _ := httputil.DumpRequest(req, true)
	// log.Println(string(dump))

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	// dump, _ := httputil.DumpResponse(res, true)
	// log.Println(string(dump))

	r, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode >= 400 {
		result := ValhallaError{}
		err = json.Unmarshal(r, &result)
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("%+v", result)
	}

	return r, err

}
