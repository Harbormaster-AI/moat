package model

import (
    "gorm.io/gorm"
)

//==============================================================
// TypeCertificate Declaration
//==============================================================
type TypeCertificate struct {
    gorm.Model
     CertificateNumber                                    string
    Authority                                    string
    ProgramId         *uint
    Program           *AircraftProgram `gorm:"foreignKey:ProgramId"`

// parent associations as their child

}

