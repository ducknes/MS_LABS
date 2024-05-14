package main

import (
	"fmt"
	"strings"
)

func main() {
	baseMartix := make([][]float64, 6)
	for i := range baseMartix {
		baseMartix[i] = make([]float64, 6)
	}
	lambda0 := 4.0
	lambda := make([]float64, 0, len(baseMartix))
	lambda = append(lambda, lambda0)
	Vi := []float64{0.5, 0.5, 0.5, 0.5, 0.5}
	K := []float64{1.0, 2.0, 2.0, 1.0, 4.0}
	m := []float64{1, 1, 1, 1, 1}
	var A []float64

	N1 := 7.0
	N2 := 4.0
	N3 := 8.0
	N4 := N1 + N2
	N5 := N1 + N3

	// Заполнение матрицы передач
	baseMartix[0][1] = 1.0
	baseMartix[1][2] = 1 / N1
	baseMartix[2][5] = 1 / N2
	baseMartix[3][2] = 1 / N3
	baseMartix[4][2] = 1 / N4
	baseMartix[5][4] = 1 / N5

	A = gauss(matrixForGauss(fillMatrix(baseMartix)))
	printMatrix(baseMartix)
	for i := range A {
		if A[i] < 0 {
			A[i] = -A[i]
		}
	}

	for i := 1; i < len(A); i++ {
		lambda = append(lambda, A[i]*lambda0)
	}

	rho := getRho(lambda, Vi, K)
	fmt.Println("\nЗагрузка каждой СМО (Pi):")
	printArray(rho, "rho", 1)

	fmt.Println("\nСреднее число занятых каналов каждой СМО (Bj):")
	b_j := getAverageNumberOfNonavailableChannels(lambda, Vi)
	printArray(b_j, "b", 1)

	fmt.Println("\nВероятности состояния сети (Poj):")
	P0 := getP0(b_j, m, K)
	printArray(P0, "P0", 1)

	fmt.Println("\nСредняя длина очереди заявок для каждой СМО:")
	l := getL(b_j, K, P0)
	printArray(l, "l", 1)

	fmt.Println("\nСреднее число заявок в каждой СМО:")
	m_i := getM(l, b_j)
	printArray(m_i, "m", 1)

	fmt.Println("\nСреднее время ожидания заявки в очереди системы Sj:")
	w := getW(l, lambda)
	printArray(w, "w", 1)

	fmt.Println("\nСреднее время пребывания заявки в системе Sj:")
	t := getT(w, Vi)
	printArray(t, "T", 1)

	fmt.Println("\nДля всей сети в целом:")
	all := getForAll(l, m_i, A, w, t)
	printMap(all)
}

func printMatrix(matrix [][]float64) {
	for _, row := range matrix {
		stringArr := make([]string, len(row))
		for i, value := range row {
			stringArr[i] = fmt.Sprintf("%f", value)
		}
		fmt.Println(strings.Join(stringArr, "  |  "))
	}
}

func printArray[T any](array []T, name string, incI int) {
	fmt.Println()
	for i, v := range array {
		fmt.Printf("%s[%d] = %v \n", name, i+incI, v)
	}
}

func printMap[T, T1 comparable](m map[T]T1) {
	for k, v := range m {
		fmt.Printf("\n %v = %v \n", k, v)
	}
}
