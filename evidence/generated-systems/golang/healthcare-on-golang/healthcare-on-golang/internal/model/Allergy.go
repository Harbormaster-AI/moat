package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Allergy Declaration
//==============================================================
type Allergy struct {
    gorm.Model
     Substance                                    string
    Reaction                                    string
    PatientId         *uint
    Patient           *Patient `gorm:"foreignKey:PatientId"`
    Severity                      AllergySeverity
    Status                      AllergyStatus

// parent associations as their child

}

