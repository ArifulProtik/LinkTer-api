package service

import (
	"LinkTer-api/internel/ent"
	"LinkTer-api/internel/utility"
	"context"
)

func (s *Service) SaveUser(usr *utility.UserInput) (*ent.User, error) {
	newUser, err := s.client.User.Create().SetName(usr.Name).
		SetUsername(usr.Username).SetEmail(usr.Email).SetPassword(usr.Password).Save(context.Background())
	if err != nil {
		s.log.Debug(err)
		return nil, err
	}
	return newUser, nil
}
