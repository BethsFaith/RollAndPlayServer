package apiserver

import (
	"RnpServer/internal/app/model"
	"RnpServer/internal/app/store"
	"encoding/json"
	"net/http"
)

func (s *server) handleSkillCreate() http.HandlerFunc {
	type request struct {
		Name             string `json:"name"`
		Icon             string `json:"icon"`
		CategoryId       int    `json:"category_id"`
		CharacteristicId int    `json:"characteristic_id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, http.StatusBadRequest, err)
			return
		}

		authUser := r.Context().Value(ctxKeyUser).(*model.User)
		skill := &model.Skill{
			Name:             req.Name,
			Icon:             req.Icon,
			CategoryId:       req.CategoryId,
			CharacteristicId: req.CharacteristicId,
			UserId:           authUser.ID,
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

		authUser := r.Context().Value(ctxKeyUser).(*model.User)
		category := &model.SkillCategory{
			Name:   req.Name,
			Icon:   req.Icon,
			UserId: authUser.ID,
		}
		if err := s.store.Skill().CreateCategory(category); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, http.StatusCreated, category)
	}
}

func (s *server) handleSkillGet() http.HandlerFunc {
	type request struct {
		CategoryId int `json:"category_id"`
		Id         int `json:"id"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			skills, err := s.store.Skill().Get()
			if err != nil {
				s.error(w, http.StatusInternalServerError, err)
				return
			}

			s.respond(w, http.StatusOK, skills)
		} else {
			if req.Id != 0 {
				skill, err := s.store.Skill().Find(req.Id)
				if err != nil {
					s.error(w, http.StatusInternalServerError, err)
					return
				}
				s.respond(w, http.StatusOK, skill)
			} else if req.CategoryId != 0 {
				skills, err := s.store.Skill().GetByCategory(req.CategoryId)
				if err != nil {
					s.error(w, http.StatusInternalServerError, err)
					return
				}
				s.respond(w, http.StatusOK, skills)
			} else {
				s.error(w, http.StatusBadRequest, nil)
			}
		}
	}
}

func (s *server) handleSkillCategoryGet() http.HandlerFunc {
	type request struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Icon string `json:"icon"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			categories, err := s.store.Skill().GetCategories()
			if err != nil {
				s.error(w, http.StatusInternalServerError, err)
				return
			}

			s.respond(w, http.StatusOK, categories)
		} else {
			cat, err := s.store.Skill().FindCategory(req.ID)
			if err != nil {
				s.error(w, http.StatusInternalServerError, err)
				return
			}
			s.respond(w, http.StatusOK, cat)
		}
	}
}

func (s *server) handleSkillUpdate() http.HandlerFunc {
	type request struct {
		ID               int    `json:"id"`
		Name             string `json:"name"`
		Icon             string `json:"icon"`
		CategoryId       int    `json:"category_id"`
		CharacteristicId int    `json:"characteristic_id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, http.StatusBadRequest, err)
			return
		}

		authUser := r.Context().Value(ctxKeyUser).(*model.User)
		skill := &model.Skill{
			ID:               req.ID,
			Name:             req.Name,
			Icon:             req.Icon,
			CategoryId:       req.CategoryId,
			CharacteristicId: req.CharacteristicId,
			UserId:           authUser.ID,
		}

		oldSkillData, err := s.store.Skill().Find(skill.ID)
		if err != nil || oldSkillData == nil {
			s.error(w, http.StatusUnprocessableEntity, store.ErrorRecordNotFound)
			return
		}

		if err := s.store.Skill().Update(skill); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, http.StatusOK, skill)
	}
}

func (s *server) handleSkillCategoryUpdate() http.HandlerFunc {
	type request struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Icon string `json:"icon"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, http.StatusBadRequest, err)
			return
		}

		authUser := r.Context().Value(ctxKeyUser).(*model.User)
		category := &model.SkillCategory{
			ID:     req.ID,
			Name:   req.Name,
			Icon:   req.Icon,
			UserId: authUser.ID,
		}

		oldSkillData, err := s.store.Skill().FindCategory(category.ID)
		if err != nil {
			s.error(w, http.StatusUnprocessableEntity, store.ErrorRecordNotFound)
		}
		if category.Icon == "" {
			category.Icon = oldSkillData.Icon
		}
		if category.Name == "" {
			category.Name = oldSkillData.Name
		}

		if err := s.store.Skill().UpdateCategory(category); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, http.StatusOK, category)
	}
}

func (s *server) handleSkillDelete() http.HandlerFunc {
	type request struct {
		ID int `json:"id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, http.StatusBadRequest, err)
			return
		}

		authUser := r.Context().Value(ctxKeyUser).(*model.User)
		skill := &model.Skill{
			ID:     req.ID,
			UserId: authUser.ID,
		}
		if err := s.store.Skill().Delete(skill.ID); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, http.StatusOK, skill)
	}
}

func (s *server) handleSkillCategoryDelete() http.HandlerFunc {
	type request struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Icon string `json:"icon"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, http.StatusBadRequest, err)
			return
		}

		authUser := r.Context().Value(ctxKeyUser).(*model.User)
		category := &model.SkillCategory{
			ID:     req.ID,
			Name:   req.Name,
			Icon:   req.Icon,
			UserId: authUser.ID,
		}
		if err := s.store.Skill().DeleteCategory(category.ID); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, http.StatusOK, category)
	}
}
