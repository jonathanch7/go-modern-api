package entity

import (
	"database/sql"
)

type Business struct {
	BusinessId  string
	Name        string
	Description string
	PublicCode  sql.NullString
	RegDate     sql.NullTime
}
