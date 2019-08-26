package team

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTeamsGetData(t *testing.T) {
	teams := Teams{}
	teams.GetData([]string{"Bayern Munich"}, "https://vintagemonster.onefootball.com/api/teams/en", 10, 100)

	assert.Equal(t, 1, len(teams.list), "should have load one team")
	assert.Equal(t, 0, len(teams.Errors), "should not have any error to display")
}

func TestTeamsGetDataInvalidTeamName(t *testing.T) {
	teams := Teams{}
	teams.GetData([]string{"invalid-team"}, "https://vintagemonster.onefootball.com/api/teams/en", 10, 100)

	assert.Equal(t, 1, len(teams.Errors), "should not have one error to display")
}

func TestTeamsGetDataInvalidAPIEndpoint(t *testing.T) {
	teams := Teams{}
	teams.GetData([]string{"Bayern Munich"}, "invalid-api-endpoint", 10, 100)

	assert.Equal(t, 1, len(teams.Errors), "should not have one error to display")
}
