package model

import (
    "gorm.io/gorm"
)

//==============================================================
// CreativeAsset Declaration
//==============================================================
type CreativeAsset struct {
    gorm.Model
     Name                                    string
    ClickUrl                                                            string
    LandingPage                                                            string
    Width                                                            string
    Height                                                            string
    DurationSeconds                                                            string
     Files           []CreativeFile `gorm:"foreignKey:FilesFromCreativeAssetId"`
     Approvals           []CreativeApproval `gorm:"foreignKey:ApprovalsFromCreativeAssetId"`
     Variations           []CreativeVariation `gorm:"foreignKey:VariationsFromCreativeAssetId"`
     LineItems           []LineItem `gorm:"foreignKey:LineItemsFromCreativeAssetId"`
    CreativeType                      CreativeType
    AdFormat                      AdFormat

// parent associations as their child

}

