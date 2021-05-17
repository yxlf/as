package app

type ValidError struct {
	Key     string
	Message string
}

type ValidErrors []ValidError

func (v *ValidError) Error() string {
	return v.Message
}

func (v ValidErrors) Error() []string {
	var errs []string
	for _, err := range v {
		errs = append(errs, err.Error())
	}

	return errs
}
