package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type FieldFillVariant struct {
	bun.BaseModel     `bun:"FieldFill_variants,alias:FieldFill_variants"`
	ID                uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name              string        `json:"name"`
	FieldID        uuid.NullUUID `bun:"type:uuid" json:"FieldId"`
	Field          *Field     `bun:"rel:belongs-to" json:"Field"`
	Order             int           `bun:"item_order" json:"order"`
	Score             int           `json:"score"`
	ShowMoreFields bool          `json:"showMoreFields"`
	//RegisterPropertyOthers          RegisterPropertyOthers `bun:"rel:has-many" json:"registerPropertyOthers"`
	//RegisterPropertyOthersForDelete []uuid.UUID            `bun:"-" json:"registerPropertyOthersForDelete"`

	AggregatedValues map[string]float64 `bun:"-" json:"aggregatedValues"`
	// RegisterQueryPercentages ResearchQueryPercentages `bun:"-" `
}

type FieldFillVariants []*FieldFillVariant

func (item *FieldFillVariant) SetIDForChildren() {
	//if len(item.RegisterPropertyOthers) == 0 {
	//	return
	//}
	//for i := range item.RegisterPropertyOthers {
	//	item.RegisterPropertyOthers[i].FieldFillVariantID = item.ID
	//}
}

func (items FieldFillVariants) SetIDForChildren() {
	if len(items) == 0 {
		return
	}
	for i := range items {
		items[i].SetIDForChildren()
	}
}

//func (items FieldFillVariants) GetRegisterPropertyOthers() RegisterPropertyOthers {
//	itemsForGet := make(RegisterPropertyOthers, 0)
//	for i := range items {
//		itemsForGet = append(itemsForGet, items[i].RegisterPropertyOthers...)
//	}
//	return itemsForGet
//}

func (items FieldFillVariants) GetRegisterPropertyOthersForDelete() []uuid.UUID {
	itemsForGet := make([]uuid.UUID, 0)
	//for i := range items {
	//	itemsForGet = append(itemsForGet, items[i].RegisterPropertyOthersForDelete...)
	//}
	return itemsForGet
}

func (item *FieldFillVariant) writeXlsxAggregatedValues(key string) {
	_, ok := item.AggregatedValues[key]
	if ok {
		item.AggregatedValues[key]++
	} else {
		item.AggregatedValues[key] = 1
	}
}

//
// func (item *FieldFillVariant) GetAggregatedPercentage() {
// 	sum := float64(0)
// 	for k, v := range item.AggregatedValues {
// 		sum += v
// 		item.RegisterQueryPercentages = append(item.RegisterQueryPercentages, &ResearchQueryPercentage{k, v})
// 	}
// 	sort.Slice(item.RegisterQueryPercentages, func(i, j int) bool {
// 		return item.RegisterQueryPercentages[i].Value > item.RegisterQueryPercentages[j].Value
// 	})
// }
//
// func (items FieldFillVariants) Include(FieldFills FieldFills) string {
// 	exists := No
// 	for _, item := range items {
// 		if len(FieldFills) == 0 {
// 			break
// 		}
// 		for _, a := range FieldFills {
// 			for _, s := range a.SelectedFieldFillVariants {
// 				if s.FieldFillVariantID == item.ID {
// 					exists = Yes
// 					break
// 				}
// 			}
// 			if exists == Yes {
// 				break
// 			}
// 		}
// 	}
// 	return exists
// }
