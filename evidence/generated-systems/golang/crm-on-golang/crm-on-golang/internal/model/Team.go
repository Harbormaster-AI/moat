package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Team Declaration
//==============================================================
type Team struct {
    gorm.Model
     Name                                    string
    OrganizationId         *uint
    Organization           *Organization `gorm:"foreignKey:OrganizationId"`
     Users           []User `gorm:"foreignKey:UsersFromTeamId"`
     Accounts           []Account `gorm:"foreignKey:AccountsFromTeamId"`
     Opportunities           []Opportunity `gorm:"foreignKey:OpportunitiesFromTeamId"`
     Cases           []Case_ `gorm:"foreignKey:CasesFromTeamId"`
     Campaigns           []Campaign `gorm:"foreignKey:CampaignsFromTeamId"`
    TeamType                      TeamType

// parent associations as their child

}

