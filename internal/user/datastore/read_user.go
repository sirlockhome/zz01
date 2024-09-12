package datastore

import (
	"context"
	"database/sql"
	"errors"
	"foxomni/internal/user/model"
	"foxomni/pkg/errs"
)

func (ds *Datastore) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	query := `
	SELECT 
		u.id, u.username, u.password_hash, u.first_name,
		u.last_name, u.gender, u.date_of_birth, u.email,
		u.phone_number, u.user_type, u.partner_id, u.last_login,
		u.created_at, u.updated_at
	FROM 
		users u
	WHERE username = @p1
	`

	var user model.User
	if err := ds.sql.DB.GetContext(ctx, &user, query, username); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.New(errs.NotFound, err)
		}

		return nil, errs.New(errs.Internal, err)
	}

	return &user, nil
}
