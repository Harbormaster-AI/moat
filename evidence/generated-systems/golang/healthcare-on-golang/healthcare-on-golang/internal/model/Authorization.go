package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Authorization Declaration
//==============================================================
type Authorization struct {
    gorm.Model
     AuthNumber                                    string
    RequestedService                                    string
    CoverageId         *uint
    Coverage           *Coverage `gorm:"foreignKey:CoverageId"`
    OrderId         *uint
    Order           *ClinicalOrder `gorm:"foreignKey:OrderId"`
    Status                      AuthorizationStatus

// parent associations as their child

}

