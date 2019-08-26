package team

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// teamResponse store the content frem the onefootball team api response
type teamResponse struct {
	Data struct {
		Team team `json:"team"`
	} `json:"data"`
}

type team struct {
	Index   int
	Name    string    `json:"name"`
	Players []*player `json:"players"`
}

// load access the api endpoint and load the data from the teamID
func load(teamID int, apiEndpoint string) (*team, error) {
	resp, err := http.Get(fmt.Sprintf("%s/%d.json", apiEndpoint, teamID))
	if err != nil {
		return nil, fmt.Errorf("error getting team info from api: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("invalid response code: %d", resp.StatusCode)
	}

	tr := &teamResponse{}
	if err := json.NewDecoder(resp.Body).Decode(tr); err != nil {
		return nil, fmt.Errorf("error in decoding the team info: %v", err)
	}

	return &tr.Data.Team, nil
}
