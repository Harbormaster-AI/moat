package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Placement Declaration
//==============================================================
type Placement struct {
    gorm.Model
     Name                                    string
    Flight                                                            string
    GoalImpressions                                                            string
    LineItemId         *uint
    LineItem           *LineItem `gorm:"foreignKey:LineItemId"`
    AdSlotId         *uint
    AdSlot           *AdSlot `gorm:"foreignKey:AdSlotId"`
    DealId         *uint
    Deal           *Deal `gorm:"foreignKey:DealId"`

// parent associations as their child

}

