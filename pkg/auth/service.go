package auth

import (
	"later/pkg/model"
	"later/pkg/repository"

	"github.com/google/uuid"

	"github.com/dgrijalva/jwt-go"
)

type Service struct {
	Repo repository.Auth
}

type Token struct {
	jwt.StandardClaims
	SessionID uuid.UUID `json:"session_id"`
}

func (t *Token) Valid() error {
	if err := t.StandardClaims.Valid(); err != nil {
		return err
	}

	return nil
}

func NewService(
	repo repository.Auth,
) Service {
	return Service{
		repo,
	}
}

var secret = "secret"

func (s *Service) CreateSession(userID uuid.UUID) (model.Session, error) {
	session := model.NewSession(userID)

	err := s.Repo.InsertSession(session)

	return session, err
}

func (s *Service) ExpireSession(id uuid.UUID) {
	s.Repo.ExpireSession(id)
}

func (s *Service) ByID(id uuid.UUID) (*model.Session, error) {
	return s.Repo.ByID(id)
}

func (s *Service) ActiveByID(id uuid.UUID) (*model.Session, error) {
	return s.Repo.ActiveByID(id)
}

func KeyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(secret), nil
}
