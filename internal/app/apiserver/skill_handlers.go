package apiserver

import (
	"RnpServer/internal/app/model"
	"encoding/json"
	"net/http"
)

func (s *server) handleSkillCreate() http.HandlerFunc {
	type request struct {
		Name       string `json:"name"`
		Icon       string `json:"icon"`
		CategoryId int    `json:"category_id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, http.StatusBadRequest, err)
			return
		}

		skill := &model.Skill{
			Name:       req.Name,
			Icon:       req.Icon,
			CategoryId: req.CategoryId,
		}
		if err := s.store.Skill().Create(skill); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, http.StatusCreated, skill)
	}
}

func (s *server) handleSkillCategoryCreate() http.HandlerFunc {
	type request struct {
		Name string `json:"name"`
		Icon string `json:"icon"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, http.StatusBadRequest, err)
			return
		}

		category := &model.SkillCategory{
			Name: req.Name,
			Icon: req.Icon,
		}
		if err := s.store.Skill().CreateCategory(category); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, http.StatusCreated, category)
	}
}
