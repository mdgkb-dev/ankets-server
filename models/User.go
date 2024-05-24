package models

import (
	"github.com/pro-assistance/pro-assister/middleware"
	baseModels "github.com/pro-assistance/pro-assister/models"

	"github.com/uptrace/bun"

	"github.com/google/uuid"
)

type User struct {
	bun.BaseModel `bun:"users,alias:users"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Position      string        `json:"position"`
	Division      string        `json:"division"`
	Role          string        `json:"role"`

	Human   *Human        `bun:"rel:belongs-to" json:"human"`
	HumanID uuid.NullUUID `bun:"type:uuid" json:"humanId"`

	UserAccountID uuid.NullUUID           `bun:"type:uuid" json:"userAccountId"`
	UserAccount   *baseModels.UserAccount `bun:"rel:belongs-to" json:"userAccount"`

	Inn                  string `json:"inn"`
	Snils                string `json:"snils"`
	PassportNum          string `json:"passportNum"`
	PassportSeria        string `json:"passportSeria"`
	PassportDivision     string `json:"passportDivision"`
	PassportDivisionCode string `json:"passportDivisionCode"`
	PassportCitzenship   string `json:"passportCitzenship"`
}

type Users []*User

type UsersWithCount struct {
	Users Users `json:"items"`
	Count int   `json:"count"`
}

func (item *User) SetJWTClaimsMap(claims map[string]interface{}) {
	claims[middleware.ClaimUserID.String()] = item.ID.UUID
}
