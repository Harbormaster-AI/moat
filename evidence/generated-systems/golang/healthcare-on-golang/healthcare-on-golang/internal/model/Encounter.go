package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Encounter Declaration
//==============================================================
type Encounter struct {
    gorm.Model
     EncounterNumber                                    string
    StartDateTime                                                            time.Time
    EndDateTime                                                            time.Time
    PatientId         *uint
    Patient           *Patient `gorm:"foreignKey:PatientId"`
    ClinicianId         *uint
    Clinician           *Clinician `gorm:"foreignKey:ClinicianId"`
    FacilityId         *uint
    Facility           *Facility `gorm:"foreignKey:FacilityId"`
    AppointmentId         *uint
    Appointment           *Appointment `gorm:"foreignKey:AppointmentId"`
     Diagnoses           []Diagnosis `gorm:"foreignKey:DiagnosesFromEncounterId"`
     Procedures           []Procedure `gorm:"foreignKey:ProceduresFromEncounterId"`
     Observations           []Observation `gorm:"foreignKey:ObservationsFromEncounterId"`
     Orders           []ClinicalOrder `gorm:"foreignKey:OrdersFromEncounterId"`
    AdmissionId         *uint
    Admission           *Admission `gorm:"foreignKey:AdmissionId"`
    DischargeId         *uint
    Discharge           *Discharge `gorm:"foreignKey:DischargeId"`
    Status                      EncounterStatus
    EncounterType                      EncounterType

// parent associations as their child

}

