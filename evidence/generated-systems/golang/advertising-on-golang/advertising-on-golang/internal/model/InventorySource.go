package model

import (
    "gorm.io/gorm"
)

//==============================================================
// InventorySource Declaration
//==============================================================
type InventorySource struct {
    gorm.Model
     Name                                    string
    Domain                                    string
    PublisherId         *uint
    Publisher           *Publisher `gorm:"foreignKey:PublisherId"`
     AdSlots           []AdSlot `gorm:"foreignKey:AdSlotsFromInventorySourceId"`
     Deals           []Deal `gorm:"foreignKey:DealsFromInventorySourceId"`
    Channel                      ChannelType
    PrimaryFormat                      AdFormat

// parent associations as their child

}

