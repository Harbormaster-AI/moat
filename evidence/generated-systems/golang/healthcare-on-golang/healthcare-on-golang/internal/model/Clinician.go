package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Clinician Declaration
//==============================================================
type Clinician struct {
    gorm.Model
     FirstName                                    string
    LastName                                    string
    LicenseNumber                                    string
     CareTeams           []CareTeam `gorm:"foreignKey:CareTeamsFromClinicianId"`
     Appointments           []Appointment `gorm:"foreignKey:AppointmentsFromClinicianId"`
     Encounters           []Encounter `gorm:"foreignKey:EncountersFromClinicianId"`
     Procedures           []Procedure `gorm:"foreignKey:ProceduresFromClinicianId"`
     ImagingReports           []ImagingReport `gorm:"foreignKey:ImagingReportsFromClinicianId"`
    ClinicianType                      ClinicianType
    Specialty                      ClinicianSpecialty

// parent associations as their child

}

