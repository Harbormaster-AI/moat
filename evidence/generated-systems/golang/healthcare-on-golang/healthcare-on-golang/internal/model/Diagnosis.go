package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Diagnosis Declaration
//==============================================================
type Diagnosis struct {
    gorm.Model
     Code                                    string
    Description                                    string
    OnsetDate                                                            time.Time
    EncounterId         *uint
    Encounter           *Encounter `gorm:"foreignKey:EncounterId"`
    PatientId         *uint
    Patient           *Patient `gorm:"foreignKey:PatientId"`
    Certainty                      DiagnosisCertainty

// parent associations as their child

}

