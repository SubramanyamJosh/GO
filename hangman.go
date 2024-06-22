package main

import (
	"fmt"
	"strings"
)

func get_keys(entries map[string]bool)(keys []string){
for k, _ := range entries{
			keys = append(keys,k)
		}

		return 
} 

func main() {
	word := "elephant"

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
		keys := []string{}
		
		fmt.Println(keys) // show the letters or words guessed till now.
		fmt.Printf("Guess a letter or the word: ")

    		// take the input
		str := ""
		fmt.Scanln(&str)

		inputList := strings.Split(str,"")
		// compare and update entries, placeholder and chances.

		newLetterAdded := false
		for i := 0; i < len(inputList); i++ {
			_,ok := entries[inputList[i]]
			if(!ok){
				entries[inputList[i]] = true 	
				newLetterAdded = true		
			}		
  	
	}
			isNewLetterCrt := false

		for index,letter := range word {
			_,ok := entries[string(letter)]
			if(ok && (placeholder[index] != string(letter))){
				placeholder[index] = string(letter)
				isNewLetterCrt = true
			}
		}
	if(newLetterAdded && !isNewLetterCrt){
			chances = chances - 1 

		}	
}
}