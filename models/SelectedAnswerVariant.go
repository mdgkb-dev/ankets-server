package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type SelectedFieldFillVariant struct {
	bun.BaseModel `bun:"selected_FieldFill_variants,alias:selected_FieldFill_variants"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	FieldFill        *FieldFill       `bun:"rel:belongs-to" json:"FieldFill"`
	FieldFillID      uuid.NullUUID `bun:"type:uuid" json:"FieldFillId"`

	FieldFillVariant   *FieldFillVariant `bun:"rel:belongs-to" json:"FieldFillVariant"`
	FieldFillVariantID uuid.NullUUID  `bun:"type:uuid" json:"FieldFillVariantId"`
	PatientID       uuid.NullUUID  `bun:"type:uuid" json:"patientId"`
}

type SelectedFieldFillVariants []*SelectedFieldFillVariant

func (item *SelectedFieldFillVariant) SetIDForChildren() {
	//if len(item.RegisterPropertyOthers) == 0 {
	//	return
	//}
	//for i := range item.RegisterPropertyOthers {
	//	item.RegisterPropertyOthers[i].FieldFillVariantID = item.ID
	//}
}

func (items SelectedFieldFillVariants) SetIDForChildren() {
	if len(items) == 0 {
		return
	}
	for i := range items {
		items[i].SetIDForChildren()
	}
}
