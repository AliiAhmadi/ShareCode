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
	statement := `INSERT INTO snippets (title, content, created, expires) 
	VALUES (?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := model.DB.Exec(statement, title, content, expires)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// Get an snippet based on its id
func (model *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

// Get 10 most recently created snippets
func (model *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
