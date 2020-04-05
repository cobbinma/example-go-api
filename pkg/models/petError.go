package models

type PetError interface {
	Error() string
	GetMessage() string
	Wrap(s string)
	UnWrap() error
}
