package model

import (
    "gorm.io/gorm"
)

//==============================================================
// WorkCenter Declaration
//==============================================================
type WorkCenter struct {
    gorm.Model
     Name                                    string
    Capability                                    string
    ProductionLineId         *uint
    ProductionLine           *ProductionLine `gorm:"foreignKey:ProductionLineId"`

// parent associations as their child

}

