package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Wallet Declaration
//==============================================================
type Wallet struct {
    gorm.Model
     Currency                                    string
    Balance                                                            string
    CustomerId         *uint
    Customer           *Customer `gorm:"foreignKey:CustomerId"`
     Transactions           []Transaction `gorm:"foreignKey:TransactionsFromWalletId"`
    Status                      WalletStatus

// parent associations as their child

}

