package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// DataProcessingActivity Declaration
//==============================================================
type DataProcessingActivity struct {
    gorm.Model
     Name                                    string
    Purpose                                    string
    StartDate                                                            time.Time
    OrganizationId         *uint
    Organization           *Organization `gorm:"foreignKey:OrganizationId"`
     DataCategories           []DataCategory `gorm:"foreignKey:DataCategoriesFromDataProcessingActivityId"`
     Systems           []System_ `gorm:"foreignKey:SystemsFromDataProcessingActivityId"`
     Records           []Record_ `gorm:"foreignKey:RecordsFromDataProcessingActivityId"`
     PrivacyNotices           []PrivacyNotice `gorm:"foreignKey:PrivacyNoticesFromDataProcessingActivityId"`
     ThirdParties           []ThirdParty `gorm:"foreignKey:ThirdPartiesFromDataProcessingActivityId"`
     Consents           []Consent `gorm:"foreignKey:ConsentsFromDataProcessingActivityId"`
     DataBreaches           []DataBreach `gorm:"foreignKey:DataBreachesFromDataProcessingActivityId"`
     DataSubjectRequests           []DataSubjectRequest `gorm:"foreignKey:DataSubjectRequestsFromDataProcessingActivityId"`
    LawfulBasis                      LawfulBasis

// parent associations as their child

}

