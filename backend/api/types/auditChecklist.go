package types

import "time"

type AuditChecklist struct {
	ACID             int       `json:"ac_id"`
	CheckName        string    `json:"check_name"`
	CheckDescription string    `json:"check_description"`
	CheckCommand     string    `json:"check_command"`
	ProfileAbility   string    `json:"profile_ability"`
	Rational         string    `json:"rational"`
	Impact           string    `json:"impact"`
	Verification     string    `json:"verification"`
	IG1              bool      `json:"IG1"`
	IG2              bool      `json:"IG2"`
	IG3              bool      `json:"IG3"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
