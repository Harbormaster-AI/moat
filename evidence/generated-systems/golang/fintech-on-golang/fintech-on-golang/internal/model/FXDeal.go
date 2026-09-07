package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// FXDeal Declaration
//==============================================================
type FXDeal struct {
    gorm.Model
     DealReference                                    string
    BaseCurrency                                    string
    QuoteCurrency                                    string
    Rate                                                            string
    Amount                                                            string
    SettlementDate                                                            time.Time
    QuoteId         *uint
    Quote           *FXQuote `gorm:"foreignKey:QuoteId"`
     PaymentOrders           []PaymentOrder `gorm:"foreignKey:PaymentOrdersFromFXDealId"`
    Status                      FXDealStatus

// parent associations as their child

}

