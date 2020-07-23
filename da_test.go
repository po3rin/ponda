package ponda_test

import (
	"testing"

	"github.com/po3rin/ponda"
)

func TestLookup(t *testing.T) {

	// da is Dobule Array for test
	// []string{ "abc", "abcd", "文字", "全角", "全角文字",}
	da := ponda.DoubleArray{
		Nodes: []ponda.Node{
			{
				Base:  1,
				Check: -1,
			},
			{
				Base:  0,
				Check: -2,
			},
			{
				Base:  -1,
				Check: -4,
			},
			{
				Base:  6,
				Check: 0,
			},
			{
				Base:  -2,
				Check: -5,
			},
			{
				Base:  -4,
				Check: -18,
			},
			{
				Base:  6,
				Check: 0,
			},
			{
				Base:  -1,
				Check: 3,
			},
			{
				Base:  7,
				Check: 0,
			},
			{
				Base:  10,
				Check: 0,
			},
			{
				Base:  8,
				Check: 6,
			},
			{
				Base:  -1,
				Check: 10,
			},
			{
				Base:  13,
				Check: 13,
			},
			{
				Base:  -10,
				Check: 8,
			},
			{
				Base:  -1,
				Check: 12,
			},
			{
				Base:  12,
				Check: 9,
			},
			{
				Base:  14,
				Check: 15,
			},
			{
				Base:  -1,
				Check: 16,
			},
		},
		Dict: ponda.RuneDic{
			23383: 7,
			100:   8,
			20840: 1,
			35282: 2,
			97:    3,
			98:    4,
			99:    5,
			25991: 6,
		},
	}

	tests := []struct {
		name  string
		da    ponda.DoubleArray
		key   []rune
		found bool
	}{
		{
			name:  "found",
			da:    da,
			key:   []rune("abc"),
			found: true,
		},
		{
			name:  "not found",
			da:    da,
			key:   []rune("aaa"),
			found: false,
		},
		{
			name:  "found ja",
			da:    da,
			key:   []rune("文字"),
			found: true,
		},
		{
			name:  "not found ja",
			da:    da,
			key:   []rune("漢字"),
			found: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got := tt.da.Lookup(tt.key)
			if got != tt.found {
				t.Errorf("key: %+v, got: %+v, want: %+v\n", string(tt.key), got, tt.found)
			}
		})
	}
}
