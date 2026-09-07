package model

import (
    "gorm.io/gorm"
)

//==============================================================
// DataProvider Declaration
//==============================================================
type DataProvider struct {
    gorm.Model
     Name                                    string
    Website                                    string
     AudienceSegments           []AudienceSegment `gorm:"foreignKey:AudienceSegmentsFromDataProviderId"`
    ProviderType                      DataProviderType

// parent associations as their child

}

