package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter seach term: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	FetchById(strings.TrimSpace(text))

}
