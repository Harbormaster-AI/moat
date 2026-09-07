package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// WorkOrder Declaration
//==============================================================
type WorkOrder struct {
    gorm.Model
     WorkOrderNumber                                    string
    PlannedStart                                                            time.Time
    PlannedEnd                                                            time.Time
    Quantity                                                            string
    Priority                                                            string
    ItemId         *uint
    Item           *Item `gorm:"foreignKey:ItemId"`
    PlantId         *uint
    Plant           *Plant `gorm:"foreignKey:PlantId"`
    RoutingId         *uint
    Routing           *Routing `gorm:"foreignKey:RoutingId"`
    BomId         *uint
    Bom           *BOM `gorm:"foreignKey:BomId"`
    ProductionScheduleId         *uint
    ProductionSchedule           *ProductionSchedule `gorm:"foreignKey:ProductionScheduleId"`
    SalesOrderId         *uint
    SalesOrder           *SalesOrder `gorm:"foreignKey:SalesOrderId"`
    Status                      WorkOrderStatus

// parent associations as their child

}

