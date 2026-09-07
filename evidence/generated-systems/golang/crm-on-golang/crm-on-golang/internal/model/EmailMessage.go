package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// EmailMessage Declaration
//==============================================================
type EmailMessage struct {
    gorm.Model
     Subject                                    string
    Body                                    string
    SentAt                                                            time.Time
    MessageId                                    string
    OrganizationId         *uint
    Organization           *Organization `gorm:"foreignKey:OrganizationId"`
    OwnerId         *uint
    Owner           *User `gorm:"foreignKey:OwnerId"`
    AccountId         *uint
    Account           *Account `gorm:"foreignKey:AccountId"`
    ContactId         *uint
    Contact           *Contact `gorm:"foreignKey:ContactId"`
    LeadId         *uint
    Lead           *Lead `gorm:"foreignKey:LeadId"`
    CaseId         *uint
    Case           *Case_ `gorm:"foreignKey:CaseId"`
    OpportunityId         *uint
    Opportunity           *Opportunity `gorm:"foreignKey:OpportunityId"`
    CampaignId         *uint
    Campaign           *Campaign `gorm:"foreignKey:CampaignId"`
    Direction                      EmailDirection
    Status                      EmailStatus

// parent associations as their child

}

