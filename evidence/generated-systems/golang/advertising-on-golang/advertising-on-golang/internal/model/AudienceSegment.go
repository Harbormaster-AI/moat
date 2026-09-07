package model

import (
    "gorm.io/gorm"
)

//==============================================================
// AudienceSegment Declaration
//==============================================================
type AudienceSegment struct {
    gorm.Model
     Name                                    string
    EstimatedReach                                                            string
    Description                                    string
    ProviderId         *uint
    Provider           *DataProvider `gorm:"foreignKey:ProviderId"`
     Campaigns           []Campaign `gorm:"foreignKey:CampaignsFromAudienceSegmentId"`
    ProviderType                      DataProviderType

// parent associations as their child

}

