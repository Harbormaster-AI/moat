package model

import (
    "gorm.io/gorm"
)

//==============================================================
// BrandSafetyPolicy Declaration
//==============================================================
type BrandSafetyPolicy struct {
    gorm.Model
      TargetingProfiles           []TargetingProfile `gorm:"foreignKey:TargetingProfilesFromBrandSafetyPolicyId"`
    Level                      BrandSafetyLevel
    ContentRatingThreshold                      ContentRating

// parent associations as their child

}

