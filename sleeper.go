package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// TODO - this data contract is invalid right now. Needs more development if to be used.
type player struct {
	position               string
	last_name              string
	pandascore_id          string // ??
	years_exp              int
	hashtag                string
	status                 string
	news_updated           string // ??
	birth_state            string // ??
	injury_body_part       string // ??
	high_school            string // ??
	depth_chart_position   string
	fantasy_positions      []string // ???????????
	yahoo_id               int
	injury_start_date      string // ??
	active                 bool
	birth_date             string
	rotoworld_id           int
	age                    int
	metadata               string // ??
	full_name              string
	fantasy_data_id        int
	sportradar_id          string
	birth_country          string // ??
	sport                  string
	number                 int
	practice_participation string // ??
	espn_id                int
	depth_chart_order      int
	birth_city             string // ??
	rotowire_id            int
	injury_notes           string // ??
	height                 string
	college                string
	gsis_id                string // ??
	first_name             string
	player_id              string
	search_rank            int
	search_last_name       string
	weight                 string
	injury_status          string // ??
	search_full_name       string
	practice_description   string // ??
	search_first_name      string
	team                   string // ??
	stats_id               int
}

const (
	// See http://golang.org/pkg/time/#Parse
	timeFormat = "2006-01-02"
)

func main() {

	sleeperAPI := "https://api.sleeper.app/v1/players/nfl"

	resp, err := http.Get(sleeperAPI)
	if err != nil {
		print(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}

	playerJSON := string(body)
	var result map[string]json.RawMessage
	err = json.Unmarshal([]byte(playerJSON), &result)

	for _, player := range result {

		var decodedPlayer map[string]string
		json.Unmarshal(player, &decodedPlayer)

		fullName := decodedPlayer["full_name"]
		position := decodedPlayer["position"]
		rawDOB := decodedPlayer["birth_date"]

		DOB, _ := time.Parse(timeFormat, rawDOB)
		age := time.Since(DOB)

		fmt.Printf("%s,%s,%f\n", fullName, position, age.Hours()/(24.0*365))
	}
}
