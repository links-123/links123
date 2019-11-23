package internal

import "github.com/google/uuid"

func (rcv *linkDomainService) generateLinkID() string {
	return uuid.New().String()
}
