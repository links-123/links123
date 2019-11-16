package repository

import (
	"errors"
	"github.com/ic2hrmk/ship_links/app/services/link/persistence/model"
)

var (
	ErrLinkNotFound = errors.New("ErrLinkNotFound")
)

type LinkRepository interface {
	Save(*model.Link) (*model.Link, error)
	SaveBulk([]*model.Link) error

	FindAll(limit, offset uint64) ([]*model.Link, error)
}
