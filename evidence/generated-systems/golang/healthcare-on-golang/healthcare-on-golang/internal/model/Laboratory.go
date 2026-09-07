package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Laboratory Declaration
//==============================================================
type Laboratory struct {
    gorm.Model
     Name                                    string
    CliaNumber                                    string
    FacilityId         *uint
    Facility           *Facility `gorm:"foreignKey:FacilityId"`
     LaboratoryOrders           []LaboratoryOrder `gorm:"foreignKey:LaboratoryOrdersFromLaboratoryId"`
     LabResults           []LabResult `gorm:"foreignKey:LabResultsFromLaboratoryId"`

// parent associations as their child

}

