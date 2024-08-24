package auth

import (
	"context"

	"github.com/ArdyJunata/grpc-go/internal/config"
	internalerror "github.com/ArdyJunata/grpc-go/internal/error"
	"github.com/sirupsen/logrus"
)

type Repository interface {
	CreateAuth(ctx context.Context, model Auth) (err error)
	GetAuthByUsername(ctx context.Context, username string) (model Auth, err error)
}

type service struct {
	repository Repository
}

func newService(repository Repository) service {
	return service{
		repository: repository,
	}
}

func (s service) register(ctx context.Context, req registerRequest) (err error) {
	authModel := req.ParseToModel()

	if err = authModel.encryptPassword(config.Cfg.App.Creds.SlatPassword); err != nil {
		logrus.Error("[register, encryptPassword] error with details : ", err.Error())
		return
	}

	isExists, err := s.repository.GetAuthByUsername(ctx, authModel.Username)
	if err != nil {
		if err != internalerror.ErrNotFound {
			logrus.Error("[register, GetAuthByUsername] error with details : ", err.Error())
			return
		}
	}

	if isExists.Id != 0 {
		err = internalerror.ErrAlreadyExists
		logrus.Error("[register, ValidateExists] error with details : ", err.Error())
		return
	}

	err = s.repository.CreateAuth(ctx, authModel)
	if err != nil {
		logrus.Error("[register, CreateAuth] error with details : ", err.Error())
		return
	}

	return
}

func (s service) login(ctx context.Context, req loginRequest) (token string, err error) {
	authModel := req.ParseToModel()

	auth, err := s.repository.GetAuthByUsername(ctx, authModel.Username)
	if err != nil {
		if err == internalerror.ErrNotFound {
			err = internalerror.ErrUnauthenticated
		}
		logrus.Error("[login, GetAuthByUsername] error with details : ", err.Error())
		return
	}

	if err = auth.VerifyPlainPassword(authModel.Password); err != nil {
		logrus.Error("[login, VerifyPlainPassword] error with details : ", err.Error())
		return
	}

	token, err = auth.GenerateJWT(config.Cfg.App.JWT.Secret, config.Cfg.App.JWT.Duration)
	if err != nil {
		logrus.Error("[login, GenerateJWT] error with details : ", err.Error())
		return
	}

	return
}
