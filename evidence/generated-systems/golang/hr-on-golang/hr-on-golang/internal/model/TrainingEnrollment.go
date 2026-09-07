package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// TrainingEnrollment Declaration
//==============================================================
type TrainingEnrollment struct {
    gorm.Model
     EnrollmentNumber                                    string
    CompletionDate                                                            time.Time
    Score                                                            string
    CourseId         *uint
    Course           *TrainingCourse `gorm:"foreignKey:CourseId"`
    EmployeeId         *uint
    Employee           *Employee `gorm:"foreignKey:EmployeeId"`
    InstructorId         *uint
    Instructor           *Employee `gorm:"foreignKey:InstructorId"`
    Status                      TrainingStatus

// parent associations as their child

}

