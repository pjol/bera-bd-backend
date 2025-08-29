package auth

import (
	"context"

	"github.com/pjol/bera-bd-backend/db/auth"
	"github.com/pjol/bera-bd-backend/logger"
)

type Service struct {
	db     *auth.Db
	logger *logger.LogCloser
}

func New(db *auth.Db, logger *logger.LogCloser) *Service {
	return &Service{db, logger}
}

func (s *Service) IsAdmin(ctx context.Context, userDid string) bool {
	isAdmin, err := s.db.IsAdmin(ctx, userDid)
	if err != nil {
		s.logger.Logf("error getting admin status for user %s: %s", userDid, err)
	}

	return isAdmin
}
