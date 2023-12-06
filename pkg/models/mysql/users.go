package mysql

import (
	"database/sql"

	"github.com/AliiAhmadi/ShareCode/pkg/models"
)

type UserModel struct {
	DB *sql.DB
}

func (model *UserModel) Insert(name string, email string, password string) error {
	return nil
}

func (model *UserModel) Authenticate(email string, password string) (int, error) {
	return 0, nil
}

func (model *UserModel) Get(id int) (*models.User, error) {
	return nil, nil
}
