package model

import (
    "gorm.io/gorm"
)

//==============================================================
// AppliedFee Declaration
//==============================================================
type AppliedFee struct {
    gorm.Model
     Amount                                                            string
    Description                                    string
    PaymentOrderId         *uint
    PaymentOrder           *PaymentOrder `gorm:"foreignKey:PaymentOrderId"`
    TransactionId         *uint
    Transaction           *Transaction `gorm:"foreignKey:TransactionId"`
    FeeType                      FeeType

// parent associations as their child

}

