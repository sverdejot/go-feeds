package validate

type Condition interface {
	Eval (filename string) error
}

func Validate(filename string, validations []Condition) []error {
	var errs []error
	for _, fn := range validations {
		if err := fn.Eval(filename); err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}

func ValidateFiles(filenames []string, fns []Condition) ([]string, map[string][]error) {
	valid		:= make([]string, 0, len(filenames))
	invalid := make(map[string][]error)
	

	for _, filename := range filenames {
		for _, fn := range fns {
			if err := fn.Eval(filename); err != nil {
				invalid[filename] = append(invalid[filename], err)
			}
		}

		if len(invalid[filename]) == 0 {
			valid = append(valid, filename)
		}
	}
	return valid, invalid
}
