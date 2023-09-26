package repl

import (
	"bufio"
	"fmt"
	"os"
)

func startRepl(process func(input string)) {
	fmt.Println("ゴミ箱へようこそ")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()
		if input == "exit" {
			break
		}
		process(input)
	}
	fmt.Println("さようなら")
}
