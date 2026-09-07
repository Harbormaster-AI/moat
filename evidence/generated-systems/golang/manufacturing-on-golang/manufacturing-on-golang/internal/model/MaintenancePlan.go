package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// MaintenancePlan Declaration
//==============================================================
type MaintenancePlan struct {
    gorm.Model
     PlanNumber                                    string
    Interval                                                            string
    LastServiceDate                                                            time.Time
    AssetId         *uint
    Asset           *Asset `gorm:"foreignKey:AssetId"`
     MaintenanceOrders           []MaintenanceOrder `gorm:"foreignKey:MaintenanceOrdersFromMaintenancePlanId"`
    Strategy                      MaintenanceStrategy

// parent associations as their child

}

