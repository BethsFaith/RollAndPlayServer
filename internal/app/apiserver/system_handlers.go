package apiserver

import (
	"RnpServer/internal/app/model"
	"RnpServer/internal/app/store"
	"encoding/json"
	"net/http"
)

func (s *server) handleSystemCreate() http.HandlerFunc {
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

		system := &model.System{
			Name: req.Name,
			Icon: req.Icon,
		}
		if err := s.store.System().Create(system); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, http.StatusCreated, system)
	}
}

func (s *server) handleSystemGetRaces() http.HandlerFunc {
	type request struct {
		Id int `json:"id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, http.StatusBadRequest, err)
			return
		}

		foundSystem, err := s.store.System().Find(req.Id)
		if err != nil || foundSystem == nil {
			s.error(w, http.StatusUnprocessableEntity, store.ErrorRecordNotFound)
			return
		}

		components, err := s.store.System().GetRaces(req.Id)
		if err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		var races []*model.Race
		for _, value := range components {
			race, err := s.store.Race().Find(value.ComponentId)

			if err != nil {
				s.error(w, http.StatusInternalServerError, err)
			}

			races = append(races, race)
		}

		s.respond(w, http.StatusOK, races)
	}
}

func (s *server) handleSystemGetClasses() http.HandlerFunc {
	type request struct {
		Id int `json:"id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, http.StatusBadRequest, err)
			return
		}

		foundSystem, err := s.store.System().Find(req.Id)
		if err != nil || foundSystem == nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		components, err := s.store.System().GetCharacterClasses(req.Id)
		if err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		var classes []*model.CharacterClass
		for _, value := range components {
			class, err := s.store.CharacterClass().Find(value.ComponentId)

			if err != nil {
				s.error(w, http.StatusInternalServerError, err)
			}

			classes = append(classes, class)
		}

		s.respond(w, http.StatusOK, classes)
	}
}

func (s *server) handleSystemGetSkillCategories() http.HandlerFunc {
	type request struct {
		Id int `json:"id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, http.StatusBadRequest, err)
			return
		}

		components, err := s.store.System().GetSkillCategories(req.Id)
		if err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		var categories []*model.SkillCategory
		for _, value := range components {
			category, err := s.store.Skill().FindCategory(value.ComponentId)

			if err != nil {
				s.error(w, http.StatusInternalServerError, err)
			}

			categories = append(categories, category)
		}

		s.respond(w, http.StatusOK, categories)
	}
}

func (s *server) handleSystemUpdate() http.HandlerFunc {
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

		system := &model.System{
			ID:   req.ID,
			Name: req.Name,
			Icon: req.Icon,
		}

		foundSys, err := s.store.System().Find(system.ID)
		if err != nil || foundSys == nil {
			s.error(w, http.StatusUnprocessableEntity, store.ErrorRecordNotFound)
			return
		}

		if err := s.store.System().Update(system); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, http.StatusOK, system)
	}
}

func (s *server) handleSystemBindRace() http.HandlerFunc {
	type request struct {
		ID     int `json:"id"`
		RaceId int `json:"race_id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, http.StatusBadRequest, err)
			return
		}

		foundSystem, err := s.store.System().Find(req.ID)
		if err != nil || foundSystem == nil {
			s.error(w, http.StatusUnprocessableEntity, store.ErrorRecordNotFound)
			return
		}
		foundRace, err := s.store.Race().Find(req.RaceId)
		if err != nil || foundRace == nil {
			s.error(w, http.StatusUnprocessableEntity, store.ErrorRecordNotFound)
			return
		}

		components, err := s.store.System().AddRace(req.ID, req.RaceId)
		if err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		var races []*model.Race
		for _, value := range components {
			race, err := s.store.Race().Find(value.ComponentId)

			if err != nil {
				s.error(w, http.StatusInternalServerError, err)
			}

			races = append(races, race)
		}

		if len(races) == 0 {
			s.error(w, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, http.StatusOK, races)
	}
}

func (s *server) handleSystemBindSkillCategory() http.HandlerFunc {
	type request struct {
		ID         int `json:"id"`
		CategoryId int `json:"skill_category_id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, http.StatusBadRequest, err)
			return
		}

		foundSystem, err := s.store.System().Find(req.ID)
		if err != nil || foundSystem == nil {
			s.error(w, http.StatusUnprocessableEntity, store.ErrorRecordNotFound)
			return
		}
		foundRace, err := s.store.Skill().FindCategory(req.CategoryId)
		if err != nil || foundRace == nil {
			s.error(w, http.StatusUnprocessableEntity, store.ErrorRecordNotFound)
			return
		}

		components, err := s.store.System().AddSkillCategory(req.ID, req.CategoryId)
		if err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		var categories []*model.SkillCategory
		for _, value := range components {
			race, err := s.store.Skill().FindCategory(value.ComponentId)

			if err != nil {
				s.error(w, http.StatusInternalServerError, err)
			}

			categories = append(categories, race)
		}

		if len(categories) == 0 {
			s.error(w, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, http.StatusOK, categories)
	}
}

func (s *server) handleSystemBindClass() http.HandlerFunc {
	type request struct {
		ID      int `json:"id"`
		ClassId int `json:"class_id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, http.StatusBadRequest, err)
			return
		}

		foundSystem, err := s.store.System().Find(req.ID)
		if err != nil || foundSystem == nil {
			s.error(w, http.StatusUnprocessableEntity, store.ErrorRecordNotFound)
			return
		}
		foundRace, err := s.store.CharacterClass().Find(req.ClassId)
		if err != nil || foundRace == nil {
			s.error(w, http.StatusUnprocessableEntity, store.ErrorRecordNotFound)
			return
		}

		components, err := s.store.System().AddCharacterClass(req.ID, req.ClassId)
		if err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		var classes []*model.CharacterClass
		for _, value := range components {
			race, err := s.store.CharacterClass().Find(value.ComponentId)

			if err != nil {
				s.error(w, http.StatusInternalServerError, err)
			}

			classes = append(classes, race)
		}

		if len(classes) == 0 {
			s.error(w, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, http.StatusOK, classes)
	}
}

func (s *server) handleSystemDelete() http.HandlerFunc {
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

		system := &model.System{
			ID:   req.ID,
			Name: req.Name,
			Icon: req.Icon,
		}

		foundRace, err := s.store.System().Find(system.ID)
		if err != nil || foundRace == nil {
			s.error(w, http.StatusUnprocessableEntity, store.ErrorRecordNotFound)
			return
		}

		if err := s.store.System().Delete(system.ID); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, http.StatusOK, system)
	}
}
