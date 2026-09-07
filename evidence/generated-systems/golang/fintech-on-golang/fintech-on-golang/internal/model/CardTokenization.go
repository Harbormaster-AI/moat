package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// CardTokenization Declaration
//==============================================================
type CardTokenization struct {
    gorm.Model
     TokenReference                                    string
    CreatedAt                                                            time.Time
    CardId         *uint
    Card           *PaymentCard `gorm:"foreignKey:CardId"`
    WalletProvider                      WalletProvider
    Status                      TokenizationStatus

// parent associations as their child

}

