package validate

import (
	"testing"
	"fmt"
	"slices"
	"reflect"
)

func TestValidate(t *testing.T) {
	t.Run("validate name starts with prefix", func(t *testing.T) {
		filename := "test_file"
		condition := StartsWith{ prefix: "test" }

		err := Validate(filename, []Condition{ condition })

		if err != nil {
			t.Errorf("validate %s should throw nothing and threw %v", filename, err)
		}
	})

	error_cases := []struct{
		Filename		string
		Conditions	[]Condition
		Errors			[]error
	}{
		{
			Filename:		"filename",
			Conditions:	[]Condition{ StartsWith{"test"} },
			Errors:			[]error{ ErrInvalidPrefix("test") },
		},
		{
			Filename:		"filename",
			Conditions:	[]Condition{ StartsWith{"test"}, StartsWith{"another-test"} },
			Errors:			[]error{ ErrInvalidPrefix("test"), ErrInvalidPrefix("another-test") },
		},
		{
			Filename:		"filename",
			Conditions: []Condition{ EndsWith{"test"}},
			Errors:			[]error{ ErrInvalidSuffix("test") },
		},
	}

	for _, test := range error_cases {
		t.Run(fmt.Sprintf("validate %s should throw %+v", test.Filename, test.Errors), func(t *testing.T) {
			err := Validate(test.Filename, test.Conditions)
			compareErrors(t, err, test.Errors)
		})
	}
}


func compareErrors(t testing.TB, got, want []error) {
	t.Helper()

	if len(got) != len(want) {
		t.Errorf("[%q] and [%q] does not contain the same count of elements", got, want)
	}
	var gotTypes, wantTypes []reflect.Type

	for i, gotErr := range got {
		gotTypes = append(gotTypes, reflect.TypeOf(gotErr))
		wantTypes = append(wantTypes, reflect.TypeOf(want[i]))
	}

	if !slices.Equal(gotTypes, wantTypes) {
		t.Errorf("[%q] and [%q] are not the same", got, want)
	}
}
