package main

import (
	"fmt"

	"github.com/drew138/text-recognition/equations"
)

func main() {

	Ab := [][]float64{
		{6, -2, 2, 4, 12},
		{12, -8, 6, 10, 34},
		{3, -13, 9, 3, 27},
		{-6, 4, 1, -18, -38},
	}

	sol, err := equations.SolveSystem(Ab)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(sol)
}
