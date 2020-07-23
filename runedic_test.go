package ponda_test

import (
	"reflect"
	"testing"

	"github.com/po3rin/ponda"
)

func TestBuildRuneDict(t *testing.T) {
	tests := []struct {
		name  string
		input [][]rune
		want  ponda.RuneDic
	}{
		{
			name: "simple",
			input: [][]rune{
				[]rune("a"),
				[]rune("b"),
				[]rune("c"),
				[]rune("d"),
				[]rune("文"),
				[]rune("字"),
				[]rune("全"),
				[]rune("角"),
			},
			want: ponda.RuneDic{
				23383: 6,
				100:   4,
				20840: 7,
				35282: 8,
				97:    1,
				98:    2,
				99:    3,
				25991: 5,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ponda.BuildRuneDict(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("want: %+v, got: %+v", tt.want, got)
			}
		})
	}
}
