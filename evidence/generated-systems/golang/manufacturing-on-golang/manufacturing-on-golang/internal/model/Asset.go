package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Asset Declaration
//==============================================================
type Asset struct {
    gorm.Model
     AssetTag                                    string
    AssetName                                    string
    CommissioningDate                                                            time.Time
    PlantId         *uint
    Plant           *Plant `gorm:"foreignKey:PlantId"`
    WorkCenterId         *uint
    WorkCenter           *WorkCenter `gorm:"foreignKey:WorkCenterId"`
     MaintenanceOrders           []MaintenanceOrder `gorm:"foreignKey:MaintenanceOrdersFromAssetId"`
     MaintenancePlans           []MaintenancePlan `gorm:"foreignKey:MaintenancePlansFromAssetId"`
    AssetStatus                      AssetStatus

// parent associations as their child

}

