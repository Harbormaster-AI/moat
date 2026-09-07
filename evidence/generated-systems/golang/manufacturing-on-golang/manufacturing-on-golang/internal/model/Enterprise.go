package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Enterprise Declaration
//==============================================================
type Enterprise struct {
    gorm.Model
     Name                                    string
    LegalName                                    string
    RegistrationCountry                                    string
    Website                                    string
    TaxId                                    string
     BusinessUnits           []BusinessUnit `gorm:"foreignKey:BusinessUnitsFromEnterpriseId"`
     Plants           []Plant `gorm:"foreignKey:PlantsFromEnterpriseId"`
     Suppliers           []Supplier `gorm:"foreignKey:SuppliersFromEnterpriseId"`
     Customers           []Customer `gorm:"foreignKey:CustomersFromEnterpriseId"`

// parent associations as their child

}

