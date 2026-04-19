package mincostclimbingstairs

import (
	"reflect"
	"testing"
)

func TestMinCostClimbingStairs(t *testing.T) {
	tests := []struct {
		name string
		cost []int
		want int
	}{
		{"example 1",
			[]int{10, 15, 20},
			15},
		{"example 2",
			[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
			6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minCostClimbingStairs(tt.cost)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minCostClimbingStairs(%v) = %v, want %v", tt.cost, got, tt.want)
			}
		})
	}
}
