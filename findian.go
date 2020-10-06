//Write a program which prompts the user to enter a string.
//The program searches through the entered string for the characters ‘i’, ‘a’, and ‘n’.
//The program should print “Found!”
//if the entered string starts with the character ‘i’,
//ends with the character ‘n’, and contains the character ‘a’.
//The program should print “Not Found!” otherwise.

//The program should not be case-sensitive, so it does not matter if the characters are upper-case or lower-case.

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

)
func main() {

	fmt.Println("!!!!!!!!!!!!!!!!!!! Welcome to findian.go !!!!!!!!!!!!!!!!!!!")
	fmt.Println("You can type exit to leave at any time.")
	fmt.Println("Enter a string of your choosing. Let's see if it begins with \"i\" and contains an  \"a\", and ends in \"n\"." )

	findian , _  := regexp.Compile("^[iI].*[aA]+.*[nN]$")

	inputScanner := bufio.NewScanner(os.Stdin)

	for {

		inputScanner.Scan()

		userResponse := inputScanner.Text()

		if userResponse == "exit" {
			os.Exit(0)
		}
		
		if findian.MatchString(userResponse) {
			fmt.Println("The word " + userResponse)
			fmt.Println("Found!")
		} else {
			fmt.Println("The word " + userResponse)
			fmt.Println("Not Found!")
		}
		
		fmt.Println("Please enter another string.")
	}
	
}
