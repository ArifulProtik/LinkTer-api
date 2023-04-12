package service

import (
	"LinkTer-api/internel/ent"
	"LinkTer-api/internel/ent/user"
	"LinkTer-api/internel/utility"
	"context"

	"github.com/google/uuid"
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

func (s *Service) GetUserByEmail(email string) (*ent.User, error) {
	usr, err := s.client.User.Query().Where(user.EmailEQ(email)).First(context.Background())
	if err != nil {
		return nil, err
	}
	return usr, nil
}

func (s *Service) GetUserByID(id uuid.UUID) (*ent.User, error) {
	usr, err := s.client.User.Get(context.Background(), id)
	if err != nil {
		return nil, err
	}
	return usr, nil
}
