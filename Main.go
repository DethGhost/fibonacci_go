package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
	"log"
	"strings"
	"strconv"
)

func main() {

	fmt.Println("Hello in Fibonachi console. \n" +
		"It's simple game where you need enter numbers in fibonacci sequence.\n" +
		"Game will be ended when you enter correct 10 numbers or make 3 mistakes.\n" +
		"You have 10 seconds for each answer. Game will be start when you enter 'start' enter 'exit' for exit")
	input := make(chan string, 1)
	fImput, _ := bufio.NewReader(os.Stdin).ReadString('\n');
	go getInput(input)
	userError := 0
	userOk := 0
start:
	if strings.TrimRight(fImput, "\n") == "exit" {
		fmt.Println("bye")
		os.Exit(1)
	} else if strings.TrimRight(fImput, "\n") == "start" {
		fmt.Println("Yuor game is begining, enter first fibonacci number")
		Fibonacci := fibonacci()
		currentNum := 0
		for i := 0; i <= 10; i++ {
			currentNum = Fibonacci()
			select {
			case res := <-input:
				enteredNum, _ := strconv.Atoi(res);
				if currentNum != enteredNum {
					userError++
					fmt.Errorf("You have enter wrong number")
					fmt.Printf("{%d:%d} \n", i, currentNum)
					fmt.Printf("You have %d tries\n", 3-userError)
					if userError == 3 {
						fmt.Println("GAME OVER")
						fmt.Println("For start game enter 'start' for exit 'exit'")
						fImput, _ = bufio.NewReader(os.Stdin).ReadString('\n');
						strings.TrimRight(fImput, "\n");
						goto start
					}
				}
				userOk++
				if userOk == 10{
					fmt.Println("Congratulation you have enter all numbers \n" +
						"YOU WIN")
					fmt.Println("For start game enter 'start' for exit 'exit'")
					fImput, _ = bufio.NewReader(os.Stdin).ReadString('\n');
					strings.TrimRight(fImput, "\n");
					goto start
				}
				break
			case <-time.After(10 * time.Second):
				fmt.Println("bye")
				os.Exit(1)
				break
			}
		}
	} else {
		fmt.Println("Your enter is wrong. Enter 'start' or 'exit'")
		fImput, _ = bufio.NewReader(os.Stdin).ReadString('\n');
		strings.TrimRight(fImput, "\n");
		goto start
	}

}

func fibonacci() func() int {
	num2, num1 := 1, 1

	return func() int {
		tempNum := num2
		num2, num1 = num1, tempNum+num1
		return tempNum
	}
}

func getInput(input chan string) {
	for {
		in := bufio.NewReader(os.Stdin)
		result, err := in.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		input <- strings.TrimRight(result, "\n")
	}
}
