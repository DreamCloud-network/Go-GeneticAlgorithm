package board

import (
	"errors"
	"log"
)

type Item int

const (
	Trash Item = iota
)

var (
	ErrorInvalidItem = errors.New("invalid item")
)

func (m Item) String() (string, error) {
	switch m {
	case Trash:
		return "Trash", nil

	default:
		log.Println("items.String - invalid item")
		err := ErrorInvalidItem
		return err.Error(), err
	}
}
