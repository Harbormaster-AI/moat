package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Forecast Declaration
//==============================================================
type Forecast struct {
    gorm.Model
     Name                                    string
    Horizon                                                            string
    ModelVersionId         *uint
    ModelVersion           *ModelVersion `gorm:"foreignKey:ModelVersionId"`
    TimeSeriesId         *uint
    TimeSeries           *TimeSeries `gorm:"foreignKey:TimeSeriesId"`
     Datasets           []DataSet `gorm:"foreignKey:DatasetsFromForecastId"`
    Granularity                      TimeGranularity

// parent associations as their child

}

