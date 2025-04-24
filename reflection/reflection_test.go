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
	t.Run("without maps", func(t *testing.T) {
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
	})

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Cow":   "Moo",
			"Sheep": "Baa",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Moo")
		assertContains(t, got, "Baa")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{27, "Amsterdam"}
			aChannel <- Profile{25, "Malaga"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Amsterdam", "Malaga"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{27, "Amsterdam"}, Profile{25, "Malaga"}
		}

		var got []string
		want := []string{"Amsterdam", "Malaga"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %v to contain %q but it didn't", haystack, needle)
	}
}
