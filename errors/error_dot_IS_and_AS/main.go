package main

import (
	"errors"
	"log"
)

type NotFoundError struct {
	msg string
}

func (e *NotFoundError) Error() string {
	return e.msg
}

type PermissionDeniedError struct {
	msg string
}

func (e *PermissionDeniedError) Error() string {
	return e.msg
}

func accessResource() error {
	e := 1
	if e == 1 {
		return &PermissionDeniedError{"Erro de permissao"}
	}
	return &NotFoundError{"Erro nao encontrado"}
}

func main() {
	if err := accessResource(); err != nil {
		var notFoundError *NotFoundError
		if errors.Is(err, &PermissionDeniedError{}) {
			log.Fatalln("Acess resource nao tem permissao")
		} else if errors.As(err, &notFoundError) {
			log.Fatalln("Acess resource nao encontrou")
		}
	}
}
