package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// UnderwritingDecision Declaration
//==============================================================
type UnderwritingDecision struct {
    gorm.Model
     Notes                                    string
    DecisionDate                                                            time.Time
    QuoteId         *uint
    Quote           *Quote `gorm:"foreignKey:QuoteId"`
    UnderwriterId         *uint
    Underwriter           *Underwriter `gorm:"foreignKey:UnderwriterId"`
    Decision                      UnderwritingDecisionType

// parent associations as their child

}

