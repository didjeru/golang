package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var numberMin, numberMax, compNumber, countTry int

func getRandomNumber(numberMin, numberMax int) int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("\nYou are chiter!\n")
			menu()
		}
	}()
	return numberMin + rand.Intn((numberMax)-(numberMin))
}

func inputData(msg string) string {
	var data string
	fmt.Print(msg)
	fmt.Scanln(&data)
	if data == "q" {
		os.Exit(3)
	}
	if data == "m" {
		menu()
	}
	return data
}

func print(value string) {
	switch value {
	case "win":
		fmt.Println("\nWin witn", countTry, "try. And number -", compNumber, "\n")
	case "comp":
		fmt.Printf("\nYour number %d?\n", compNumber)
	case "userMinus":
		fmt.Println("\nYour number < computer number. Try again!")
	case "userPlus":
		fmt.Println("\nYour number > computer number. Try again!")
	case "err":
		fmt.Println("\nWRONG INPUT!")
	default:
		fmt.Println("error")
	}
}

func variant1() {
	fmt.Println("\nThink of a number in the range of 1 to about 100")
	print("comp")
	for {
		inputSymbol := inputData("\nType 'm' for menu, 'q' for quit.\nEnter symbol < = >: ")
		if inputSymbol == "<" || inputSymbol == "=" || inputSymbol == ">" {
			countTry++
		}
		switch inputSymbol {
		case "=":
			print("win")
			menu()
		case ">":
			numberMin = compNumber + 1
			compNumber = getRandomNumber(numberMin, numberMax)
			print("comp")
		case "<":
			numberMax = compNumber
			compNumber = getRandomNumber(numberMin, numberMax)
			print("comp")
		default:
			print("err")
		}
	}
}

func variant2() {
	fmt.Println("\nComputer think of a number in the range of 0 to about 100")
	compNumber = getRandomNumber(numberMin, numberMax)
	for {
		userNumber, err := strconv.Atoi(inputData("\nType 'm' for menu, 'q' for quit.\nEnter your numer: "))
		if err != nil {
			print("err")
		}
		countTry++
		if compNumber == userNumber {
			print("win")
			menu()
		}
		if compNumber > userNumber {
			print("userMinus")
		}
		if compNumber < userNumber {
			print("userPlus")
		}
	}
}

func menu() {
	numberMin, numberMax = 1, 100
	compNumber = getRandomNumber(numberMin, numberMax)
	countTry = 0
	for {
		inputVariant := inputData("Type 'q' for quit.\nEnter 1 if you think a number. Enter 2 if computer think a number: ")
		switch inputVariant {
		case "1":
			variant1()
		case "2":
			variant2()
		default:
			print("err")
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Guessing game!")
	menu()
}
