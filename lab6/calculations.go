package main

import "math"

func getRho(lambda, v, k []float64) []float64 {
	rho := make([]float64, 0, len(v))
	for i := range lambda {
		if i == 0 {
			continue
		}
		rho_i := (lambda[i] * v[i-1]) / k[i-1]
		for rho_i > 1 {
			v[i-1] /= 2
			rho_i = (lambda[i] * v[i-1]) / k[i-1]
		}
		rho = append(rho, rho_i)
	}

	return rho
}

func getAverageNumberOfNonavailableChannels(lambda, v []float64) []float64 {
	result := make([]float64, 0, len(v))
	for i := range v {
		result = append(result, v[i]*lambda[i+1])
	}

	return result
}

func factorial(n float64) float64 {
	if n < 0 {
		return -1 // Факториал отрицательных чисел не определен
	}
	result := 1.0
	for i := 2.0; i <= n; i++ {
		result *= i
	}

	return result
}

func getP0(b, m, k []float64) []float64 {
	result := make([]float64, 0, len(b))

	for i := range b {
		sumBdivM := 0.0
		if k[i]-1 == 0.0 {
			sumBdivM += math.Pow(b[i], m[i]) / factorial(m[i])
		} else {
			for j := 0; j < int(k[i]); j++ {
				sumBdivM += math.Pow(b[j], m[j]) / factorial(m[j])
			}
		}

		unionSum := math.Pow(sumBdivM+(math.Pow(b[i], k[i])/(factorial(k[i])*(1-(b[i]/k[i])))), -1)
		result = append(result, unionSum)
	}

	return result
}

func getL(b, k, p0 []float64) []float64 {
	result := make([]float64, 0, len(b))

	for i := range b {
		l := ((math.Pow(b[i], 1+k[i])) / (factorial(k[i]) * k[i] * math.Pow((1-(b[i]/k[i])), 2))) * p0[i]
		result = append(result, l)
	}

	return result
}

func getM(l, b []float64) []float64 {
	result := make([]float64, 0, len(l))

	for i := range l {
		m := l[i] + b[i]
		result = append(result, m)
	}

	return result
}

func getW(l, lambda []float64) []float64 {
	result := make([]float64, 0, len(l))

	for i := range l {
		w := l[i] / lambda[i+1]
		result = append(result, w)
	}

	return result
}

func getT(w, vi []float64) []float64 {
	result := make([]float64, 0, len(w))

	for i := range w {
		t := w[i] + vi[i]
		result = append(result, t)
	}

	return result
}

func getForAll(l, m, a, w, t []float64) map[string]float64 {
	result := make(map[string]float64)

	for i := range l {
		result["L"] += l[i]
		result["N"] += m[i]
		result["W"] += a[i+1] * w[i]
		result["T"] += a[i+1] * t[i]
	}

	return result
}