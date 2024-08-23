package types

import "time"

type Object struct {
	ID        string    `db:"id"`
	Name      string    `db:"name"`
	Content   string    `db:"content"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
