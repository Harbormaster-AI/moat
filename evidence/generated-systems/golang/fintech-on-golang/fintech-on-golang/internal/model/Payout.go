package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Payout Declaration
//==============================================================
type Payout struct {
    gorm.Model
     PayoutReference                                    string
    Amount                                                            string
    Currency                                    string
    ScheduledDate                                                            time.Time
    PaidDate                                                            time.Time
    MerchantId         *uint
    Merchant           *Merchant `gorm:"foreignKey:MerchantId"`
    SettlementBatchId         *uint
    SettlementBatch           *SettlementBatch `gorm:"foreignKey:SettlementBatchId"`
    DestinationAccountId         *uint
    DestinationAccount           *Account `gorm:"foreignKey:DestinationAccountId"`
    Status                      PayoutStatus

// parent associations as their child

}

