package repositories

import (
	"barebone-go-crud/src/models/entity"
	"context"
	"database/sql"
	"errors"
)

type UserRepository interface {
	// GetAllUser(ctx context.Context) (*entity.User, error)
	GetUserById(ctx context.Context, id int64) (*entity.User, error)
	CreateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	UpdateUser(ctx context.Context, id int64, user *entity.User) (*entity.User, error)
	DeleteUser(ctx context.Context, id int64) error
}

type userRepository struct {
	db *sql.DB
}

// masih gapaham sama ini
func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (userRepo *userRepository) GetUserById(ctx context.Context, id int64) (*entity.User, error) {
	row := userRepo.db.QueryRowContext(ctx, "SELECT id, name, email FROM users WHERE id = ?", id)

	var user entity.User
	err := row.Scan(&user.Id, &user.Name, user.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (userRepo *userRepository) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	stmt, err := userRepo.db.PrepareContext(ctx, "INSERT INTO users (name, email) VALUES (?, ?)")

	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, user.Name, user.Email)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	createUser, err := userRepo.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}

	return createUser, nil
}

func (userRepo *userRepository) UpdateUser(ctx context.Context, id int64, user *entity.User) (*entity.User, error) {
	stmt, err := userRepo.db.PrepareContext(ctx, "UPDATE users SET name = ?, email = ?, WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, user.Name, user.Email, id)
	if err != nil {
		return nil, nil
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return nil, nil
	}
	if affected == 0 {
		return nil, sql.ErrNoRows
	}

	UpdateUser, err := userRepo.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}

	return UpdateUser, nil
}

func (userRepo *userRepository) DeleteUser(ctx context.Context, id int64) error {
	stmt, err := userRepo.db.PrepareContext(ctx, "DELETE FROM users WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, id)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
