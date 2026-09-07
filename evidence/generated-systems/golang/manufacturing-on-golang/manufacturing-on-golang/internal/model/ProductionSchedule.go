package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// ProductionSchedule Declaration
//==============================================================
type ProductionSchedule struct {
    gorm.Model
     ScheduleNumber                                    string
    HorizonStart                                                            time.Time
    HorizonEnd                                                            time.Time
    PlantId         *uint
    Plant           *Plant `gorm:"foreignKey:PlantId"`
     WorkOrders           []WorkOrder `gorm:"foreignKey:WorkOrdersFromProductionScheduleId"`
    Status                      ScheduleStatus

// parent associations as their child

}

