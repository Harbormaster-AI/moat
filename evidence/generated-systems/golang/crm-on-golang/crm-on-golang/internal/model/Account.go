package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Account Declaration
//==============================================================
type Account struct {
    gorm.Model
     Name                                    string
    AccountNumber                                    string
    Industry                                    string
    BillingAddress                                                            string
    ShippingAddress                                                            string
    Website                                                            string
    Phone                                                            string
    AsActive                                    bool
    OrganizationId         *uint
    Organization           *Organization `gorm:"foreignKey:OrganizationId"`
    ParentAccountId         *uint
    ParentAccount           *Account `gorm:"foreignKey:ParentAccountId"`
     ChildAccounts           []Account `gorm:"foreignKey:ChildAccountsFromAccountId"`
     Contacts           []Contact `gorm:"foreignKey:ContactsFromAccountId"`
     Opportunities           []Opportunity `gorm:"foreignKey:OpportunitiesFromAccountId"`
     Cases           []Case_ `gorm:"foreignKey:CasesFromAccountId"`
    OwnerId         *uint
    Owner           *User `gorm:"foreignKey:OwnerId"`
    TerritoryId         *uint
    Territory           *Territory `gorm:"foreignKey:TerritoryId"`
     Activities           []Activity `gorm:"foreignKey:ActivitiesFromAccountId"`
     Campaigns           []Campaign `gorm:"foreignKey:CampaignsFromAccountId"`
     Quotes           []Quote `gorm:"foreignKey:QuotesFromAccountId"`
     Orders           []Order `gorm:"foreignKey:OrdersFromAccountId"`
     Contracts           []Contract `gorm:"foreignKey:ContractsFromAccountId"`
     Notes           []Note `gorm:"foreignKey:NotesFromAccountId"`
     EmailMessages           []EmailMessage `gorm:"foreignKey:EmailMessagesFromAccountId"`
    AccountType                      AccountType
    LifecycleStage                      AccountLifecycleStage

// parent associations as their child

}

