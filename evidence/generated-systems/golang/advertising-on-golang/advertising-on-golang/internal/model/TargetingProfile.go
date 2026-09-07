package model

import (
    "gorm.io/gorm"
)

//==============================================================
// TargetingProfile Declaration
//==============================================================
type TargetingProfile struct {
    gorm.Model
     Name                                    string
     AudienceSegments           []AudienceSegment `gorm:"foreignKey:AudienceSegmentsFromTargetingProfileId"`
     GeoRegions           []GeoRegion `gorm:"foreignKey:GeoRegionsFromTargetingProfileId"`
     ContentCategories           []ContentCategory `gorm:"foreignKey:ContentCategoriesFromTargetingProfileId"`
    BrandSafetyPolicyId         *uint
    BrandSafetyPolicy           *BrandSafetyPolicy `gorm:"foreignKey:BrandSafetyPolicyId"`
     DeviceCriteria           []DeviceCriterion `gorm:"foreignKey:DeviceCriteriaFromTargetingProfileId"`

// parent associations as their child

}

