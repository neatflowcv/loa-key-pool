package main

import (
	"context"

	"github.com/neatflowcv/loa-key-pool/cmd/loa-key-pool/api"
)

var _ api.StrictServerInterface = (*Service)(nil)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) CreateKey( //nolint:ireturn
	ctx context.Context,
	request api.CreateKeyRequestObject,
) (api.CreateKeyResponseObject, error) {
	panic("unimplemented")
}

func (s *Service) ListKeys( //nolint:ireturn
	ctx context.Context,
	request api.ListKeysRequestObject,
) (api.ListKeysResponseObject, error) {
	panic("unimplemented")
}

func (s *Service) DeleteKey( //nolint:ireturn
	ctx context.Context,
	request api.DeleteKeyRequestObject,
) (api.DeleteKeyResponseObject, error) {
	panic("unimplemented")
}

func (s *Service) PickKey( //nolint:ireturn
	ctx context.Context,
	request api.PickKeyRequestObject,
) (api.PickKeyResponseObject, error) {
	panic("unimplemented")
}
