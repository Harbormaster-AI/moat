package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// InspectionResult Declaration
//==============================================================
type InspectionResult struct {
    gorm.Model
     ResultValue                                                            string
    RecordedOn                                                            time.Time
    Notes                                    string
    InspectionLotId         *uint
    InspectionLot           *InspectionLot `gorm:"foreignKey:InspectionLotId"`
    CharacteristicId         *uint
    Characteristic           *InspectionCharacteristic `gorm:"foreignKey:CharacteristicId"`
    ResultStatus                      InspectionResultStatus

// parent associations as their child

}

