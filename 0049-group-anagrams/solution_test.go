package groupanagrams

import (
	"reflect"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want interface{}
	}{
		{
			"example 1",
			[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
			nil, /* TODO: fill expected */
		},
		{
			"example 2",
			[]string{""},
			nil, /* TODO: fill expected */
		},
		{
			"example 3",
			[]string{"a"},
			nil, /* TODO: fill expected */
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := groupAnagrams(tt.strs)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupAnagrams(%v) = %v, want %v", tt.strs, got, tt.want)
			}
		})
	}
}
