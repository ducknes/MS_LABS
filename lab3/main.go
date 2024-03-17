package main

import (
	"fmt"
)

// Состояния автомата
const (
	StateA = iota
	StateB
	StateC
)

// Входы автомата
const (
	Input0 = 1 << iota
	Input1
	Input2
	Input3
	Input4
	Input5
	Input6
	Input7
	Input8
	Input9
	Input10
	Input11
)

// Таблица переходов автомата
var transitionTable = [][]int{
	{StateB, StateA, StateA, StateA, StateA, StateA, StateA, StateA, StateA, StateA, StateA, StateA},
	{StateC, StateC, StateB, StateB, StateA, StateA, StateA, StateA, StateA, StateA, StateA, StateA},
	{StateC, StateC, StateC, StateB, StateB, StateA, StateA, StateA, StateA, StateA, StateA, StateA},
}

// Таблица выходов автомата
var outputTable = []bool{
	false,
	true,
	false,
}

func main() {
	// Печатаем заголовок таблицы
	fmt.Printf("%-10s | ", "Состояние")
	for i := 0; i < 12; i++ {
		fmt.Printf("%-5d ", i)
	}
	fmt.Println()

	// Печатаем таблицу переходов
	for state, row := range transitionTable {
		switch state {
		case StateA:
			fmt.Printf("%-10s | ", "A")
		case StateB:
			fmt.Printf("%-10s | ", "B")
		case StateC:
			fmt.Printf("%-10s | ", "C")
		}
		for _, nextState := range row {
			switch nextState {
			case StateA:
				fmt.Printf("%-5s ", "A")
			case StateB:
				fmt.Printf("%-5s ", "B")
			case StateC:
				fmt.Printf("%-5s ", "C")
			}
		}
		fmt.Println()
	}

	// Печатаем таблицу выходов
	fmt.Println("Выходы:")
	for state, output := range outputTable {
		switch state {
		case StateA:
			fmt.Printf("%s: ", "A")
		case StateB:
			fmt.Printf("%s: ", "B")
		case StateC:
			fmt.Printf("%s: ", "C")
		}
		if output {
			fmt.Println("1")
		} else {
			fmt.Println("0")
		}
	}
}
