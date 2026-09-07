package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Visualization Declaration
//==============================================================
type Visualization struct {
    gorm.Model
     Title                                    string
    Options                                                            string
    DashboardId         *uint
    Dashboard           *Dashboard `gorm:"foreignKey:DashboardId"`
    ReportId         *uint
    Report           *Report `gorm:"foreignKey:ReportId"`
     Metrics           []Metric `gorm:"foreignKey:MetricsFromVisualizationId"`
     Dimensions           []Dimension `gorm:"foreignKey:DimensionsFromVisualizationId"`
     Datasets           []DataSet `gorm:"foreignKey:DatasetsFromVisualizationId"`
    ChartType                      ChartType

// parent associations as their child

}

