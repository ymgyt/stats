package stats_test

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/ymgyt/stats"
)

func TestSlice_Mean(t *testing.T) {
	tests := []struct {
		data []float64
		want float64
	}{
		{[]float64{0}, 0},
		{[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5.5},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v", tc.data), func(t *testing.T) {
			s := stats.Slice(tc.data)
			got, want := s.Mean(), tc.want
			if got != want {
				t.Errorf("got: %f, want: %f", got, want)
			}
		})
	}
}

func TestSlice_Median(t *testing.T) {
	tests := []struct {
		data []float64
		want float64
	}{
		{[]float64{0}, 0},
		{[]float64{1, 2, 3}, 2},
		{[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5.5},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v", tc.data), func(t *testing.T) {
			s := stats.Slice(tc.data)
			got, want := s.Mean(), tc.want
			if got != want {
				t.Errorf("got: %f, want: %f", got, want)
			}
		})
	}

}
func TestSlice_Mode(t *testing.T) {
	tests := []struct {
		data []float64
		want []float64
	}{
		{[]float64{0}, []float64{0}},
		{[]float64{1, 1, 2}, []float64{1}},
		{[]float64{1, 1, 2, 2, 3, 3}, []float64{1, 2, 3}},
		{[]float64{1, 2, 3, 4, 5}, []float64{1, 2, 3, 4, 5}},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v", tc.data), func(t *testing.T) {
			s := stats.Slice(tc.data)
			got, want := s.Mode(), tc.want
			if diff := cmp.Diff(got, want); diff != "" {
				t.Errorf("(-got +want)%s", diff)
			}
		})
	}
}

func TestSlice_Variance(t *testing.T) {
	tests := []struct {
		data []float64
		want float64
	}{
		{[]float64{0}, 0},
		{[]float64{10, 10, 40}, 200},
		{[]float64{42, 69, 56, 41, 57, 48, 65, 49, 65, 58}, 86},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v", tc.data), func(t *testing.T) {
			s := stats.Slice(tc.data)
			got, want := s.Variance(), tc.want
			if diff := cmp.Diff(got, want); diff != "" {
				t.Errorf("(-got +want)%s", diff)
			}
		})
	}
}

func TestSlice_Percentile(t *testing.T) {
	tests := []struct {
		data       []float64
		percentile float64
		want       float64
	}{
		{[]float64{0}, 0, 0},
		{frange(1, 100), 0, 1},
		{frange(1, 100), 1, 100},
		{frange(1, 100), 0.5, 50},
		{frange(1, 100), 0.99, 99},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%f", tc.percentile), func(t *testing.T) {
			s := stats.Slice(tc.data)
			got, want := s.Percentile(tc.percentile), tc.want
			if diff := cmp.Diff(got, want); diff != "" {
				t.Errorf("(-got +want)%s", diff)
			}
		})
	}
}

func TestSlice_Sort(t *testing.T) {
	tests := []struct {
		data stats.Slice
		want stats.Slice
	}{
		{
			stats.Slice([]float64{5, 3, 1, 2, 7}),
			stats.Slice([]float64{1, 2, 3, 5, 7}),
		},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v", tc.data), func(t *testing.T) {
			tc.data.Sort()
			if diff := cmp.Diff(tc.data, tc.want); diff != "" {
				t.Errorf("(-got +want)%s", diff)
			}
		})
	}
}

func TestSlice_SortReverse(t *testing.T) {
	tests := []struct {
		data stats.Slice
		want stats.Slice
	}{
		{
			stats.Slice([]float64{5, 3, 1, 2, 7}),
			stats.Slice([]float64{7, 5, 3, 2, 1}),
		},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v", tc.data), func(t *testing.T) {
			tc.data.SortReverse()
			if diff := cmp.Diff(tc.data, tc.want); diff != "" {
				t.Errorf("(-got +want)%s", diff)
			}
		})
	}
}

func frange(start, end int) []float64 {
	var fs []float64
	for i := start; i <= end; i++ {
		fs = append(fs, float64(i))
	}
	return fs
}
