package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Mikita"},
			[]string{"Mikita"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Mikita", "Amsterdam"},
			[]string{"Mikita", "Amsterdam"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Mikita", 27},
			[]string{"Mikita"},
		},
		{
			"nested fields",
			Person{
				"Mikita",
				Profile{27, "Amsterdam"},
			},
			[]string{"Mikita", "Amsterdam"},
		},
		{
			"pointers to things",
			&Person{
				"Mikita",
				Profile{27, "Amsterdam"},
			},
			[]string{"Mikita", "Amsterdam"},
		},
		{
			"slices",
			[]Profile{
				{25, "Malaga"},
				{27, "Amsterdam"},
			},
			[]string{"Malaga", "Amsterdam"},
		},
		{
			"arrays",
			[2]Profile{
				{25, "Malaga"},
				{27, "Amsterdam"},
			},
			[]string{"Malaga", "Amsterdam"},
		},
		{
			"maps",
			map[string]string{
				"Cow":   "Moo",
				"Sheep": "Baa",
			},
			[]string{"Moo", "Baa"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}
