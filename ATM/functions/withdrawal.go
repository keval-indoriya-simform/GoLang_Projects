package Functions

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

// This function will Debit ammount from account
func Withdrawal(user *User) {
	var inputReader = bufio.NewReader(os.Stdin)
	typeOfTransaction := "Debit"
	openingBalance := user.Balance
	fmt.Println(strings.Repeat("-", 58))
	fmt.Println("Enter amount you want to Withdrawal :")
	amtStr, _ := inputReader.ReadString('\n')
	amt, ok := AmmountToFloat(amtStr)
	if ok {
		if amt >= 0 {
			if amt >= 500 {
				if amt > 50000 {
					fmt.Println(strings.Repeat("-", 58))
					fmt.Println("You cant not withdrawal amount grater than 50000 Rs.")
					fmt.Println(strings.Repeat("-", 58))
				} else if math.Mod(amt, 500) == 0 {
					// To check if ammount is in multiple of 100

					if amt <= user.Balance-500 {
						// To check user has atleast 500 Rs. after completion of Debit
						// if yes do debit other wise cancle that
						user.Balance -= amt
						CreateLog(user, typeOfTransaction, amt, openingBalance, user.Balance)
						fmt.Println(strings.Repeat("-", 58))
						fmt.Println("Transaction Successful")
						ViewBalance(user)
					} else {
						fmt.Println(strings.Repeat("-", 58))
						fmt.Println("Transaction Failed")
						fmt.Println(strings.Repeat("-", 58))
						fmt.Println("Insufficiant Balance!!\nyou also need to maintain minimum 500 in account")
						fmt.Println(strings.Repeat("-", 58))
					}

				} else {
					fmt.Println(strings.Repeat("-", 58))
					fmt.Println("Transaction Failed")
					fmt.Println("Amount sholud be multiple of 500")
					fmt.Println(strings.Repeat("-", 58))
				}

			} else {
				fmt.Println(strings.Repeat("-", 58))
				fmt.Println("Transaction Failed")
				fmt.Println("Minimum Withdrawal have to be 500 Rs.")
				fmt.Println(strings.Repeat("-", 58))
			}

		} else {
			fmt.Println(strings.Repeat("-", 58))
			fmt.Println("Transaction Failed")
			fmt.Println("Invalid Amount")
			fmt.Println(strings.Repeat("-", 58))
		}
	}

	if IsContinue() {
		return
	} else {
		os.Exit(0)
	}
}
