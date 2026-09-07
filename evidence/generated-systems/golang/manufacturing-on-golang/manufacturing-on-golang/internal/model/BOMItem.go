package model

import (
    "gorm.io/gorm"
)

//==============================================================
// BOMItem Declaration
//==============================================================
type BOMItem struct {
    gorm.Model
     LineNumber                                                            string
    Quantity                                                            string
    ScrapPercent                                                            string
    BomId         *uint
    Bom           *BOM `gorm:"foreignKey:BomId"`
    ComponentId         *uint
    Component           *Item `gorm:"foreignKey:ComponentId"`

// parent associations as their child

}

