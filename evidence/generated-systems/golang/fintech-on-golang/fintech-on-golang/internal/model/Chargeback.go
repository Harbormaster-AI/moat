package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Chargeback Declaration
//==============================================================
type Chargeback struct {
    gorm.Model
     ChargebackReference                                    string
    Amount                                                            string
    PostedAt                                                            time.Time
    DisputeId         *uint
    Dispute           *Dispute `gorm:"foreignKey:DisputeId"`
    TransactionId         *uint
    Transaction           *Transaction `gorm:"foreignKey:TransactionId"`
    Stage                      ChargebackStage
    Status                      ChargebackStatus

// parent associations as their child

}

