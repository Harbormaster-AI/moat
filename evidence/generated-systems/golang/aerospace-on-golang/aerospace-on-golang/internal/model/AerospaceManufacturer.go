package model

import (
    "gorm.io/gorm"
)

//==============================================================
// AerospaceManufacturer Declaration
//==============================================================
type AerospaceManufacturer struct {
    gorm.Model
     Name                                    string
    LegalName                                    string
    HeadquartersCountry                                    string
    Website                                    string
     Programs           []AircraftProgram `gorm:"foreignKey:ProgramsFromAerospaceManufacturerId"`
     Plants           []Plant `gorm:"foreignKey:PlantsFromAerospaceManufacturerId"`
     Suppliers           []Supplier `gorm:"foreignKey:SuppliersFromAerospaceManufacturerId"`
     ProductionCertificates           []ProductionCertificate `gorm:"foreignKey:ProductionCertificatesFromAerospaceManufacturerId"`

// parent associations as their child

}

