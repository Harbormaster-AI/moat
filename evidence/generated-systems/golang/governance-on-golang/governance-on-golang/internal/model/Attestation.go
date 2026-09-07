package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Attestation Declaration
//==============================================================
type Attestation struct {
    gorm.Model
     Statement                                    string
    Attestor                                    string
    DateSigned                                                            time.Time
    ControlId         *uint
    Control           *Control `gorm:"foreignKey:ControlId"`
    PolicyId         *uint
    Policy           *Policy `gorm:"foreignKey:PolicyId"`
    ComplianceProgramId         *uint
    ComplianceProgram           *ComplianceProgram `gorm:"foreignKey:ComplianceProgramId"`
    Result                      AttestationResult

// parent associations as their child

}

