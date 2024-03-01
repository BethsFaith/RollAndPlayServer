package apiserver

import (
	"RnpServer/internal/app/model"
	"RnpServer/internal/app/store"
	"encoding/json"
	"net/http"
)

func (s *server) handleCharacteristicCreate() http.HandlerFunc {
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
		characteristic := &model.Characteristic{
			Name:   req.Name,
			Icon:   req.Icon,
			UserId: authUser.ID,
		}
		if err := s.store.Characteristic().Create(characteristic); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, http.StatusCreated, characteristic)
	}
}

func (s *server) handleCharacteristicGet() http.HandlerFunc {
	type request struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Icon string `json:"icon"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			characteristics, err := s.store.Characteristic().Get()
			if err != nil {
				s.error(w, http.StatusInternalServerError, err)
				return
			}

			s.respond(w, http.StatusOK, characteristics)
		} else {

			characteristic, err := s.store.Characteristic().Find(req.ID)
			if err != nil {
				s.error(w, http.StatusInternalServerError, err)
				return
			}
			s.respond(w, http.StatusOK, characteristic)
		}
	}
}

func (s *server) handleCharacteristicUpdate() http.HandlerFunc {
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
		c := &model.Characteristic{
			ID:     req.ID,
			Name:   req.Name,
			Icon:   req.Icon,
			UserId: authUser.ID,
		}

		oldCharacteristicData, err := s.store.Characteristic().Find(c.ID)
		if err != nil || oldCharacteristicData == nil {
			s.error(w, http.StatusUnprocessableEntity, store.ErrorRecordNotFound)
			return
		}

		if err := s.store.Characteristic().Update(c); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, http.StatusOK, c)
	}
}

func (s *server) handleCharacteristicDelete() http.HandlerFunc {
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
		c := &model.Characteristic{
			ID:     req.ID,
			Name:   req.Name,
			Icon:   req.Icon,
			UserId: authUser.ID,
		}
		if err := s.store.Characteristic().Delete(c.ID); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, http.StatusOK, c)
	}
}
