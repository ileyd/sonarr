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

func (sc *SonarrClient) SeriesLookup(term string) (SeriesLookupResponse, error) {

	lookupUrl := *sc.address

	parameters := url.Values{}
	parameters.Add("term", term)

	lookupUrl.RawQuery = parameters.Encode()
	lookupUrl.Path += "series/lookup"

	req, err := http.NewRequest("GET", lookupUrl.String(), nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("X-Api-Key", sc.apiKey)

	response, err := sc.HttpClient.Do(req)

	if err != nil {
		return nil, err
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return nil, errors.New(fmt.Sprintf("Status code %v", response.StatusCode))
	}

	var rv SeriesLookupResponse

	err = json.NewDecoder(response.Body).Decode(&rv)

	if err != nil {
		return nil, err
	}

	return rv, nil
}