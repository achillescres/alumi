package mentoring_serv

import (
	"context"
	"fmt"
	"github.com/achillescres/pkg/utils"
	"itamconnect/ent"
	"itamconnect/ent/skill"
	"itamconnect/ent/user"
	"itamconnect/internal/domain/valueobject"
)

type SearchMentorsOpts struct {
	Skills []string `binding:"required,dive,required"`
}

type MentoringService interface {
	GetMentors(ctx context.Context, opts SearchMentorsOpts) ([]*ent.User, error)
}

type mentoringService struct {
	c *ent.Client
}

func New(c *ent.Client) MentoringService {
	return &mentoringService{
		c: c,
	}
}

func (m *mentoringService) GetMentors(ctx context.Context, opts SearchMentorsOpts) ([]*ent.User, error) {
	ew := utils.NewErrorWrapper("mentoringService - GetMentorsForMe")

	// заглушка
	var skills = user.HasSkillsWith(skill.Name(opts.Skills[0]))
	for _, sk := range opts.Skills {
		skills = user.And(skills, user.HasSkillsWith(skill.Name(sk)))
	}
	mentors, err := m.c.User.Query().Where(
		user.TypeEQ(valueobject.UserTypeMentor),
		skills,
	).All(ctx)
	if err != nil {
		return nil, ew(fmt.Errorf("query matching mentors: %w", err))
	}

	return mentors, nil
}
