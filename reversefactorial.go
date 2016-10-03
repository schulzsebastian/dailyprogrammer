package main

import (
	"bufio"
	"fmt"
	"os"
    "strconv"
)

func reverse_factorial(text string) string {
    x, _ := strconv.Atoi(text)
    i := 1
    for x > 1 {
        x = x / i
        i += 1
        if x == 1 {
            return text + " = " + strconv.Itoa(i - 1) + "!"
        }
    }
    return text + " NONE"
}

func main() {
	fileHandle, _ := os.Open("reversefactorial_input.txt")
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)
	for fileScanner.Scan() {
        fmt.Println(reverse_factorial(fileScanner.Text()))
	}
}
