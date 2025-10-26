package flow

import (
	"context"
	"errors"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

// CreateKey returns key id
func (s *Service) CreateKey(ctx context.Context, key string, description string) (*Key, error) {
	return nil, errors.New("not implemented")
}

func (s *Service) ListKeys(ctx context.Context) ([]*Key, error) {
	return nil, errors.New("not implemented")
}

func (s *Service) DeleteKey(ctx context.Context, id string) error {
	return errors.New("not implemented")
}

func (s *Service) PickKey(ctx context.Context) (*Key, error) {
	return nil, errors.New("not implemented")
}
