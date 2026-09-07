package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Appointment Declaration
//==============================================================
type Appointment struct {
    gorm.Model
     AppointmentDate                                                            time.Time
    Reason                                    string
    PatientId         *uint
    Patient           *Patient `gorm:"foreignKey:PatientId"`
    ClinicianId         *uint
    Clinician           *Clinician `gorm:"foreignKey:ClinicianId"`
    FacilityId         *uint
    Facility           *Facility `gorm:"foreignKey:FacilityId"`
    EncounterId         *uint
    Encounter           *Encounter `gorm:"foreignKey:EncounterId"`
    Status                      AppointmentStatus
    Priority                      Priority

// parent associations as their child

}

