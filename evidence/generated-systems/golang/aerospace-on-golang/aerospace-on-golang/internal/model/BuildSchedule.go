package model

import (
    "gorm.io/gorm"
)

//==============================================================
// BuildSchedule Declaration
//==============================================================
type BuildSchedule struct {
    gorm.Model
     ScheduleNumber                                    string
     ProductionOrders           []ProductionOrder `gorm:"foreignKey:ProductionOrdersFromBuildScheduleId"`
    Status                      ScheduleStatus

// parent associations as their child

}

