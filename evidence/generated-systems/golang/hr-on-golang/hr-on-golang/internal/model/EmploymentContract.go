package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// EmploymentContract Declaration
//==============================================================
type EmploymentContract struct {
    gorm.Model
     ContractNumber                                    string
    StartDate                                                            time.Time
    EndDate                                                            time.Time
    WorkHoursPerWeek                                                            string
    EmployeeId         *uint
    Employee           *Employee `gorm:"foreignKey:EmployeeId"`
    CompensationPackageId         *uint
    CompensationPackage           *CompensationPackage `gorm:"foreignKey:CompensationPackageId"`
    WorkScheduleId         *uint
    WorkSchedule           *WorkSchedule `gorm:"foreignKey:WorkScheduleId"`
    LocationId         *uint
    Location           *Location `gorm:"foreignKey:LocationId"`
    PayrollCalendarId         *uint
    PayrollCalendar           *PayrollCalendar `gorm:"foreignKey:PayrollCalendarId"`
    EmploymentType                      EmploymentType
    Status                      ContractStatus
    PayFrequency                      PayFrequency

// parent associations as their child

}

