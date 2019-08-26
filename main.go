package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"github.com/andreluzz/team-players-serializer/team"
)

var (
	teamsAPIEndpoint = flag.String("endpoint", "https://vintagemonster.onefootball.com/api/teams/en", "API endpoint")
	filePath         = flag.String("teams", "teams.json", "Path to the json file with the teams name array.")
	maxAPITeamsIndex = flag.Int("max-api-team-index", 500, "Defines the max index to try to find a team on the API.")
	maxConcurrency   = flag.Int("max-concurrency", 5, "Defines the max number of concurrency API request.")
)

func main() {
	flag.Parse()
	start := time.Now()

	// Load teams from the file path defined
	jsonBytes, err := ioutil.ReadFile(*filePath)
	if err != nil {
		log.Fatalf("error reading %s. Error: %v", *filePath, err)
	}

	teamsList := []string{}
	json.Unmarshal(jsonBytes, &teamsList)

	log.Printf("Start processing %d teams\n", len(teamsList))

	teams := team.Teams{}
	teams.GetData(teamsList, *teamsAPIEndpoint, *maxConcurrency, *maxAPITeamsIndex)

	fmt.Println("")
	fmt.Println(teams.Builder.String())
	if len(teams.Errors) > 0 {
		fmt.Println("Errors:")
		fmt.Println(strings.Join(teams.Errors, "\n"))
		fmt.Println("")
	}

	elapsed := time.Since(start)
	log.Printf("Processing ended after %s", elapsed)
}
