package addtwonumbers

import (
	"reflect"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		name string
		l1   *ListNode
		l2   *ListNode
		want *ListNode
	}{
		{"example 1",
			nil, /* TODO: unsupported type "ListNode", fill manually */
			nil, /* TODO: unsupported type "ListNode", fill manually */
			nil /* TODO: fill expected */},
		{"example 2",
			nil, /* TODO: unsupported type "ListNode", fill manually */
			nil, /* TODO: unsupported type "ListNode", fill manually */
			nil /* TODO: fill expected */},
		{"example 3",
			nil, /* TODO: unsupported type "ListNode", fill manually */
			nil, /* TODO: unsupported type "ListNode", fill manually */
			nil /* TODO: fill expected */},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addTwoNumbers(tt.l1, tt.l2)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers(%v, %v) = %v, want %v", tt.l1, tt.l2, got, tt.want)
			}
		})
	}
}
