package sqlstore

import (
	"RnpServer/internal/app/model"
	"RnpServer/internal/app/store"
	"database/sql"
	"errors"
	"strconv"
	"strings"
)

type UserRepository struct {
	store *Store
}

// Create ...
func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	return r.store.CreateRetId(
		InsertQ+UsersT+UsersP+"values ($1, $2, $3) RETURNING id", u.Email, u.Nickname, u.EncryptedPassword,
	).Scan(&u.ID)
}

// FindByEmail ...
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}

	if err := r.store.SelectRow(
		SelectQ+UsersT+"WHERE email = $1", email,
	).Scan(
		&u.ID,
		&u.Email,
		&u.Nickname,
		&u.EncryptedPassword,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, store.ErrorRecordNotFound
		}
		return nil, err
	}

	return u, nil
}

// Find ...
func (r *UserRepository) Find(id int) (*model.User, error) {
	u := &model.User{}

	if err := r.store.SelectRow(
		SelectQ+UsersT+"WHERE id = $1", id,
	).Scan(
		&u.ID,
		&u.Email,
		&u.Nickname,
		&u.EncryptedPassword,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, store.ErrorRecordNotFound
		}
		return nil, err
	}

	return u, nil
}

func (r *UserRepository) Update(u *model.User) error {
	lastU, err := r.Find(u.ID)
	if err != nil {
		return store.ErrorRecordNotFound
	}

	err = u.Validate()
	if err != nil {
		return err
	}

	var fields []string
	var vars []any
	if lastU.Email != u.Email {
		fields = append(fields, "email")
		vars = append(vars, u.Email)
	}
	if lastU.Nickname != u.Nickname {
		fields = append(fields, "nickname")
		vars = append(vars, u.Nickname)
	}
	if !lastU.ComparePassword(u.Password) {
		fields = append(fields, "password")
		vars = append(vars, u.Password)
	}

	for i := 0; i < len(fields); i++ {
		fields[i] += " = $" + strconv.Itoa(i+1)
	}

	query := strings.Join(fields, ", ")
	query = query + " WHERE id = $" + strconv.Itoa(len(fields)+1)

	vars = append(vars, u.ID)

	_, err = r.store.Update(
		UpdateQ+UsersT+"SET "+query, vars...,
	)

	return err
}
