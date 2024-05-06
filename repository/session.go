package repository

import (
	"goreact/model"
	"time"
	"gorm.io/gorm"
)

type SessionRepo interface {
	AddSessions(session model.Session) error
	DeleteSession(token string) error
	UpdateSessions(session model.Session) error
	SessionAvailEmail(email string) (model.Session, error)
	SessionAvailToken(token string) (model.Session, error)
	TokenExpired(session model.Session) bool
}

type sessionRepo struct {
	db *gorm.DB
}

func NewSessionRepo(db *gorm.DB) *sessionRepo {
	return &sessionRepo{db}
}

func (s *sessionRepo) AddSessions(session model.Session) error {	
	err := s.db.Create(&session)
	return err.Error
}

func (s *sessionRepo) DeleteSession(token string) error {			
	err := s.db.Where("token = ?", token).Delete(&model.Session{})
	return err.Error
}

func (s *sessionRepo) UpdateSessions(session model.Session) error {
	err := s.db.Where("email = ?", session.Email).Updates(session)
	return err.Error
}

func (s *sessionRepo) SessionAvailEmail(email string) (model.Session, error) {
	var result model.Session
	err := s.db.Where("email = ?", email).First(&result).Error
	if err != nil{
		return model.Session{}, err
	}
	return result, nil
}

func (s *sessionRepo) SessionAvailToken(token string) (model.Session, error) {
	var result model.Session
	err := s.db.Where("token = ?", token).First(&result).Error
	if err != nil{
		return model.Session{}, err 
	}
	return result, nil
}

func (s *sessionRepo) TokenValidity(token string) (model.Session, error) {
	session, err := s.SessionAvailToken(token)
	if err != nil {
		return model.Session{}, err
	}

	if s.TokenExpired(session) {
		err := s.DeleteSession(token)
		if err != nil {
			return model.Session{}, err
		}
		return model.Session{}, err
	}

	return session, nil
}

func (s *sessionRepo) TokenExpired(session model.Session) bool {
	return session.Expiry.Before(time.Now())
}