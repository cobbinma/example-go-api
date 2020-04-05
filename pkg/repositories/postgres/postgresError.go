package postgres

import "fmt"

type postgresError struct {
	err     error
	code    int
	message string
}

func (pe *postgresError) Error() string {
	return pe.err.Error()
}

func (pe *postgresError) Wrap(s string) {
	pe.err = fmt.Errorf("%s : %w", s, pe.err)
}

func (pe *postgresError) UnWrap() error {
	return pe.err
}

func (pe *postgresError) GetMessage() string {
	return pe.message
}

func (pe *postgresError) GetCode() int {
	return pe.code
}
