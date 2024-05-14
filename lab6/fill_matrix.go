package main

import "math/rand"

func fillMatrix(m [][]float64) [][]float64 {
	temp := make([][]float64, len(m))
	for i := range temp {
		temp[i] = make([]float64, len(m[i]))
		copy(temp[i], m[i]) // Копирование каждого внутреннего среза
	}
	copy(temp, m)
	for _, row := range temp {
		nonZeroSum := sum(filter(row, func(x float64) bool { return x != 0 }))
		zeroCount := len(row) - len(filter(row, func(x float64) bool { return x != 0 }))
		if zeroCount > 0 {
			remainingSum := 1 - nonZeroSum
			if remainingSum == 0 {
				continue
			}
			randomValues := generateRandomValues(zeroCount, remainingSum)
			for i, x := range row {
				if x == 0 {
					row[i] = randomValues[0]
					randomValues = randomValues[1:]
				}
			}
		}
	}

	return temp
}

func generateRandomValues(count int, remainingSum float64) []float64 {
	values := make([]float64, count)
	for i := range values {
		values[i] = rand.Float64() * remainingSum
	}
	total := sum(values)
	for i := range values {
		values[i] /= total
		values[i] *= remainingSum
	}
	return values
}

func sum(values []float64) float64 {
	total := 0.0
	for _, v := range values {
		total += v
	}
	return total
}

func filter(values []float64, predicate func(float64) bool) []float64 {
	result := make([]float64, 0)
	for _, v := range values {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}