package models

import "time"

type Role struct {
	ID          string    `db:"id"`
	TenantID    string    `db:"tenant_id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
