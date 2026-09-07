package model

import (
    "gorm.io/gorm"
)

//==============================================================
// HealthSystem Declaration
//==============================================================
type HealthSystem struct {
    gorm.Model
     Name                                    string
    LegalName                                    string
    HeadquartersCountry                                    string
    Website                                    string
     Facilities           []Facility `gorm:"foreignKey:FacilitiesFromHealthSystemId"`
     Suppliers           []MedicalSupplier `gorm:"foreignKey:SuppliersFromHealthSystemId"`

// parent associations as their child

}

