package model

import (
    "gorm.io/gorm"
)

//==============================================================
// TimeSeries Declaration
//==============================================================
type TimeSeries struct {
    gorm.Model
     Name                                    string
    Timezone                                    string
     Datasets           []DataSet `gorm:"foreignKey:DatasetsFromTimeSeriesId"`
     Forecasts           []Forecast `gorm:"foreignKey:ForecastsFromTimeSeriesId"`
     Anomalies           []Anomaly `gorm:"foreignKey:AnomaliesFromTimeSeriesId"`
    Granularity                      TimeGranularity

// parent associations as their child

}

