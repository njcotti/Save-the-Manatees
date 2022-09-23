package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type gender int64

type manatee struct {
	id     int
	age    int
	weight int
	mf     gender
}

func newManatee(identifier int, age int, weight int, mf gender) manatee {
	p := manatee{id: identifier, age: age, weight: weight, mf: mf}
	return p
}

func manateeInput(s string) []string {
	num := strings.Fields(s)
	return num
}
func main() {

	fmt.Println("Input number of manatees in a row")
	var rowNum int
	fmt.Scanln(&rowNum)
	fmt.Println("Input manatee array")

	scanner := bufio.NewScanner(os.Stdin)
	num := make([][]string, 0) // Declare a slice to receive other slices inside it

	for scanner.Scan() { // Scrolls all typed data to true

		// If the user does not type anything, that is, if he presses Enter an interrupt will occur
		if scanner.Text() == "" {
			break
		} else {
			num = append(num, manateeInput(scanner.Text())) // Adding the slice inside list
		}
	}

	fmt.Println(num) // print list

}

/*
, secondLine, thirdLine, fourthLine
case 2:
			secondLine = manateeInput(scanner.Text())
		case 3:
			thirdLine = manateeInput(scanner.Text())
		case 4:
			fourthLine = manateeInput(scanner.Text())
fmt.Println(secondLine)
	fmt.Println(thirdLine)
	fmt.Println(fourthLine)
	// probably won't need this if our lists are kept separate. Removing means we take the gender field out of Manatee
	const (
		male gender = iota
		female
	)

	constant is necessary atm, final product may use a slice instead of an array, just bc i don't know yet if a
	   constant can be set through user input
	const numManateesInEachGender int = 4
	var i, row, col, l int = 0, 0, 0, 0

	// commented out because not currently in use
	//var stack []manatee

	var input = [numManateesInEachGender][numManateesInEachGender]int{
		3 2 1 2
		2 3 4 3
		2 1 2 1
		2 2 1 3
3 2 1 2
2 3 4 3
2 1 2 1
2 2 1 3

	var ListOfFemales []manatee
	var ListOfMales []manatee
*/
/*
	numManateesInRow := len(input[0]) / 2
	numManateesInCol := numManateesInEachGender / numManateesInRow
	half := len(input) / 2 // how many rows before we switch from female to male

	i = 1
	for i <= numManateesInEachGender {
		for row = 0; row < numManateesInCol; row++ {
			for col = 0; col <= numManateesInRow; col += 2 {
				ListOfFemales = append(ListOfFemales, newManatee(i, input[row][col], input[row][col+1], female))
				ListOfMales = append(ListOfMales, newManatee(i, input[row+half][col], input[row+half][col+1], male))
				i += 1
			}
		}
	}
	// print out the list of females and then the list of males
	fmt.Println("Females: ")
	for l = 0; l < len(ListOfFemales); l++ {
		fmt.Println(ListOfFemales[l])
	}
	fmt.Println("Males: ")
	for l = 0; l < len(ListOfMales); l++ {
		fmt.Println(ListOfMales[l])
	}
*/
