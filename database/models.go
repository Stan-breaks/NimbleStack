// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package database

import (
	"database/sql"
)

type User struct {
	ID        int64
	Username  string
	Email     string
	Firstname string
	Lastname  string
	Password  string
	CreatedAt sql.NullTime
}
