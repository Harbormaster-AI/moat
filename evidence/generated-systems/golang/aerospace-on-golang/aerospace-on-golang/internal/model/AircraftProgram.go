package model

import (
    "gorm.io/gorm"
)

//==============================================================
// AircraftProgram Declaration
//==============================================================
type AircraftProgram struct {
    gorm.Model
     Name                                    string
    ProgramCode                                    string
    EntryIntoServiceYear                                                            string
    ManufacturerId         *uint
    Manufacturer           *AerospaceManufacturer `gorm:"foreignKey:ManufacturerId"`
     AircraftFamilies           []AircraftFamily `gorm:"foreignKey:AircraftFamiliesFromAircraftProgramId"`
    TypeCertificateId         *uint
    TypeCertificate           *TypeCertificate `gorm:"foreignKey:TypeCertificateId"`
     KeySuppliers           []Supplier `gorm:"foreignKey:KeySuppliersFromAircraftProgramId"`
    Status                      ProgramStatus

// parent associations as their child

}

