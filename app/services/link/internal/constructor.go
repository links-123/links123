package internal

import (
	"fmt"
	"github.com/ic2hrmk/ship_links/app/services/link/persistence/repository"
	"net"

	linkPb "github.com/ic2hrmk/ship_links/app/services/link/pb/link"

	"github.com/ic2hrmk/ship_links/app"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type linkDomainService struct {
	linkRepository repository.LinkRepository
}

func NewLinkDomainService(
	linkRepository repository.LinkRepository,
) app.MicroService {
	return &linkDomainService{
		linkRepository: linkRepository,
	}
}

func (rcv *linkDomainService) Serve(address string) error {
	return rcv.serve(address)
}

func (rcv *linkDomainService) serve(address string) error {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to acquire address [%s]: %s", address, err)
	}

	s := grpc.NewServer()
	linkPb.RegisterLinkDomainServiceServer(s, rcv)

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}

	return nil
}
