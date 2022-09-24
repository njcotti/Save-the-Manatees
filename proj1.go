// Author: Justin Swanson, jswanson2020@my.fit.edu
// Author: Naomi Cotti, ncotti2020@my.fit.edu
// Course: CSE 4250, Fall 2022
// Project: Proj1, Manatee Evacuation
// Implementation: go version xgcc (Ubuntu 4.9.3-13ubuntu2) 4.9.3 linux/amd64

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

//------------- Begin Make Types------------
type manatee struct {
	id int
	age int
	weight int
	position int
}

type pair struct {
	back manatee 	// the female
	front manatee 	// the male
	totalAge int
}
//------------- End Make Types------------

func test(male manatee, female manatee) bool {
	return male.weight <= female.weight
}

func pop(list []manatee) manatee {
	// returns the last element of a slice (or an array?)
	return list[len(list) - 1]
}

//-------------- Begin Create Instances--------------
func newManatee(identifier int, age int, weight int, position int) manatee {
	p := manatee{id: identifier, age: age, weight: weight, position: position}
	return p
}

func newPair(male manatee, female manatee) pair {
	p := pair{front: male, back: female, totalAge: male.age + female.age}
	return p
}
//-------------- End Create Instances--------------

func createPairings(males []manatee, females []manatee) []pair {
	var stack []manatee // to hold properly placed male manatees
	var holding []manatee // for manatees not yet placeable in stack (their position is too high)
	var pairings []pair // output array for pairings, to keep the manatees together
	i, j, k := 0, 0, 0


	for len(males) > 0 {
		for k = 0; k < len(holding); k++ {
			if (holding[k].position == i) {
				// finally push the holding into the stack
				stack = append(stack, holding[k])
				i += 1
			}
		}
		male := pop(males) // get male
		males = males[:len(males) - 1] // shorten the males list to give illusion of popping
		passed := test(male, females[i]) // test if male is smaller than female
		if(passed) {
			// push into stack
			male.position = i
			stack = append(stack, male)
			i += 1
		} else {
			// you didnt pass the test, and could be approaching bounds
			j = tryNext(i, male, females) // searched through remaining position to see if works
			if (j != -1) {
				// there is a valid position for the male, update his position and add to holding
				male.position = j
				holding = append(holding, male)
			} else {
				// not a valid position in the remaining spaces, backtrack, then try to see if male fits where i was
				newPos := backtrack(i, male, females, stack) // newPos is where you should backtrack to
				if (newPos == -1) {
					fmt.Println("impossible")
					return nil
				}
				for (newPos + 1 <= len(stack)) {
					// while loop to do backtracking
					poppedMale := pop(stack) // backtrack and see if the would be a better fit in poppedMale's spot
					poppedMale.position = -1
					males = append(males, poppedMale) // add the poppedMale back to males
					// swap back with front to avoid circular backtracking (if it ever happens)
					males[0], males[len(males) - 1] = males[len(males) - 1], males[0]
				}
				male.position = newPos
				stack = append(stack, male)
			}
		}
	}

	for k = 0; k < len(stack); k++ {
		// make new pairings using our properly ordered stack, and female slice
		pairings = append(pairings, newPair(stack[k], females[k]))
	}
	return pairings
}

func backtrack(i int, male manatee, females []manatee, stack []manatee) int {
	// returns the integer position where backtracking succeeded, or -1 if not
	for (len(stack) > 0) {
		pop(stack)
		stack = stack[:len(stack) - 1]
		i -= 1
		passed := test(male, females[i])
		if (passed) {
			return i
		}
	}
	// length of stack is 0, check to see if male fits in 0, if not return -1
	if (test(male, females[0]) == true) {
		return 0
	}
	return -1
}

func tryNext(i int, male manatee, females []manatee) int {
	for (i < len(females)) {
		pass := test(male, females[i])
		if (pass) {
			return i
		}
		i += 1
	}
	return -1
}

// creates and returns a row of the 2D slice, that we'll later use to make manatees
func manateeInput(s string) []int {
	var mTee []int
	for _, f := range strings.Fields(s) {
		i, err := strconv.Atoi(f)
		if err == nil {
			mTee = append(mTee, i)
		}
	}
	//num := strings.Fields(s)
	return mTee
}

func main() {

	var i, row, col  int = 0, 0, 0

	// make pair stack for later
	var pairStack []pair

	// tables used in the printing function
	var outputFemale []manatee
	var outputMale []manatee

	//------------Begin Handle Input------------------
	fmt.Print("Input number of manatees of each gender: ")
	var rowNum int
	fmt.Scanln(&rowNum) // this lowkey does not matter at all
	fmt.Println("Input manatee array: ")

	scanner := bufio.NewScanner(os.Stdin)
	num := make([][]int, 0) // Declare a slice to receive other slices inside it

	for scanner.Scan() { // Scrolls all typed data to true

		// If the user does not type anything, that is, if they press Enter an interrupt will occur
		if scanner.Text() == "" {
			break
		} else {
			num = append(num, manateeInput(scanner.Text())) // Adding the slice inside list
		}
	}
	/*for _, value := range num {
		fmt.Printf("%d\n", value)
	} */
	//------------End Handle Input--------------------

	var ListOfFemales []manatee
	var ListOfMales []manatee

	//------------Begin Calculate Row Lengths---------------
	numManateesInEachGender := rowNum
	numManateesInRow := len(num[0])/2
	numManateesInCol := numManateesInEachGender/numManateesInRow
	half := len(num)/2
	//------------End Calculate Row Lengths---------------

	//------------Begin Instantiate Manatees--------------
	i = 1
	for i <= numManateesInEachGender {
		for row = 0; row < numManateesInCol; row++ {
			for col = 0; col <= numManateesInRow; col +=2 {
				ListOfFemales = append(ListOfFemales, newManatee(i, num[row][col], num[row][col+1], -1))
				ListOfMales = append(ListOfMales, newManatee(i, num[row+ half][col], num[row+ half][col+1], -1))
				i += 1
			}
		}
	}
	//------------End Instantiate Manatees----------------

	//------------Start Making Pairings-------------------
	pairStack = createPairings(ListOfMales, ListOfFemales)
	//------------End Making Pairings---------------------

	// sort pairStack by total age to get the lowest aged sets of manatees to the left
	sort.SliceStable(pairStack, func(i, j int) bool { return pairStack[i].totalAge < pairStack[j].totalAge })

	// print out the list of females and then the list of males
	if (pairStack != nil) {
		for i = 0; i < len(pairStack); i++ {
			outputFemale = append(outputFemale, pairStack[i].back)
			outputMale = append(outputMale, pairStack[i].front)
		}
		for i = 0; i < len(pairStack); i++ {
			fmt.Printf("%d ", outputFemale[i].id)
		}
		fmt.Println()
		for i = 0; i < len(pairStack); i++ {
			fmt.Printf("%d ", outputMale[i].id)
		}
		fmt.Println()
	}
}


