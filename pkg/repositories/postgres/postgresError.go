package postgres

import "fmt"

type postgresError struct {
	err     error
	code    int
	message string
}

func newPetError(err error, message string, code int) *postgresError {
	return &postgresError{
		err:     err,
		code:    code,
		message: message,
	}
}

func (pe *postgresError) Error() string {
	return pe.err.Error()
}

func (pe *postgresError) Wrap(s string) {
	pe.err = fmt.Errorf("%s : %w", s, pe.err)
}

func (pe *postgresError) Unwrap() error {
	return pe.err
}

func (pe *postgresError) GetMessage() string {
	return pe.message
}

func (pe *postgresError) GetCode() int {
	return pe.code
}
