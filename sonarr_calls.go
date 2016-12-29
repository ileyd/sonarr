package go_sonarr

import "fmt"

func (sc *SonarrClient) SeriesLookup(term string) (SeriesLookupGetResponse, error) {

	rv := &SeriesLookupGetResponse{}

	err := sc.DoRequest("GET", "series/lookup", map[string]string{"term":term}, nil, rv)

	if err != nil {
		return nil, err
	}

	return *rv, nil
}


func (sc *SonarrClient) CreateSeries(tvdbId int) (SeriesPostResponse, error) {
	//first lookup data
	slr, err := sc.SeriesLookup(fmt.Sprintf("tvdb:%v", tvdbId))

	if err != nil {
		return SeriesPostResponse{}, err
	}

	var spr_out SeriesPostResponse

	//then pass the whole lot back
	err = sc.DoRequest("POST", "series", nil, slr[0], &spr_out)

	if err != nil {
		return SeriesPostResponse{}, err
	}

	return spr_out, nil
}