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
		condition := StartsWith{ Prefix: "test" }

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

func TestValidateFiles(t *testing.T) {
	t.Run("all valid files", func(t *testing.T) {
		filenames := []string{
			"valid-file-1",
			"valid-file-2",
		}

		valid, rejected := ValidateFiles(filenames, []Condition{ StartsWith{"valid-"} })

		allFilesShouldBeValid(t, filenames, valid, rejected)
	})

	t.Run("one invalid file", func(t *testing.T) {
		filenames := []string{
			"valid-file-1",
			"invalid-file-1",
		}

		_, rejected := ValidateFiles(filenames, []Condition{ StartsWith{"valid-"} })

		if rejectErrors := rejected["invalid-file-1"]; rejectErrors == nil {
			t.Errorf("validate should have thrown an error but got nothing")
		}
	})


	t.Run("all valid files with multiple conditions", func(t *testing.T) {
		filenames := []string{
			"valid-file-1-suffix",
			"valid-file-2-suffix",
		}
		conditions := []Condition{
			StartsWith{"valid-"},
			EndsWith{"-suffix"},
		}

		valid, rejected := ValidateFiles(filenames, conditions)
		
		allFilesShouldBeValid(t, filenames, valid, rejected)
	})
}

func allFilesShouldBeValid(t testing.TB, filenames, valid []string, rejected map[string][]error) {
	t.Helper()
	
	if len(rejected) != 0 {
		t.Errorf("no files should be rejected but got %d %v", len(rejected), rejected)
	}

	if !slices.Equal(valid, filenames) {
		t.Errorf("all %d files %v should be valid but only %d files %v are", len(filenames), filenames, len(valid), valid)
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
