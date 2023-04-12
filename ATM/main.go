package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const LengthOfSeperator = 58 // This is lenght of seperator used for seperating input and outputs

func main() {

	var (
		inputReader     = bufio.NewReader(os.Stdin)
		try         int = 1 // This variable is used to look how many times user enterd pin
	)
	fmt.Println("Changed")

	// This slice will store user structs
	users := []User{
		{555544443333, "Keval", 7123, 50000, []Log{
			{1, "08/10/2022 15:04:05", "Credit", 25000.00, 0.00, 25000.00},
			{2, "08/12/2022 12:40:25", "Credit", 25000.00, 25000.00, 50000.00},
		}},
		{111122223333, "Meet", 3665, 45000.00, []Log{
			{1, "10/02/2023 15:04:05", "Credit", 25000.00, 0.00, 25000.00},
			{2, "18/03/2023 12:40:44", "Credit", 20000.00, 25000.00, 45000.00},
		}},
		{123443211234, "Hari", 6454, 80000.00, []Log{
			{1, "08/10/2022 18:04:05", "Credit", 35000.00, 0.00, 35000.00},
			{2, "08/12/2022 13:15:30", "Credit", 45000.00, 35000.00, 80000.00},
		}},
	}

cardNumDiv:
	fmt.Println(strings.Repeat("-", LengthOfSeperator))
	fmt.Println("Enter Card Number : ")
	cardStr, _ := inputReader.ReadString('\n') // This will read string from consol
	cardStr = strings.Trim(cardStr, "\n")
	cardStr = strings.Trim(cardStr, " ")
	cardNo, ok := StrToNum(cardStr) // This will convert string to int
	if ok {
		index := UserIndex(users, cardNo) // This function will get the index of user who has that card number
		if index >= 0 {

			// This is pin entering portion
		pinDiv:
			fmt.Println(strings.Repeat("-", LengthOfSeperator))
			fmt.Println("Enter 4 Digit PIN")
			pinStr, _ := inputReader.ReadString('\n') // This will read string from consol
			pin, ok := StrToNum(pinStr)               // This will convert string to int
			pinStr = strings.Trim(pinStr, "\n")
			pinStr = strings.Trim(pinStr, " ")
			if ok {
				// Checking if given pin is correct or not
				if pin == users[index].PIN {
					for {
						// This portion will display menu on consol
						fmt.Println(strings.Repeat("=", LengthOfSeperator))
						fmt.Printf("%27s, %-27s\n", "Hello", users[index].UserName)
						fmt.Println(strings.Repeat("=", LengthOfSeperator))
						fmt.Printf("%32s\n", "Main menu")
						fmt.Println(strings.Repeat("=", LengthOfSeperator))
						fmt.Println("1. >> View Balance")
						fmt.Println("2. >> View Transactions")
						fmt.Println("3. >> Deposit Money")
						fmt.Println("4. >> withdrawal Money")
						fmt.Println("5. >> Exit")
						fmt.Println(strings.Repeat("=", LengthOfSeperator))

						// This portion will call perticular functions that operation will be performed
					choiseDiv:
						fmt.Println("Enter which action you want to perform :")
						choiseStr, _ := inputReader.ReadString('\n') // This will read string from consol
						choise, ok := StrToNum(choiseStr)            // This will convert string to int
						if ok {
							switch choise {
							case 1:
								ViewBalance(&users[index]) // This function will display the user balance
								if IsContinue() {
									continue
								} else {
									os.Exit(0)
								}
							case 2:
								ViewTransactions(&users[index]) // This function will display all the transaction on consol
							case 3:
								Deposit(&users[index]) // This function will Credit ammount into your account
							case 4:
								Withdrawal(&users[index]) // This function will Debit ammount from your account
							case 5:
								fmt.Println(strings.Repeat("-", LengthOfSeperator))
								fmt.Println("Thank you for chossing our Bank")
								fmt.Println(strings.Repeat("-", LengthOfSeperator))
								os.Exit(0) // to exit from program
							default:
								fmt.Println(strings.Repeat("-", LengthOfSeperator))
								fmt.Println("You Entered Wrong choise\nPlease enter only values whice are present in Menu")
								fmt.Println(strings.Repeat("-", LengthOfSeperator))
								goto choiseDiv
							}
						} else {
							// if you enter somethingthat is not in option this will take you back to main menu section
							goto choiseDiv
						}

					}
				} else {
					// This else part will keep track of pin tries if you give your pin wrong three times
					// then it will exit the program
					try++
					if try == 4 {
						fmt.Println(strings.Repeat("-", LengthOfSeperator))
						fmt.Println("You have exceeded the PIN trying limit")
						fmt.Println(strings.Repeat("-", LengthOfSeperator))
						fmt.Println("Thank you for chossing our Bank")
						fmt.Println(strings.Repeat("-", LengthOfSeperator))
						os.Exit(0)
					} else {
						fmt.Println(strings.Repeat("-", LengthOfSeperator))
						if len(pinStr) != 4 {
							fmt.Println("PIN length should be 4")
							fmt.Println(strings.Repeat("-", LengthOfSeperator))
						}
						fmt.Println("This is Try number", try, "(After this", 3-try, "try will be left)")

					}
					goto pinDiv
				}
			} else {
				try++
				if try == 4 {
					fmt.Println(strings.Repeat("-", LengthOfSeperator))
					fmt.Println("You have exceeded the PIN trying limit")
					fmt.Println(strings.Repeat("-", LengthOfSeperator))
					fmt.Println("Thank you for chossing our Bank")
					fmt.Println(strings.Repeat("-", LengthOfSeperator))
					os.Exit(0)
				}
				fmt.Println("This is Try number", try, "(After this", 3-try, "try will be left)")
				goto pinDiv
			}
		} else if len(cardStr) != 12 {
			// To check you entered valid length of numbers in card number
			fmt.Println(strings.Repeat("-", LengthOfSeperator))
			fmt.Println("Please enter Valid 12 digit Card Number")
			goto cardNumDiv
		} else {
			// If you enter any wrong card number porgram will get close
			fmt.Println(strings.Repeat("-", LengthOfSeperator))
			fmt.Println("Card Number is Not Found!! (User Not Found)")
			fmt.Println(strings.Repeat("-", LengthOfSeperator))
			fmt.Println("Thank you for chossing our Bank")
			fmt.Println(strings.Repeat("-", LengthOfSeperator))
			os.Exit(0)
		}

	}

}
