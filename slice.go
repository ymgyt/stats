package stats

import (
	"math"
	"sort"
)

type Slice []float64

func (s Slice) Mean() float64 {
	return s.Sum() / float64(s.Len())
}

func (s Slice) Sum() float64 {
	var sum float64
	for _, v := range s {
		sum += v
	}
	return sum
}

func (s Slice) Median() float64 {
	n := s.Len()
	if n == 0 {
		return 0
	}
	s.Sort()
	var median float64
	if n%2 == 0 {
		m0 := s[n/2-1]
		m1 := s[n/2]
		median = m0 + m1/2.0
	} else {
		median = s[n/2]
	}
	return median
}

func (s Slice) Variance() float64 {
	mean := s.Mean()
	deviation := s.Map(func(x float64) float64 { return math.Pow((x - mean), 2) })
	return deviation.Mean()
}

func (s Slice) StandardDeviation() float64 {
	return math.Sqrt(s.Variance())
}

func (s Slice) Percentile(p float64) float64 {
	n := float64(len(s))
	if n == 0 {
		return 0
	}
	if p > 1.0 {
		return 0
	}
	for i, v := range s {
		if float64(i+1)/n >= p {
			return v
		}
	}
	return 0
}

func (s Slice) Mode() []float64 {
	m := make(map[float64]int)
	for _, v := range s {
		m[v]++
	}
	var max int
	for _, c := range m {
		if c > max {
			max = c
		}
	}
	var modes []float64
	for v, c := range m {
		if c == max {
			modes = append(modes, v)
		}
	}
	sort.Float64s(modes) // for testability
	return modes
}

func (s Slice) Map(f func(float64) float64) Slice {
	ss := make(Slice, len(s))
	for i := range s {
		ss[i] = f(s[i])
	}
	return ss
}

func (s Slice) Len() int           { return len(s) }
func (s Slice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s Slice) Less(i, j int) bool { return s[i] < s[j] }
func (s Slice) Sort()              { sort.Sort(s) }
func (s Slice) SortReverse()       { sort.Sort(sort.Reverse(s)) }
