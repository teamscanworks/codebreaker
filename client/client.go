package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Client is a simple wrapper for performing http requests to the contracts registry and
// parsing the corresponding response
// todo: add query capabilities to the wasm module of the supported chains
type Client struct {
	registryUrl string
}

func New(registryUrl string) (*Client, error) {
	_, err := url.Parse(registryUrl)
	if err != nil {
		return nil, err
	}
	return &Client{registryUrl: registryUrl}, nil
}

func (c Client) Chains() ([]string, error) {
	bz, err := c.get(fmt.Sprintf("%s/v1/chains", c.registryUrl))
	if err != nil {
		return nil, err
	}
	var chains []string
	err = json.Unmarshal(bz, &chains)
	if err != nil {
		return nil, err
	}

	return chains, nil
}

// todo: write queries for contract info calls

func (c Client) get(query string) ([]byte, error) {
	resp, err := http.Get(query)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == http.StatusNotFound {
		return nil, errors.New("resource not found")
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bodyBytes, nil
}
