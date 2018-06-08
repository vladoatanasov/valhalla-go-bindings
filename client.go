package valhalla

import (
	"io"
	"io/ioutil"
	"net/http"
)

type Client struct {
	Endpoint string
}

func New(endpoint string) *Client {
	return &Client{
		Endpoint: endpoint,
	}
}

func (c *Client) request(method, resource string, body io.Reader) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, c.Endpoint+"/"+resource, body)
	if err != nil {
		return nil, err
	}

	// todo: hide debug information once I'm done
	// dump, _ := httputil.DumpRequest(req, true)
	// log.Println(string(dump))

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	// dump, _ := httputil.DumpResponse(res, true)
	// log.Println(string(dump))

	return ioutil.ReadAll(res.Body)

}
