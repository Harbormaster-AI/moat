package model

import (
    "gorm.io/gorm"
)

//==============================================================
// ProductionCertificate Declaration
//==============================================================
type ProductionCertificate struct {
    gorm.Model
     CertificateNumber                                    string
    Authority                                    string
    ManufacturerId         *uint
    Manufacturer           *AerospaceManufacturer `gorm:"foreignKey:ManufacturerId"`

// parent associations as their child

}

