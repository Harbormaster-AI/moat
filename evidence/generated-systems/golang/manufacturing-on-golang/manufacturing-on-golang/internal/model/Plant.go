package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Plant Declaration
//==============================================================
type Plant struct {
    gorm.Model
     Name                                    string
    PlantCode                                    string
    Address                                                            string
    TimeZone                                    string
    EnterpriseId         *uint
    Enterprise           *Enterprise `gorm:"foreignKey:EnterpriseId"`
     ProductionLines           []ProductionLine `gorm:"foreignKey:ProductionLinesFromPlantId"`
     WorkCenters           []WorkCenter `gorm:"foreignKey:WorkCentersFromPlantId"`
     Warehouses           []Warehouse `gorm:"foreignKey:WarehousesFromPlantId"`
     Assets           []Asset `gorm:"foreignKey:AssetsFromPlantId"`
     ProductionSchedules           []ProductionSchedule `gorm:"foreignKey:ProductionSchedulesFromPlantId"`

// parent associations as their child

}

