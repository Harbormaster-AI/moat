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
    Leader                                    string
    OrganizationId         *uint
    Organization           *Organization `gorm:"foreignKey:OrganizationId"`
     Audits           []AuditEngagement `gorm:"foreignKey:AuditsFromBusinessUnitId"`

// parent associations as their child

}

