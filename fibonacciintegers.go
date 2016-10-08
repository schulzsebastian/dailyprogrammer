package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func fib(i int) []int {
    l := []int{}
	a, b := 1, 1
	for a <= i {
		a, b = b, a + b
		if a <= i {
			l = append(l, a)
		}
	}
    return l
}

func comb(iterable []int, r int) [][]int {
	l := [][]int{}
	if r > len(iterable) {
        return l
	}
	indices := []int{}
	for i := 0; i < r; i++ {
		indices = append(indices, i)
	}
	temp := []int{}
	rev := []int{}
	for i := 0; i < len(indices); i++ {
		temp = append(temp, iterable[i])
		rev = append(rev, indices[len(indices)-1-i])
	}
	l = append(l, temp)
	for {
		f := false
		x := 0
		for i := 0; i < len(rev); i++ {
			if indices[i] != i + len(iterable) - r {
                x = i
                f = true
                break
			}
		}
		if !f {
            return l
		}
		indices[x] += 1
		for j := x + 1; j < r; j++ {
			indices[j] = indices[j-1] + 1
		}
		temp := []int{}
		for i := 0; i < len(indices); i++ {
			temp = append(temp, indices[i])
		}
		l = append(l, temp)
	}
}

func main() {
	fileHandle, _ := os.Open("fibonacciintegers_input.txt")
    defer fileHandle.Close()
    fileScanner := bufio.NewScanner(fileHandle)
	stat := false
    for fileScanner.Scan() {
		if !stat {
			stat = true
			continue
		}
		n, _ := strconv.Atoi(fileScanner.Text())
        l := fib(n)
		fmt.Println(comb(l, 3))
	}
}
