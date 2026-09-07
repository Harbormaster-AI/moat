package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// MedicationDispense Declaration
//==============================================================
type MedicationDispense struct {
    gorm.Model
     DispenseNumber                                    string
    Quantity                                                            string
    WhenPrepared                                                            time.Time
    MedicationOrderId         *uint
    MedicationOrder           *MedicationOrder `gorm:"foreignKey:MedicationOrderId"`
    PharmacyId         *uint
    Pharmacy           *Pharmacy `gorm:"foreignKey:PharmacyId"`
    PatientId         *uint
    Patient           *Patient `gorm:"foreignKey:PatientId"`
    Status                      DispenseStatus

// parent associations as their child

}

