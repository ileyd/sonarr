package go_sonarr

import "time"

type SeriesPostResponse struct {
	Title string `json:"title"`
	SeasonCount int `json:"seasonCount"`
	EpisodeCount int `json:"episodeCount"`
	EpisodeFileCount int `json:"episodeFileCount"`
	Status string `json:"status"`
	Overview string `json:"overview"`
	NextAiring time.Time `json:"nextAiring"`
	Network string `json:"network"`
	AirTime string `json:"airTime"`
	Images []struct {
		CoverType string `json:"coverType"`
		URL string `json:"url"`
	} `json:"images"`
	Seasons []struct {
		SeasonNumber int `json:"seasonNumber"`
		Monitored bool `json:"monitored"`
	} `json:"seasons"`
	Year int `json:"year"`
	Path string `json:"path"`
	QualityProfileID int `json:"qualityProfileId"`
	SeasonFolder bool `json:"seasonFolder"`
	Monitored bool `json:"monitored"`
	UseSceneNumbering bool `json:"useSceneNumbering"`
	Runtime int `json:"runtime"`
	TvdbID int `json:"tvdbId"`
	TvRageID int `json:"tvRageId"`
	FirstAired time.Time `json:"firstAired"`
	LastInfoSync time.Time `json:"lastInfoSync"`
	SeriesType string `json:"seriesType"`
	CleanTitle string `json:"cleanTitle"`
	ImdbID string `json:"imdbId"`
	TitleSlug string `json:"titleSlug"`
	ID int `json:"id"`
}

type SeriesLookupGetResponse []struct {
	Title string `json:"title"`
	SeasonCount int `json:"seasonCount"`
	EpisodeCount int `json:"episodeCount"`
	EpisodeFileCount int `json:"episodeFileCount"`
	Status string `json:"status"`
	Overview string `json:"overview"`
	Network string `json:"network"`
	Images []struct {
		CoverType string `json:"coverType"`
		URL string `json:"url"`
	} `json:"images"`
	RemotePoster string `json:"remotePoster"`
	Seasons []struct {
		SeasonNumber int `json:"seasonNumber"`
		Monitored bool `json:"monitored"`
	} `json:"seasons"`
	Year int `json:"year"`
	QualityProfileID int `json:"qualityProfileId"`
	SeasonFolder bool `json:"seasonFolder"`
	Monitored bool `json:"monitored"`
	UseSceneNumbering bool `json:"useSceneNumbering"`
	Runtime int `json:"runtime"`
	TvdbID int `json:"tvdbId"`
	TvRageID int `json:"tvRageId"`
	SeriesType string `json:"seriesType"`
	CleanTitle string `json:"cleanTitle"`
	ImdbID string `json:"imdbId"`
	TitleSlug string `json:"titleSlug"`
}