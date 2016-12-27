package go_sonarr

type SeriesLookupResponse []struct {
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