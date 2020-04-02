package repository

import (
	"errors"
	"github.com/links-123/links123/app/services/link/persistence/model"
)

var (
	ErrLinkNotFound   = errors.New("ErrLinkNotFound")
	ErrStorageFailure = errors.New("StorageFailure")
)

type LinkRepository interface {
	Save(link *model.Link) (*model.Link, error)
	Delete(linkID, userID string) error

	FindByUserID(userID string, limit, offset uint64) (links []*model.Link, allFound uint64, err error)
}
