package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Opportunity Declaration
//==============================================================
type Opportunity struct {
    gorm.Model
     Name                                    string
    Amount                                                            string
    CloseDate                                                            time.Time
    Probability                                                            string
    Description                                    string
    OrganizationId         *uint
    Organization           *Organization `gorm:"foreignKey:OrganizationId"`
    AccountId         *uint
    Account           *Account `gorm:"foreignKey:AccountId"`
    OwnerId         *uint
    Owner           *User `gorm:"foreignKey:OwnerId"`
     Contacts           []Contact `gorm:"foreignKey:ContactsFromOpportunityId"`
     LineItems           []OpportunityLineItem `gorm:"foreignKey:LineItemsFromOpportunityId"`
     StageHistory           []OpportunityStageHistory `gorm:"foreignKey:StageHistoryFromOpportunityId"`
     Quotes           []Quote `gorm:"foreignKey:QuotesFromOpportunityId"`
     Orders           []Order `gorm:"foreignKey:OrdersFromOpportunityId"`
     Campaigns           []Campaign `gorm:"foreignKey:CampaignsFromOpportunityId"`
     Activities           []Activity `gorm:"foreignKey:ActivitiesFromOpportunityId"`
     Teams           []Team `gorm:"foreignKey:TeamsFromOpportunityId"`
    Stage                      OpportunityStage
    Type                      OpportunityType
    ForecastCategory                      ForecastCategory

// parent associations as their child

}

