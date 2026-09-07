package model

import (
    "gorm.io/gorm"
)

//==============================================================
// MedicationOrder Declaration
//==============================================================
type MedicationOrder struct {
    gorm.Model
     MedicationCode                                    string
    Dose                                                            string
    Frequency                                    string
    Duration                                    string
    OrderId         *uint
    Order           *ClinicalOrder `gorm:"foreignKey:OrderId"`
    PharmacyId         *uint
    Pharmacy           *Pharmacy `gorm:"foreignKey:PharmacyId"`
     Dispenses           []MedicationDispense `gorm:"foreignKey:DispensesFromMedicationOrderId"`
    Route                      RouteOfAdministration

// parent associations as their child

}

