package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/gonum/stat/distuv"
)

func inverseTransformSampling(n int) []float64 {
	samples := make([]float64, n)
	for i := 0; i < n; i++ {
		samples[i] = -math.Log(1-rand.Float64()) - 1
	}
	return samples
}

func calculateMeanVariance(samples []float64) (float64, float64) {
	mean := stat.Mean(samples, nil)
	variance := stat.Variance(samples, nil)
	return mean, variance
}

func calculateKsStatistic(samples []float64) float64 {
	cdf := distuv.Exponential{Lambda: 1}.CDF
	sort.Float64s(samples)
	n := len(samples)
	dMax := math.Inf(-1)
	for i := 0; i < n; i++ {
		d := math.Max(math.Abs(cdf(samples[i])-float64(i+1)/float64(n)), math.Abs(cdf(samples[i])-float64(i)/float64(n)))
		if d > dMax {
			dMax = d
		}
	}
	return math.Sqrt(n) * dMax
}

func plotHistogram(samples []float64, bins int) {
	hist, _ := charts.NewHistogram()
	hist.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "Гистограмма",
			Subtitle: "Эмпирическое и теоретическое распределения",
		}),
		charts.WithXAxisOpts(opts.XAxis{
			Name: "Значения",
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name: "Плотность вероятности",
		}),
	)

	hist.AddSeries("Эмпирическое", samples)
	hist.SetSeriesOptions(
		charts.WithHistogramSeriesOptions(
			opts.HistogramSeries{
				BinCount: bins,
			},
		),
	)

	x := make([]float64, 100)
	y := make([]float64, 100)
	for i := 0; i < 100; i++ {
		x[i] = math.E*float64(i)/100 - 1
		y[i] = 1 / (math.E - x[i])
	}
	hist.AddSeries("Теоретическое", y).SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{
		XAxis: x,
	}))

	f, _ := hist.Render()
	fmt.Println(string(f))
}

func main() {
	sampleSizes := []int{50, 100, 1000, 10000, 100000}

	for _, sampleSize := range sampleSizes {
		samples := inverseTransformSampling(sampleSize)

		mean, variance := calculateMeanVariance(samples)
		fmt.Printf("Размер выборки: %d\n", sampleSize)
		fmt.Printf("Математическое ожидание: %f, Дисперсия: %f\n", mean, variance)
		fmt.Printf("Истинное математическое ожидание: %f, Истинная дисперсия: %f\n", 1/math.E, (math.E-2)/(math.E*math.E))

		ksStatistic := calculateKsStatistic(samples)
		fmt.Printf("Статистика критерия Колмогорова: %f\n", ksStatistic)

		plotHistogram(samples, 50)
	}
}
