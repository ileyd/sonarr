package sonarr

import (
	"log"
	"strconv"
	"time"

	"github.com/juju/errors"
)

type Episode struct {
	SeriesID              int       `json:"seriesId"`
	EpisodeFileID         int       `json:"episodeFileId"`
	SeasonNumber          int       `json:"seasonNumber"`
	EpisodeNumber         int       `json:"episodeNumber"`
	Title                 string    `json:"title"`
	AirDate               string    `json:"airDate"`
	AirDateUtc            time.Time `json:"airDateUtc"`
	Overview              string    `json:"overview"`
	HasFile               bool      `json:"hasFile"`
	Monitored             bool      `json:"monitored"`
	SceneEpisodeNumber    int       `json:"sceneEpisodeNumber"`
	SceneSeasonNumber     int       `json:"sceneSeasonNumber"`
	TvDbEpisodeID         int       `json:"tvDbEpisodeId"`
	AbsoluteEpisodeNumber int       `json:"absoluteEpisodeNumber"`
	ID                    int       `json:"id"`
}

func (sc *SonarrClient) GetAllEpisodes(seriesID int) (e []Episode, err error) {
	log.Println("Received seriesID", seriesID)
	err = sc.DoRequest("GET", "episode", map[string]string{"seriesId": strconv.Itoa(seriesID)}, nil, &e)

	if err != nil {
		return []Episode{}, errors.Annotate(err, "Failed to get episodes for series "+strconv.Itoa(seriesID))
	}

	return e, nil
}

func (sc *SonarrClient) GetEpisode(eID int) (e Episode, err error) {
	err = sc.DoRequest("GET", "episode", map[string]string{"id": strconv.Itoa(eID)}, nil, &e)

	if err != nil {
		return Episode{}, errors.Annotate(err, "Failed to get episode file "+strconv.Itoa(eID))
	}

	return e, nil
}

func (sc *SonarrClient) UpdateEpisode(e Episode) (er Episode, err error) {
	err = sc.DoRequest("PUT", "episode", nil, e, &er)

	if err != nil {
		return Episode{}, errors.Annotate(err, "Failed to update episode "+strconv.Itoa(e.ID))
	}

	return er, nil
}
