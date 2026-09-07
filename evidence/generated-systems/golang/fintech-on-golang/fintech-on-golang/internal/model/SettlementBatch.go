package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// SettlementBatch Declaration
//==============================================================
type SettlementBatch struct {
    gorm.Model
     BatchId                                    string
    PeriodStart                                                            time.Time
    PeriodEnd                                                            time.Time
    TotalVolume                                                            string
    TotalCount                                                            string
    ProcessorId         *uint
    Processor           *PaymentProcessor `gorm:"foreignKey:ProcessorId"`
    MerchantId         *uint
    Merchant           *Merchant `gorm:"foreignKey:MerchantId"`
     Payouts           []Payout `gorm:"foreignKey:PayoutsFromSettlementBatchId"`
     Transactions           []Transaction `gorm:"foreignKey:TransactionsFromSettlementBatchId"`
    Status                      SettlementStatus

// parent associations as their child

}

