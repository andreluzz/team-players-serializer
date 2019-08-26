package team

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayerToString(t *testing.T) {
	player := player{
		Name:  "player name",
		Age:   "20",
		Teams: []string{"team 01", "team 02", "team 03"},
	}

	assert.Equal(t, "player name; 20; team 01, team 02, team 03\n", player.ToString(), "they should be equal")
}

func TestPlayersSort(t *testing.T) {
	sortedPlayers := []player{
		{Name: "AAA"},
		{Name: "BBB"},
		{Name: "CCC"},
		{Name: "DDD"},
	}

	unsortedPlayers := []player{
		{Name: "AAA"},
		{Name: "CCC"},
		{Name: "DDD"},
		{Name: "BBB"},
	}

	sort.Sort(ByName(unsortedPlayers))

	assert.Equal(t, sortedPlayers, unsortedPlayers, "they should be equal")
}
