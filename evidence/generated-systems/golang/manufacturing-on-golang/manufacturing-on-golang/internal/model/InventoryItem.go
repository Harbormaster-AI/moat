package model

import (
    "gorm.io/gorm"
)

//==============================================================
// InventoryItem Declaration
//==============================================================
type InventoryItem struct {
    gorm.Model
     QuantityOnHand                                                            string
    QuantityReserved                                                            string
    LotNumber                                                            string
    SerialNumber                                                            string
    ItemId         *uint
    Item           *Item `gorm:"foreignKey:ItemId"`
    LocationId         *uint
    Location           *Location `gorm:"foreignKey:LocationId"`

// parent associations as their child

}

