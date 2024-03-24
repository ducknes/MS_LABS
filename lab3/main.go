package main

import "fmt"

// Статусы автомата
const (
	StatusA = iota
	StatusB
	StatusC
	StatusD
	StatusE
	StatusF
	StatusG
	StatusH
	StatusI
)

// Входы автомата
const (
	Input1 = 1 + iota
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
	Input12
)

// Таблица переходов автомата
var table = [9][12]func(int) int{
	{func(int) int { return StatusB }, func(int) int { return StatusC }, func(int) int { return StatusA }, func(int) int { return StatusD }, func(int) int { return StatusE }, func(int) int { return StatusA }, func(int) int { return StatusF }, func(int) int { return StatusA }, func(int) int { return StatusG }, func(int) int { return StatusA }, func(int) int { return StatusH }, func(int) int { return StatusA }}, // StatusA
	{func(int) int { return StatusA }, func(int) int { return StatusB }, func(int) int { return StatusA }, func(int) int { return StatusA }, func(int) int { return StatusA }, func(int) int { return StatusC }, func(int) int { return StatusA }, func(int) int { return StatusD }, func(int) int { return StatusA }, func(int) int { return StatusE }, func(int) int { return StatusA }, func(int) int { return StatusF }}, // StatusB
	{func(int) int { return StatusA }, func(int) int { return StatusA }, func(int) int { return StatusB }, func(int) int { return StatusC }, func(int) int { return StatusD }, func(int) int { return StatusA }, func(int) int { return StatusE }, func(int) int { return StatusA }, func(int) int { return StatusF }, func(int) int { return StatusA }, func(int) int { return StatusG }, func(int) int { return StatusA }}, // StatusC
	{func(int) int { return StatusA }, func(int) int { return StatusB }, func(int) int { return StatusC }, func(int) int { return StatusD }, func(int) int { return StatusA }, func(int) int { return StatusE }, func(int) int { return StatusA }, func(int) int { return StatusF }, func(int) int { return StatusA }, func(int) int { return StatusG }, func(int) int { return StatusA }, func(int) int { return StatusH }}, // StatusD
	{func(int) int { return StatusA }, func(int) int { return StatusB }, func(int) int { return StatusC }, func(int) int { return StatusD }, func(int) int { return StatusE }, func(int) int { return StatusA }, func(int) int { return StatusF }, func(int) int { return StatusA }, func(int) int { return StatusG }, func(int) int { return StatusA }, func(int) int { return StatusH }, func(int) int { return StatusA }}, // StatusE
	{func(int) int { return StatusA }, func(int) int { return StatusB }, func(int) int { return StatusC }, func(int) int { return StatusD }, func(int) int { return StatusE }, func(int) int { return StatusF }, func(int) int { return StatusA }, func(int) int { return StatusG }, func(int) int { return StatusA }, func(int) int { return StatusH }, func(int) int { return StatusA }, func(int) int { return StatusI }}, // StatusF
	{func(int) int { return StatusA }, func(int) int { return StatusB }, func(int) int { return StatusC }, func(int) int { return StatusD }, func(int) int { return StatusE }, func(int) int { return StatusF }, func(int) int { return StatusG }, func(int) int { return StatusA }, func(int) int { return StatusH }, func(int) int { return StatusA }, func(int) int { return StatusI }, func(int) int { return StatusA }}, // StatusG
	{func(int) int { return StatusA }, func(int) int { return StatusB }, func(int) int { return StatusC }, func(int) int { return StatusD }, func(int) int { return StatusE }, func(int) int { return StatusF }, func(int) int { return StatusG }, func(int) int { return StatusH }, func(int) int { return StatusA }, func(int) int { return StatusI }, func(int) int { return StatusA }, func(int) int { return StatusA }}, // StatusH
	{func(int) int { return StatusA }, func(int) int { return StatusB }, func(int) int { return StatusC }, func(int) int { return StatusD }, func(int) int { return StatusE }, func(int) int { return StatusF }, func(int) int { return StatusG }, func(int) int { return StatusH }, func(int) int { return StatusI }, func(int) int { return StatusA }, func(int) int { return StatusA }, func(int) int { return StatusA }}, // StatusI
}

func printTable() {
	fmt.Println("   | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 10 | 11 | 12 |")
	fmt.Println("---+---+---+---+---+---+---+---+---+---+----+----+----+")
	for i := 0; i < 9; i++ {
		fmt.Printf("%c  | ", 'A'+i)
		for j := 0; j < 12; j++ {
			status := table[i][j](i)
			if j < 9 {
				fmt.Printf("%c | ", 'A'+status)
				continue
			}
			fmt.Printf("%c  | ", 'A'+status)
		}
		fmt.Println()
		fmt.Println("---+---+---+---+---+---+---+---+---+---+----+----+----+")
	}
}

func main() {
	printTable()
}
