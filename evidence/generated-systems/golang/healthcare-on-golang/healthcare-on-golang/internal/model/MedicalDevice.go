package model

import (
    "gorm.io/gorm"
)

//==============================================================
// MedicalDevice Declaration
//==============================================================
type MedicalDevice struct {
    gorm.Model
     Udi                                    string
    Manufacturer                                    string
    PatientId         *uint
    Patient           *Patient `gorm:"foreignKey:PatientId"`
     Observations           []Observation `gorm:"foreignKey:ObservationsFromMedicalDeviceId"`
     SoftwareUpdates           []SoftwareUpdate `gorm:"foreignKey:SoftwareUpdatesFromMedicalDeviceId"`
    DeviceType                      DeviceType
    ConnectivityStatus                      DeviceConnectivityStatus

// parent associations as their child

}

