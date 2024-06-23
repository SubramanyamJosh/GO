package main

import (
	"fmt"
	"strings"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

func get_keys(entries map[string]bool)(keys []string){
for k, _ := range entries{
			keys = append(keys,k)
		}

		return 
} 

func get_words() string{
resp, err := http.Get("https://random-word-api.herokuapp.com/word?number=5")
if err != nil {
	return "elephant"
}
defer resp.Body.Close()
body, err := ioutil.ReadAll(resp.Body)

var words []string
err = json.Unmarshal(body,&words)

if err != nil {
	return "elephant"
}

for _,word := range(words) {
	if len(word) > 4 && len(word) < 9 {
		return word
	}
}
return words[0]
}

func main() {
	word := get_words()

	// lookup for entries made by the user.
	entries :=   map[string]bool{}
  
  	// list of "_" corrosponding to the number of letters in the word. [ _ _ _ _ _ ]
	placeholder := []string{}
	// placeholder := make([]string, len(word))
	for i:=0 ; i<len(word) ; i++ {
		placeholder = append(placeholder,"_")
	}
	
	chances := 8

	for {
		// evaluate a loss! If user guesses a wrong letter or the wrong word, they lose a chance.
		userInput := strings.Join(placeholder,"")
		if chances == 0 && userInput != word {
			fmt.Println("Game Over! Try again")
			break
		}
	
		// evaluate a win!
			if  userInput == word {
			fmt.Println("You Win")
			break
		}
	
    		// Console display
		fmt.Println("\n")
		fmt.Println(placeholder) // render the placeholder
		fmt.Printf("Chances: %d\n",chances) // render the chances left		
		fmt.Println(get_keys(entries)) // show the letters or words guessed till now.
		fmt.Printf("Guess a letter or the word: ")

    		// take the input
		str := ""
		fmt.Scanln(&str)

		if len(str) > 2{
			if(str == word ){
				fmt.Println("You Win")
				break
			}else{
				entries[str] = true
				chances -= 1
				continue
		}
				
		}

		// compare and update entries, placeholder and chances.

		_, ok := entries[str]
		fmt.Println(ok)
		if ok {
			continue
		}

		entries[str] = true
		
		isInputFound := false

		for index,letter := range word {
			if str == string(letter) {
				placeholder[index] = string(letter)
				isInputFound = true
			}
		}
		
		if(!isInputFound){
			chances = chances - 1
		}	
}
}