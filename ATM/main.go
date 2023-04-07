package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

// This is a struct for storing user details
type User struct {
	CardNum  int
	UserName string
	PIN      int
	Balance  float64
	logs     []Log
}

// This is a struct for storing log for user's transaction
type Log struct {
	TransactionId     int
	DateTime          string
	TypeOfTransaction string
	Amount            float64
	OpeningBalance    float64
	ClosingBalance    float64
}

// This methods are for sorting the struct slice by transactionId
type ById []Log

func (a ById) Len() int           { return len(a) }
func (a ById) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ById) Less(i, j int) bool { return a[i].TransactionId > a[j].TransactionId }

// This is an reader which takes input from user
var inputReader = bufio.NewReader(os.Stdin)

const (
	lengthOfSeperator = 58 // This is lenght of seperator used for seperating input and outputs
)

func main() {

	var (
		try int = 1 // This variable is used to look how many times user enterd pin
	)

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
	fmt.Println(strings.Repeat("-", lengthOfSeperator))
	fmt.Println("Enter Card Number : ")
	cardStr, _ := inputReader.ReadString('\n') // This will read string from consol
	cardStr = strings.Trim(cardStr, "\n")
	cardStr = strings.Trim(cardStr, " ")
	cardNo, ok := strToNum(cardStr) // This will convert string to int
	if ok {
		index := userIndex(users, cardNo) // This function will get the index of user who has that card number
		if index >= 0 {

			// This is pin entering portion
		pinDiv:
			fmt.Println(strings.Repeat("-", lengthOfSeperator))
			fmt.Println("Enter 4 Digit PIN")
			pinStr, _ := inputReader.ReadString('\n') // This will read string from consol
			pin, ok := strToNum(pinStr)               // This will convert string to int
			pinStr = strings.Trim(pinStr, "\n")
			pinStr = strings.Trim(pinStr, " ")
			if ok {
				// Checking if given pin is correct or not
				if pin == users[index].PIN {
					for {
						// This portion will display menu on consol
						fmt.Println(strings.Repeat("=", lengthOfSeperator))
						fmt.Printf("%27s, %-27s\n", "Hello", users[index].UserName)
						fmt.Println(strings.Repeat("=", lengthOfSeperator))
						fmt.Printf("%32s\n", "Main menu")
						fmt.Println(strings.Repeat("=", lengthOfSeperator))
						fmt.Println("1. >> View Balance")
						fmt.Println("2. >> View Transactions")
						fmt.Println("3. >> Deposit Money")
						fmt.Println("4. >> withdrawal Money")
						fmt.Println("5. >> Exit")
						fmt.Println(strings.Repeat("=", lengthOfSeperator))

						// This portion will call perticular functions that operation will be performed
					choiseDiv:
						fmt.Println("Enter which action you want to perform :")
						choiseStr, _ := inputReader.ReadString('\n') // This will read string from consol
						choise, ok := strToNum(choiseStr)            // This will convert string to int
						if ok {
							switch choise {
							case 1:
								viewBalance(&users[index]) // This function will display the user balance
								if isContinue() {
									continue
								} else {
									os.Exit(0)
								}
							case 2:
								viewTransactions(&users[index]) // This function will display all the transaction on consol
							case 3:
								Deposit(&users[index]) // This function will Credit ammount into your account
							case 4:
								withdrawal(&users[index]) // This function will Debit ammount from your account
							case 5:
								fmt.Println(strings.Repeat("-", lengthOfSeperator))
								fmt.Println("Thank you for chossing our Bank")
								fmt.Println(strings.Repeat("-", lengthOfSeperator))
								os.Exit(0) // to exit from program
							default:
								fmt.Println(strings.Repeat("-", lengthOfSeperator))
								fmt.Println("You Entered Wrong choise\nPlease enter only values whice are present in Menu")
								fmt.Println(strings.Repeat("-", lengthOfSeperator))
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
						fmt.Println(strings.Repeat("-", lengthOfSeperator))
						fmt.Println("You have exceeded the PIN trying limit")
						fmt.Println(strings.Repeat("-", lengthOfSeperator))
						fmt.Println("Thank you for chossing our Bank")
						fmt.Println(strings.Repeat("-", lengthOfSeperator))
						os.Exit(0)
					} else {
						fmt.Println(strings.Repeat("-", lengthOfSeperator))
						if len(pinStr) != 4 {
							fmt.Println("PIN length should be 4")
							fmt.Println(strings.Repeat("-", lengthOfSeperator))
						}
						fmt.Println("This is Try number", try, "(After this", 3-try, "try will be left)")
						// fmt.Println(strings.Repeat("-", lengthOfSeperator))

					}
					goto pinDiv
				}
			} else {
				try++
				if try == 4 {
					fmt.Println(strings.Repeat("-", lengthOfSeperator))
					fmt.Println("You have exceeded the PIN trying limit")
					fmt.Println(strings.Repeat("-", lengthOfSeperator))
					fmt.Println("Thank you for chossing our Bank")
					fmt.Println(strings.Repeat("-", lengthOfSeperator))
					os.Exit(0)
				}
				fmt.Println("This is Try number", try, "(After this", 3-try, "try will be left)")
				// fmt.Println(strings.Repeat("-", lengthOfSeperator))
				goto pinDiv
			}
		} else if len(cardStr) != 12 {
			// To check you entered valid length of numbers in card number
			fmt.Println(strings.Repeat("-", lengthOfSeperator))
			fmt.Println("Please enter Valid 12 digit Card Number")
			goto cardNumDiv
		} else {
			// If you enter any wrong card number porgram will get close
			fmt.Println(strings.Repeat("-", lengthOfSeperator))
			fmt.Println("Card Number is Not Found!! (User Not Found)")
			fmt.Println(strings.Repeat("-", lengthOfSeperator))
			fmt.Println("Thank you for chossing our Bank")
			fmt.Println(strings.Repeat("-", lengthOfSeperator))
			os.Exit(0)
		}

	}

}

// This function will convert given string to int
func strToNum(str string) (int, bool) {
	str = strings.Trim(str, "\n")
	str = strings.Trim(str, " ")
	num, err := strconv.Atoi(str)
	if err == nil {
		return num, true
	}
	fmt.Println(strings.Repeat("-", lengthOfSeperator))
	fmt.Println("You Entered Something else but only numbers are required")
	fmt.Println(strings.Repeat("-", lengthOfSeperator))
	return 0, false
}

// This function will convert given string to float64
func ammountToFloat(str string) (float64, bool) {
	str = strings.Trim(str, "\n")
	str = strings.Trim(str, " ")
	num, err := strconv.ParseFloat(str, 64)
	if err == nil {
		return num, true
	}
	fmt.Println(strings.Repeat("-", lengthOfSeperator))
	fmt.Println("You Entered Something else but only numbers are required")
	fmt.Println(strings.Repeat("-", lengthOfSeperator))
	return 0, false
}

// This function will return the index where your card number is matched
// if not found it will return -1
func userIndex(users []User, cardNo int) int {
	for i, user := range users {
		if user.CardNum == cardNo {
			return i
		}
	}
	return -1
}

// This function will display the balance of that perticular user in consol
// user is passed by reference
func viewBalance(user *User) {
	fmt.Println(strings.Repeat("-", lengthOfSeperator))
	fmt.Printf("Your Balance is : %.2f\n", user.Balance)
	fmt.Println(strings.Repeat("-", lengthOfSeperator))
}

// This function will display all the logs of that perticular user in consol
func viewTransactions(user *User) {
	fmt.Println(strings.Repeat("-", 110))
	fmt.Printf("%-25s %-20s %20s %20s %20s\n", "Date and Time", "Type of Transaction", "Amount", "Opening Balance", "Closing Balance")
	fmt.Println(strings.Repeat("-", 110))
	sort.Sort(ById(user.logs))
	for _, log := range user.logs {
		fmt.Printf("%-25s %-20s %20.2f %20.2f %20.2f\n", log.DateTime, log.TypeOfTransaction, log.Amount, log.OpeningBalance, log.ClosingBalance)
	}
	fmt.Println(strings.Repeat("-", 110))
	if isContinue() {
		return
	} else {
		os.Exit(0)
	}
}

// This function will Credit ammount into perticular user's account
func Deposit(user *User) {
depositDiv:
	typeOfTransaction := "Credit"
	openingBalance := user.Balance
	fmt.Println(strings.Repeat("-", lengthOfSeperator))
	fmt.Println("Enter amount you want to Deposite :")
	amtStr, _ := inputReader.ReadString('\n')
	amt, ok := ammountToFloat(amtStr)
	if ok {
		if amt > 0 {
			if amt > 50000 {
				fmt.Println(strings.Repeat("-", lengthOfSeperator))
				fmt.Println("Deposite ammount sholud be less than 50000 Rs.")
				fmt.Println(strings.Repeat("-", lengthOfSeperator))
			} else if amt >= 500 {
				if math.Mod(amt, 500) == 0 {
					user.Balance += amt
					createLog(user, typeOfTransaction, amt, openingBalance, user.Balance)
					fmt.Println(strings.Repeat("-", lengthOfSeperator))
					fmt.Println("Transaction Successful")
					viewBalance(user)
					if isContinue() {
						return
					} else {
						os.Exit(0)
					}
				} else {
					fmt.Println(strings.Repeat("-", lengthOfSeperator))
					fmt.Println("Transaction Failed")
					fmt.Println("Amount sholud be multiple of 500")
					fmt.Println(strings.Repeat("-", lengthOfSeperator))
				}
			} else {
				fmt.Println(strings.Repeat("-", lengthOfSeperator))
				fmt.Println("Transaction Failed")
				fmt.Println("You have to Deposite Minimum 500 Rs.")
				fmt.Println(strings.Repeat("-", lengthOfSeperator))
			}
		} else {
			fmt.Println(strings.Repeat("-", lengthOfSeperator))
			fmt.Println("Transaction Failed")
			fmt.Println("Invalid Amount")
			fmt.Println(strings.Repeat("-", lengthOfSeperator))
		}
	}
	if isContinue() {
		goto depositDiv
	} else {
		os.Exit(0)
	}
}

// This function will Debit ammount from account
func withdrawal(user *User) {
withdrawalDiv:
	typeOfTransaction := "Debit"
	openingBalance := user.Balance
	fmt.Println(strings.Repeat("-", lengthOfSeperator))
	fmt.Println("Enter amount you want to Withdrawal :")
	amtStr, _ := inputReader.ReadString('\n')
	amt, ok := ammountToFloat(amtStr)
	if ok {
		if amt > 0 {
			if amt >= 500 {
				if amt > 50000 {
					fmt.Println(strings.Repeat("-", lengthOfSeperator))
					fmt.Println("You cant not withdrawal amount grater than 50000 Rs.")
					fmt.Println(strings.Repeat("-", lengthOfSeperator))
				} else if math.Mod(amt, 500) == 0 {
					// To check if ammount is in multiple of 100

					if amt <= user.Balance-500 {
						// To check user has atleast 500 Rs. after completion of Debit
						// if yes do debit other wise cancle that
						user.Balance -= amt
						createLog(user, typeOfTransaction, amt, openingBalance, user.Balance)
						fmt.Println(strings.Repeat("-", lengthOfSeperator))
						fmt.Println("Transaction Successful")
						viewBalance(user)
						if isContinue() {
							return
						} else {
							os.Exit(0)
						}
					} else {
						fmt.Println(strings.Repeat("-", lengthOfSeperator))
						fmt.Println("Transaction Failed")
						fmt.Println(strings.Repeat("-", lengthOfSeperator))
						fmt.Println("Insufficiant Balance!!\nyou also need to maintain minimum 500 in account")
						fmt.Println(strings.Repeat("-", lengthOfSeperator))
					}

				} else {
					fmt.Println(strings.Repeat("-", lengthOfSeperator))
					fmt.Println("Transaction Failed")
					fmt.Println("Amount sholud be multiple of 500")
					fmt.Println(strings.Repeat("-", lengthOfSeperator))
				}

			} else {
				fmt.Println(strings.Repeat("-", lengthOfSeperator))
				fmt.Println("Transaction Failed")
				fmt.Println("Minimum Withdrawal have to be 500 Rs.")
				fmt.Println(strings.Repeat("-", lengthOfSeperator))
			}

		} else {
			fmt.Println(strings.Repeat("-", lengthOfSeperator))
			fmt.Println("Transaction Failed")
			fmt.Println("Invalid Amount")
			fmt.Println(strings.Repeat("-", lengthOfSeperator))
		}
	}

	if isContinue() {
		goto withdrawalDiv
	} else {
		os.Exit(0)
	}
}

// This function will do the enry in log slice if Cradit or Debit is occured
func createLog(user *User, typeOfTransaction string, amount float64, openingBalance float64, closingBalance float64) {
	date := time.Now()
	newdate := date.Format("01/02/2006 15:04:05")
	sort.Sort(ById(user.logs))
	newId := user.logs[0].TransactionId + 1
	newlog := Log{newId, newdate, typeOfTransaction, amount, openingBalance, closingBalance}
	user.logs = append(user.logs, newlog)
}

// This function is used when any transcation is get finishit and we wnat to ask user
// if he want to continue the another transaction or hw want to exit the program
func isContinue() bool {
continueTransaction:
	fmt.Println("Do you want to Continue (yes,y / no,n) :")
	val, _ := inputReader.ReadString('\n')
	val = strings.Trim(val, "\n")
	val = strings.Trim(val, " ")
	val = strings.ToLower(val)

	switch val {
	case "yes", "y":
		return true
	case "no", "n":
		fmt.Println(strings.Repeat("-", lengthOfSeperator))
		fmt.Println("Thank you for chossing our Bank")
		fmt.Println(strings.Repeat("-", lengthOfSeperator))
		return false
	default:
		fmt.Println(strings.Repeat("-", lengthOfSeperator))
		fmt.Println("You Entered something else please answer in yes and no")
		fmt.Println(strings.Repeat("-", lengthOfSeperator))
		goto continueTransaction
	}
}
