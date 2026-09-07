package model

import (
    "gorm.io/gorm"
)

//==============================================================
// ContentCategory Declaration
//==============================================================
type ContentCategory struct {
    gorm.Model
     Code                                    string
    Name                                    string

// parent associations as their child

}

