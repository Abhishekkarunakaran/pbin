package domain

import (
	"fmt"

	"github.com/gofrs/uuid"
)

type Payload struct {
	Content  string `json:"content"`
	Password string `json:"password"`
}

func (p Payload) ValidErr() error {
	switch {
	case len(p.Content) == 0 || p.Content == "":
		return fmt.Errorf("content :  required")
	case len(p.Password) == 0 || p.Password == "":
		return fmt.Errorf("password : required")
	case len(p.Password) < 8:
		return fmt.Errorf("password : should have min. 8 characters")
	default:
		return nil
	}
}

type ReplyPayload struct {
	Link string `json:"link"`
}

type Data struct {
	Password  []byte `json:"password"`
	Content   []byte `json:"content"`
	CreatedAt int64  `json:"createdat"`
}


type DataRequest struct {
	Id uuid.UUID `param:"id"`
}

type Content string

func (c *Content) String() string {
	return string(*c)
}
