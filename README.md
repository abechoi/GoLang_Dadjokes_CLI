<h1 align="center">
  Dadjokes CLI with GoLang
</h1>

## 1. Installation

```
go get -u github.com/spf13/cobra
cobra init --pkg-name github.com/spf13/dadjokes
go mod init github.com/example/dadjokes
```

```
go get github.com/mitchellh/go-homedir
go get github.com/spf13/cobra
go get github.com/spf13/viper
```

test the icanhazdadjoke api

```
curl -H "Accept: application/json" https://icanhazdadjoke.com/
// returns json with id, joke, and status
output: {"id":"O7haxA5Tfxc","joke":"Where do cats write notes?\r\nScratch Paper!","status":200}
```

## 2. Create command: random

```
cobra add random
```

random.go

```
// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "randomCmd returns a random dad joke.",
	Long:  "This command fetches a random dad joke from the icanhazdadjoke api.",
  // Run executes by "go run main.go random"
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
// execute this function var randomCmd
func getRandomJoke() {
	fmt.Println("This is a random joke.")
}
```

## 3. Create 2 Functions: getRandomJoke and getJokeData

```
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
```

## 4. Run random command

```
go run main.go random
```

## Finished random.go

```
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
	Long:  `This command fetches a random dad joke from the icanhazdadjoke api.`,
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

	// ioutil.ReadAll(response.Body) returns a responseBytes and error.
	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Cannot read body of response - %v", err)
	}

	return responseBytes
}
```
