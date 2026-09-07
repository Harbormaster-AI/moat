package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Customer Declaration
//==============================================================
type Customer struct {
    gorm.Model
     Name                                    string
    CustomerCode                                    string
    Address                                                            string
     Enterprises           []Enterprise `gorm:"foreignKey:EnterprisesFromCustomerId"`
     SalesOrders           []SalesOrder `gorm:"foreignKey:SalesOrdersFromCustomerId"`
    CustomerType                      CustomerType

// parent associations as their child

}

