package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// MaintenanceAppointment Declaration
//==============================================================
type MaintenanceAppointment struct {
    gorm.Model
     AppointmentDate                                                            time.Time
    AircraftId         *uint
    Aircraft           *Aircraft `gorm:"foreignKey:AircraftId"`
    MroFacilityId         *uint
    MroFacility           *MROFacility `gorm:"foreignKey:MroFacilityId"`
    WorkOrderId         *uint
    WorkOrder           *MaintenanceWorkOrder `gorm:"foreignKey:WorkOrderId"`
    Status                      AppointmentStatus

// parent associations as their child

}

