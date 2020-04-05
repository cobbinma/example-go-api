package models

import "fmt"

type PetError interface {
	Error() string
	GetMessage() string
	GetCode() int
	Wrap(s string)
	UnWrap() error
}

type petError struct {
	err     error
	code    int
	message string
}

func (pe *petError) Error() string {
	return pe.err.Error()
}

func (pe *petError) Wrap(s string) {
	pe.err = fmt.Errorf("%s : %w", s, pe.err)
}

func (pe *petError) UnWrap() error {
	return pe.err
}

func (pe *petError) GetMessage() string {
	return pe.message
}

func (pe *petError) GetCode() int {
	return pe.code
}
