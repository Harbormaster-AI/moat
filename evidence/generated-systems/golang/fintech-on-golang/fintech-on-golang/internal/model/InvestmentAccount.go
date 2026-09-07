package model

import (
    "gorm.io/gorm"
)

//==============================================================
// InvestmentAccount Declaration
//==============================================================
type InvestmentAccount struct {
    gorm.Model
     AccountNumber                                                            string
    BaseCurrency                                    string
    Balance                                                            string
    PortfolioId         *uint
    Portfolio           *InvestmentPortfolio `gorm:"foreignKey:PortfolioId"`
     Trades           []Trade `gorm:"foreignKey:TradesFromInvestmentAccountId"`
     Orders           []TradeOrder `gorm:"foreignKey:OrdersFromInvestmentAccountId"`
    AccountType                      InvestmentAccountType
    Status                      AccountStatus

// parent associations as their child

}

