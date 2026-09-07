package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// InvestmentPortfolio Declaration
//==============================================================
type InvestmentPortfolio struct {
    gorm.Model
     PortfolioCode                                    string
    BaseCurrency                                    string
    CreatedAt                                                            time.Time
    CustomerId         *uint
    Customer           *Customer `gorm:"foreignKey:CustomerId"`
     Accounts           []InvestmentAccount `gorm:"foreignKey:AccountsFromInvestmentPortfolioId"`
     Orders           []TradeOrder `gorm:"foreignKey:OrdersFromInvestmentPortfolioId"`
     Holdings           []Position `gorm:"foreignKey:HoldingsFromInvestmentPortfolioId"`
    Status                      PortfolioStatus

// parent associations as their child

}

