package models

import (
	"time"

	"github.com/uptrace/bun"

	"github.com/google/uuid"
)

type UserResearch struct {
	bun.BaseModel `bun:"users_researches,alias:users_researches"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `

	Research   *Research     `bun:"rel:belongs-to" json:"research"`
	ResearchID uuid.NullUUID `bun:"type:uuid" json:"researchId"`

	ResearchResults ResearchResults `bun:"rel:has-many" json:"researchResults"`

	User   *User         `bun:"rel:belongs-to" json:"user"`
	UserID uuid.NullUUID `bun:"type:uuid" json:"userID"`
	Num    string        `json:"num"`

	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"createdAt"`
}

type UsersResearches []*UserResearch

type UsersResearchesWithCount struct {
	UsersResearches UsersResearches `json:"items"`
	Count           int             `json:"count"`
}
