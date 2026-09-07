package model

import (
    "gorm.io/gorm"
)

//==============================================================
// TrainingCourse Declaration
//==============================================================
type TrainingCourse struct {
    gorm.Model
     Code                                    string
    Title                                    string
    DurationHours                                                            string
     Prerequisites           []TrainingCourse `gorm:"foreignKey:PrerequisitesFromTrainingCourseId"`
     Enrollments           []TrainingEnrollment `gorm:"foreignKey:EnrollmentsFromTrainingCourseId"`
     JobProfiles           []JobProfile `gorm:"foreignKey:JobProfilesFromTrainingCourseId"`
    DeliveryMethod                      DeliveryMethod

// parent associations as their child

}

