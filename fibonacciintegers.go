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
		for i := range rev {
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
	}
	/*
	indices := []int{}
	temp := []int{}
	rev := []int{}
    for i := 0; i < r; i++ {
		indices = append(indices, i)
		temp = append(temp, iterable[i])
	}
	for i:= range indices {
		rev = append(rev, indices[r-1-i])
	}
    l = append(l, temp)
    for{
		x := 0
		f := false
        for i := range rev {
			x = i
			if indices[i] != i + len(iterable) - r {
				f = true
                break
			}
        }
		if !f {
			return l
		}
        indices[x] += 1
        for j := x+1; j < r; j++ {
            indices[j] = indices[j-1] + 1
		}
		y := []int{}
		for i := range indices {
			y = append(y, iterable[i])
		}
        l = append(l, y)
	}
	*/
	return l
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
		fmt.Println(comb(l, 2))
	}
}
