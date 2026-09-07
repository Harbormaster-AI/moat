package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Discharge Declaration
//==============================================================
type Discharge struct {
    gorm.Model
     DischargeDateTime                                                            time.Time
    EncounterId         *uint
    Encounter           *Encounter `gorm:"foreignKey:EncounterId"`
    Disposition                      DischargeDisposition

// parent associations as their child

}

