package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// LabResult Declaration
//==============================================================
type LabResult struct {
    gorm.Model
     ResultCode                                    string
    IssuedDate                                                            time.Time
    LaboratoryOrderId         *uint
    LaboratoryOrder           *LaboratoryOrder `gorm:"foreignKey:LaboratoryOrderId"`
     Observations           []Observation `gorm:"foreignKey:ObservationsFromLabResultId"`
    LaboratoryId         *uint
    Laboratory           *Laboratory `gorm:"foreignKey:LaboratoryId"`
    Status                      ResultStatus

// parent associations as their child

}

