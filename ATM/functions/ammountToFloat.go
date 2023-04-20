package functions

import (
	"fmt"
	"strconv"
	"strings"
)

// This function will convert given string to float64
func AmmountToFloat(str string) (float64, bool) {
	str = strings.Trim(str, "\n")
	str = strings.Trim(str, " ")
	num, err := strconv.ParseFloat(str, 64)
	if err == nil {
		return num, true
	}
	fmt.Println(strings.Repeat("-", 58))
	fmt.Println("You Entered Something else but only numbers are required")
	fmt.Println(strings.Repeat("-", 58))
	return 0, false
}
