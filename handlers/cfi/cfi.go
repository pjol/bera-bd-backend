package cfi

import (
	"net/http"

	"github.com/pjol/bera-bd-backend/db/cfi"
	"github.com/pjol/bera-bd-backend/logger"
)

type Service struct {
	db     *cfi.Db
	logger *logger.LogCloser
}

func New(db *cfi.Db, logger *logger.LogCloser) *Service {
	return &Service{db, logger}
}

func (s *Service) NewCfiApplication(w http.ResponseWriter, r *http.Request) {

}

func (s *Service) ApproveCfiApplication(w http.ResponseWriter, r *http.Request) {

}

func (s *Service) GetCfiApplicationsAuthed(w http.ResponseWriter, r *http.Request) {

}

func (s *Service) GetApprovedCfiApplications(w http.ResponseWriter, r *http.Request) {

}
