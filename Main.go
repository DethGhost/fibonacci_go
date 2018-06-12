package main

import (
	"bufio"
	"fmt"
	"os"
	"log"
	"strings"
	"strconv"
	"time"
)

func main() {

	fmt.Println("Hello in Fibonachi console. \n" +
		"It's simple game where you need enter numbers in fibonacci sequence.\n" +
		"Game will be ended when you enter correct 10 numbers or make 3 mistakes.\n" +
		"You have 10 seconds for each answer. Game will be start when you enter 'start' enter 'exit' for exit")

	gameFlow()

}

func fibonacci() func() int {
	num2, num1 := 1, 1

	return func() int {
		tempNum := num2
		num2, num1 = num1, tempNum+num1
		return tempNum
	}
}

func gameFlow() {

	fImput, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	fImput = strings.TrimRight(fImput, "\r\n")
	userError := 0
	userOk := 0

	if strings.TrimRight(fImput, "\r\n") == "exit" {
		fmt.Println("bye")
		os.Exit(1)
	} else if strings.TrimRight(fImput, "\r\n") == "start" {
		fmt.Println("Yuor game is begining, enter first fibonacci number")
		Fibonacci := fibonacci()
		position := 0
		for {
			position++
			a := game(userError, userOk, position, Fibonacci())
			userError = a[1]
			userOk = a[0]
			if userOk == 10 {
				fmt.Println("Congratulation you have enter all numbers \n" +
					"YOU WIN")
				fmt.Println("For start game enter 'start' for exit 'exit'")
break
			} else if userError == 3 {
				fmt.Println("GAME OVER")
				fmt.Println("For start game enter 'start' for exit 'exit'")
				break
			}
		}

	} else {
		fmt.Println("Your enter is wrong. Enter 'start' or 'exit'")
	}
	gameFlow()
}

func game(userError int, userOk int, position int, currentNum int) []int {

	input := make(chan string, 1)
	go getInput(input)
	fmt.Println(currentNum)
	select {
	case res := <-input:
		enteredNum, _ := strconv.Atoi(res);
		fmt.Println(enteredNum)
		if currentNum != enteredNum {
			userError++
			fmt.Println("You have enter wrong number")
			fmt.Printf("{%d:%d} \n", position, currentNum)
			fmt.Printf("You have %d tries\n", 3-userError)
		} else {
			userOk++
		}
		break
	case <-time.After(10 * time.Second):
		userError++
		fmt.Println("Times up")
		fmt.Printf("{%d:%d} \n", position, currentNum)
		fmt.Printf("You have %d tries\n", 3-userError)
		break
	}
	return []int{userOk, userError}
}

func getInput(input chan string) {
	for {
		in := bufio.NewReader(os.Stdin)
		result, err := in.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		input <- strings.TrimRight(result, "\r\n")
	}
}
