package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Operation Declaration
//==============================================================
type Operation struct {
    gorm.Model
     OperationNumber                                    string
    Name                                    string
    SetupTime                                                            string
    StandardCycleTime                                                            string
    RoutingId         *uint
    Routing           *Routing `gorm:"foreignKey:RoutingId"`
    WorkCenterId         *uint
    WorkCenter           *WorkCenter `gorm:"foreignKey:WorkCenterId"`
    InspectionPlanId         *uint
    InspectionPlan           *InspectionPlan `gorm:"foreignKey:InspectionPlanId"`
    OperationType                      OperationType

// parent associations as their child

}

