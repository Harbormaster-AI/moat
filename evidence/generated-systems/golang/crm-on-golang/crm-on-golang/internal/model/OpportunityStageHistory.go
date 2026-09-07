package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// OpportunityStageHistory Declaration
//==============================================================
type OpportunityStageHistory struct {
    gorm.Model
     ChangedAt                                                            time.Time
    Comment                                    string
    OpportunityId         *uint
    Opportunity           *Opportunity `gorm:"foreignKey:OpportunityId"`
    ChangedById         *uint
    ChangedBy           *User `gorm:"foreignKey:ChangedById"`
    FromStage                      OpportunityStage
    ToStage                      OpportunityStage

// parent associations as their child

}

