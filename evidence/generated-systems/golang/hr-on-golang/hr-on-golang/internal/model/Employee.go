package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Employee Declaration
//==============================================================
type Employee struct {
    gorm.Model
     EmployeeNumber                                    string
    Name                                                            string
    WorkEmail                                                            string
    WorkPhone                                                            string
    DateOfHire                                                            time.Time
    NationalId                                                            string
    ManagerId         *uint
    Manager           *Employee `gorm:"foreignKey:ManagerId"`
     DirectReports           []Employee `gorm:"foreignKey:DirectReportsFromEmployeeId"`
    DepartmentId         *uint
    Department           *Department `gorm:"foreignKey:DepartmentId"`
    PrimaryLocationId         *uint
    PrimaryLocation           *Location `gorm:"foreignKey:PrimaryLocationId"`
    CostCenterId         *uint
    CostCenter           *CostCenter `gorm:"foreignKey:CostCenterId"`
     EmploymentAssignments           []EmploymentAssignment `gorm:"foreignKey:EmploymentAssignmentsFromEmployeeId"`
     Contracts           []EmploymentContract `gorm:"foreignKey:ContractsFromEmployeeId"`
     BenefitEnrollments           []BenefitEnrollment `gorm:"foreignKey:BenefitEnrollmentsFromEmployeeId"`
     Timesheets           []Timesheet `gorm:"foreignKey:TimesheetsFromEmployeeId"`
     LeaveRequests           []LeaveRequest `gorm:"foreignKey:LeaveRequestsFromEmployeeId"`
     PerformanceReviews           []PerformanceReview `gorm:"foreignKey:PerformanceReviewsFromEmployeeId"`
     TrainingEnrollments           []TrainingEnrollment `gorm:"foreignKey:TrainingEnrollmentsFromEmployeeId"`
     WorkAuthorizations           []WorkAuthorization `gorm:"foreignKey:WorkAuthorizationsFromEmployeeId"`
    Status                      EmploymentStatus

// parent associations as their child

}

