package apiserver

import (
	"RnpServer/internal/app/model"
	"RnpServer/internal/app/store"
	"encoding/json"
	"net/http"
)

func (s *server) handleRaceCreate() http.HandlerFunc {
	type request struct {
		Name  string `json:"name"`
		Model string `json:"model"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, http.StatusBadRequest, err)
			return
		}

		race := &model.Race{
			Name:  req.Name,
			Model: req.Model,
		}
		if err := s.store.Race().Create(race); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, http.StatusCreated, race)
	}
}

func (s *server) handleRaceUpdate() http.HandlerFunc {
	type request struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Model string `json:"model"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, http.StatusBadRequest, err)
			return
		}

		race := &model.Race{
			ID:    req.ID,
			Name:  req.Name,
			Model: req.Model,
		}

		oldSkillData, err := s.store.Race().Find(race.ID)
		if err != nil || oldSkillData == nil {
			s.error(w, http.StatusUnprocessableEntity, store.ErrorRecordNotFound)
			return
		}

		if err := s.store.Race().Update(race); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, http.StatusOK, race)
	}
}

func (s *server) handleRaceDelete() http.HandlerFunc {
	type request struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Model string `json:"model"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, http.StatusBadRequest, err)
			return
		}

		race := &model.Race{
			ID:    req.ID,
			Name:  req.Name,
			Model: req.Model,
		}
		if err := s.store.Race().Delete(race.ID); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, http.StatusOK, race)
	}
}
