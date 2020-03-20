package representation

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/pkg/errors"
)

type LinkEntityResponse struct {
	LinkID  string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type LinkCreateRequest struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

func (rcv *LinkCreateRequest) Validate() error {
	if rcv == nil {
		return errors.New("request is empty")
	}

	return validation.ValidateStruct(rcv,
		validation.Field(&rcv.Address, validation.Length(1, 2048)),
	)
}

type LinkDeleteRequest struct {
	LinkID string `json:"id"`
}

func (rcv *LinkDeleteRequest) Validate() error {
	if rcv == nil {
		return errors.New("request is empty")
	}

	return validation.ValidateStruct(rcv,
		validation.Field(&rcv.LinkID, validation.Required),
	)
}

type LinkListResponse struct {
	Items            []*LinkEntityResponse `json:"items"`
	TotalLinksNumber uint64                `json:"found"`
}
