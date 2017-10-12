package sonarr

import "github.com/juju/errors"

type ParsedEpisodeInfo struct {
	Title             string `json:"title"`
	ParsedEpisodeInfo struct {
		ReleaseTitle    string `json:"releaseTitle"`
		SeriesTitle     string `json:"seriesTitle"`
		SeriesTitleInfo struct {
			Title            string `json:"title"`
			TitleWithoutYear string `json:"titleWithoutYear"`
			Year             int    `json:"year"`
		} `json:"seriesTitleInfo"`
		Quality struct {
			Quality struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"quality"`
			Revision struct {
				Version int `json:"version"`
				Real    int `json:"real"`
			} `json:"revision"`
		} `json:"quality"`
		SeasonNumber             int           `json:"seasonNumber"`
		EpisodeNumbers           []int         `json:"episodeNumbers"`
		AbsoluteEpisodeNumbers   []interface{} `json:"absoluteEpisodeNumbers"`
		Language                 string        `json:"language"`
		FullSeason               bool          `json:"fullSeason"`
		Special                  bool          `json:"special"`
		ReleaseGroup             string        `json:"releaseGroup"`
		ReleaseHash              string        `json:"releaseHash"`
		IsDaily                  bool          `json:"isDaily"`
		IsAbsoluteNumbering      bool          `json:"isAbsoluteNumbering"`
		IsPossibleSpecialEpisode bool          `json:"isPossibleSpecialEpisode"`
	} `json:"parsedEpisodeInfo"`
	Series struct {
	} `json:"series"`
	Episodes []interface{} `json:"episodes"`
}

func (sc *SonarrClient) ParseEpisodeByRelease(rel string) (info ParsedEpisodeInfo, err error) {
	err = sc.DoRequest("GET", "parse", map[string]string{"title": rel}, nil, &info)

	if err != nil {
		return ParsedEpisodeInfo{}, errors.Annotate(err, "Failed to parse release "+rel)
	}

	return info, nil
}
