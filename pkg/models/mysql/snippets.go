package mysql

import (
	"database/sql"

	"github.com/AliiAhmadi/ShareCode/pkg/models"
)

type SnippetModel struct {
	DB *sql.DB
}

// Insert a new snippet to database.
func (model *SnippetModel) Insert(title string, content string, expires string) (int, error) {
	return 0, nil
}

// Get an snippet based on its id
func (model *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

// Get 10 most recently created snippets
func (model *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
