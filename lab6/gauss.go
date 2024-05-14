package main

import "math"

func gauss(matrix [][]float64) []float64 {
	// Размерность матрицы
	n := len(matrix)

	// Добавление столбца свободных членов
	for i := range matrix {
		matrix[i] = append(matrix[i], -1)
	}

	// Метод Гаусса с выбором главного элемента
	for k := 0; k < n-1; k++ {
		maxRow := k
		maxVal := math.Abs(matrix[k][k])

		for i := k + 1; i < n; i++ {
			if math.Abs(matrix[i][k]) > maxVal {
				maxRow = i
				maxVal = math.Abs(matrix[i][k])
			}
		}

		if maxRow != k {
			matrix[k], matrix[maxRow] = matrix[maxRow], matrix[k]
		}

		for i := k + 1; i < n; i++ {
			coef := matrix[i][k] / matrix[k][k]

			for j := k; j < n+1; j++ {
				matrix[i][j] = matrix[i][j] - coef*matrix[k][j]
			}
		}
	}

	// Вычисление значений a_i
	a := make([]float64, n)
	for i := n - 1; i >= 0; i-- {
		sum := 0.0
		for j := i + 1; j < n; j++ {
			sum += matrix[i][j] * a[j]
		}
		a[i] = (matrix[i][n] - sum) / matrix[i][i]
	}

	return a
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