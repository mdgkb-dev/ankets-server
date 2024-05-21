package models

import (
	"github.com/uptrace/bun"

	"github.com/google/uuid"
)

type UserResearch struct {
	bun.BaseModel `bun:"users,alias:users"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `

	Research   *Research     `bun:"rel:belongs-to" json:"research"`
	ResearchID uuid.NullUUID `bun:"type:uuid" json:"researchId"`

	Num string `json:"num"`
}

type UsersResearches []*UserResearch

type UsersResearchesWithCount struct {
	UsersResearches UsersResearches `json:"items"`
	Count           int             `json:"count"`
}
