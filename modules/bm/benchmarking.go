package bm

import (
	"errors"
)

type Benchmarking struct {
	Standards map[string]PerformanceStandard
}

type PerformanceStandard struct {
	MinimumValue int
	MaximumValue int
	Measurement  string
}

func NewPerformanceStandard(minimumValue int, maximumValue int, measurement string) *PerformanceStandard {
	return &PerformanceStandard{MinimumValue: minimumValue, MaximumValue: maximumValue, Measurement: measurement}
}

func NewBenchmarking(standards map[string]PerformanceStandard) *Benchmarking {
	return &Benchmarking{Standards: standards}
}

func (b *Benchmarking) Benchmark(performer Performer) (int, error) {
	if len(b.Standards) == 0 {
		return 0, errors.New("no performance standards defined")
	}

	performanceData := performer.GetPerformance()
	if len(performanceData) == 0 {
		return 0, errors.New("no performance data available")
	}

	totalScore := 0
	for _, data := range performanceData {
		standard, ok := b.Standards[data.KPI]
		if !ok {
			return 0, errors.New("no performance standard defined for KPI: " + data.KPI)
		}

		score := calculateScore(data.Value, standard.MinimumValue, standard.MaximumValue)
		totalScore += score
	}

	averageScore := totalScore / len(performanceData)
	return averageScore, nil
}

func calculateScore(value, minimum, maximum int) int {
	if value < minimum {
		return 0
	}
	if value > maximum {
		return 100
	}
	rangeValue := maximum - minimum
	rangeScore := 100
	if rangeValue > 0 {
		rangeScore = 100 / rangeValue
	}
	return ((value - minimum) * rangeScore)
}
