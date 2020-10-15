package mysql

import (
	"database/sql"

	"letsgo.net/snippetbox/pkg/models"
)

//SnippetModel ...
type SnippetModel struct {
	DB *sql.DB
}

//This ... will insert a new snippet into the database.
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires)
	VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`
	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}
	// Use the LastInsertId() method on the result object to get the ID // newly inserted record in the snippets table.
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil

}

// This will return a specific snippet based on its id.
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

// This will return the 10 most recently created snippets.
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}