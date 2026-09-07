package model

import (
    "gorm.io/gorm"
)

//==============================================================
// ImagingCenter Declaration
//==============================================================
type ImagingCenter struct {
    gorm.Model
     Name                                    string
    FacilityId         *uint
    Facility           *Facility `gorm:"foreignKey:FacilityId"`
     ImagingOrders           []ImagingOrder `gorm:"foreignKey:ImagingOrdersFromImagingCenterId"`
     ImagingReports           []ImagingReport `gorm:"foreignKey:ImagingReportsFromImagingCenterId"`

// parent associations as their child

}

