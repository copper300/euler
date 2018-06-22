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

func calcRow(g [][]int) int {
	var maxPrd int
	var prd int
	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[i])-4; j++ {
			prd = g[i][j] * g[i][j+1] * g[i][j+2] * g[i][j+3]
			if prd > maxPrd {
				maxPrd = prd
			}
		}
	}
	return maxPrd
}

func calcCol(g [][]int) int {
	var maxPrd int
	var prd int
	for i := 0; i < len(g)-4; i++ {
		for j := 0; j < len(g[i]); j++ {
			prd = g[i][j] * g[i+1][j] * g[i+2][j] * g[i+3][j]
			if prd > maxPrd {
				maxPrd = prd
			}
		}
	}
	return maxPrd
}

func calcDiag(g [][]int) int {
	var maxPrd int
	var prd int
	for i := 0; i < len(g)-4; i++ {
		for j := 0; j < len(g[i])-4; j++ {
			prd = g[i][j] * g[i+1][j+1] * g[i+2][j+2] * g[i+3][j+3]
			if prd > maxPrd {
				maxPrd = prd
			}
		}
	}
	return maxPrd
}

func calcRevDiag(g [][]int) int {
	var maxPrd int
	var prd int
	for i := 0; i < len(g)-4; i++ {
		for j := 0; j < len(g[i])-4; j++ {
			prd = g[i+3][j] * g[i+2][j+1] * g[i+1][j+2] * g[i][j+3]
			if prd > maxPrd {
				maxPrd = prd
			}
		}
	}
	return maxPrd
}

func main() {
	var maxP int
	numArray := [][]int{}
	numArray = createGrid(numArray)

	rowV := calcRow(numArray)

	colV := calcCol(numArray)

	diagV := calcDiag(numArray)

	revDiagV := calcRevDiag(numArray)

	if rowV > maxP {
		maxP = rowV
	}
	if colV > maxP {
		maxP = colV
	}
	if diagV > maxP {
		maxP = diagV
	}
	if revDiagV > maxP {
		maxP = revDiagV
	}

	for i := 0; i < len(numArray); i++ {
		for j := 0; j < len(numArray[i]); j++ {
			fmt.Printf("%02d ", numArray[i][j])
		}
		fmt.Printf("\n")
	}
	fmt.Println("max row product is ", rowV)
	fmt.Println("max column product is ", colV)
	fmt.Println("max diagonal product is ", diagV)
	fmt.Println("max reverse diagonal product is ", revDiagV)
	fmt.Println("max product is ", maxP)
}
