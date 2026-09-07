package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Document Declaration
//==============================================================
type Document struct {
    gorm.Model
     Name                                    string
    FileUrl                                    string
    UploadedDate                                                            time.Time
    CandidateId         *uint
    Candidate           *Candidate `gorm:"foreignKey:CandidateId"`
    EmployeeId         *uint
    Employee           *Employee `gorm:"foreignKey:EmployeeId"`
    DocumentType                      DocumentType

// parent associations as their child

}

