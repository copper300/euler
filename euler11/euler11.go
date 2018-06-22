package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func createGrid(g [][]int) [][]int {
	var l int
	fi, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()
	r := bufio.NewReader(fi)
	for {
		gridL := []int{}
		line, _, err := r.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		arr := strings.Split(string(line), " ")
		for i := 0; i < len(arr); i++ {
			num, err := strconv.Atoi(arr[i])
			if err != nil {
				log.Println(err)
				os.Exit(2)
			}
			gridL = append(gridL, num)
		}
		g = append(g, gridL)
	}
	return g
}

func main() {
	numArray := [][]int{}
	numArray = createGrid(numArray)
	for i := 0; i < len(numArray); i++ {
		for j := 0; j < len(numArray[i]); j++ {
			fmt.Printf("%d ", numArray[i][j])
			fmt.Printf("\n")
		}
	}
}
