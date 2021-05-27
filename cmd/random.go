/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "randomCmd returns a random dad joke.",
	Long:  "This command fetches a random dad joke from the icanhazdadjoke api.",
	Run: func(cmd *cobra.Command, args []string) {
		getRandomJoke()
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)
}

// create Joke struct
type Joke struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

func getRandomJoke() {
	url := "https://icanhazdadjoke.com/"
	responseBytes := getJokeData(url)

	joke := Joke{}

	// json.Unmarshall takes 2 arguments: responseBytes
	if err := json.Unmarshal(responseBytes, &joke); err != nil {
		log.Printf("Could not unmarshal response - %v", err)
	}

	fmt.Println(string(joke.Joke))
}

func getJokeData(baseAPI string) []byte {
	// http.NewRequest takes 3 arguments: http.MethodGet, baseAPI, and nil.
	// returns request and error
	request, err := http.NewRequest(
		http.MethodGet,
		baseAPI,
		nil,
	)
	if err != nil {
		log.Printf("Could not request a dad joke - %v", err)
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "Dadjoke CLI (github.com/example/dadjoke)")

	// http.DefaultClient.Do(request) returns a response and error.
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("Could not return a response - %v", err)
	}

	// util.ReadAll(response.Body) returns a responseBytes and error.
	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Cannot read body of response - %v", err)
	}

	return responseBytes
}
