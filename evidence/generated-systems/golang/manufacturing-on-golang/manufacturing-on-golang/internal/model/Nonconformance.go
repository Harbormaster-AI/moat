package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Nonconformance Declaration
//==============================================================
type Nonconformance struct {
    gorm.Model
     NcNumber                                    string
    Description                                    string
    ContainmentAction                                    string
    ItemId         *uint
    Item           *Item `gorm:"foreignKey:ItemId"`
    WorkOrderId         *uint
    WorkOrder           *WorkOrder `gorm:"foreignKey:WorkOrderId"`
    InspectionLotId         *uint
    InspectionLot           *InspectionLot `gorm:"foreignKey:InspectionLotId"`
    CorrectiveActionId         *uint
    CorrectiveAction           *CorrectiveAction `gorm:"foreignKey:CorrectiveActionId"`
    NcType                      NonconformanceType
    Severity                      QualitySeverity
    Status                      NonconformanceStatus

// parent associations as their child

}

