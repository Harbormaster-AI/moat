package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// PurchaseAgreement Declaration
//==============================================================
type PurchaseAgreement struct {
    gorm.Model
     AgreementNumber                                    string
    EffectiveDate                                                            time.Time
    AircraftOrderId         *uint
    AircraftOrder           *AircraftOrder `gorm:"foreignKey:AircraftOrderId"`

// parent associations as their child

}

