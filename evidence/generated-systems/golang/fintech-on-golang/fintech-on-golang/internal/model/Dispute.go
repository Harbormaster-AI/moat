package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Dispute Declaration
//==============================================================
type Dispute struct {
    gorm.Model
     DisputeReference                                    string
    OpenedAt                                                            time.Time
    ClosedAt                                                            time.Time
    TransactionId         *uint
    Transaction           *Transaction `gorm:"foreignKey:TransactionId"`
    CardId         *uint
    Card           *PaymentCard `gorm:"foreignKey:CardId"`
    MerchantId         *uint
    Merchant           *Merchant `gorm:"foreignKey:MerchantId"`
     Chargebacks           []Chargeback `gorm:"foreignKey:ChargebacksFromDisputeId"`
    Reason                      DisputeReason
    Status                      DisputeStatus

// parent associations as their child

}

