package model

import (
    "gorm.io/gorm"
)

//==============================================================
// AuditProgram Declaration
//==============================================================
type AuditProgram struct {
    gorm.Model
     Name                                    string
    Scope                                    string
    OrganizationId         *uint
    Organization           *Organization `gorm:"foreignKey:OrganizationId"`
     Engagements           []AuditEngagement `gorm:"foreignKey:EngagementsFromAuditProgramId"`
    Cycle                      AuditCycle
    Status                      AuditStatus

// parent associations as their child

}

