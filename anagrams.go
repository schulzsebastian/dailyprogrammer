package main

import (
	"bufio"
	"fmt"
	"os"
    "strings"
    "sort"
)

func check_equal(x, y []string) bool {
    sort.Strings(x)
    sort.Strings(y)
    for i := range x {
        if x[i] != y[i] {
            return false
        }
    }
    return true
}

func main() {
    fileHandle, _ := os.Open("anagrams_input.txt")
    defer fileHandle.Close()
    fileScanner := bufio.NewScanner(fileHandle)
    for fileScanner.Scan() {
        t := strings.Split(fileScanner.Text(), "?")
        r := strings.NewReplacer("\"", "", " ", "", "'", "")
        s := strings.Split(strings.ToLower(r.Replace(fileScanner.Text())), "?")
        if check_equal(strings.Split(s[0], ""), strings.Split(s[1], "")) {
            fmt.Println(strings.TrimSpace(t[0]), "is anagram of", strings.TrimSpace(t[1]))
        } else {
            fmt.Println(strings.TrimSpace(t[0]), "is NOT anagram of", strings.TrimSpace(t[1]))
        }
	}
}
