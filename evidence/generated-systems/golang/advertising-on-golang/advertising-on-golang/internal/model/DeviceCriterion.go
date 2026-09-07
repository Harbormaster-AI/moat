package model

import (
    "gorm.io/gorm"
)

//==============================================================
// DeviceCriterion Declaration
//==============================================================
type DeviceCriterion struct {
    gorm.Model
     TargetingProfileId         *uint
    TargetingProfile           *TargetingProfile `gorm:"foreignKey:TargetingProfileId"`
    DeviceType                      DeviceType
    PlatformType                      PlatformType
    Operator                      TargetingOperator

// parent associations as their child

}

