package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// SoftwareUpdate Declaration
//==============================================================
type SoftwareUpdate struct {
    gorm.Model
     Version                                    string
    AppliedDate                                                            time.Time
    DeviceId         *uint
    Device           *MedicalDevice `gorm:"foreignKey:DeviceId"`
    UpdateType                      SoftwareUpdateType

// parent associations as their child

}

