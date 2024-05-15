package main

import "math"

// gauss решает систему линейных уравнений методом Гаусса.
func gauss(matrix [][]float64) []float64 {
	//n := len(matrix)
	augmentedMatrix := createAugmentedMatrix(matrix)

	forwardElimination(augmentedMatrix)

	return backSubstitution(augmentedMatrix)
}

// createAugmentedMatrix создает расширенную матрицу, добавляя столбец свободных членов.
func createAugmentedMatrix(matrix [][]float64) [][]float64 {
	n := len(matrix)
	augmentedMatrix := make([][]float64, n)
	for i := range matrix {
		augmentedMatrix[i] = append([]float64(nil), matrix[i]...)
		augmentedMatrix[i] = append(augmentedMatrix[i], -1) // Добавление столбца свободных членов
	}
	return augmentedMatrix
}

// forwardElimination выполняет прямой ход метода Гаусса.
func forwardElimination(matrix [][]float64) {
	n := len(matrix)
	for k := 0; k < n-1; k++ {
		pivot(matrix, k)
		for i := k + 1; i < n; i++ {
			coef := matrix[i][k] / matrix[k][k]
			for j := k; j <= n; j++ {
				matrix[i][j] -= coef * matrix[k][j]
			}
		}
	}
}

// pivot находит и применяет опорный элемент для текущего шага k.
func pivot(matrix [][]float64, k int) {
	n := len(matrix)
	maxRow := k
	maxVal := math.Abs(matrix[k][k])
	for i := k + 1; i < n; i++ {
		if currentVal := math.Abs(matrix[i][k]); currentVal > maxVal {
			maxRow = i
			maxVal = currentVal
		}
	}
	if maxRow != k {
		matrix[k], matrix[maxRow] = matrix[maxRow], matrix[k]
	}
}

// backSubstitution выполняет обратный ход метода Гаусса.
func backSubstitution(matrix [][]float64) []float64 {
	n := len(matrix)
	solution := make([]float64, n)
	for i := n - 1; i >= 0; i-- {
		sum := 0.0
		for j := i + 1; j < n; j++ {
			sum += matrix[i][j] * solution[j]
		}
		solution[i] = (matrix[i][n] - sum) / matrix[i][i]
	}
	return solution
}

func matrixForGauss(m [][]float64) [][]float64 {
	result := make([][]float64, 6)
	for i := range result {
		result[i] = make([]float64, 0, 6)
	}

	for i := range m {
		for j := range m {
			result[i] = append(result[i], m[j][i])
		}
	}

	return result
}
