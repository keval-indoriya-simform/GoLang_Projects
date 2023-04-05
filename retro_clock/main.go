package main

import (
	"fmt"
	"time"
)

func main() {
	type placeholder [5]string

	zero := placeholder{
		"▓▓▓",
		"▓ ▓",
		"▓ ▓",
		"▓ ▓",
		"▓▓▓",
	}
	one := placeholder{
		"▓▓ ",
		" ▓ ",
		" ▓ ",
		" ▓ ",
		"▓▓▓",
	}
	two := placeholder{
		"▓▓▓",
		"  ▓",
		"▓▓▓",
		"▓  ",
		"▓▓▓",
	}

	three := placeholder{
		"▓▓▓",
		"  ▓",
		"▓▓▓",
		"  ▓",
		"▓▓▓",
	}
	four := placeholder{
		"▓ ▓",
		"▓ ▓",
		"▓▓▓",
		"  ▓",
		"  ▓",
	}
	five := placeholder{
		"▓▓▓",
		"▓  ",
		"▓▓▓",
		"  ▓",
		"▓▓▓",
	}
	six := placeholder{
		"▓▓▓",
		"▓  ",
		"▓▓▓",
		"▓ ▓",
		"▓▓▓",
	}
	seven := placeholder{
		"▓▓▓",
		"  ▓",
		"  ▓",
		"  ▓",
		"  ▓",
	}
	eight := placeholder{
		"▓▓▓",
		"▓ ▓",
		"▓▓▓",
		"▓ ▓",
		"▓▓▓",
	}
	nine := placeholder{
		"▓▓▓",
		"▓ ▓",
		"▓▓▓",
		"  ▓",
		"▓▓▓",
	}

	colon := placeholder{
		" ",
		"░",
		" ",
		"░",
		" ",
	}

	digits := [...]placeholder{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}

	for {
		currentTime := time.Now()

		hour, min, sec := currentTime.Hour(), currentTime.Minute(), currentTime.Second()

		clock := [...]placeholder{
			digits[hour/10], digits[hour%10],
			colon,
			digits[min/10], digits[min%10],
			colon,
			digits[sec/10], digits[sec%10],
		}

		for line := range clock[0] {

			for index, digit := range clock {
				next := clock[index][line]
				if digit == colon && sec%2 == 0 {
					next = " "
				}
				fmt.Print(next, " ")
				
			}
			fmt.Println("")
		}
		// if sec%10 == 0 {
		// 	fmt.Println("Alarm")
		// }
		fmt.Println("")

		time.Sleep(time.Second)

	}
}
