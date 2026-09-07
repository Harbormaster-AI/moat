package model

import (
    "gorm.io/gorm"
)

//==============================================================
// AirworthinessDirective Declaration
//==============================================================
type AirworthinessDirective struct {
    gorm.Model
     DirectiveNumber                                    string
    Title                                    string
     WorkOrders           []MaintenanceWorkOrder `gorm:"foreignKey:WorkOrdersFromAirworthinessDirectiveId"`

// parent associations as their child

}

