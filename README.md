# team-players-serializer

## Options:
- endpoint: Defines the API endpoint to retrieve the data.
- teams: file path to the json array with the teams.
- max-api-team-index: Defines the maximum team_id value to seek for teams. Default value: 500
- max-concurrency: Defines the max number of concurrency API request. Default value: 5

## Executing the code:
```
go get github.com/stretchr/testify/assert
go build
./team-players-serializer
```

### Observations:
1. Lib github.com/stretchr/testify/assert was used to simplify test assetions
2. When executing the code it will render a string with dots and numbers. Where the dots are requests that does not match the team name and numbers are the amount of correct request that matched a tem name in the defined array.

```
2019/08/26 16:50:06 Start processing 10 teams
.1....2.3...4............5....6.........7...........8.................9...................................10
```
