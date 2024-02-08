package apiserver

import (
	"RnpServer/internal/app/model"
	"RnpServer/internal/app/store"
	"encoding/json"
	"net/http"
)

func (s *server) handleClassCreate() http.HandlerFunc {
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
		class := &model.CharacterClass{
			Name:   req.Name,
			Icon:   req.Icon,
			UserId: authUser.ID,
		}
		if err := s.store.CharacterClass().Create(class); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, http.StatusCreated, class)
	}
}

func (s *server) handleClassGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		classes, err := s.store.CharacterClass().Get()
		if err != nil {
			s.error(w, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, http.StatusOK, classes)
	}
}

func (s *server) handleClassUpdate() http.HandlerFunc {
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
		class := &model.CharacterClass{
			ID:     req.ID,
			Name:   req.Name,
			Icon:   req.Icon,
			UserId: authUser.ID,
		}

		oldClassData, err := s.store.CharacterClass().Find(class.ID)
		if err != nil || oldClassData == nil {
			s.error(w, http.StatusUnprocessableEntity, store.ErrorRecordNotFound)
			return
		}

		if err := s.store.CharacterClass().Update(class); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, http.StatusOK, class)
	}
}

func (s *server) handleClassDelete() http.HandlerFunc {
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
		class := &model.CharacterClass{
			ID:     req.ID,
			Name:   req.Name,
			Icon:   req.Icon,
			UserId: authUser.ID,
		}

		if err := s.store.CharacterClass().Delete(class.ID); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, http.StatusOK, class)
	}
}

func (s *server) handleClassBonusCreate() http.HandlerFunc {
	type request struct {
		ClassId int `json:"class_id"`
		SkillId int `json:"skill_id"`
		Bonus   int `json:"bonus"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, http.StatusBadRequest, err)
			return
		}

		bonus := &model.CharacterClassBonus{
			ClassId: req.ClassId,
			SkillId: req.SkillId,
			Bonus:   req.Bonus,
		}
		if err := s.store.CharacterClassBonus().Create(bonus); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, http.StatusCreated, bonus)
	}
}

func (s *server) handleClassBonuses() http.HandlerFunc {
	type request struct {
		ClassId int `json:"class_id"`
		Bonus   int `json:"bonus"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, http.StatusBadRequest, err)
			return
		}

		bonuses, err := s.store.CharacterClassBonus().FindByClassId(req.ClassId)
		if err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, http.StatusOK, bonuses)
	}
}

func (s *server) handleClassBonusUpdate() http.HandlerFunc {
	type request struct {
		ClassId int `json:"class_id"`
		SkillId int `json:"skill_id"`
		Bonus   int `json:"bonus"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, http.StatusBadRequest, err)
			return
		}

		bonus := &model.CharacterClassBonus{
			ClassId: req.ClassId,
			SkillId: req.SkillId,
			Bonus:   req.Bonus,
		}

		oldBonusData, err := s.store.CharacterClassBonus().Find(bonus.ClassId, bonus.SkillId)
		if err != nil || oldBonusData == nil {
			s.error(w, http.StatusUnprocessableEntity, store.ErrorRecordNotFound)
			return
		}

		if err := s.store.CharacterClassBonus().Update(bonus); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, http.StatusOK, bonus)
	}
}

func (s *server) handleClassBonusDelete() http.HandlerFunc {
	type request struct {
		ClassId int `json:"class_id"`
		SkillId int `json:"skill_id"`
		Bonus   int `json:"bonus"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, http.StatusBadRequest, err)
			return
		}

		bonus := &model.CharacterClassBonus{
			ClassId: req.ClassId,
			SkillId: req.SkillId,
			Bonus:   req.Bonus,
		}

		foundBonus, err := s.store.CharacterClassBonus().Find(bonus.ClassId, bonus.SkillId)
		if err != nil || foundBonus == nil {
			s.error(w, http.StatusUnprocessableEntity, store.ErrorRecordNotFound)
			return
		}

		if err := s.store.CharacterClassBonus().Delete(bonus.ClassId, bonus.SkillId); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, http.StatusOK, bonus)
	}
}
