package auth_serv

import (
	"context"
	"fmt"
	"github.com/achillescres/pkg/security/ajwt"
	"github.com/achillescres/pkg/security/passlib"
	"github.com/achillescres/pkg/utils"
	"itamconnect/ent"
	"itamconnect/ent/user"
)

type Auth struct {
	c      *ent.Client
	hasher passlib.HashManager
	jwt    ajwt.JWTManager
}

func New(c *ent.Client, hasher passlib.HashManager, jwt ajwt.JWTManager) *Auth {
	return &Auth{c: c, hasher: hasher, jwt: jwt}
}

func (a *Auth) Check(ctx context.Context, jwt string) error {
	ew := utils.NewErrorWrapper("Auth - Check")

	u, err := a.jwt.ParseUser(jwt)
	if err != nil {
		return ew(fmt.Errorf("parse user from jwt: %w", err))
	}
	if _, err := a.c.User.Query().Where(user.Login(u.Login)).Only(ctx); err != nil {
		return ew(fmt.Errorf("query user: %w", err))
	}

	return nil
}
