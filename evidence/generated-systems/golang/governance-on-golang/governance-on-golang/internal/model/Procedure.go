package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Procedure Declaration
//==============================================================
type Procedure struct {
    gorm.Model
     Title                                    string
    VersionLabel                                    string
    PolicyId         *uint
    Policy           *Policy `gorm:"foreignKey:PolicyId"`
     Controls           []Control `gorm:"foreignKey:ControlsFromProcedureId"`
    Status                      DocumentStatus

// parent associations as their child

}

