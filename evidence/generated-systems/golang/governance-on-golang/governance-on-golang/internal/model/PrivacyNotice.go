package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// PrivacyNotice Declaration
//==============================================================
type PrivacyNotice struct {
    gorm.Model
     Title                                    string
    Audience                                    string
    VersionLabel                                    string
    PublicationDate                                                            time.Time
    PublicationUrl                                                            string
     ProcessingActivities           []DataProcessingActivity `gorm:"foreignKey:ProcessingActivitiesFromPrivacyNoticeId"`
    OrganizationId         *uint
    Organization           *Organization `gorm:"foreignKey:OrganizationId"`
     Consents           []Consent `gorm:"foreignKey:ConsentsFromPrivacyNoticeId"`
    Status                      DocumentStatus

// parent associations as their child

}

