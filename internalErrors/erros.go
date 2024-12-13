package internalErrors

import (
	"errors"
	"gorm.io/gorm"
)

var ErrInternal error = errors.New("Internal Error")

// validado erros do gorm 404 - se ele for diferente do gorm ele vai retorna 500 se nao vai retornar o error que foi passado para ele
func ProcessErrorToReturn(err error) error {
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrInternal
	}
	return err
}
