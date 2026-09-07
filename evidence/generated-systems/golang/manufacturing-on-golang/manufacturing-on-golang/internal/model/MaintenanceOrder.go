package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// MaintenanceOrder Declaration
//==============================================================
type MaintenanceOrder struct {
    gorm.Model
     OrderNumber                                    string
    Priority                                                            string
    RequestedDate                                                            time.Time
    CompletionDate                                                            time.Time
    AssetId         *uint
    Asset           *Asset `gorm:"foreignKey:AssetId"`
    PlanId         *uint
    Plan           *MaintenancePlan `gorm:"foreignKey:PlanId"`
    WorkCenterId         *uint
    WorkCenter           *WorkCenter `gorm:"foreignKey:WorkCenterId"`
    Status                      MaintenanceOrderStatus

// parent associations as their child

}

