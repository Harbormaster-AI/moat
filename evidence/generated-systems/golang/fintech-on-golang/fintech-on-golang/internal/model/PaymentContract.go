package model

import (
    "gorm.io/gorm"
)

//==============================================================
// PaymentContract Declaration
//==============================================================
type PaymentContract struct {
    gorm.Model
     ContractNumber                                    string
    PricingPlanCode                                    string
    MerchantId         *uint
    Merchant           *Merchant `gorm:"foreignKey:MerchantId"`
    AcquirerId         *uint
    Acquirer           *PaymentProcessor `gorm:"foreignKey:AcquirerId"`
    Status                      ContractStatus

// parent associations as their child

}

