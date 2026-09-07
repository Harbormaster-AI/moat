package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Rate Declaration
//==============================================================
type Rate struct {
    gorm.Model
     UnitPrice                                                            string
    RateCardId         *uint
    RateCard           *RateCard `gorm:"foreignKey:RateCardId"`
    AdSlotId         *uint
    AdSlot           *AdSlot `gorm:"foreignKey:AdSlotId"`
    AdFormat                      AdFormat
    PricingModel                      PricingModel

// parent associations as their child

}

