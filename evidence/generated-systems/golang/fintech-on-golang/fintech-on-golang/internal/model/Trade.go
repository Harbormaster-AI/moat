package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Trade Declaration
//==============================================================
type Trade struct {
    gorm.Model
     ExecutedAt                                                            time.Time
    Quantity                                                            string
    Price                                                            string
    Fees                                                            string
    SettlementDate                                                            time.Time
    OrderId         *uint
    Order           *TradeOrder `gorm:"foreignKey:OrderId"`
    SecurityId         *uint
    Security           *Security `gorm:"foreignKey:SecurityId"`
    InvestmentAccountId         *uint
    InvestmentAccount           *InvestmentAccount `gorm:"foreignKey:InvestmentAccountId"`

// parent associations as their child

}

