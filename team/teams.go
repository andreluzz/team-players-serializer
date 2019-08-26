package team

import (
	"fmt"
	"log"
	"sort"
	"strings"
)

// Teams define a list of teams and a builder with the players ordered
type Teams struct {
	list    map[string]*team
	Builder strings.Builder
	Errors  []string
}

// GetData acess the api and get the data for each team in the list
// maxAPIIndex is an integer used to define the maximum index to try to find a team at the API
func (teams *Teams) GetData(teamNames []string, apiEndpoint string, maxConcurrency, maxAPIIndex int) {
	teams.list = make(map[string]*team, len(teamNames))
	for _, teamName := range teamNames {
		teams.list[teamName] = nil
	}

	processedTeams := 0
	totalIndexloaded := 0

	done := make(chan bool, 1)
	indexesToProcess := make(chan int, maxAPIIndex)
	results := make(chan *team, maxConcurrency)

	for w := 1; w <= maxConcurrency; w++ {
		go worker(apiEndpoint, indexesToProcess, results, done)
	}

	for i := 1; i <= maxAPIIndex; i++ {
		indexesToProcess <- i
	}
	close(indexesToProcess)

	for {
		select {
		case loadedTeam := <-results:
			totalIndexloaded++
			fmt.Print(".")
			if _, ok := teams.list[loadedTeam.Name]; ok {
				teams.list[loadedTeam.Name] = loadedTeam
				processedTeams++
				fmt.Print(processedTeams)
			}
			//fmt.Printf("Total: %d | Teams: %d", totalIndexloaded, processedTeams)
			if processedTeams == len(teamNames) || totalIndexloaded == maxAPIIndex {
				done <- true
				teams.process()
				return
			}
		}
	}
}

func worker(apiEndpoint string, indexes <-chan int, results chan<- *team, done <-chan bool) {
	for {
		select {
		case i := <-indexes:
			if i <= 0 {
				continue
			}
			loadedTeam, err := load(i, apiEndpoint)
			if err != nil {
				log.Printf("error loading team_id: %d - error: %v\n", i, err)
				results <- &team{Index: i}
				continue
			}
			loadedTeam.Index = i
			results <- loadedTeam
		case <-done:
			return
		}
	}
}

// process render the players from all teams ordered by name to the builder
// any error will be stored on teams errors slice
func (teams *Teams) process() {
	teams.Builder.Reset()
	players := make(map[string]*player)
	for name, t := range teams.list {
		if t == nil {
			teams.Errors = append(teams.Errors, fmt.Sprintf("team: %s not found", name))
			continue
		}

		for _, p := range t.Players {
			id := p.ID
			if _, ok := players[id]; ok {
				players[id].Teams = append(players[id].Teams, t.Name)
				continue
			}
			p.Teams = append(p.Teams, t.Name)
			players[id] = p
		}
	}

	sortedPlayers := []player{}
	for _, p := range players {
		sortedPlayers = append(sortedPlayers, *p)
	}

	sort.Sort(ByName(sortedPlayers))

	for index, p := range sortedPlayers {
		str := fmt.Sprintf("%03d. %s", index+1, p.ToString())
		teams.Builder.WriteString(str)
	}
}
