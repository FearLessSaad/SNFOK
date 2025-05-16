package auth

import (
	"time"

	"github.com/uptrace/bun"
)

type AuditFields struct {
	CreatedBy string       `bun:"type:varchar(150)"`
	UpdatedBy string       `bun:"type:varchar(150)"`
	CreatedAt time.Time    `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt bun.NullTime `bun:",nullzero"`
}

type UserStatus string

const (
	UserStatusActive      UserStatus = "ACTIVE"
	UserStatusPending     UserStatus = "PENDING"
	UserStatusDeactivated UserStatus = "DEACTIVATED"
)

type Users struct {
	bun.BaseModel `bun:"table:auth.users,alias:u"`

	ID          string `bun:",pk,type:uuid,default:gen_random_uuid()"`
	FirstName   string
	LastName    string
	Email       string `bun:",unique,"`
	Designation string
	Token       string
	Status      UserStatus `bun:",type:varchar(20),notnull,default:'PENDING'"`

	AuditFields
}

const UsersTableName = "auth.users"
