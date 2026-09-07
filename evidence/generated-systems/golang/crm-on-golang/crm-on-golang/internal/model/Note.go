package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Note Declaration
//==============================================================
type Note struct {
    gorm.Model
     Title                                    string
    Content                                    string
    CreatedAt                                                            time.Time
    UpdatedAt                                                            time.Time
    OrganizationId         *uint
    Organization           *Organization `gorm:"foreignKey:OrganizationId"`
    OwnerId         *uint
    Owner           *User `gorm:"foreignKey:OwnerId"`
    AccountId         *uint
    Account           *Account `gorm:"foreignKey:AccountId"`
    ContactId         *uint
    Contact           *Contact `gorm:"foreignKey:ContactId"`
    OpportunityId         *uint
    Opportunity           *Opportunity `gorm:"foreignKey:OpportunityId"`
    CaseId         *uint
    Case           *Case_ `gorm:"foreignKey:CaseId"`
    LeadId         *uint
    Lead           *Lead `gorm:"foreignKey:LeadId"`

// parent associations as their child

}

