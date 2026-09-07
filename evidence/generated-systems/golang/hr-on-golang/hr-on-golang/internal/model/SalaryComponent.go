package model

import (
    "gorm.io/gorm"
)

//==============================================================
// SalaryComponent Declaration
//==============================================================
type SalaryComponent struct {
    gorm.Model
     Amount                                                            string
    Recurring                                    bool
    CompensationPackageId         *uint
    CompensationPackage           *CompensationPackage `gorm:"foreignKey:CompensationPackageId"`
    ComponentType                      SalaryComponentType

// parent associations as their child

}

