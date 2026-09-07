package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// DataSubjectRequest Declaration
//==============================================================
type DataSubjectRequest struct {
    gorm.Model
     ReceivedDate                                                            time.Time
    DueDate                                                            time.Time
    RequesterCountry                                    string
    OrganizationId         *uint
    Organization           *Organization `gorm:"foreignKey:OrganizationId"`
     ProcessingActivities           []DataProcessingActivity `gorm:"foreignKey:ProcessingActivitiesFromDataSubjectRequestId"`
     Records           []Record_ `gorm:"foreignKey:RecordsFromDataSubjectRequestId"`
    RequestType                      DataSubjectRequestType
    Status                      RequestStatus

// parent associations as their child

}

