package entity

import (
	"errors"
	"time"

	"golang.org/x/crypto/blake2b"
)

type Paste struct {
	ID                uint64
	Author            string
	CreationDate      time.Time
	EditDate          time.Time
	Content           string
	Language          string
	PasswordProtected bool
	Password          string
}

func CreatePaste(author string, content string, language string, password string) (*Paste, error) {

	if content == "" || author == "" {
		return nil, errors.New("Malformed author or content.")
	}

	if password == "" {
		return &Paste{
			Author:            author,
			Content:           content,
			CreationDate:      time.Now(),
			EditDate:          time.Now(),
			Language:          language,
			PasswordProtected: false,
		}, nil
	}

	hasher, err := blake2b.New256(nil)

	if err != nil {
		return nil, err
	}

	hash, err := hasher.Write([]byte(password))

	if err != nil {
		return nil, err
	}

	return &Paste{
		Author:            author,
		Content:           content,
		CreationDate:      time.Now(),
		EditDate:          time.Now(),
		Language:          language,
		PasswordProtected: true,
		Password:          string(hash),
	}, nil

}

func (p *Paste) GetTimeString() string {
	return p.CreationDate.Format("Wed Feb 25 11:06:39 PST 2015")
}
