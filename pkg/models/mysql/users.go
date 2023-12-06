package mysql

import (
	"database/sql"
	"strings"

	"github.com/AliiAhmadi/ShareCode/pkg/models"
	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	DB *sql.DB
}

func (model *UserModel) Insert(name string, email string, password string) (int, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return 0, err
	}

	statement := `INSERT INTO users (name, email, hashed_password, created) VALUES (?, ?, ?, UTC_TIMESTAMP())`
	result, err := model.DB.Exec(statement, name, email, string(hashedPassword))
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 && strings.Contains(mysqlErr.Message, "users_uc_email") {
				return 0, models.ErrDuplicateEmail
			}
		}
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (model *UserModel) Authenticate(email string, password string) (int, error) {
	user := &models.User{}

	row := model.DB.QueryRow("SELECT id, hashed_password FROM users WHERE email = ?", email)
	err := row.Scan(&user.ID, &user.HashedPassword)
	if err == sql.ErrNoRows {
		return 0, models.ErrInvalidCredentials
	} else if err != nil {
		return 0, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return 0, models.ErrInvalidCredentials
	} else if err != nil {
		return 0, nil
	}

	return user.ID, nil
}

func (model *UserModel) Get(id int) (*models.User, error) {
	return nil, nil
}
