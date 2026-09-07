package model

import (
    "gorm.io/gorm"
)

//==============================================================
// WorkCenter Declaration
//==============================================================
type WorkCenter struct {
    gorm.Model
     Name                                    string
    Code                                    string
    CapacityPerHour                                                            string
    OeeTarget                                                            string
    ProductionLineId         *uint
    ProductionLine           *ProductionLine `gorm:"foreignKey:ProductionLineId"`
     Assets           []Asset `gorm:"foreignKey:AssetsFromWorkCenterId"`
     MaintenanceOrders           []MaintenanceOrder `gorm:"foreignKey:MaintenanceOrdersFromWorkCenterId"`
    WorkCenterType                      WorkCenterType

// parent associations as their child

}

