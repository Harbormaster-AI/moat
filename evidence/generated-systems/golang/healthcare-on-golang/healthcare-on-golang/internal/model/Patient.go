package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Patient Declaration
//==============================================================
type Patient struct {
    gorm.Model
     FirstName                                    string
    LastName                                    string
    Mrn                                                            string
    DateOfBirth                                                            time.Time
    Address                                                            string
    PrimaryLanguage                                    string
     Appointments           []Appointment `gorm:"foreignKey:AppointmentsFromPatientId"`
     Encounters           []Encounter `gorm:"foreignKey:EncountersFromPatientId"`
     CarePlans           []CarePlan `gorm:"foreignKey:CarePlansFromPatientId"`
     Allergies           []Allergy `gorm:"foreignKey:AllergiesFromPatientId"`
     Conditions           []Condition `gorm:"foreignKey:ConditionsFromPatientId"`
     MedicationOrders           []MedicationOrder `gorm:"foreignKey:MedicationOrdersFromPatientId"`
     LabOrders           []LaboratoryOrder `gorm:"foreignKey:LabOrdersFromPatientId"`
     ImagingOrders           []ImagingOrder `gorm:"foreignKey:ImagingOrdersFromPatientId"`
     Coverages           []Coverage `gorm:"foreignKey:CoveragesFromPatientId"`
     Claims           []Claim `gorm:"foreignKey:ClaimsFromPatientId"`
     Devices           []MedicalDevice `gorm:"foreignKey:DevicesFromPatientId"`
     Observations           []Observation `gorm:"foreignKey:ObservationsFromPatientId"`
    SexAtBirth                      AdministrativeSex
    BloodType                      BloodType

// parent associations as their child

}

