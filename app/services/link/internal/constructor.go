package internal

import (
	"github.com/links-123/links123/app/services/link/persistence/repository"
)

type linkDomainService struct {
	linkRepository repository.LinkRepository
}

func NewLinkDomainService(
	linkRepository repository.LinkRepository,
) *linkDomainService {
	return &linkDomainService{
		linkRepository: linkRepository,
	}
}
