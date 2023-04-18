package Functions

import (
	"sort"
	"time"
)

// This function will do the enry in log slice if Cradit or Debit is occured
func CreateLog(user *User, typeOfTransaction string, amount float64, openingBalance float64, closingBalance float64) {
	date := time.Now()
	newdate := date.Format("01/02/2006 15:04:05")
	sort.Sort(ById(user.logs))
	newId := user.logs[0].TransactionId + 1
	newlog := Log{newId, newdate, typeOfTransaction, amount, openingBalance, closingBalance}
	user.logs = append(user.logs, newlog)
}
