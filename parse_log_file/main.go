package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	var sum map[string]int
	domains := []string{}
	var total int
	var lines int

	sum = make(map[string]int)

	for in.Scan() {
		lines++
		fields := strings.Fields(in.Text())

		if len(fields) != 2 {
			fmt.Printf("Wrong Input: %v (line #%d\n", fields, lines)
			return
		}

		domain := fields[0]
		visits, err := strconv.Atoi(fields[1])
		if visits < 0 || err != nil {
			fmt.Printf("Wrong Input: %v (line #%d\n", fields[1], lines)
			return
		}

		if _, ok := sum[domain]; !ok {
			domains = append(domains, domain)
		}
		total += visits
		sum[domain] += visits
	}

	fmt.Printf("%-30s %10s", "DOMAIN", "VISITS\n")
	fmt.Println(strings.Repeat("-", 50))

	sort.Strings(domains)
	for _, domain := range domains {
		visits := sum[domain]
	fmt.Printf("%-30s %10d\n", domain, visits)

	}
	fmt.Printf("\n%-30s %10d\n", "Total", total)

	if err :=  in.Err(); err != nil {
		fmt.Println("> Err:", err)
	}
}
