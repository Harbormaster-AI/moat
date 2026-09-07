package model

import (
    "gorm.io/gorm"
)

//==============================================================
// CreativeVariation Declaration
//==============================================================
type CreativeVariation struct {
    gorm.Model
     Name                                    string
    Language                                    string
    Headline                                    string
    BodyText                                    string
    CallToAction                                    string
    CreativeAssetId         *uint
    CreativeAsset           *CreativeAsset `gorm:"foreignKey:CreativeAssetId"`

// parent associations as their child

}

