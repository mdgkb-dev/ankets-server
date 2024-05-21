package models

import (
	"github.com/uptrace/bun"

	"github.com/google/uuid"
)

type Human struct {
	bun.BaseModel `bun:"humans,alias:humans"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string        `json:"name"`
	Surname       string        `json:"surname"`
	Patronymic    string        `json:"patronymic"`

	// Photo   *baseModels.FileInfo `bun:"rel:belongs-to" json:"photo"`
	// PhotoID uuid.NullUUID        `bun:"type:uuid" json:"photoId"`
	//
	// PhotoMini   *baseModels.FileInfo `bun:"rel:belongs-to" json:"photoMini"`
	// PhotoMiniID uuid.NullUUID        `bun:"type:uuid" json:"photoMiniId"`
}

type Humans []*Human

type HumansWithCount struct {
	Humans Humans `json:"items"`
	Count  int    `json:"count"`
}
