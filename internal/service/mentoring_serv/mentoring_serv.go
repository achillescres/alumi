package mentoring_serv

import (
	"context"
	"fmt"
	"github.com/achillescres/pkg/utils"
	"itamconnect/ent"
	"itamconnect/ent/match"
	"itamconnect/ent/menti"
	"itamconnect/ent/mentor"
	"itamconnect/ent/skill"
	"itamconnect/ent/user"
	"itamconnect/internal/domain/valueobject"
)

type SearchMentorsOpts struct {
	Skills []string `binding:"required,dive,required"`
}

type MentoringService interface {
	GetMentors(ctx context.Context, opts SearchMentorsOpts) ([]*ent.User, error)
	Match(ctx context.Context, mentiID, mentorID int) error
	SolveMatch(ctx context.Context, mentorID, requestID int, accept bool) error
	GetMatches(ctx context.Context, user *ent.User) ([]*ent.Match, error)
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

func (m *mentoringService) Match(ctx context.Context, mentiID, mentorID int) error {
	ew := utils.NewErrorWrapper("mentoringService - Match")

	_, err := m.c.Match.Create().
		SetMentiID(mentiID).
		SetMentorID(mentorID).
		SetStatus(valueobject.MatchStatusPending).
		Save(ctx)
	if err != nil {
		return ew(fmt.Errorf("create match: %w", err))
	}

	return nil
}

func (m *mentoringService) SolveMatch(ctx context.Context, mentorID, matchID int, accept bool) error {
	ew := utils.NewErrorWrapper("mentoringService - SolveMatch")

	acceptMatch, err := m.c.Match.Get(ctx, matchID)
	if err != nil {
		return ew(fmt.Errorf("get match: %w", err))
	}
	if acceptMatch.MentorID != mentorID {
		return ew(fmt.Errorf("unauthorized mentor"))
	}

	var newStatus = valueobject.MatchStatusAccepted
	if !accept {
		newStatus = valueobject.MatchStatusDeclined
	}
	_, err = acceptMatch.Update().SetStatus(newStatus).Save(ctx)
	if err != nil {
		return ew(fmt.Errorf("update match status to accepted: %w", err))
	}

	return nil
}

func (m *mentoringService) GetMatches(ctx context.Context, u *ent.User) ([]*ent.Match, error) {
	switch u.Type {
	case valueobject.UserTypeMenti:
		return m.c.Match.Query().Where(match.HasMentiWith(menti.HasUserWith(user.ID(u.ID)))).All(ctx)
	case valueobject.UserTypeMentor:
		return m.c.Match.Query().Where(match.HasMentorWith(mentor.HasUserWith(user.ID(u.ID)))).All(ctx)
	default:
		return nil, fmt.Errorf("invalid user type")
	}
}
