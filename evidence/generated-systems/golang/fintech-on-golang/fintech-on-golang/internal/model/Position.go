package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Position Declaration
//==============================================================
type Position struct {
    gorm.Model
     Quantity                                                            string
    AverageCost                                                            string
    MarketValue                                                            string
    PortfolioId         *uint
    Portfolio           *InvestmentPortfolio `gorm:"foreignKey:PortfolioId"`
    SecurityId         *uint
    Security           *Security `gorm:"foreignKey:SecurityId"`

// parent associations as their child

}

