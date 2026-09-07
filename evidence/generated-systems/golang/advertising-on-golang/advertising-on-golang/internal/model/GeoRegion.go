package model

import (
    "gorm.io/gorm"
)

//==============================================================
// GeoRegion Declaration
//==============================================================
type GeoRegion struct {
    gorm.Model
     Code                                    string
    Name                                    string
    ParentId         *uint
    Parent           *GeoRegion `gorm:"foreignKey:ParentId"`
     Children           []GeoRegion `gorm:"foreignKey:ChildrenFromGeoRegionId"`
    RegionType                      GeoRegionType

// parent associations as their child

}

