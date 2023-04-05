package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	link  = "http://"
	nlink = len(link)
	mask  = '*'
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	str, _ := inputReader.ReadString('\n')
	var (
		size   = len(str)
		buf = []byte(str)
		// buf    = make([]byte, 0, size)
		inLink bool
	)

	// buf = buf[:0]

	for i := 0; i < size; i++ {

		if len(str[i:]) >= nlink && str[i:i+nlink] == link {
			inLink = true

			// buf = append(buf, link...)
			i += nlink
		}

		c := str[i]

		switch c {
		case ' ', '\t', '\n':
			inLink = false
		}

		if inLink {
			// c = mask
			buf[i] = mask
		}

		// buf = append(buf, c)
	}
	fmt.Println(string(buf))
}
