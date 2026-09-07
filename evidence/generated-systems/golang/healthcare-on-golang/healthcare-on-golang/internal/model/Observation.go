package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Observation Declaration
//==============================================================
type Observation struct {
    gorm.Model
     Code                                    string
    Value                                    string
    Unit                                    string
    EffectiveDateTime                                                            time.Time
    EncounterId         *uint
    Encounter           *Encounter `gorm:"foreignKey:EncounterId"`
    PatientId         *uint
    Patient           *Patient `gorm:"foreignKey:PatientId"`
    DeviceId         *uint
    Device           *MedicalDevice `gorm:"foreignKey:DeviceId"`
    LabResultId         *uint
    LabResult           *LabResult `gorm:"foreignKey:LabResultId"`
    Interpretation                      ObservationInterpretation

// parent associations as their child

}

