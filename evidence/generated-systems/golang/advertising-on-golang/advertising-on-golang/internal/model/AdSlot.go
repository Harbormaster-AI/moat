package model

import (
    "gorm.io/gorm"
)

//==============================================================
// AdSlot Declaration
//==============================================================
type AdSlot struct {
    gorm.Model
     SlotCode                                    string
    Width                                                            string
    Height                                                            string
    FloorPrice                                                            string
    InventorySourceId         *uint
    InventorySource           *InventorySource `gorm:"foreignKey:InventorySourceId"`
     Placements           []Placement `gorm:"foreignKey:PlacementsFromAdSlotId"`
     Rates           []Rate `gorm:"foreignKey:RatesFromAdSlotId"`
    Format                      AdFormat

// parent associations as their child

}

