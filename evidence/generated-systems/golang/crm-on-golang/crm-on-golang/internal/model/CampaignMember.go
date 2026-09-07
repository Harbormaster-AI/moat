package model

import (
    "gorm.io/gorm"
)

//==============================================================
// CampaignMember Declaration
//==============================================================
type CampaignMember struct {
    gorm.Model
     Responded                                    bool
    CampaignId         *uint
    Campaign           *Campaign `gorm:"foreignKey:CampaignId"`
    LeadId         *uint
    Lead           *Lead `gorm:"foreignKey:LeadId"`
    ContactId         *uint
    Contact           *Contact `gorm:"foreignKey:ContactId"`
    Status                      CampaignMemberStatus
    MemberType                      CampaignMemberType

// parent associations as their child

}

