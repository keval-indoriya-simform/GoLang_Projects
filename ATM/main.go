package main

import (
	"bufio"
	"fmt"
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
	transactionId     int
	DateTime          string
	TypeOfTransaction string
	Amount            float64
	OpeningBalance    float64
	ClosingBalance    float64
}

type ById []Log

func (a ById) Len() int           { return len(a) }
func (a ById) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ById) Less(i, j int) bool { return a[i].transactionId > a[j].transactionId }

var inputReader = bufio.NewReader(os.Stdin)

func main() {
	var (
		try int = 1
	)

	users := []User{
		{555544443333, "Keval", 7123, 50000, []Log{
			{1, "08/10/2022", "Credit", 25000.00, 0.00, 25000.00},
			{2, "08/12/2022", "Credit", 25000.00, 25000.00, 50000.00},
		}},
		{111122223333, "Meet", 365, 45000.00, []Log{
			{1, "10/02/2023", "Credit", 25000.00, 0.00, 25000.00},
			{2, "18/03/2023", "Credit", 20000.00, 25000.00, 45000.00},
		}},
		{123443211234, "Hari", 6454, 80000.00, []Log{
			{1, "08/10/2022", "Credit", 35000.00, 0.00, 35000.00},
			{2, "08/12/2022", "Credit", 45000.00, 35000.00, 80000.00},
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
			fmt.Println(pin)
			if ok {
				if pin == users[index].PIN {
					for {
						fmt.Println(strings.Repeat("=", 30))
						fmt.Println("Hello,", users[index].UserName)
						fmt.Println(strings.Repeat("=", 30))
						fmt.Println("Main Menu")
						fmt.Println(strings.Repeat("=", 30))
						fmt.Println("1. >> View Balance")
						fmt.Println("2. >> View Transactions")
						fmt.Println("3. >> Deposit Money")
						fmt.Println("4. >> withdrawal Money")
						fmt.Println("5. >> Exit")
						fmt.Println(strings.Repeat("=", 30))

					choiseDiv:
						fmt.Println("Enter which action you want to perform :")
						choiseStr, _ := inputReader.ReadString('\n')
						choise, ok := strToNum(choiseStr)
						if ok {
							switch choise {
							case 1:
								viewBalance(&users[index])
							case 2:
								viewTransactions(&users[index])
							case 3:
								Deposit(&users[index])
							case 4:
								withdrawal(&users[index])
							case 5:
								os.Exit(0)
							default:
								fmt.Println(strings.Repeat("-", 30))
								fmt.Println("You Enterd Wrong choise\nplease enter only values whice are present in screen")
								fmt.Println(strings.Repeat("-", 30))
								goto choiseDiv
							}
						} else {
							goto choiseDiv
						}

					}
				} else {
					try++
					if try == 4 {
						fmt.Println(strings.Repeat("-", 30))
						fmt.Println("You have exceeded the pin trying limit")
						fmt.Println(strings.Repeat("-", 30))

						os.Exit(0)
					} else {
						fmt.Println(strings.Repeat("-", 30))
						fmt.Println("This is Try number", try)
						fmt.Println(strings.Repeat("-", 30))

					}
					goto pinDiv
				}
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
		// 				fmt.Println(strings.Repeat("=", 30))
		// 				fmt.Println("MAIN MENU")
		// 				fmt.Println(strings.Repeat("=", 30))
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
	fmt.Println(strings.Repeat("-", 30))
	fmt.Println("You Enter alphabets but required number only")
	fmt.Println(strings.Repeat("-", 30))
	return 0, false
}

func ammountToFloat(str string) (float64, bool) {
	str = strings.Trim(str, "\n")
	str = strings.Trim(str, " ")
	num, err := strconv.ParseFloat(str, 64)
	if err == nil {
		return num, true
	}
	fmt.Println(strings.Repeat("-", 30))
	fmt.Println("You Enter alphabets but required number only")
	fmt.Println(strings.Repeat("-", 30))
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
	fmt.Println(strings.Repeat("-", 30))
	fmt.Println("Your Balance is :", user.Balance)
	fmt.Println(strings.Repeat("-", 30))

}

func viewTransactions(user *User) {
	fmt.Println(strings.Repeat("-", 110))
	fmt.Printf("%-20s %-20s %20s %20s %20s\n", "Date", "Type of Transaction", "Amount", "Opening Balance", "Closing Balance")
	fmt.Println(strings.Repeat("-", 110))
	sort.Sort(ById(user.logs))
	for _, log := range user.logs {
		fmt.Printf("%-20s %-20s %20v %20v %20v\n", log.DateTime, log.TypeOfTransaction, log.Amount, log.OpeningBalance, log.ClosingBalance)
	}
	fmt.Println(strings.Repeat("-", 110))
}

func Deposit(user *User) {
	typeOfTransaction := "Credit"
	openingBalance := user.Balance
	fmt.Println(strings.Repeat("-", 30))
	fmt.Println("Enter amount you want to Deposite :")
	amtStr, _ := inputReader.ReadString('\n')
	amt, ok := ammountToFloat(amtStr)
	if ok {
		if amt > 0 {
			if amt > 100 {
				user.Balance += amt
				closingBalance := user.Balance
				createLog(user, typeOfTransaction, amt, openingBalance, closingBalance)
				viewBalance(user)
			} else {
				fmt.Println(strings.Repeat("-", 30))
				fmt.Println("You Have to Deposite Minimum 100 Rs.")
				fmt.Println(strings.Repeat("-", 30))
			}
		}
	}
}

func withdrawal(user *User) {
	fmt.Println("withdrawal")
}

func createLog(user *User,typeOfTransaction string, amount float64 ,openingBalance float64, closingBalance float64) {
	date = time.Now()
	fmt.Println(date.Format("01/02/2006 15:04:05"))
	// sort.Sort(ById(user.logs))
	// newId := user.logs[0].transactionId + 1
	// log := {newId, }
}
