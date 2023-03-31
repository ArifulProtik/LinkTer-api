package service

import (
	"LinkTer-api/internel/ent"
	"LinkTer-api/internel/utility"
	"context"
)

func (s *Service) SaveUser(usr *utility.UserInput) (*ent.User, error) {
	newPass, err := utility.HashBeforeSave(usr.Password)
	if err != nil {
		return nil, err
	}
	newUser, err := s.client.User.Create().SetName(usr.Name).
		SetUsername(usr.Username).SetEmail(usr.Email).SetPassword(string(newPass)).Save(context.Background())
	if err != nil {
		s.log.Debug(err)
		return nil, err
	}
	return newUser, nil
}
