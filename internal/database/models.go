// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"database/sql"
)

type GoCrmUser struct {
	UserID         int32
	UserAccount    string
	UserPassword   string
	UserSalt       string
	UserLoginTime  sql.NullTime
	UserLogoutTime sql.NullTime
	UserLoginIp    sql.NullString
	UserCreatedAt  sql.NullTime
	UserUpdatedAt  sql.NullTime
}
