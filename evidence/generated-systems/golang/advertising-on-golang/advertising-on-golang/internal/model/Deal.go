package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Deal Declaration
//==============================================================
type Deal struct {
    gorm.Model
     FloorPrice                                                            string
    PublisherId         *uint
    Publisher           *Publisher `gorm:"foreignKey:PublisherId"`
     InventorySources           []InventorySource `gorm:"foreignKey:InventorySourcesFromDealId"`
     Placements           []Placement `gorm:"foreignKey:PlacementsFromDealId"`
    DealType                      DealType

// parent associations as their child

}

