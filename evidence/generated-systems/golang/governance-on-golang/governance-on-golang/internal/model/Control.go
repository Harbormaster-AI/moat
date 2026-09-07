package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Control Declaration
//==============================================================
type Control struct {
    gorm.Model
     Name                                    string
    Objective                                    string
    OwnerDepartment                                    string
    PolicyId         *uint
    Policy           *Policy `gorm:"foreignKey:PolicyId"`
     ControlTests           []ControlTest_ `gorm:"foreignKey:ControlTestsFromControlId"`
     Evidence           []Evidence `gorm:"foreignKey:EvidenceFromControlId"`
     Risks           []Risk `gorm:"foreignKey:RisksFromControlId"`
     Obligations           []Obligation `gorm:"foreignKey:ObligationsFromControlId"`
     Procedures           []Procedure `gorm:"foreignKey:ProceduresFromControlId"`
     Issues           []Issue `gorm:"foreignKey:IssuesFromControlId"`
    ControlType                      ControlType
    Frequency                      ControlFrequency
    Status                      ControlStatus

// parent associations as their child

}

