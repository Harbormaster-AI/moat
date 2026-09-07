package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// MRPRun Declaration
//==============================================================
type MRPRun struct {
    gorm.Model
     RunNumber                                    string
    RunDateTime                                                            time.Time
    PlanningHorizonDays                                                            string
    PlantId         *uint
    Plant           *Plant `gorm:"foreignKey:PlantId"`
     PlannedOrders           []PlannedOrder `gorm:"foreignKey:PlannedOrdersFromMRPRunId"`
    Status                      MRPRunStatus

// parent associations as their child

}

