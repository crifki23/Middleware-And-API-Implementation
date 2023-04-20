package user_pg

import (
	"chapter3-sesi2/entity"
	"chapter3-sesi2/pkg/errs"
	"chapter3-sesi2/repository/user_repository"
	"database/sql"
)

const (
	createUserQuery = `
		INSERT INTO "users"
		(
			"email",
			"password"
		)
		VALUES($1, $2)
	`
	getUserByEmailQuery = `
		SELECT id, email, password, level from "users"
		WHERE email = $1;
	`
)

type userPG struct {
	db *sql.DB
}

func NewUserPG(db *sql.DB) user_repository.UserRepository {
	return &userPG{
		db: db,
	}
}
func (u *userPG) CreateNewUser(user entity.User) errs.MessageErr {
	_, err := u.db.Exec(createUserQuery, user.Email, user.Password)
	if err != nil {
		return errs.NewInternalServerError("something went wrong")
	}
	return nil
}

// GetUserByEmail implements user_repository.UserRepository
func (u *userPG) GetUserByEmail(userEmail string) (*entity.User, errs.MessageErr) {
	row := u.db.QueryRow(getUserByEmailQuery, userEmail)
	var user entity.User
	err := row.Scan(&user.Id, &user.Email, &user.Password, &user.Level)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("user not found")
		}
		return nil, errs.NewInternalServerError("something went wrong")
	}
	return &user, nil
}

// GetUserById implements user_repository.UserRepository
func (*userPG) GetUserById(userId int) (*entity.User, errs.MessageErr) {
	return nil, nil
}
