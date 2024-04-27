package auth_serv

import (
	"context"
	"errors"
	"fmt"
	"github.com/achillescres/pkg/utils"
	"itamconnect/ent/user"
)

var ErrInvalidLogin = errors.New("invalid login")

type Login struct {
	Login    string
	Password string
}

func (a *Auth) Login(ctx context.Context, l Login) (string, error) {
	ew := utils.NewErrorWrapper("Auth - Login")

	u, err := a.c.User.Query().Where(user.Login(l.Login)).Only(ctx)
	if err != nil {
		return "", ew(fmt.Errorf("query user: %w", err))
	}
	if err := a.hasher.Compare(u.HashedPassword, l.Password); err != nil {
		return "", ew(fmt.Errorf("compare password: %w", err))
	}

	jwt, err := a.jwt.NewUser(l.Login)

	return jwt, nil
}
