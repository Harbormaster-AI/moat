package model

import (
    "gorm.io/gorm"
)

//==============================================================
// ProductionLine Declaration
//==============================================================
type ProductionLine struct {
    gorm.Model
     Name                                    string
    PlantId         *uint
    Plant           *Plant `gorm:"foreignKey:PlantId"`
     WorkCenters           []WorkCenter `gorm:"foreignKey:WorkCentersFromProductionLineId"`
    LineType                      ProductionLineType

// parent associations as their child

}

