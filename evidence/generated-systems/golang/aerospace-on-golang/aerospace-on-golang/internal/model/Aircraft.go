package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Aircraft Declaration
//==============================================================
type Aircraft struct {
    gorm.Model
     Msn                                                            string
    DeliveryDate                                                            time.Time
    VariantId         *uint
    Variant           *AircraftVariant `gorm:"foreignKey:VariantId"`
    OperatorId         *uint
    Operator           *Operator `gorm:"foreignKey:OperatorId"`
    RegistrationId         *uint
    Registration           *Registration `gorm:"foreignKey:RegistrationId"`
    WarrantyId         *uint
    Warranty           *Warranty `gorm:"foreignKey:WarrantyId"`
     MaintenanceRecords           []MaintenanceWorkOrder `gorm:"foreignKey:MaintenanceRecordsFromAircraftId"`
    ConnectedAircraftId         *uint
    ConnectedAircraft           *ConnectedAircraft `gorm:"foreignKey:ConnectedAircraftId"`
    CabinLayoutId         *uint
    CabinLayout           *CabinLayout `gorm:"foreignKey:CabinLayoutId"`

// parent associations as their child

}

