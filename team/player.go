package team

import (
	"fmt"
	"strings"
)

type player struct {
	ID    string `json:"id"`
	Age   string `json:"age"`
	Name  string `json:"name"`
	Teams []string
}

// ToString return the player data formated
func (p *player) ToString() string {
	return fmt.Sprintf("%s; %s; %s\n", p.Name, p.Age, strings.Join(p.Teams, ", "))
}

// ByName implements sort.Interface for []player based on the Name field
type ByName []player

func (p ByName) Len() int           { return len(p) }
func (p ByName) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p ByName) Less(i, j int) bool { return p[i].Name < p[j].Name }
