package ramps

import (
	"net/http"

	"github.com/pjol/bera-bd-backend/db/ramps"
	"github.com/pjol/bera-bd-backend/logger"
)

type Service struct {
	db     *ramps.Db
	logger *logger.LogCloser
}

func New(db *ramps.Db, logger *logger.LogCloser) *Service {
	return &Service{db, logger}
}

func (s *Service) NewRampApplication(w http.ResponseWriter, r *http.Request) {

}

func (s *Service) ApproveRampApplication(w http.ResponseWriter, r *http.Request) {

}

func (s *Service) GetRampApplicationsAuthed(w http.ResponseWriter, r *http.Request) {

}

func (s *Service) GetApprovedRampApplications(w http.ResponseWriter, r *http.Request) {

}
