package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Admission Declaration
//==============================================================
type Admission struct {
    gorm.Model
     AdmitDateTime                                                            time.Time
    Bed                                    string
    EncounterId         *uint
    Encounter           *Encounter `gorm:"foreignKey:EncounterId"`
    FacilityId         *uint
    Facility           *Facility `gorm:"foreignKey:FacilityId"`
    AdmissionType                      AdmissionType

// parent associations as their child

}

