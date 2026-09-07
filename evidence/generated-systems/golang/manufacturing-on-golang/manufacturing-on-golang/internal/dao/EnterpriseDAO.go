package dao

import (
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing EnterpriseDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateEnterprise - creates a new db entry
//----------------------------------------------------------------------------
func CreateEnterprise(obj model.Enterprise)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var createMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	result := utils.GetDB().Create(&obj).Error

	if result == nil {
	    createMsg = fmt.Sprintf( "Created a Enterprise with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Enterprise", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateEnterprise", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetEnterprise - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetEnterprise(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Enterprise

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Enterprise with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Enterprise using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Enterprise using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetEnterprise", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllEnterprise - returns all
//----------------------------------------------------------------------------
func GetAllEnterprise()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Enterprise

	//----------------------------------------------------------------------------
	// Request the ORM to find all Enterprise
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Enterprise" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Enterprise", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllEnterprise", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateEnterprise - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateEnterprise(obj model.Enterprise)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var updateMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to save
	//----------------------------------------------------------------------------
	result := utils.GetDB().Save(&obj).Error

	if result == nil {
	    updateMsg = fmt.Sprintf( "Updated a Enterprise using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Enterprise using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateEnterprise", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteEnterprise - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteEnterprise(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Enterprise with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetEnterprise(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Enterprise so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Enterprise)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Enterprise using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Enterprise using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteEnterprise", requestResult.Data}

	}

	return requestResult
}



//----------------------------------------------------------------------------
// adds one or more businessUnitsIds as a BusinessUnits to a Enterprise
//----------------------------------------------------------------------------
func AddBusinessUnitsToEnterprise ( enterpriseId uint64, businessUnitsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Enterprise with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEnterprise(enterpriseId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Enterprise so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Enterprise)

		// slice the ids on comma with no spaces
		ids := strings.Split( businessUnitsIds, ",")

		for _, businessUnitsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.BusinessUnit

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a BusinessUnit
			// with a matching businessUnitsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , businessUnitsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the BusinessUnits using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("BusinessUnits").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "BusinessUnits", businessUnitsId )
				return utils.RequestResult{false, msg, "unassignBusinessUnits", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Enterprise from the gorm
		//----------------------------------------------------------------------------
		return GetEnterprise(enterpriseId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more businessUnitsIds as a BusinessUnits from a Enterprise
//----------------------------------------------------------------------------
func RemoveBusinessUnitsFromEnterprise( enterpriseId uint64, businessUnitsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Enterprise with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEnterprise(enterpriseId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Enterprise so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Enterprise)

		// slice the ids on comma with no spaces
		ids := strings.Split( businessUnitsIds, ",")

		for _, businessUnitsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.BusinessUnit

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a BusinessUnit
			// with a matching businessUnitsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , businessUnitsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove BusinessUnitObj from the BusinessUnits array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("BusinessUnits").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "BusinessUnits", businessUnitsId )
				return utils.RequestResult{false, msg, "removeBusinessUnits", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Enterprise from the gorm
		//----------------------------------------------------------------------------
		return GetEnterprise(enterpriseId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more plantsIds as a Plants to a Enterprise
//----------------------------------------------------------------------------
func AddPlantsToEnterprise ( enterpriseId uint64, plantsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Enterprise with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEnterprise(enterpriseId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Enterprise so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Enterprise)

		// slice the ids on comma with no spaces
		ids := strings.Split( plantsIds, ",")

		for _, plantsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Plant

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Plant
			// with a matching plantsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , plantsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Plants using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Plants").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Plants", plantsId )
				return utils.RequestResult{false, msg, "unassignPlants", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Enterprise from the gorm
		//----------------------------------------------------------------------------
		return GetEnterprise(enterpriseId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more plantsIds as a Plants from a Enterprise
//----------------------------------------------------------------------------
func RemovePlantsFromEnterprise( enterpriseId uint64, plantsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Enterprise with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEnterprise(enterpriseId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Enterprise so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Enterprise)

		// slice the ids on comma with no spaces
		ids := strings.Split( plantsIds, ",")

		for _, plantsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Plant

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Plant
			// with a matching plantsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , plantsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove PlantObj from the Plants array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Plants").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Plants", plantsId )
				return utils.RequestResult{false, msg, "removePlants", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Enterprise from the gorm
		//----------------------------------------------------------------------------
		return GetEnterprise(enterpriseId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more suppliersIds as a Suppliers to a Enterprise
//----------------------------------------------------------------------------
func AddSuppliersToEnterprise ( enterpriseId uint64, suppliersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Enterprise with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEnterprise(enterpriseId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Enterprise so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Enterprise)

		// slice the ids on comma with no spaces
		ids := strings.Split( suppliersIds, ",")

		for _, suppliersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Supplier

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Supplier
			// with a matching suppliersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , suppliersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Suppliers using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Suppliers").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Suppliers", suppliersId )
				return utils.RequestResult{false, msg, "unassignSuppliers", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Enterprise from the gorm
		//----------------------------------------------------------------------------
		return GetEnterprise(enterpriseId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more suppliersIds as a Suppliers from a Enterprise
//----------------------------------------------------------------------------
func RemoveSuppliersFromEnterprise( enterpriseId uint64, suppliersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Enterprise with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEnterprise(enterpriseId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Enterprise so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Enterprise)

		// slice the ids on comma with no spaces
		ids := strings.Split( suppliersIds, ",")

		for _, suppliersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Supplier

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Supplier
			// with a matching suppliersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , suppliersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove SupplierObj from the Suppliers array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Suppliers").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Suppliers", suppliersId )
				return utils.RequestResult{false, msg, "removeSuppliers", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Enterprise from the gorm
		//----------------------------------------------------------------------------
		return GetEnterprise(enterpriseId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more customersIds as a Customers to a Enterprise
//----------------------------------------------------------------------------
func AddCustomersToEnterprise ( enterpriseId uint64, customersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Enterprise with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEnterprise(enterpriseId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Enterprise so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Enterprise)

		// slice the ids on comma with no spaces
		ids := strings.Split( customersIds, ",")

		for _, customersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Customer

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Customer
			// with a matching customersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , customersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Customers using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Customers").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Customers", customersId )
				return utils.RequestResult{false, msg, "unassignCustomers", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Enterprise from the gorm
		//----------------------------------------------------------------------------
		return GetEnterprise(enterpriseId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more customersIds as a Customers from a Enterprise
//----------------------------------------------------------------------------
func RemoveCustomersFromEnterprise( enterpriseId uint64, customersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Enterprise with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEnterprise(enterpriseId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Enterprise so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Enterprise)

		// slice the ids on comma with no spaces
		ids := strings.Split( customersIds, ",")

		for _, customersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Customer

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Customer
			// with a matching customersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , customersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove CustomerObj from the Customers array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Customers").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Customers", customersId )
				return utils.RequestResult{false, msg, "removeCustomers", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Enterprise from the gorm
		//----------------------------------------------------------------------------
		return GetEnterprise(enterpriseId)

	} else {
		return parentRequestResult
	}
}

