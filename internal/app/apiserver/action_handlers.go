package apiserver

import (
	"RnpServer/internal/app/model"
	"RnpServer/internal/app/store"
	"encoding/json"
	"net/http"
)

func (s *server) handleActionCreate() http.HandlerFunc {
	type request struct {
		Name    string `json:"name"`
		Icon    string `json:"icon"`
		SkillId int    `json:"skill_id"`
		Points  int    `json:"points"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, http.StatusBadRequest, err)
			return
		}

		authUser := r.Context().Value(ctxKeyUser).(*model.User)
		action := &model.Action{
			Name:    req.Name,
			Icon:    req.Icon,
			SkillId: req.SkillId,
			Points:  req.Points,
			UserId:  authUser.ID,
		}
		if err := s.store.Action().Create(action); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, http.StatusCreated, action)
	}
}

func (s *server) handleActionGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		actions, err := s.store.Action().Get()
		if err != nil {
			s.error(w, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, http.StatusOK, actions)
	}
}

func (s *server) handleActionUpdate() http.HandlerFunc {
	type request struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Icon    string `json:"icon"`
		SkillId int    `json:"skill_id"`
		Points  int    `json:"points"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, http.StatusBadRequest, err)
			return
		}

		authUser := r.Context().Value(ctxKeyUser).(*model.User)
		action := &model.Action{
			ID:      req.ID,
			Name:    req.Name,
			Icon:    req.Icon,
			SkillId: req.SkillId,
			Points:  req.Points,
			UserId:  authUser.ID,
		}

		oldActionData, err := s.store.Action().Find(action.ID)
		if err != nil || oldActionData == nil {
			s.error(w, http.StatusUnprocessableEntity, store.ErrorRecordNotFound)
			return
		}

		if err := s.store.Action().Update(action); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, http.StatusOK, action)
	}
}

func (s *server) handleActionDelete() http.HandlerFunc {
	type request struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Icon    string `json:"icon"`
		SkillId int    `json:"skill_id"`
		Points  int    `json:"points"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, http.StatusBadRequest, err)
			return
		}

		authUser := r.Context().Value(ctxKeyUser).(*model.User)
		action := &model.Action{
			ID:      req.ID,
			Name:    req.Name,
			Icon:    req.Icon,
			SkillId: req.SkillId,
			Points:  req.Points,
			UserId:  authUser.ID,
		}
		if err := s.store.Action().Delete(action.ID); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, http.StatusOK, action)
	}
}
