package service

import (
	"TaxiAppDriver/internal/broker"
	"TaxiAppDriver/internal/model"
	"context"

	"golang.org/x/crypto/bcrypt"
)

type Repository interface {
	AddNewDriver(ctx context.Context, driver model.Driver) error
	GetDriverPhoneAndPasswordByPhone(ctx context.Context, phone string) (model.Driver, error)
	Exists(ctx context.Context, user model.Login) (bool, error)
	GetAllDrivers(ctx context.Context) ([]model.Driver, error)
	GetAllFreeDrivers(ctx context.Context) ([]model.Driver, error)
	UpdateDriverStatus(ctx context.Context, user model.Driver) error
}

type Broker interface {
	SendMessage(value string) error
}
type Service struct {
	Repos  Repository
	Broker Broker
}

func NewService(repos Repository, broker broker.KafkaClient) *Service {
	return &Service{Repos: repos, Broker: &broker}
}

func (s *Service) Save(ctx context.Context, driver model.Driver) error {
	exists, err := s.Repos.Exists(ctx, model.Login{Phone: driver.Phone, Password: driver.Password})
	if exists {
		return ErrUserAlreadyExists
	} else {
		if err != nil {
			return err
		}
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(driver.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}
	driver.Password = string(hash)
	err = s.Repos.AddNewDriver(ctx, driver)
	if err != nil {
		return err
	}
	err = s.Broker.SendMessage("Driver " + driver.Phone + " created account")
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Authorize(ctx context.Context, driver model.Login) error {
	searchUser := model.Driver{
		Phone:    driver.Phone,
		Password: driver.Password,
	}
	exists, err := s.Repos.Exists(ctx, driver)
	if !exists {
		if err != ErrUserNotFound {
			return err
		}
	}

	comparedUser, err := s.Repos.GetDriverPhoneAndPasswordByPhone(ctx, searchUser.Phone)
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(comparedUser.Password), []byte(searchUser.Password))
	if err != nil {
		return ErrWrongPassword
	}
	return nil
}

func (s *Service) GetAllDrivers(ctx context.Context) ([]model.Driver, error) {
	drivers, err := s.Repos.GetAllDrivers(ctx)
	if err != nil {
		return nil, err
	}
	return drivers, nil
}

func (s *Service) GetAllFreeDrivers(ctx context.Context) ([]model.Driver, error) {
	drivers, err := s.Repos.GetAllFreeDrivers(ctx)
	if err != nil {
		return nil, err
	}
	return drivers, nil
}
func (s *Service) SetDriverStatus(ctx context.Context, phone string) error {

	driver, err := s.Repos.GetDriverPhoneAndPasswordByPhone(ctx, phone)
	if err != nil {
		return err
	}
	err = s.Repos.UpdateDriverStatus(ctx, driver)
	if err != nil {
		return err
	}
	return nil
}
