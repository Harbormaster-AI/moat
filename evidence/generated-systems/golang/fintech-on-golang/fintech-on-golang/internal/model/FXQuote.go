package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// FXQuote Declaration
//==============================================================
type FXQuote struct {
    gorm.Model
     BaseCurrency                                    string
    QuoteCurrency                                    string
    Rate                                                            string
    QuotedAt                                                            time.Time
    ExpiresAt                                                            time.Time
    RequestedById         *uint
    RequestedBy           *Customer `gorm:"foreignKey:RequestedById"`
    PriceType                      FXPriceType

// parent associations as their child

}

