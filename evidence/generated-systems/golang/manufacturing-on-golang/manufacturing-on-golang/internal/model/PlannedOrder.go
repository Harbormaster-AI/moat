package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// PlannedOrder Declaration
//==============================================================
type PlannedOrder struct {
    gorm.Model
     PlannedOrderNumber                                    string
    Quantity                                                            string
    DueDate                                                            time.Time
    MrpRunId         *uint
    MrpRun           *MRPRun `gorm:"foreignKey:MrpRunId"`
    ItemId         *uint
    Item           *Item `gorm:"foreignKey:ItemId"`
    PlantId         *uint
    Plant           *Plant `gorm:"foreignKey:PlantId"`
    OrderType                      PlannedOrderType
    Status                      PlannedOrderStatus

// parent associations as their child

}

