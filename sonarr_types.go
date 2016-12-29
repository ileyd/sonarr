package go_sonarr

import "time"

type SeriesPostResponse struct {
	Title string `json:"title,omitempty"`
	SeasonCount int `json:"seasonCount,omitempty"`
	EpisodeCount int `json:"episodeCount,omitempty"`
	EpisodeFileCount int `json:"episodeFileCount,omitempty"`
	Status string `json:"status,omitempty"`
	Overview string `json:"overview,omitempty"`
	NextAiring time.Time `json:"nextAiring,omitempty"`
	Network string `json:"network,omitempty"`
	AirTime string `json:"airTime,omitempty"`
	Images []struct {
		CoverType string `json:"coverType,omitempty"`
		URL string `json:"url,omitempty"`
	} `json:"images,omitempty"`
	Seasons []struct {
		SeasonNumber int `json:"seasonNumber,omitempty"`
		Monitored bool `json:"monitored,omitempty"`
	} `json:"seasons,omitempty"`
	Year int `json:"year,omitempty"`
	Path string `json:"path,omitempty"`
	QualityProfileID int `json:"qualityProfileId,omitempty"`
	SeasonFolder bool `json:"seasonFolder,omitempty"`
	Monitored bool `json:"monitored,omitempty"`
	UseSceneNumbering bool `json:"useSceneNumbering,omitempty"`
	Runtime int `json:"runtime,omitempty"`
	TvdbID int `json:"tvdbId,omitempty"`
	TvRageID int `json:"tvRageId,omitempty"`
	FirstAired time.Time `json:"firstAired,omitempty"`
	LastInfoSync time.Time `json:"lastInfoSync,omitempty"`
	SeriesType string `json:"seriesType,omitempty"`
	CleanTitle string `json:"cleanTitle,omitempty"`
	ImdbID string `json:"imdbId,omitempty"`
	TitleSlug string `json:"titleSlug,omitempty"`
	ID int `json:"id,omitempty"`
}

type SeriesLookupGetResponse []struct {
	Title string `json:"title,omitempty"`
	SeasonCount int `json:"seasonCount,omitempty"`
	EpisodeCount int `json:"episodeCount,omitempty"`
	EpisodeFileCount int `json:"episodeFileCount,omitempty"`
	Status string `json:"status,omitempty"`
	Overview string `json:"overview,omitempty"`
	Network string `json:"network,omitempty"`
	Images []struct {
		CoverType string `json:"coverType,omitempty"`
		URL string `json:"url,omitempty"`
	} `json:"images,omitempty"`
	RemotePoster string `json:"remotePoster,omitempty"`
	Seasons []struct {
		SeasonNumber int `json:"seasonNumber,omitempty"`
		Monitored bool `json:"monitored,omitempty"`
	} `json:"seasons,omitempty"`
	Year int `json:"year,omitempty"`
	QualityProfileID int `json:"qualityProfileId,omitempty"`
	SeasonFolder bool `json:"seasonFolder,omitempty"`
	Monitored bool `json:"monitored,omitempty"`
	UseSceneNumbering bool `json:"useSceneNumbering,omitempty"`
	Runtime int `json:"runtime,omitempty"`
	TvdbID int `json:"tvdbId,omitempty"`
	TvRageID int `json:"tvRageId,omitempty"`
	SeriesType string `json:"seriesType,omitempty"`
	CleanTitle string `json:"cleanTitle,omitempty"`
	ImdbID string `json:"imdbId,omitempty"`
	TitleSlug string `json:"titleSlug,omitempty"`
}