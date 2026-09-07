package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// PaymentOrder Declaration
//==============================================================
type PaymentOrder struct {
    gorm.Model
     OrderReference                                    string
    RequestedExecutionDate                                                            time.Time
    Purpose                                    string
    SourceAccountId         *uint
    SourceAccount           *Account `gorm:"foreignKey:SourceAccountId"`
    DestinationAccountId         *uint
    DestinationAccount           *Account `gorm:"foreignKey:DestinationAccountId"`
    BeneficiaryId         *uint
    Beneficiary           *Beneficiary `gorm:"foreignKey:BeneficiaryId"`
     Transactions           []Transaction `gorm:"foreignKey:TransactionsFromPaymentOrderId"`
    FxDealId         *uint
    FxDeal           *FXDeal `gorm:"foreignKey:FxDealId"`
     Fees           []AppliedFee `gorm:"foreignKey:FeesFromPaymentOrderId"`
    PaymentMethod                      PaymentMethod
    Status                      PaymentOrderStatus
    Priority                      PaymentPriority

// parent associations as their child

}

