package solver

import (
	"errors"
)

type Matrix [][]float64

func gaussianElimination(Ab Matrix) error {
	n := len(Ab)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if Ab[i][i] == 0 {
				return errors.New("zero division encountered when processing input matrix")
			}
			mult := Ab[j][i] / Ab[i][i]
			for k := i; k < n+1; k++ {
				Ab[j][k] -= mult * Ab[i][k]
			}
		}
	}
	return nil
}

func regressiveSubstitution(Ab Matrix) []float64 {

	n := len(Ab)
	x := make([]float64, n)
	x[n-1] = Ab[n-1][n] / Ab[n-1][n-1]
	for i := n - 2; i >= 0; i-- {
		sum := 0.0
		for j := i + 1; j < n; j++ {
			sum += Ab[i][j] * x[j]
		}
		x[i] = (Ab[i][n] - sum) / Ab[i][i]
	}
	return x
}

func SolveSystem(Ab Matrix) ([]float64, error) {
	err := gaussianElimination(Ab)
	if err != nil {
		return nil, err
	}
	sol := regressiveSubstitution(Ab)
	return sol, nil
}
