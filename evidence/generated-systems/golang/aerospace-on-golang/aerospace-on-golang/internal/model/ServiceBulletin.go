package model

import (
    "gorm.io/gorm"
)

//==============================================================
// ServiceBulletin Declaration
//==============================================================
type ServiceBulletin struct {
    gorm.Model
     BulletinNumber                                    string
     WorkOrders           []MaintenanceWorkOrder `gorm:"foreignKey:WorkOrdersFromServiceBulletinId"`
     Variants           []AircraftVariant `gorm:"foreignKey:VariantsFromServiceBulletinId"`
    Category                      ServiceBulletinCategory

// parent associations as their child

}

