package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// WorkAuthorization Declaration
//==============================================================
type WorkAuthorization struct {
    gorm.Model
     Country                                    string
    ExpirationDate                                                            time.Time
    EmployeeId         *uint
    Employee           *Employee `gorm:"foreignKey:EmployeeId"`
     Documents           []Document `gorm:"foreignKey:DocumentsFromWorkAuthorizationId"`
    Status                      WorkAuthorizationStatus

// parent associations as their child

}

