package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Quote Declaration
//==============================================================
type Quote struct {
    gorm.Model
     QuoteNumber                                    string
    TotalPremium                                                            string
    RatingDate                                                            time.Time
    AsBound                                    bool
    ApplicationId         *uint
    Application           *Application `gorm:"foreignKey:ApplicationId"`
     UnderwritingDecisions           []UnderwritingDecision `gorm:"foreignKey:UnderwritingDecisionsFromQuoteId"`
    PolicyId         *uint
    Policy           *Policy `gorm:"foreignKey:PolicyId"`

// parent associations as their child

}

