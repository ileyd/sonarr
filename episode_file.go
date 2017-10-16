package sonarr

import (
	"strconv"
	"time"

	"github.com/juju/errors"
)

type EpisodeFile struct {
	SeriesID     int       `json:"seriesId"`
	SeasonNumber int       `json:"seasonNumber"`
	Path         string    `json:"path"`
	Size         int64     `json:"size"`
	DateAdded    time.Time `json:"dateAdded"`
	SceneName    string    `json:"sceneName"`
	Quality      struct {
		Quality struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"quality"`
		Proper bool `json:"proper"`
	} `json:"quality"`
	ID int `json:"id"`
}

func (sc *SonarrClient) GetAllEpisodeFiles(seriesID int) (ef []EpisodeFile, err error) {
	err = sc.DoRequest("GET", "episodeFile", map[string]string{"seriesId": strconv.Itoa(seriesID)}, nil, &ef)

	if err != nil {
		return []EpisodeFile{}, errors.Annotate(err, "Failed to get episode files for series "+strconv.Itoa(seriesID))
	}

	return ef, nil
}

func (sc *SonarrClient) GetEpisodeFile(efID int) (ef EpisodeFile, err error) {
	err = sc.DoRequest("GET", "episodeFile", map[string]string{"id": strconv.Itoa(efID)}, nil, &ef)

	if err != nil {
		return EpisodeFile{}, errors.Annotate(err, "Failed to get episode file "+strconv.Itoa(efID))
	}

	return ef, nil
}

func (sc *SonarrClient) DeleteEpisodeFile(efID int) (err error) {
	err = sc.DoRequest("DELETE", "episodeFile", map[string]string{"id": strconv.Itoa(efID)}, nil, nil)

	if err != nil {
		return errors.Annotate(err, "Failed to delete episode file "+strconv.Itoa(efID))
	}

	return nil
}
