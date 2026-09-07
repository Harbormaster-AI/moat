package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// ClaimPayment Declaration
//==============================================================
type ClaimPayment struct {
    gorm.Model
     PaymentNumber                                    string
    Amount                                                            string
    PaymentDate                                                            time.Time
    ClaimId         *uint
    Claim           *Claim `gorm:"foreignKey:ClaimId"`
    ExposureId         *uint
    Exposure           *Exposure `gorm:"foreignKey:ExposureId"`
    BeneficiaryId         *uint
    Beneficiary           *Beneficiary `gorm:"foreignKey:BeneficiaryId"`
    ServiceProviderId         *uint
    ServiceProvider           *ServiceProvider `gorm:"foreignKey:ServiceProviderId"`
    CustomerId         *uint
    Customer           *Customer `gorm:"foreignKey:CustomerId"`
    PayeeType                      PayeeType
    Method                      PaymentMethod
    Status                      PaymentStatus

// parent associations as their child

}

