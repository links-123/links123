package internal

import (
	"github.com/sirupsen/logrus"

	"github.com/links-123/links123/app/services/link/persistence/repository"
)

type linkDomainService struct {
	log            *logrus.Logger
	linkRepository repository.LinkRepository
}

func NewLinkDomainService(
	logger *logrus.Logger,
	linkRepository repository.LinkRepository,
) *linkDomainService {
	return &linkDomainService{
		log:            logger,
		linkRepository: linkRepository,
	}
}
