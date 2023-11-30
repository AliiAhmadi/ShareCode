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

	statement := `SELECT id, title, content, created, expires FROM snippets
	WHERE expires > UTC_TIMESTAMP() AND id = ?`

	row := model.DB.QueryRow(statement, id)

	// Initialize a pointer to a new zeroed Snippet struct.
	snippet := &models.Snippet{}

	err := row.Scan(&snippet.ID, &snippet.Title, &snippet.Content, &snippet.Created, &snippet.Expires)

	if err == sql.ErrNoRows {
		return nil, models.ErrorNoRecord
	} else if err != nil {
		return nil, err
	}

	return snippet, nil
}

// Get 10 most recently created snippets
func (model *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
