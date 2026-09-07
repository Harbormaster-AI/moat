package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Pharmacy Declaration
//==============================================================
type Pharmacy struct {
    gorm.Model
     Name                                    string
    FacilityId         *uint
    Facility           *Facility `gorm:"foreignKey:FacilityId"`
     MedicationDispenses           []MedicationDispense `gorm:"foreignKey:MedicationDispensesFromPharmacyId"`
     MedicationOrders           []MedicationOrder `gorm:"foreignKey:MedicationOrdersFromPharmacyId"`

// parent associations as their child

}

