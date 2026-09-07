package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Security Declaration
//==============================================================
type Security struct {
    gorm.Model
     Symbol                                    string
    Isin                                    string
    Cusip                                    string
    Currency                                    string
     Positions           []Position `gorm:"foreignKey:PositionsFromSecurityId"`
     Trades           []Trade `gorm:"foreignKey:TradesFromSecurityId"`
     Orders           []TradeOrder `gorm:"foreignKey:OrdersFromSecurityId"`
    SecurityType                      SecurityType

// parent associations as their child

}

