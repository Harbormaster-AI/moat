package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Campaign Declaration
//==============================================================
type Campaign struct {
    gorm.Model
     Name                                    string
    StartDate                                                            time.Time
    EndDate                                                            time.Time
    Budget                                                            string
    ActualCost                                                            string
    ExpectedRevenue                                                            string
    OrganizationId         *uint
    Organization           *Organization `gorm:"foreignKey:OrganizationId"`
    ParentCampaignId         *uint
    ParentCampaign           *Campaign `gorm:"foreignKey:ParentCampaignId"`
     ChildCampaigns           []Campaign `gorm:"foreignKey:ChildCampaignsFromCampaignId"`
     Members           []CampaignMember `gorm:"foreignKey:MembersFromCampaignId"`
     Opportunities           []Opportunity `gorm:"foreignKey:OpportunitiesFromCampaignId"`
     Accounts           []Account `gorm:"foreignKey:AccountsFromCampaignId"`
     Leads           []Lead `gorm:"foreignKey:LeadsFromCampaignId"`
     Contacts           []Contact `gorm:"foreignKey:ContactsFromCampaignId"`
     Teams           []Team `gorm:"foreignKey:TeamsFromCampaignId"`
     Activities           []Activity `gorm:"foreignKey:ActivitiesFromCampaignId"`
    Status                      CampaignStatus
    Type                      CampaignType

// parent associations as their child

}

