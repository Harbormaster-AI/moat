package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Consent Declaration
//==============================================================
type Consent struct {
    gorm.Model
     SubjectIdentifier                                    string
    CaptureDate                                                            time.Time
    ExpiryDate                                                            time.Time
     ProcessingActivities           []DataProcessingActivity `gorm:"foreignKey:ProcessingActivitiesFromConsentId"`
    PrivacyNoticeId         *uint
    PrivacyNotice           *PrivacyNotice `gorm:"foreignKey:PrivacyNoticeId"`
    ConsentType                      ConsentType
    Status                      ConsentStatus

// parent associations as their child

}

