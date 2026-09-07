package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Evidence Declaration
//==============================================================
type Evidence struct {
    gorm.Model
     Title                                    string
    LocationUrl                                                            string
    ReceivedDate                                                            time.Time
    ControlTestId         *uint
    ControlTest           *ControlTest_ `gorm:"foreignKey:ControlTestId"`
    ControlId         *uint
    Control           *Control `gorm:"foreignKey:ControlId"`
    ObligationId         *uint
    Obligation           *Obligation `gorm:"foreignKey:ObligationId"`
    WorkpaperId         *uint
    Workpaper           *AuditWorkpaper `gorm:"foreignKey:WorkpaperId"`
    EvidenceType                      EvidenceType

// parent associations as their child

}

