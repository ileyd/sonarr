package go_sonarr

import (
	"net/http"
	"net/url"
	"encoding/json"
	"fmt"
	"errors"
	"strings"
)

type SonarrClient struct {
	address *url.URL
	apiKey string
	HttpClient *http.Client
}

func NewSonarrClient(address string, apiKey string) (*SonarrClient, error) {

	if address == "" {
		return nil, errors.New("No address specified")
	}

	addressUrl, err := url.Parse(address)

	path := addressUrl.Path
	//
	if !strings.HasSuffix(path, "/") {
		path += "/"
	}

	if !strings.HasSuffix(path, "api/") {
		path += "api/"
	}

	addressUrl.Path = path

	if err != nil {
		return nil, err
	}

	return &SonarrClient{
		address:addressUrl,
		apiKey:apiKey,
		HttpClient:http.DefaultClient,
	}, nil
}

func (sc *SonarrClient) DoRequest(action, path string, params map[string]string, reqData, resData interface{}) error {
	lookupUrl := *sc.address

	parameters := url.Values{}

	if params != nil {
		for k, v := range params {
			parameters.Add(k, v)
		}
	}

	lookupUrl.RawQuery = parameters.Encode()
	lookupUrl.Path += path

	req, err := http.NewRequest(action, lookupUrl.String(), nil)

	if err != nil {
		return err
	}

	req.Header.Add("X-Api-Key", sc.apiKey)

	response, err := sc.HttpClient.Do(req)

	if err != nil {
		return err
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return errors.New(fmt.Sprintf("Status code %v", response.StatusCode))
	}

	err = json.NewDecoder(response.Body).Decode(resData)

	if err != nil {
		return err
	}

	return nil
}