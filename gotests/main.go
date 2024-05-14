package main

import (
	"fmt"
)

func determinant(m [][]float64) float64 {
	switch len(m) {
	case 2:
		return m[0][0]*m[1][1] - m[0][1]*m[1][0]
	default:
		det := 0.0
		sign := 1.0
		for i := 0; i < len(m); i++ {
			subMatrix := make([][]float64, len(m)-1)
			for j := 1; j < len(m); j++ {
				subMatrix[j-1] = make([]float64, len(m)-1)
				copy(subMatrix[j-1], m[j][:i])
				copy(subMatrix[j-1][i:], m[j][i+1:])
			}
			det += sign * m[0][i] * determinant(subMatrix)
			sign *= -1.0
		}
		return det
	}
}

func main() {
	// Задаем матрицу P
	P := [][]float64{
		{0.000000, 1.000000, 0.000000, 0.000000, 0.000000, 0.000000},
		{0.000000, 0.857143, 0.142857, 0.000000, 0.000000, 0.000000},
		{0.000000, 0.000000, 0.000000, 0.000000, 0.750000, 0.250000},
		{0.000000, 0.875000, 0.125000, 0.000000, 0.000000, 0.000000},
		{0.000000, 0.000000, 0.090909, 0.000000, 0.909091, 0.000000},
		{0.933333, 0.000000, 0.000000, 0.000000, 0.066667, 0.000000},
	}

	// Создаем вектор-столбец из единиц и вычитаем из него диагональную матрицу P
	A := make([][]float64, 6)
	for i := 0; i < 6; i++ {
		A[i] = make([]float64, 6)
		for j := 0; j < 6; j++ {
			if i == j {
				A[i][j] = 1 - P[i][j]
			} else {
				A[i][j] = -P[i][j]
			}
		}
	}

	// Создаем вектор-столбец единиц
	b := make([]float64, 6)
	for i := 0; i < 6; i++ {
		b[i] = 1.0
	}

	// Вычисляем определитель матрицы A
	detA := determinant(A)

	// Вычисляем вектор-столбец значений a_i с помощью метода Крамера
	a := make([]float64, 6)
	if detA != 0 {
		for i := 0; i < 6; i++ {
			temp := make([][]float64, 6)
			copy(temp, A)
			for j := 0; j < 6; j++ {
				temp[j][i] = b[j]
			}
			a[i] = determinant(temp) / detA
		}
	} else {
		fmt.Println("Определитель матрицы A равен нулю. Невозможно вычислить значения a_i.")
		return
	}

	// Выводим значения a_i
	fmt.Println("Значения a_i:", a)
}
