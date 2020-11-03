package db

import (
	"errors"

	"github.com/FBreuer2/pastebin-micro/entity"
)

type PasteDB interface {
	SavePaste(paste *entity.Paste) (*entity.Paste, *error)
	GetPaste(id uint64) (*entity.Paste, *error)
	DeletePaste(id uint64)
}

var PasteNotFound = errors.New("Paste could not be found.")
