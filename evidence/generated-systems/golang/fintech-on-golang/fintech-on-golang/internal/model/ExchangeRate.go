package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// ExchangeRate Declaration
//==============================================================
type ExchangeRate struct {
    gorm.Model
     BaseCurrency                                    string
    QuoteCurrency                                    string
    Rate                                                            string
    AsOf                                                            time.Time
    Source                                    string
     UsedByQuotes           []FXQuote `gorm:"foreignKey:UsedByQuotesFromExchangeRateId"`

// parent associations as their child

}

