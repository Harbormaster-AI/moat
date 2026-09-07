package model

import (
    "gorm.io/gorm"
)

//==============================================================
// ImagingOrder Declaration
//==============================================================
type ImagingOrder struct {
    gorm.Model
     BodySite                                    string
    Contrast                                    bool
    OrderId         *uint
    Order           *ClinicalOrder `gorm:"foreignKey:OrderId"`
    ImagingCenterId         *uint
    ImagingCenter           *ImagingCenter `gorm:"foreignKey:ImagingCenterId"`
     Reports           []ImagingReport `gorm:"foreignKey:ReportsFromImagingOrderId"`
    Modality                      ImagingModality

// parent associations as their child

}

