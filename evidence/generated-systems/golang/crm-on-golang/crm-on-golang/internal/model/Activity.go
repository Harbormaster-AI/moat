package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Activity Declaration
//==============================================================
type Activity struct {
    gorm.Model
     Subject                                    string
    DueDate                                                            time.Time
    StartAt                                                            time.Time
    EndAt                                                            time.Time
    Location                                    string
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
    OpportunityId         *uint
    Opportunity           *Opportunity `gorm:"foreignKey:OpportunityId"`
    CaseId         *uint
    Case           *Case_ `gorm:"foreignKey:CaseId"`
    CampaignId         *uint
    Campaign           *Campaign `gorm:"foreignKey:CampaignId"`
    ActivityType                      ActivityType
    Status                      ActivityStatus
    Priority                      ActivityPriority

// parent associations as their child

}

