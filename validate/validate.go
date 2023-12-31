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
