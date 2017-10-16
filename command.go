package sonarr

import "github.com/juju/errors"

func (sc *SonarrClient) RescanSeries(seriesID int) (err error) {
	err = sc.DoRequest("POST", "command", nil, map[string]interface{}{"name": "RescanSeries", "seriesId": seriesID}, nil)

	if err != nil {
		return errors.Annotate(err, "Failed to request rescanning of series with ID "+string(seriesID))
	}

	return nil
}
