package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Agreement Declaration
//==============================================================
type Agreement struct {
    gorm.Model
     AgreementNumber                                    string
    EffectiveDate                                                            time.Time
    CustomerId         *uint
    Customer           *Customer `gorm:"foreignKey:CustomerId"`
    ProductOfferingId         *uint
    ProductOffering           *ProductOffering `gorm:"foreignKey:ProductOfferingId"`
    AgreementType                      AgreementType
    Status                      AgreementStatus

// parent associations as their child

}

