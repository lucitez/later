package auth

import (
	"errors"
	"github.com/lucitez/later/api/src/pkg/model"
	"github.com/lucitez/later/api/src/pkg/repository"

	"github.com/google/uuid"

	"github.com/dgrijalva/jwt-go"
)

type Service struct {
	Repo     repository.Auth
	UserRepo repository.User
}

type Token struct {
	jwt.StandardClaims
	UserSessionID uuid.UUID `json:"user_session_id"`
}

func (t *Token) Valid() error {
	if err := t.StandardClaims.Valid(); err != nil {
		return err
	}

	return nil
}

func NewService(
	repo repository.Auth,
	userRepo repository.User,
) Service {
	return Service{
		repo,
		userRepo,
	}
}

var secret = "secret"

func (s *Service) CreateUserSession(userID uuid.UUID) (model.UserSession, error) {
	userSession := model.NewUserSession(userID)

	err := s.Repo.InsertUserSession(userSession)

	return userSession, err
}

func (s *Service) ExpireUserSession(id uuid.UUID) {
	s.Repo.ExpireUserSession(id)
}

func (s *Service) ByID(id uuid.UUID) (*model.UserSession, error) {
	return s.Repo.ByID(id)
}

func (s *Service) ActiveByID(id uuid.UUID) (*model.UserSession, error) {
	return s.Repo.ActiveByID(id)
}

func (s *Service) CheckConflicts(
	phoneNumber string,
	username string,
) error {
	existingUser := s.UserRepo.ByIdentifiers(
		phoneNumber,
		username,
	)

	if existingUser == nil {
		return nil
	}

	switch {
	case existingUser.PhoneNumber == phoneNumber:
		return errors.New("Phone Number is already in use")
	case existingUser.Username == username:
		return errors.New("Username is already in use")
	default:
		return errors.New("Conflict")
	}
}

func KeyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(secret), nil
}
