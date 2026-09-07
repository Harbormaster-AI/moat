package model

import (
    "gorm.io/gorm"
)

//==============================================================
// MaintenanceWorkOrder Declaration
//==============================================================
type MaintenanceWorkOrder struct {
    gorm.Model
     WorkOrderNumber                                    string
    AircraftId         *uint
    Aircraft           *Aircraft `gorm:"foreignKey:AircraftId"`
    AirworthinessDirectiveId         *uint
    AirworthinessDirective           *AirworthinessDirective `gorm:"foreignKey:AirworthinessDirectiveId"`
    ServiceBulletinId         *uint
    ServiceBulletin           *ServiceBulletin `gorm:"foreignKey:ServiceBulletinId"`
    Status                      WorkOrderStatus

// parent associations as their child

}

