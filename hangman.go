package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var dev = flag.Bool("dev", false, "dev mode")

func get_keys(entries map[string]bool) (keys []string) {
	for k, _ := range entries {
		keys = append(keys, k)
	}
	return
}

func get_word() string {
	if *dev {
		return "elephant"
	}
	resp, err := http.Get("https://random-word-api.herokuapp.com/word?number=5")
	if err != nil {
		return "elephant"
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var words []string
	err = json.Unmarshal(body, &words)
	if err != nil {
		// handle error
		return "elephant"
	}

	//fmt.Println(words)
	for _, word := range words {
		if len(word) > 4 && len(word) < 9 {
			return word
		}
	}

	return words[0]
}

func main() {
	flag.Parse()
	word := get_word()

	// lookup for entries made by the user.
	entries := map[string]bool{}

	// list of "_" corrosponding to the number of letters in the word. [ _ _ _ _ _ ]
	placeholder := []string{}
	//placeholder := make([]string, len(word), len(word))

	// get length of the word len(word)
	// initialize slice with each element as "_"
	for i := 0; i < len(word); i++ {
		placeholder = append(placeholder, "_")
		//placeholder[i] = "_"
	}

	t := time.NewTimer(1 * time.Minute)

	chances := 8

	// true: win
	// false: lose
	result := make(chan bool)
	go func() {
		for {
			// evaluate a loss! If user guesses a wrong letter or the wrong word, they lose a chance.
			userInput := strings.Join(placeholder, "")
			if chances == 0 && userInput != word {
				result <- false
				fmt.Println("You're out of chances")
				fmt.Println("Word was: ", word)
				fmt.Println("Game Over! Try again")
				return
			}

			// evaluate a win!
			if userInput == word {
				result <- true
				fmt.Println("You win!!")
				return
			}

			// Console display
			fmt.Println()
			fmt.Println(placeholder)                       // render the placeholder
			fmt.Printf("Chances: %d\n", chances)           // render the chances left
			fmt.Printf("Entries: %v\n", get_keys(entries)) // show the letters or words guessed till now.
			fmt.Printf("Guess a letter or the word: ")
			// Addon validation: Allow only alpha!
			// Addon validation: manage case!

			// take the input
			str := ""
			fmt.Scanln(&str)

			// if len(str) > 2, compare the word with the str
			if len(str) > 2 {
				if str == word {
					result <- true
					fmt.Println("You win!!")
					return
				} else {
					// you lose a chance
					entries[str] = true
					chances -= 1
					continue
				}
			}

			// compare and update entries, placeholder and chances.
			_, ok := entries[str]
			if ok {
				// key exists already; duplicate
				continue
			}

			entries[str] = true
			// check if letter exists in the word!
			found := false
			for i, v := range word {
				if str == string(v) {
					placeholder[i] = string(v)
					found = true
				}
			}

			/*
				temp := strings.Split(word, "")
				for i, v := range temp {
					if v == str {
						// update the placeholder indices.
						placeholder[i] = v
						found = true
					}
				}
			*/
			if !found {
				chances -= 1
			}
		}
	}()

	// infinite for loop is for reference incase later you need to read from the channels again
	// In our case, we need to read ONLY ONCE
	for {
		select {
		case <-result:
			fmt.Println("...")
			goto END
		case <-t.C:
			fmt.Println("Timedout... too bad!")
			goto END
		}
	}
END:
	fmt.Println("Play again..")
}
