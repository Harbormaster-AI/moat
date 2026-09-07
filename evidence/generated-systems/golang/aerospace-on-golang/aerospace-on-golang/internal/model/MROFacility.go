package model

import (
    "gorm.io/gorm"
)

//==============================================================
// MROFacility Declaration
//==============================================================
type MROFacility struct {
    gorm.Model
     Name                                    string
    ApprovalScope                                    string
    Address                                                            string
     Appointments           []MaintenanceAppointment `gorm:"foreignKey:AppointmentsFromMROFacilityId"`
     WorkOrders           []MaintenanceWorkOrder `gorm:"foreignKey:WorkOrdersFromMROFacilityId"`

// parent associations as their child

}

