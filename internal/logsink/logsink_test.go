package logsink

import (
	"fmt"
	"testing"
)

func TestNumDigits(t *testing.T) {
	cases := []struct {
		in  uint64
		num int
	}{
		{0, 1},
		{1, 1},
		{2, 1},
		{3, 1},
		{8, 1},
		{9, 1},
		{10, 2},
		{11, 2},
		{99, 2},
		{100, 3},
		{101, 3},
		{234567, 6},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("%d-%d", tc.in, tc.num), func(t *testing.T) {
			if got := numDigits(tc.in); got != tc.num {
				t.Errorf("numDigits(%d) = %d; want %d", tc.in, got, tc.num)
			}
		})
	}
}
