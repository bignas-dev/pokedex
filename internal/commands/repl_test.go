package commands

import (
	"slices"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct{
		input string
		expected []string
	}{
		{
			" hello world  ",
			[]string{"hello", "world"},
		},
		{
			"  ",
			[]string{},
		},
		{
			"  IM THE BIGGEST BIRD",
			[]string{"im", "the", "biggest", "bird"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if !slices.Equal(actual, c.expected) {
			t.Errorf("actual: %v,  expected: %v", actual, c.expected)
		}
    }


}
