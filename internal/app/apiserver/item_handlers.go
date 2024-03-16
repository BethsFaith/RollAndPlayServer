package apiserver

import (
	"RnpServer/internal/app/model"
	"RnpServer/internal/app/store"
	"encoding/json"
	"net/http"
)

func (s *server) handleItemCreate() http.HandlerFunc {
	type request struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
		TypeId      int    `json:"type_id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, http.StatusBadRequest, err)
			return
		}

		authUser := r.Context().Value(ctxKeyUser).(*model.User)
		item := &model.Item{
			Name:        req.Name,
			Description: req.Description,
			Icon:        req.Icon,
			TypeId:      req.TypeId,
			UserId:      authUser.ID,
		}
		if err := s.store.Item().Create(item); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, http.StatusCreated, item)
	}
}

func (s *server) handleItemGet() http.HandlerFunc {
	type request struct {
		ID int `json:"id"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			Items, err := s.store.Item().Get()
			if err != nil {
				s.error(w, http.StatusInternalServerError, err)
				return
			}

			s.respond(w, http.StatusOK, Items)
		} else if req.ID == 0 {
			Items, err := s.store.Item().Get()
			if err != nil {
				s.error(w, http.StatusInternalServerError, err)
				return
			}

			s.respond(w, http.StatusOK, Items)
		} else {
			Item, err := s.store.Item().Find(req.ID)
			if err != nil {
				s.error(w, http.StatusInternalServerError, err)
				return
			}
			s.respond(w, http.StatusOK, Item)
		}
	}
}

func (s *server) handleItemUpdate() http.HandlerFunc {
	type request struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
		TypeId      int    `json:"type_id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, http.StatusBadRequest, err)
			return
		}

		authUser := r.Context().Value(ctxKeyUser).(*model.User)
		item := &model.Item{
			ID:          req.ID,
			Name:        req.Name,
			Description: req.Description,
			Icon:        req.Icon,
			TypeId:      req.TypeId,
			UserId:      authUser.ID,
		}

		oldItemData, err := s.store.Item().Find(item.ID)
		if err != nil || oldItemData == nil {
			s.error(w, http.StatusUnprocessableEntity, store.ErrorRecordNotFound)
			return
		}
		if oldItemData.UserId != authUser.ID {
			s.error(w, http.StatusMethodNotAllowed, store.ErrorNoAccess)
			return
		}

		if err := s.store.Item().Update(item); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, http.StatusOK, item)
	}
}

func (s *server) handleItemDelete() http.HandlerFunc {
	type request struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
		Count       int    `json:"count"`
		TypeId      int    `json:"type_id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, http.StatusBadRequest, err)
			return
		}

		authUser := r.Context().Value(ctxKeyUser).(*model.User)
		item, err := s.store.Item().Find(req.ID)
		if err != nil {
			s.error(w, http.StatusUnprocessableEntity, store.ErrorRecordNotFound)
			return
		}
		if item.UserId != authUser.ID {
			s.error(w, http.StatusMethodNotAllowed, store.ErrorNoAccess)
			return
		}

		if err := s.store.Item().Delete(item.ID); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, http.StatusOK, item)
	}
}

func (s *server) handleItemTypeCreate() http.HandlerFunc {
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
		itemType := &model.ItemType{
			Name:   req.Name,
			Icon:   req.Icon,
			UserId: authUser.ID,
		}
		if err := s.store.Item().CreateType(itemType); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, http.StatusCreated, itemType)
	}
}

func (s *server) handleItemTypeGet() http.HandlerFunc {
	type request struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Icon string `json:"icon"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			Items, err := s.store.Item().GetTypes()
			if err != nil {
				s.error(w, http.StatusInternalServerError, err)
				return
			}

			s.respond(w, http.StatusOK, Items)
		} else {
			Item, err := s.store.Item().FindType(req.ID)
			if err != nil {
				s.error(w, http.StatusInternalServerError, err)
				return
			}
			s.respond(w, http.StatusOK, Item)
		}
	}
}

func (s *server) handleItemTypeUpdate() http.HandlerFunc {
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
		itemType := &model.ItemType{
			ID:     req.ID,
			Name:   req.Name,
			Icon:   req.Icon,
			UserId: authUser.ID,
		}

		oldItemData, err := s.store.Item().FindType(itemType.ID)
		if err != nil || oldItemData == nil {
			s.error(w, http.StatusUnprocessableEntity, store.ErrorRecordNotFound)
			return
		}
		if oldItemData.UserId != authUser.ID {
			s.error(w, http.StatusMethodNotAllowed, store.ErrorNoAccess)
			return
		}

		if err := s.store.Item().UpdateType(itemType); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, http.StatusOK, itemType)
	}
}

func (s *server) handleItemTypeDelete() http.HandlerFunc {
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
		itemType, err := s.store.Item().FindType(req.ID)
		if err != nil || itemType == nil {
			s.error(w, http.StatusUnprocessableEntity, store.ErrorRecordNotFound)
			return
		}
		if itemType.UserId != authUser.ID {
			s.error(w, http.StatusMethodNotAllowed, store.ErrorNoAccess)
			return
		}

		if err := s.store.Item().DeleteType(itemType.ID); err != nil {
			s.error(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, http.StatusOK, itemType)
	}
}
