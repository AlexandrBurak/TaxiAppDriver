package test

import (
	"TaxiAppDriver/internal/model"
	srv "TaxiAppDriver/internal/service"
	"context"
	"errors"
	"testing"

	"github.com/bmizerany/assert"
	"github.com/golang/mock/gomock"
)

func TestService_Save(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repos := NewMockRepository(ctrl)
	ctx := context.TODO()
	user := model.Driver{Phone: "123123123", Email: "123@qawe", Password: "zxczxccz"}
	repos.EXPECT().Exists(gomock.Any(), model.Login{Phone: user.Phone, Password: user.Password}).Return(false, nil)
	repos.EXPECT().AddNewDriver(gomock.Any(), gomock.Any()).Return(nil)
	broker := NewMockBroker(ctrl)
	broker.EXPECT().SendMessage(gomock.Any()).Return(nil)
	service := srv.Service{Repos: repos, Broker: broker}

	err := service.Save(ctx, user)
	assert.Equal(t, nil, err, "didn't save user with correct data")
}

func TestService_Save_UserExistsError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.TODO()
	repos := NewMockRepository(ctrl)
	user := model.Driver{Phone: "123123123", Password: "zxczxccz"}
	repos.EXPECT().Exists(gomock.Any(), model.Login{Phone: user.Phone, Password: user.Password}).Return(true, nil)

	service := srv.Service{Repos: repos}

	err := service.Save(ctx, user)
	assert.Equal(t, srv.ErrUserAlreadyExists, err, "didn't catch that user already exists")
}

func TestService_Save_UnhandledError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.TODO()
	repos := NewMockRepository(ctrl)

	user := model.Driver{Phone: "123123123", Email: "123@qawe", Password: "zxczxccz"}
	repos.EXPECT().Exists(gomock.Any(), model.Login{Phone: user.Phone, Password: user.Password}).Return(false, errors.New("zxzc"))

	service := srv.Service{Repos: repos}

	err := service.Save(ctx, user)
	assert.NotEqual(t, nil, err, "didn't catch unhandled error")
}

func TestService_Authorize_UnhandledError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.TODO()
	repos := NewMockRepository(ctrl)

	user := model.Driver{Phone: "123123123", Password: "2345w891"}
	repos.EXPECT().Exists(gomock.Any(), model.Login{Phone: user.Phone, Password: user.Password}).Return(true, nil)
	repos.EXPECT().GetDriverPhoneAndPasswordByPhone(gomock.Any(), user.Phone).Return(model.Driver{Phone: "123123123", Password: ""}, nil)
	service := srv.Service{Repos: repos}

	err := service.Authorize(ctx, model.Login{Phone: user.Phone, Password: "2345w891"})
	assert.NotEqual(t, nil, err, "didn't catch unhandled error")
}

func TestService_Authorize(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.TODO()
	repos := NewMockRepository(ctrl)

	user := model.Driver{Phone: "123123123", Password: "2345w891"}
	repos.EXPECT().Exists(gomock.Any(), model.Login{Phone: user.Phone, Password: user.Password}).Return(true, nil)
	repos.EXPECT().GetDriverPhoneAndPasswordByPhone(gomock.Any(), user.Phone).Return(model.Driver{Phone: "123123123", Password: "$2a$04$ddAHC4UK4JF0iId7AZR3wOccM4d.sasMvX0Ldsr/.ZrZXYsj7CZeW"}, nil)
	service := srv.Service{Repos: repos}

	err := service.Authorize(ctx, model.Login{Phone: user.Phone, Password: "2345w891"})
	assert.Equal(t, nil, err, "didn't log in with correct data")
}
func TestService_Authorize_Wrong_Password(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.TODO()
	repos := NewMockRepository(ctrl)

	user := model.Driver{Phone: "123123123", Password: "2345w892"}
	repos.EXPECT().GetDriverPhoneAndPasswordByPhone(gomock.Any(), user.Phone).Return(model.Driver{Phone: "123123123", Password: "$2a$04$ddAHC4UK4JF0iId7AZR3wOccM4d.sasMvX0Ldsr/.ZrZXYsj7CZeW"}, nil)
	repos.EXPECT().Exists(gomock.Any(), model.Login{Phone: user.Phone, Password: user.Password}).Return(true, nil)
	service := srv.Service{Repos: repos}

	err := service.Authorize(ctx, model.Login{Phone: user.Phone, Password: "2345w892"})
	assert.Equal(t, srv.ErrWrongPassword, err, "logged in with incorrect password")
}
