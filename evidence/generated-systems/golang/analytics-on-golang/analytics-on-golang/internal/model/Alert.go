package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Alert Declaration
//==============================================================
type Alert struct {
    gorm.Model
     Title                                    string
    CreatedAt                                                            time.Time
    MetricId         *uint
    Metric           *Metric `gorm:"foreignKey:MetricId"`
    DashboardId         *uint
    Dashboard           *Dashboard `gorm:"foreignKey:DashboardId"`
    DatasetId         *uint
    Dataset           *DataSet `gorm:"foreignKey:DatasetId"`
    RuleId         *uint
    Rule           *QualityRule `gorm:"foreignKey:RuleId"`
     Anomalies           []Anomaly `gorm:"foreignKey:AnomaliesFromAlertId"`
     Subscribers           []Subscriber `gorm:"foreignKey:SubscribersFromAlertId"`
    Severity                      AlertSeverity
    Status                      AlertStatus

// parent associations as their child

}

