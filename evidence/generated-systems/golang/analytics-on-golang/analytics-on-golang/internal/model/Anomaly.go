package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Anomaly Declaration
//==============================================================
type Anomaly struct {
    gorm.Model
     OccurredAt                                                            time.Time
    Details                                    string
    TimeSeriesId         *uint
    TimeSeries           *TimeSeries `gorm:"foreignKey:TimeSeriesId"`
    AlertId         *uint
    Alert           *Alert `gorm:"foreignKey:AlertId"`
    DatasetId         *uint
    Dataset           *DataSet `gorm:"foreignKey:DatasetId"`
    AnomalyType                      AnomalyType
    Severity                      AlertSeverity

// parent associations as their child

}

