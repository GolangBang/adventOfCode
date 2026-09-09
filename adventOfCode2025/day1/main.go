package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./text.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

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
		if leftOrRight == "R" {
			newNumber = (currentNumber + directionNumber) % 99
		} else if leftOrRight == "L" {
			newNumber = ((currentNumber - directionNumber) + 99) % 99
		}

	} else {
		newNumber = currentNumber
	}

	return newNumber
}

// function that tests dialAlgorithm,
func testDialAlgorithm() {
	currentNum := 50
	testList := [...]string{"L55, R60, L3, R55, R4, L7, L78"}
	for _, val := range testList {
		fmt.Println("CurrentNum is: ", currentNum)
		fmt.Println("Num is getting subtracted by 55, the correct answer should be 94.")
		newNum := dialAlgorithm(val, currentNum)
	}

}
