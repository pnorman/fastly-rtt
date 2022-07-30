/*
Copyright © 2022 Paul Norman

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type Result struct {
	PopCode string  `json:"popCode"`
	Rtt     float32 `json:"rtt"`
}

type Response struct {
	Results []Result `json:"popResults"`
}

func getRecommend(host string) {
	fmt.Printf("Querying Fastly API for host %s\n", host)
	apiUrl := "https://developer.fastly.com/api/internal/shieldRecommend?dest=" + url.QueryEscape(host)
	resp, err := http.Get(apiUrl)
	if err != nil {
		log.Fatalf("Error fetching from Fastly API: %s", err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Received status %s from Fastly API", resp.Status)
	}
	dec := json.NewDecoder(resp.Body)
	var jsonResponse Response
	err = dec.Decode(&jsonResponse)
	if err != nil {
		log.Fatalf("Error decoding JSON: %s", err)
	}
	for _, result := range jsonResponse.Results {
		fmt.Printf("%s %.0f\n", result.PopCode, result.Rtt)
	}
}
