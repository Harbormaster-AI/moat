package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// BOM Declaration
//==============================================================
type BOM struct {
    gorm.Model
     BomNumber                                    string
    Revision                                    string
    EffectivityStart                                                            time.Time
    EffectivityEnd                                                            time.Time
    ParentItemId         *uint
    ParentItem           *Item `gorm:"foreignKey:ParentItemId"`
     BomItems           []BOMItem `gorm:"foreignKey:BomItemsFromBOMId"`
    Status                      BOMStatus

// parent associations as their child

}

