package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Screening Declaration
//==============================================================
type Screening struct {
    gorm.Model
     Score                                                            string
    ScreenedAt                                                            time.Time
    KycProfileId         *uint
    KycProfile           *KYCProfile `gorm:"foreignKey:KycProfileId"`
     Alerts           []ComplianceAlert `gorm:"foreignKey:AlertsFromScreeningId"`
    ScreeningType                      ScreeningType
    Status                      ScreeningStatus

// parent associations as their child

}

