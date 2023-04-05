package main

import "fmt"

func main() {
	students := [...][3]float64{
		{25, 28, 24},
		{21, 21, 18},
		{21, 18, 18},
	}

	for _, student := range students {
		var sum float64 = 0
		for _, marks := range student {
			sum += marks
		}
		avg := sum / float64(len(student))
		fmt.Printf("%.2f\n",avg)
	}

}
