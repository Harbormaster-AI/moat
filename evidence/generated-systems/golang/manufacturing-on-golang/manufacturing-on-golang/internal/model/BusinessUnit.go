package model

import (
    "gorm.io/gorm"
)

//==============================================================
// BusinessUnit Declaration
//==============================================================
type BusinessUnit struct {
    gorm.Model
     Name                                    string
    Code                                    string
    EnterpriseId         *uint
    Enterprise           *Enterprise `gorm:"foreignKey:EnterpriseId"`
     Items           []Item `gorm:"foreignKey:ItemsFromBusinessUnitId"`
     Plants           []Plant `gorm:"foreignKey:PlantsFromBusinessUnitId"`
    Category                      BusinessUnitCategory

// parent associations as their child

}

