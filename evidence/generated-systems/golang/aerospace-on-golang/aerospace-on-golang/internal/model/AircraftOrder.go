package model

import (
    "gorm.io/gorm"
)

//==============================================================
// AircraftOrder Declaration
//==============================================================
type AircraftOrder struct {
    gorm.Model
     OrderNumber                                    string
    TotalAmount                                                            string
    OperatorId         *uint
    Operator           *Operator `gorm:"foreignKey:OperatorId"`
    VariantId         *uint
    Variant           *AircraftVariant `gorm:"foreignKey:VariantId"`
    QuoteId         *uint
    Quote           *Quote `gorm:"foreignKey:QuoteId"`
    PurchaseAgreementId         *uint
    PurchaseAgreement           *PurchaseAgreement `gorm:"foreignKey:PurchaseAgreementId"`
    Status                      AircraftOrderStatus

// parent associations as their child

}

