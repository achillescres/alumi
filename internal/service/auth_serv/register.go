package auth_serv

import (
	"context"
	"fmt"
	"github.com/achillescres/pkg/utils"
	"itamconnect/internal/domain/valueobject"
)

type Register struct {
	Login         string               `binding:"required"`
	Password      string               `binding:"required"`
	Email         valueobject.Email    `binding:"required,email"`
	Bio           string               `binding:"required"`
	EducationInfo string               `binding:"required"`
	Phone         string               `binding:"required,e164"`
	Telegram      string               `binding:"required"`
	OtherContacts string               `binding:""`
	Skills        []string             `binding:"required"`
	Type          valueobject.UserType `binding:"required"`
}

func (a *Auth) Register(ctx context.Context, r Register) error {
	ew := utils.NewErrorWrapper("Auth - Register")

	hashedPass, err := a.hasher.Hash(r.Password)
	if err != nil {
		return ew(fmt.Errorf("hash password: %w", err))
	}

	_, err = a.c.User.Create().
		SetLogin(r.Login).SetHashedPassword(hashedPass).
		SetEmail(r.Email).
		SetBio(r.Bio).
		SetEducationInfo(r.EducationInfo).
		SetPhone(r.Phone).SetTelegram(r.Telegram).SetOtherContacts(r.OtherContacts).
		SetSkills(r.Skills).
		SetType(r.Type).
		Save(ctx)
	if err != nil {
		return ew(fmt.Errorf("create user: %w", err))
	}

	return nil
}

/*
field.String("login"),
		field.String("hashed_password"),
		field.String("email").GoType(valueobject.Email("")),
		field.String("bio"),
		field.String("education_info"),
		field.String("phone").Optional(),
		field.String("telegram").Optional(),
		field.String("other_contacts").Optional(),
		field.Strings("skills"),
		field.Enum("type").GoType(valueobject.UserType("")),
*/
