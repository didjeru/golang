package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

var numberMin, numberMax = 0, 100
var compNumber = getRandomNumber(numberMin, numberMax)
var countTry = 0

func getRandomNumber(numberMin, numberMax int) int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("You are chiter!")
			os.Exit(3)
		}
	}()
	return numberMin + rand.Intn((numberMax)-(numberMin))
}

func inputData(msg string) string {
	var data string
	fmt.Print(msg)
	fmt.Scanln(&data)
	if data == "exit" {
		os.Exit(3)
	}
	return data
}

func variant1(){
	fmt.Println("Think of a number in the range of 0 to about 100")
	fmt.Println("Your number", compNumber, "?")
	for {
		inputSymbol := inputData("Enter symbol <, =, > or exit:")
		if inputSymbol == "<" || inputSymbol == "=" || inputSymbol == ">" {
			countTry++
		}
		switch inputSymbol {
		case "=":
			fmt.Println("Win witn", countTry, "try. And number -", compNumber)
			os.Exit(3)
		case ">":
			numberMin = compNumber + 1
			compNumber = getRandomNumber(numberMin, numberMax)
			fmt.Println("Your number", compNumber, "?")
		case "<":
			numberMax = compNumber
			compNumber = getRandomNumber(numberMin, numberMax)
			fmt.Println("Your number", compNumber, "?")
		default:
			fmt.Println("Wrong symbol")
		}
	}
}

func variant2() {
	fmt.Println("Computer think of a number in the range of 0 to about 100")
	compNumber = getRandomNumber(numberMin, numberMax)
	userNumber := inputData("Enter your number or exit")
	countTry++
	fmt.Println(countTry, compNumber,userNumber)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Guessing game!")
	inputVariant := inputData("Enter 1 if you think a number. Enter 2 if computer think a number: ")
	switch inputVariant {
	case "1":
		variant1()
	case "2":
		variant2()
	default:
		fmt.Println("Wrong input number")
	}
}
