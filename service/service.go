package service

import "begingo/dao"

type Service interface {
	Users() UserSrv
}

type service struct {
	store dao.Factory
}

// NewService returns Service interface.
func NewService(store dao.Factory) Service {
	return &service{
		store: store,
	}
}

func (s *service) Users() UserSrv {
	return newUsers(s)
}
