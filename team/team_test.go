package team

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTeamLoad(t *testing.T) {
	team, err := load(1, "https://vintagemonster.onefootball.com/api/teams/en")

	assert.NoError(t, err, "should not return any error")
	assert.NotNil(t, team, "should not be nil")
}

func TestTeamLoadInvalidIndex(t *testing.T) {
	team, err := load(-1, "https://vintagemonster.onefootball.com/api/teams/en")

	assert.Error(t, err, "should return error")
	assert.Nil(t, team, "should be nil")
}

func TestTeamLoadInvalidURL(t *testing.T) {
	team, err := load(-1, "wrong-url-endpoint")

	assert.Error(t, err, "should return error")
	assert.Nil(t, team, "should be nil")
}
