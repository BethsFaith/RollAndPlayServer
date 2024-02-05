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

		authUser := r.Context().Value(ctxKeyUser).(*model.User)
		race := &model.Race{
			Name:   req.Name,
			Model:  req.Model,
			UserId: authUser.ID,
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

		authUser := r.Context().Value(ctxKeyUser).(*model.User)
		race := &model.Race{
			ID:     req.ID,
			Name:   req.Name,
			Model:  req.Model,
			UserId: authUser.ID,
		}

		foundRace, err := s.store.Race().Find(race.ID)
		if err != nil || foundRace == nil {
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

		authUser := r.Context().Value(ctxKeyUser).(*model.User)
		race := &model.Race{
			ID:     req.ID,
			Name:   req.Name,
			Model:  req.Model,
			UserId: authUser.ID,
		}

		foundRace, err := s.store.Race().Find(race.ID)
		if err != nil || foundRace == nil {
			s.error(w, http.StatusUnprocessableEntity, store.ErrorRecordNotFound)
			return
		}

		if err := s.store.Race().Delete(race.ID); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, http.StatusOK, race)
	}
}

func (s *server) handleRaceBonusCreate() http.HandlerFunc {
	type request struct {
		RaceId  int `json:"race_id"`
		SkillId int `json:"skill_id"`
		Bonus   int `json:"bonus"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, http.StatusBadRequest, err)
			return
		}

		bonus := &model.RaceBonus{
			RaceId:  req.RaceId,
			SkillId: req.SkillId,
			Bonus:   req.Bonus,
		}
		if err := s.store.RaceBonus().Create(bonus); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, http.StatusCreated, bonus)
	}
}

func (s *server) handleRaceBonuses() http.HandlerFunc {
	type request struct {
		RaceId int `json:"race_id"`
		Bonus  int `json:"bonus"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, http.StatusBadRequest, err)
			return
		}

		bonuses, err := s.store.RaceBonus().FindByRaceId(req.RaceId)
		if err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, http.StatusOK, bonuses)
	}
}

func (s *server) handleRaceBonusUpdate() http.HandlerFunc {
	type request struct {
		RaceId  int `json:"race_id"`
		SkillId int `json:"skill_id"`
		Bonus   int `json:"bonus"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, http.StatusBadRequest, err)
			return
		}

		bonus := &model.RaceBonus{
			RaceId:  req.RaceId,
			SkillId: req.SkillId,
			Bonus:   req.Bonus,
		}

		oldBonusData, err := s.store.RaceBonus().Find(bonus.RaceId, bonus.SkillId)
		if err != nil || oldBonusData == nil {
			s.error(w, http.StatusUnprocessableEntity, store.ErrorRecordNotFound)
			return
		}

		if err := s.store.RaceBonus().Update(bonus); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, http.StatusOK, bonus)
	}
}

func (s *server) handleRaceBonusDelete() http.HandlerFunc {
	type request struct {
		RaceId  int `json:"race_id"`
		SkillId int `json:"skill_id"`
		Bonus   int `json:"bonus"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, http.StatusBadRequest, err)
			return
		}

		bonus := &model.RaceBonus{
			RaceId:  req.RaceId,
			SkillId: req.SkillId,
			Bonus:   req.Bonus,
		}

		foundBonus, err := s.store.RaceBonus().Find(bonus.RaceId, bonus.SkillId)
		if err != nil || foundBonus == nil {
			s.error(w, http.StatusUnprocessableEntity, store.ErrorRecordNotFound)
			return
		}

		if err := s.store.RaceBonus().Delete(bonus.RaceId, bonus.SkillId); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, http.StatusOK, bonus)
	}
}
