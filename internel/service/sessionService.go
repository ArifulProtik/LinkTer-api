package service

import (
	"LinkTer-api/internel/ent"
	"LinkTer-api/internel/ent/session"
	"LinkTer-api/internel/utility"
	"context"

	"github.com/google/uuid"
)

func (s *Service) SaveSession(data utility.SessionInput) error {
	_, err := s.client.Session.Create().SetUserID(data.UserID).
		SetToken(data.Token).SetIP(data.IP).Save(context.Background())
	if err != nil {
		return err
	}
	return nil

}
func (s *Service) DeleteSession(id uuid.UUID) error {
	err := s.client.Session.DeleteOneID(id).Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}
func (s *Service) GetSesssionByUser(id uuid.UUID) (*ent.Session, error) {
	existings, err := s.client.Session.Query().Where(session.UserIDEQ(id)).First(context.Background())
	if err != nil {
		return nil, err
	}
	return existings, nil
}
