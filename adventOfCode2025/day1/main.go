package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	//file handling
	file, err := os.Open("./text.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	zeroCount := 0
	currentNum := 50

	scanner := bufio.NewScanner(file)

	//main meat and potatoes - get all the times that it gets to zero which is the code for the elf
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		newNum := dialAlgorithm(scanner.Text(), currentNum)
		if newNum == 0 {
			zeroCount++
		}
		currentNum = newNum
	}

	//error handling for scanner
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("The elf code is: ", zeroCount)
}

// dial starts at 50 and goes either left or right
func dialAlgorithm(phrase string, currentNumber int) int {
	//split the string into 2 parts
	// L36 --> 'L' and '36'
	newNumber := 0
	runes := []rune(phrase)

	//check if string is empty, these mfers will get you like that
	if len(runes) > 0 {
		leftOrRight := string(runes[0])

		directionNumber, err := strconv.Atoi(string(runes[1:]))
		if err != nil {
			fmt.Println("Error with strconv.Atoi")
			return currentNumber
		}

		//checks if left or right movement and the algorithm uses modulus to "wrap" around 99
		//lowkey might be better as a switch statement
		if leftOrRight == "R" {
			newNumber = (currentNumber + directionNumber) % 100
			//this is the tricky one. If the number > 100 then we use %100 before adding 100
			//if we use modulus 100 first we get all the rotations out of the way and then when we add 100 we get to where we are supposed to be
			//the extra %100 at the end is incase we get 142 and we need 42.
		} else if leftOrRight == "L" {
			newNumber = (((currentNumber - directionNumber) % 100) + 100) % 100
		}

	} else {
		newNumber = currentNumber
	}

	return newNumber
}

// function that tests dialAlgorithm. Looks likle it works so far
func testDialAlgorithm() {
	currentNum := 50

	testList := [...]string{"L55", "R60", "L3", "R55", "R4", "L7", "L78"}

	for _, val := range testList {
		fmt.Println("Starting Number: ", currentNum)
		fmt.Println("Testing: ", val)

		newNum := dialAlgorithm(val, currentNum)
		fmt.Println("After dialAlgorithm: ", newNum)
		fmt.Println("---------------")

		currentNum = newNum
	}
}
