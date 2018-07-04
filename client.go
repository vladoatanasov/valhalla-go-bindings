package valhalla

import (
	"encoding/json"
	"fmt"
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
