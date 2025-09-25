// Program to print index and value of each element in os.Args
package main

import (
	"fmt"
	"os"
)

func main() {
	for index, value := range os.Args {
		fmt.Printf("%d\t%s\n", index, value)
	}
}
