package errors

type errorString struct {
	s string
}

func New(s string) error {
	return &errorString{s}
}

func (e *errorString) Error() string {
	return e.s
}
