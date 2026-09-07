package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Procedure Declaration
//==============================================================
type Procedure struct {
    gorm.Model
     ProcedureCode                                    string
    StartDateTime                                                            time.Time
    EndDateTime                                                            time.Time
    EncounterId         *uint
    Encounter           *Encounter `gorm:"foreignKey:EncounterId"`
    PerformerId         *uint
    Performer           *Clinician `gorm:"foreignKey:PerformerId"`
    ProcedureOrderId         *uint
    ProcedureOrder           *ProcedureOrder `gorm:"foreignKey:ProcedureOrderId"`
    Status                      ProcedureStatus

// parent associations as their child

}

