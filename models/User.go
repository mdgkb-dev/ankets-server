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

	UserAccountID uuid.NullUUID           `bun:"type:uuid" json:"userAccountId"`
	UserAccount   *baseModels.UserAccount `bun:"rel:belongs-to" json:"userAccount"`
}

type Users []*User

type UsersWithCount struct {
	Users Users `json:"items"`
	Count int   `json:"count"`
}

func (item *User) SetJWTClaimsMap(claims map[string]interface{}) {
	claims[middleware.ClaimUserID.String()] = item.ID.UUID
}
