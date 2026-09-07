package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Case_ Declaration
//==============================================================
type Case_ struct {
    gorm.Model
     CaseNumber                                    string
    Subject                                    string
    Description                                    string
    SlaDue                                                            time.Time
    OrganizationId         *uint
    Organization           *Organization `gorm:"foreignKey:OrganizationId"`
    AccountId         *uint
    Account           *Account `gorm:"foreignKey:AccountId"`
    ContactId         *uint
    Contact           *Contact `gorm:"foreignKey:ContactId"`
    OwnerId         *uint
    Owner           *User `gorm:"foreignKey:OwnerId"`
    TeamId         *uint
    Team           *Team `gorm:"foreignKey:TeamId"`
     Activities           []Activity `gorm:"foreignKey:ActivitiesFromCase_Id"`
     CaseComments           []Note `gorm:"foreignKey:CaseCommentsFromCase_Id"`
     Emails           []EmailMessage `gorm:"foreignKey:EmailsFromCase_Id"`
     RelatedOpportunities           []Opportunity `gorm:"foreignKey:RelatedOpportunitiesFromCase_Id"`
    Status                      CaseStatus
    Priority                      CasePriority
    Origin                      CaseOrigin
    Severity                      CaseSeverity

// parent associations as their child

}

