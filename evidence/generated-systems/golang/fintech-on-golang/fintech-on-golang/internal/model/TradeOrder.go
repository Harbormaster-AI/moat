package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// TradeOrder Declaration
//==============================================================
type TradeOrder struct {
    gorm.Model
     OrderId                                    string
    Quantity                                                            string
    LimitPrice                                                            string
    PlacedAt                                                            time.Time
    PortfolioId         *uint
    Portfolio           *InvestmentPortfolio `gorm:"foreignKey:PortfolioId"`
    SecurityId         *uint
    Security           *Security `gorm:"foreignKey:SecurityId"`
     Trades           []Trade `gorm:"foreignKey:TradesFromTradeOrderId"`
    Side                      OrderSide
    Type                      OrderType
    Status                      OrderStatus
    TimeInForce                      TimeInForce

// parent associations as their child

}

