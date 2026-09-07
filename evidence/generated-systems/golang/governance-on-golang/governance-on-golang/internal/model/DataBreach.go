package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// DataBreach Declaration
//==============================================================
type DataBreach struct {
    gorm.Model
     IncidentDate                                                            time.Time
    Description                                    string
    RecordsAffected                                                            string
    NotificationRequired                                    bool
    OrganizationId         *uint
    Organization           *Organization `gorm:"foreignKey:OrganizationId"`
     ProcessingActivities           []DataProcessingActivity `gorm:"foreignKey:ProcessingActivitiesFromDataBreachId"`
     DataCategories           []DataCategory `gorm:"foreignKey:DataCategoriesFromDataBreachId"`
     ThirdParties           []ThirdParty `gorm:"foreignKey:ThirdPartiesFromDataBreachId"`
    MatterId         *uint
    Matter           *Matter `gorm:"foreignKey:MatterId"`
    Severity                      BreachSeverity
    Status                      IncidentStatus

// parent associations as their child

}

