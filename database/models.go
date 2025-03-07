// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package database

import (
	"database/sql"
)

type SqliteSchema struct {
	Type     string
	Name     string
	TblName  string
	Rootpage int64
	Sql      string
}

type User struct {
	ID        int64
	Username  string
	Email     string
	Firstname string
	Lastname  string
	Password  string
	CreatedAt sql.NullTime
}
