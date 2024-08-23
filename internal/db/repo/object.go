package repo

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/Mungrel/objects/internal/types"
)

type ObjectStore struct {
	db *sqlx.DB
}

func (s ObjectStore) Create(ctx context.Context, object types.Object) error {
	const insert = `
		INSERT INTO object (
			id,
			name,
			content,
			created_at,
			updated_at
		) VALUES (
			:id,
			:name,
			:content,
			:created_at,
			:updated_at
		)`

	if _, err := s.db.NamedExecContext(ctx, insert, object); err != nil {
		return fmt.Errorf("ObjectStore#Create: %w", err)
	}

	return nil
}

func (s ObjectStore) List(ctx context.Context) ([]types.Object, error) {
	const list = `
		SELECT 
			id,
			name,
			content,
			created_at,
			updated_at
		FROM
			object
		ORDER BY
			updated_at DESC`

	var objects []types.Object
	if err := s.db.SelectContext(ctx, &objects, list); err != nil {
		return nil, fmt.Errorf("ObjectStore#List: %w", err)
	}

	return objects, nil
}
