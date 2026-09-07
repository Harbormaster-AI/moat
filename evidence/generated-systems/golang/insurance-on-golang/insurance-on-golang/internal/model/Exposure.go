package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Exposure Declaration
//==============================================================
type Exposure struct {
    gorm.Model
     ClaimId         *uint
    Claim           *Claim `gorm:"foreignKey:ClaimId"`
    PolicyCoverageId         *uint
    PolicyCoverage           *PolicyCoverage `gorm:"foreignKey:PolicyCoverageId"`
    InsuredObjectId         *uint
    InsuredObject           *InsuredObject `gorm:"foreignKey:InsuredObjectId"`
     Reserves           []ClaimReserve `gorm:"foreignKey:ReservesFromExposureId"`
     Payments           []ClaimPayment `gorm:"foreignKey:PaymentsFromExposureId"`
    ExposureType                      ExposureType
    Status                      ExposureStatus

// parent associations as their child

}

