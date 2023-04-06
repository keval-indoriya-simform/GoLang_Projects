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

type User struct {
	CardNum  int
	UserName string
	PIN      int
	Balance  float64
	logs     []Log
}

type Log struct {
	TransactionId     int
	DateTime          string
	TypeOfTransaction string
	Amount            float64
	OpeningBalance    float64
	ClosingBalance    float64
}

type ById []Log

func (a ById) Len() int           { return len(a) }
func (a ById) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ById) Less(i, j int) bool { return a[i].TransactionId > a[j].TransactionId }

var inputReader = bufio.NewReader(os.Stdin)

const (
	lengthOfSeperator = 40
)

func main() {
	var (
		try int = 1
	)

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

	fmt.Println("Enter Card Number : ")
	cardStr, _ := inputReader.ReadString('\n')

	cardNo, ok := strToNum(cardStr)
	if ok {
		index := userIndex(users, cardNo)
		if index >= 0 {

		pinDiv:
			fmt.Println("Enter 4 Digit PIN")
			pinStr, _ := inputReader.ReadString('\n')
			pin, ok := strToNum(pinStr)
			pinStr = strings.Trim(pinStr, "\n")
			pinStr = strings.Trim(pinStr, " ")
			if ok {
				if pin == users[index].PIN {
					for {
						fmt.Println(strings.Repeat("=", lengthOfSeperator))
						fmt.Println("Hello,", users[index].UserName)
						fmt.Println(strings.Repeat("=", lengthOfSeperator))
						fmt.Println("Main Menu")
						fmt.Println(strings.Repeat("=", lengthOfSeperator))
						fmt.Println("1. >> View Balance")
						fmt.Println("2. >> View Transactions")
						fmt.Println("3. >> Deposit Money")
						fmt.Println("4. >> withdrawal Money")
						fmt.Println("5. >> Exit")
						fmt.Println(strings.Repeat("=", lengthOfSeperator))

					choiseDiv:
						fmt.Println("Enter which action you want to perform :")
						choiseStr, _ := inputReader.ReadString('\n')
						choise, ok := strToNum(choiseStr)
						if ok {
							switch choise {
							case 1:
								viewBalance(&users[index])
								if isContinue() {
									continue
								} else {
									os.Exit(0)
								}
							case 2:
								viewTransactions(&users[index])
							case 3:
								Deposit(&users[index])
							case 4:
								withdrawal(&users[index])
							case 5:
								os.Exit(0)
							default:
								fmt.Println(strings.Repeat("-", lengthOfSeperator))
								fmt.Println("You Enterd Wrong choise\nplease enter only values whice are present in screen")
								fmt.Println(strings.Repeat("-", lengthOfSeperator))
								goto choiseDiv
							}
						} else {
							goto choiseDiv
						}

					}
				} else {

					try++
					if try == 4 {
						fmt.Println(strings.Repeat("-", lengthOfSeperator))
						fmt.Println("You have exceeded the pin trying limit")
						fmt.Println(strings.Repeat("-", lengthOfSeperator))

						os.Exit(0)
					} else {
						fmt.Println(strings.Repeat("-", lengthOfSeperator))
						if len(pinStr) != 4 {
							fmt.Println("PIN length should be 4")
						}
						fmt.Println("This is Try number", try, "(After this", 3-try, "try will be left)")
						fmt.Println(strings.Repeat("-", lengthOfSeperator))

					}
					goto pinDiv
				}
			} else {
				try++
				if try == 4 {
					fmt.Println(strings.Repeat("-", lengthOfSeperator))
					fmt.Println("You have exceeded the pin trying limit")
					fmt.Println(strings.Repeat("-", lengthOfSeperator))

					os.Exit(0)
				}
				fmt.Println("This is Try number", try, "(After this", 3-try, "try will be left)")
				fmt.Println(strings.Repeat("-", lengthOfSeperator))
				goto pinDiv
			}
		} else {
			fmt.Println("Card Number is Not Found!! (User Not Found)")
		}

		// for _, user := range users {
		// 	if cardNo == user.CardNum {
		// 		fmt.Println("Enter 4 Digit PIN")
		// 		pinStr, _ := inputReader.ReadString('\n')
		// 		pin, ok := strToNum(pinStr)
		// 		if ok {
		// 			if pin == user.PIN {
		// 				fmt.Println("Hello,", user.UserName)
		// 				fmt.Println(strings.Repeat("=", lengthOfSeperator))
		// 				fmt.Println("MAIN MENU")
		// 				fmt.Println(strings.Repeat("=", lengthOfSeperator))
		// 				break
		// 			} else {
		// 				fmt.Println("Your Pin is wrong")
		// 			}
		// 		}
		// 	}
		// }
	}

}

func strToNum(str string) (int, bool) {
	str = strings.Trim(str, "\n")
	str = strings.Trim(str, " ")
	num, err := strconv.Atoi(str)
	if err == nil {
		return num, true
	}
	fmt.Println(strings.Repeat("-", lengthOfSeperator))
	fmt.Println("You Enter Something else but only numbers are required")
	fmt.Println(strings.Repeat("-", lengthOfSeperator))
	return 0, false
}

func ammountToFloat(str string) (float64, bool) {
	str = strings.Trim(str, "\n")
	str = strings.Trim(str, " ")
	str = strings.ReplaceAll(str, ",", "")
	num, err := strconv.ParseFloat(str, 64)
	if err == nil {
		return num, true
	}
	fmt.Println(strings.Repeat("-", lengthOfSeperator))
	fmt.Println("You Enter alphabets but required number only")
	fmt.Println(strings.Repeat("-", lengthOfSeperator))
	return 0, false
}

func userIndex(users []User, cardNo int) int {
	for i, user := range users {
		if user.CardNum == cardNo {
			return i
		}
	}
	return -1
}

func viewBalance(user *User) {
	fmt.Println(strings.Repeat("-", lengthOfSeperator))
	fmt.Printf("Your Balance is : %.2f\n", user.Balance)
	fmt.Println(strings.Repeat("-", lengthOfSeperator))
}

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

func Deposit(user *User) {
	typeOfTransaction := "Credit"
	openingBalance := user.Balance
	fmt.Println(strings.Repeat("-", lengthOfSeperator))
	fmt.Println("Enter amount you want to Deposite :")
	amtStr, _ := inputReader.ReadString('\n')
	amt, ok := ammountToFloat(amtStr)
	if ok {
		if amt > 0 {
			if amt > 100 {
				user.Balance += amt
				closingBalance := user.Balance
				createLog(user, typeOfTransaction, amt, openingBalance, closingBalance)
				fmt.Println(strings.Repeat("-", lengthOfSeperator))
				fmt.Println("Transaction Successful")
				viewBalance(user)
			} else {
				fmt.Println(strings.Repeat("-", lengthOfSeperator))
				fmt.Println("Transaction Failed")
				fmt.Println("You have to Deposite Minimum 100 Rs.")
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
		return
	} else {
		os.Exit(0)
	}
}

func withdrawal(user *User) {
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
				} else if math.Mod(amt, 100) == 0 {
					if amt <= user.Balance-500 {
						user.Balance -= amt
						closingBalance := user.Balance
						createLog(user, typeOfTransaction, amt, openingBalance, closingBalance)
						fmt.Println(strings.Repeat("-", lengthOfSeperator))
						fmt.Println("Transaction Successful")
						viewBalance(user)
					} else {
						fmt.Println(strings.Repeat("-", lengthOfSeperator))
						fmt.Println("Transaction Failed")
						fmt.Println("Insufficiant Balance!!\nyou also need to maintain minimum 500 in account")
						fmt.Println(strings.Repeat("-", lengthOfSeperator))
					}
				} else {
					fmt.Println(strings.Repeat("-", lengthOfSeperator))
					fmt.Println("Transaction Failed")
					fmt.Println("Amount sholud be multiple of 100")
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
		return
	} else {
		os.Exit(0)
	}
}

func createLog(user *User, typeOfTransaction string, amount float64, openingBalance float64, closingBalance float64) {
	date := time.Now()
	newdate := date.Format("01/02/2006 15:04:05")
	sort.Sort(ById(user.logs))
	newId := user.logs[0].TransactionId + 1
	newlog := Log{newId, newdate, typeOfTransaction, amount, openingBalance, closingBalance}
	user.logs = append(user.logs, newlog)
}

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
		return false
	default:
		fmt.Println(strings.Repeat("-", lengthOfSeperator))
		fmt.Println("You Entered something else please answer in yes and no")
		fmt.Println(strings.Repeat("-", lengthOfSeperator))
		goto continueTransaction
	}
}
