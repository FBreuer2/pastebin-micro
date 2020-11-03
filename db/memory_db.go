package db

import (
	"github.com/FBreuer2/pastebin-micro/entity"
)

type InMemoryPasteDB struct {
	pastes              map[uint64]*entity.Paste
	currentHighestIndex uint64
}

func NewInMemoryPasteDB() (*InMemoryPasteDB, error) {
	return &InMemoryPasteDB{
		pastes:              make(map[uint64]*entity.Paste, 1),
		currentHighestIndex: 1,
	}, nil
}

func (db *InMemoryPasteDB) SavePaste(paste *entity.Paste) (*entity.Paste, error) {
	paste.ID = db.currentHighestIndex

	db.pastes[paste.ID] = paste

	db.currentHighestIndex++

	return paste, nil
}

func (db *InMemoryPasteDB) GetPaste(id uint64) (*entity.Paste, error) {
	if paste := db.pastes[id]; paste != nil {
		return paste, nil
	}

	return nil, PasteNotFound
}

func (db *InMemoryPasteDB) DeletePaste(id uint64) {
	db.pastes[id] = nil
}
