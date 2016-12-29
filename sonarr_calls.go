package go_sonarr

import "fmt"

func (sc *SonarrClient) SeriesLookup(term string) (SeriesLookupGetResponse, error) {

	var rv *SeriesLookupGetResponse

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

	//copy what we need
	spr_in := &SeriesPostResponse{
		TvdbID: tvdbId,
		Title: slr[0].Title,
		QualityProfileID: slr[0].QualityProfileID,
		TitleSlug: slr[0].TitleSlug,
		Images: slr[0].Images,
	}

	var spr_out SeriesPostResponse

	//then add it
	err = sc.DoRequest("POST", "series", nil, spr_in, &spr_out)

	if err != nil {
		return SeriesPostResponse{}, err
	}

	return spr_out, nil
}