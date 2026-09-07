package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Forecast Declaration
//==============================================================
type Forecast struct {
    gorm.Model
     ForecastNumber                                    string
    ForecastHorizonStart                                                            time.Time
    ForecastHorizonEnd                                                            time.Time
     Lines           []ForecastLine `gorm:"foreignKey:LinesFromForecastId"`
    Method                      ForecastMethod

// parent associations as their child

}

