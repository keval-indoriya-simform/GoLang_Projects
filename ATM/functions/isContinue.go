package functions

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// This function is used when any transcation is get finishit and we want to ask user
// if he want to continue the another transaction or hw want to exit the program
func IsContinue() bool {
	var inputReader = bufio.NewReader(os.Stdin)
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
		fmt.Println(strings.Repeat("-", 58))
		fmt.Println("Thank you for chossing our Bank")
		fmt.Println(strings.Repeat("-", 58))
		return false
	default:
		fmt.Println(strings.Repeat("-", 58))
		fmt.Println("You Entered something else please answer in yes and no")
		fmt.Println(strings.Repeat("-", 58))
		goto continueTransaction
	}
}
