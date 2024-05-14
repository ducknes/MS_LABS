package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Параметры автомата
const (
	n = 2 // Количество состояний
	m = 3 // Количество входных символов
	p = 5 // Количество выходных символов
)

// Матрица переходов автомата
var A = [n][m][2]float64{
	{
		{1.0 / 3, 1.0 / 3},
		{1.0 / 3, 1.0 / 3},
		{1.0 / 3, 1.0 / 3},
	},
	{
		{1.0 / 2, 1.0 / 2},
		{1.0 / 2, 1.0 / 2},
		{0, 1},
	},
}

// Функция, вычисляющая новое состояние автомата
func getNewState(zOld, x int) int {
	probabilities := A[zOld][x]
	r1 := rand.Float64()
	sum := 0.0
	for i, probability := range probabilities {
		sum += probability
		if r1 < sum {
			return i
		}
	}
	return len(probabilities) - 1
}

// Функция, вычисляющая выходной сигнал автомата
func getOutputSignal(zNew int, r2 float64) int {
	if zNew == n {
		return p - 1
	}
	if r2 < A[zNew][0][1] {
		return 0
	}
	return 1
}

// Главная функция
func main() {
	fmt.Println("x\tzOld\tr1\tzNew\tr2\ty")
	fmt.Println("---\t---\t---\t---\t---\t---")

	// Инициализация генератора случайных чисел
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// Начальное состояние автомата
	zOld := 0

	// Счетчики выходных сигналов
	outputSignals := make([]int, p)

	// Моделирование работы автомата
	for i := 0; i < 10; i++ {
		// Генерация входного символа
		x := rand.Intn(m)

		// Генерация случайных чисел для перехода и выходного сигнала
		r1 := rand.Float64()
		r2 := rand.Float64()

		// Вычисление нового состояния и выходного сигнала
		zNew := getNewState(zOld, x)
		y := getOutputSignal(zNew, r2)

		// Вывод информации о состоянии автомата
		fmt.Printf("%d\t%d\t%.2f\t%d\t%.2f\t%d\n", x, zOld, r1, zNew, r2, y)

		// Обновление счетчиков выходных сигналов
		outputSignals[y]++

		// Обновление текущего состояния
		zOld = zNew
	}

	// Вывод статистики выходных сигналов
	fmt.Println("\nOutput signals statistics:")
	for i := 0; i < p; i++ {
		fmt.Printf("y%d:\t%d\t%.2f%%\n", i, outputSignals[i], float64(outputSignals[i])/1000*100)
	}
}
