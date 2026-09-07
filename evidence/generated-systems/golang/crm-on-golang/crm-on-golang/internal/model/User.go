package model

import (
    "gorm.io/gorm"
)

//==============================================================
// User Declaration
//==============================================================
type User struct {
    gorm.Model
     Username                                    string
    FullName                                    string
    Email                                                            string
    Locale                                                            string
    OrganizationId         *uint
    Organization           *Organization `gorm:"foreignKey:OrganizationId"`
     Teams           []Team `gorm:"foreignKey:TeamsFromUserId"`
     Activities           []Activity `gorm:"foreignKey:ActivitiesFromUserId"`
     OwnedAccounts           []Account `gorm:"foreignKey:OwnedAccountsFromUserId"`
     OwnedLeads           []Lead `gorm:"foreignKey:OwnedLeadsFromUserId"`
     OwnedOpportunities           []Opportunity `gorm:"foreignKey:OwnedOpportunitiesFromUserId"`
     OwnedCases           []Case_ `gorm:"foreignKey:OwnedCasesFromUserId"`
     Quotes           []Quote `gorm:"foreignKey:QuotesFromUserId"`
     Orders           []Order `gorm:"foreignKey:OrdersFromUserId"`
     Contracts           []Contract `gorm:"foreignKey:ContractsFromUserId"`
     EmailMessages           []EmailMessage `gorm:"foreignKey:EmailMessagesFromUserId"`
    Role                      UserRole
    Status                      UserStatus

// parent associations as their child

}

