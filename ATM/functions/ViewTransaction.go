package functions

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

// This function will display all the logs of that perticular user in consol
func ViewTransactions(user *User) {
	fmt.Println(strings.Repeat("-", 110))
	fmt.Printf("%-25s %-20s %20s %20s %20s\n", "Date and Time", "Type of Transaction", "Amount", "Opening Balance", "Closing Balance")
	fmt.Println(strings.Repeat("-", 110))
	sort.Sort(ById(user.logs))
	for _, log := range user.logs {
		fmt.Printf("%-25s %-20s %20.2f %20.2f %20.2f\n", log.DateTime, log.TypeOfTransaction, log.Amount, log.OpeningBalance, log.ClosingBalance)
	}
	fmt.Println(strings.Repeat("-", 110))
	if IsContinue() {
		return
	} else {
		os.Exit(0)
	}
}
