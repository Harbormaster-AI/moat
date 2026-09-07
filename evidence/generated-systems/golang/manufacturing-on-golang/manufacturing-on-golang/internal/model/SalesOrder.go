package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// SalesOrder Declaration
//==============================================================
type SalesOrder struct {
    gorm.Model
     OrderNumber                                    string
    OrderDate                                                            time.Time
    TotalAmount                                                            string
    CustomerId         *uint
    Customer           *Customer `gorm:"foreignKey:CustomerId"`
    PlantId         *uint
    Plant           *Plant `gorm:"foreignKey:PlantId"`
     Lines           []SalesOrderLine `gorm:"foreignKey:LinesFromSalesOrderId"`
     WorkOrders           []WorkOrder `gorm:"foreignKey:WorkOrdersFromSalesOrderId"`
    Status                      SalesOrderStatus

// parent associations as their child

}

