package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Lead Declaration
//==============================================================
type Lead struct {
    gorm.Model
     FirstName                                    string
    LastName                                    string
    Company                                    string
    Email                                                            string
    Phone                                                            string
    Converted                                    bool
    OrganizationId         *uint
    Organization           *Organization `gorm:"foreignKey:OrganizationId"`
    OwnerId         *uint
    Owner           *User `gorm:"foreignKey:OwnerId"`
     Activities           []Activity `gorm:"foreignKey:ActivitiesFromLeadId"`
     Campaigns           []Campaign `gorm:"foreignKey:CampaignsFromLeadId"`
    ConvertedAccountId         *uint
    ConvertedAccount           *Account `gorm:"foreignKey:ConvertedAccountId"`
    ConvertedContactId         *uint
    ConvertedContact           *Contact `gorm:"foreignKey:ConvertedContactId"`
    ConvertedOpportunityId         *uint
    ConvertedOpportunity           *Opportunity `gorm:"foreignKey:ConvertedOpportunityId"`
     Notes           []Note `gorm:"foreignKey:NotesFromLeadId"`
     EmailMessages           []EmailMessage `gorm:"foreignKey:EmailMessagesFromLeadId"`
    Status                      LeadStatus
    Source                      LeadSource
    Rating                      LeadRating

// parent associations as their child

}

