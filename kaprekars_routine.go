package main

import (
    "fmt"
    "strconv"
    "sort"
    "strings"
)

func largest_digit(num int) int {
    text := strconv.Itoa(num)
    splitted := strings.Split(text, "")
    sort.Sort(sort.Reverse(sort.StringSlice(splitted)))
    output, _ := strconv.Atoi(splitted[0])
    return output
}

func asc_digits(num int) int {
    text := strconv.Itoa(num)
    splitted := strings.Split(text, "")
    for i := len(splitted); i < 4; i++ {
        splitted = append(splitted, "0")
    }
    sort.Strings(splitted)
    joined := strings.Join(splitted, "")
    integer, _ := strconv.Atoi(joined)
    return integer
}

func desc_digits(num int) int {
    text := strconv.Itoa(num)
    splitted := strings.Split(text, "")
    sort.Sort(sort.Reverse(sort.StringSlice(splitted)))
    for i := len(splitted); i < 4; i++ {
        splitted = append(splitted, "0")
    }
    joined := strings.Join(splitted, "")
    integer, _ := strconv.Atoi(joined)
    return integer
}

func kaprekar(num int) int {
    i := 0
    for num != 0 && num != 6174 {
        num = desc_digits(num) - asc_digits(num)
        i++
    }
    return i
}

func secret_number() int {
    max := 0
    for i := 1000; i < 10000; i++ {
        if kaprekar(i) > max {
            max = kaprekar(i)
        }
    }
    return max
}

func main() {
    fmt.Println(secret_number())
}
