package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// SubrogationRecovery Declaration
//==============================================================
type SubrogationRecovery struct {
    gorm.Model
     RecoveryReference                                    string
    Amount                                                            string
    RecoveryDate                                                            time.Time
    ClaimId         *uint
    Claim           *Claim `gorm:"foreignKey:ClaimId"`
    ExposureId         *uint
    Exposure           *Exposure `gorm:"foreignKey:ExposureId"`
    CounterpartyId         *uint
    Counterparty           *ThirdParty `gorm:"foreignKey:CounterpartyId"`
    Status                      SubrogationStatus

// parent associations as their child

}

