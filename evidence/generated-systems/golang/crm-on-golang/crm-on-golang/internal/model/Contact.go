package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Contact Declaration
//==============================================================
type Contact struct {
    gorm.Model
     FirstName                                    string
    LastName                                    string
    Title                                    string
    Email                                                            string
    Phone                                                            string
    Mobile                                                            string
    MailingAddress                                                            string
    OrganizationId         *uint
    Organization           *Organization `gorm:"foreignKey:OrganizationId"`
    AccountId         *uint
    Account           *Account `gorm:"foreignKey:AccountId"`
    OwnerId         *uint
    Owner           *User `gorm:"foreignKey:OwnerId"`
     Activities           []Activity `gorm:"foreignKey:ActivitiesFromContactId"`
     Opportunities           []Opportunity `gorm:"foreignKey:OpportunitiesFromContactId"`
     Cases           []Case_ `gorm:"foreignKey:CasesFromContactId"`
     Campaigns           []Campaign `gorm:"foreignKey:CampaignsFromContactId"`
     Notes           []Note `gorm:"foreignKey:NotesFromContactId"`
     EmailMessages           []EmailMessage `gorm:"foreignKey:EmailMessagesFromContactId"`
    PreferredContactMethod                      ContactMethod

// parent associations as their child

}

