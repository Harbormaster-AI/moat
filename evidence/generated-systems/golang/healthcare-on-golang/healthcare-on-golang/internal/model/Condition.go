package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Condition Declaration
//==============================================================
type Condition struct {
    gorm.Model
     Code                                    string
    OnsetDate                                                            time.Time
    AbatementDate                                                            time.Time
    PatientId         *uint
    Patient           *Patient `gorm:"foreignKey:PatientId"`
    ClinicalStatus                      ConditionStatus
    VerificationStatus                      DiagnosisCertainty

// parent associations as their child

}

