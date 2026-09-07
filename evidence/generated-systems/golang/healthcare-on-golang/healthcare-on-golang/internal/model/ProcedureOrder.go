package model

import (
    "gorm.io/gorm"
)

//==============================================================
// ProcedureOrder Declaration
//==============================================================
type ProcedureOrder struct {
    gorm.Model
     ProcedureCode                                    string
    ConsentObtained                                    bool
    OrderId         *uint
    Order           *ClinicalOrder `gorm:"foreignKey:OrderId"`
    FacilityId         *uint
    Facility           *Facility `gorm:"foreignKey:FacilityId"`
    ProcedureId         *uint
    Procedure           *Procedure `gorm:"foreignKey:ProcedureId"`
    AnesthesiaType                      AnesthesiaType

// parent associations as their child

}

