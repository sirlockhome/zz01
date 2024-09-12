package service

import (
	"context"
	"foxomni/internal/user/model"
	"foxomni/pkg/errs"
	"foxomni/pkg/security"
)

const (
	ErrPassword = "password failed"
)

func (svc *Service) Login(ctx context.Context, userLogin model.UserLogin) (*model.UserAuth, error) {
	foundUser, err := svc.ds.GetUserByUsername(ctx, userLogin.Username)
	if err != nil {
		return nil, err
	}

	if !security.CheckPasswordHash(userLogin.Password, foundUser.PasswordHash) {
		return nil, errs.New(errs.Unauthorized, ErrPassword)
	}

	accessToken, refreshToken, err := svc.jwt.GenerateTokenPair(foundUser.ID)
	if err != nil {
		return nil, err
	}

	return &model.UserAuth{
		UserData:     foundUser,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		AccessExp:    svc.conf.Auth.AccessExp,
		RefreshExp:   svc.conf.Auth.RefreshExp,
	}, nil
}
