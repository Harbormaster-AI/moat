package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// QualityCheck Declaration
//==============================================================
type QualityCheck struct {
    gorm.Model
     CheckedAt                                                            time.Time
    ObservedValue                                                            string
    SampleSize                                                            string
    RuleId         *uint
    Rule           *QualityRule `gorm:"foreignKey:RuleId"`
    DatasetId         *uint
    Dataset           *DataSet `gorm:"foreignKey:DatasetId"`
    Status                      QualityStatus

// parent associations as their child

}

