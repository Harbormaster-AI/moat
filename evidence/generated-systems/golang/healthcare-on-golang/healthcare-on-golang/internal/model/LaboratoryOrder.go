package model

import (
    "gorm.io/gorm"
)

//==============================================================
// LaboratoryOrder Declaration
//==============================================================
type LaboratoryOrder struct {
    gorm.Model
     TestCode                                    string
    FastingRequired                                    bool
    OrderId         *uint
    Order           *ClinicalOrder `gorm:"foreignKey:OrderId"`
    LaboratoryId         *uint
    Laboratory           *Laboratory `gorm:"foreignKey:LaboratoryId"`
     Results           []LabResult `gorm:"foreignKey:ResultsFromLaboratoryOrderId"`
    SpecimenType                      SpecimenType

// parent associations as their child

}

