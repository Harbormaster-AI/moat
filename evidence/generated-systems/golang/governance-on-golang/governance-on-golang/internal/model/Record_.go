package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Record_ Declaration
//==============================================================
type Record_ struct {
    gorm.Model
     Title                                    string
    CreationDate                                                            time.Time
    RepositoryId         *uint
    Repository           *RecordsRepository `gorm:"foreignKey:RepositoryId"`
    RetentionScheduleId         *uint
    RetentionSchedule           *RetentionSchedule `gorm:"foreignKey:RetentionScheduleId"`
     ProcessingActivities           []DataProcessingActivity `gorm:"foreignKey:ProcessingActivitiesFromRecord_Id"`
     DataCategories           []DataCategory `gorm:"foreignKey:DataCategoriesFromRecord_Id"`
     LegalHolds           []LegalHold `gorm:"foreignKey:LegalHoldsFromRecord_Id"`
     DataSubjectRequests           []DataSubjectRequest `gorm:"foreignKey:DataSubjectRequestsFromRecord_Id"`
    RecordType                      RecordType
    Classification                      DataClassificationLevel
    Status                      RecordStatus

// parent associations as their child

}

