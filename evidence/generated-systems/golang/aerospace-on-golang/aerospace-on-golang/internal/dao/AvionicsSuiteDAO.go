package dao

import (
    "aerospace-on-golang/internal/model"
    "aerospace-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing AvionicsSuiteDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateAvionicsSuite - creates a new db entry
//----------------------------------------------------------------------------
func CreateAvionicsSuite(obj model.AvionicsSuite)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a AvionicsSuite with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a AvionicsSuite", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateAvionicsSuite", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetAvionicsSuite - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetAvionicsSuite(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.AvionicsSuite

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a AvionicsSuite with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a AvionicsSuite using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a AvionicsSuite using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetAvionicsSuite", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllAvionicsSuite - returns all
//----------------------------------------------------------------------------
func GetAllAvionicsSuite()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.AvionicsSuite

	//----------------------------------------------------------------------------
	// Request the ORM to find all AvionicsSuite
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all AvionicsSuite" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all AvionicsSuite", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllAvionicsSuite", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateAvionicsSuite - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateAvionicsSuite(obj model.AvionicsSuite)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a AvionicsSuite using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a AvionicsSuite using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateAvionicsSuite", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteAvionicsSuite - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteAvionicsSuite(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the AvionicsSuite with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetAvionicsSuite(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AvionicsSuite so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.AvionicsSuite)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a AvionicsSuite using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a AvionicsSuite using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteAvionicsSuite", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Supplier on a AvionicsSuite
//----------------------------------------------------------------------------
func AssignSupplierToAvionicsSuite( avionicsSuiteId uint64, supplierId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the AvionicsSuite with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAvionicsSuite(avionicsSuiteId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AvionicsSuite so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AvionicsSuite)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Supplier

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Supplier with a
		// matching supplierId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, supplierId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Supplier	to the AvionicsSuite
			//----------------------------------------------------------------------------
			parentObj.Supplier = &childObj

			//----------------------------------------------------------------------------
			// save the AvionicsSuite
			//----------------------------------------------------------------------------
			return UpdateAvionicsSuite(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Supplier", supplierId )
			return utils.RequestResult{false, msg, "assignSupplier", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Supplier on a AvionicsSuite
//----------------------------------------------------------------------------
func UnassignSupplierFromAvionicsSuite(avionicsSuiteId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AvionicsSuite with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAvionicsSuite(avionicsSuiteId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AvionicsSuite so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AvionicsSuite)

		//----------------------------------------------------------------------------
		// assign an empty Supplier to the Supplier
		//----------------------------------------------------------------------------
		parentObj.Supplier = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Supplier
		//----------------------------------------------------------------------------
		parentObj.SupplierId = nil;

		//----------------------------------------------------------------------------
		// save the AvionicsSuite
		//----------------------------------------------------------------------------
		return UpdateAvionicsSuite(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more variantsIds as a Variants to a AvionicsSuite
//----------------------------------------------------------------------------
func AddVariantsToAvionicsSuite ( avionicsSuiteId uint64, variantsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AvionicsSuite with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAvionicsSuite(avionicsSuiteId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AvionicsSuite so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AvionicsSuite)

		// slice the ids on comma with no spaces
		ids := strings.Split( variantsIds, ",")

		for _, variantsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.AircraftVariant

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a AircraftVariant
			// with a matching variantsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , variantsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Variants using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Variants").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Variants", variantsId )
				return utils.RequestResult{false, msg, "unassignVariants", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified AvionicsSuite from the gorm
		//----------------------------------------------------------------------------
		return GetAvionicsSuite(avionicsSuiteId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more variantsIds as a Variants from a AvionicsSuite
//----------------------------------------------------------------------------
func RemoveVariantsFromAvionicsSuite( avionicsSuiteId uint64, variantsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the AvionicsSuite with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAvionicsSuite(avionicsSuiteId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AvionicsSuite so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AvionicsSuite)

		// slice the ids on comma with no spaces
		ids := strings.Split( variantsIds, ",")

		for _, variantsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.AircraftVariant

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a AircraftVariant
			// with a matching variantsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , variantsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove AircraftVariantObj from the Variants array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Variants").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Variants", variantsId )
				return utils.RequestResult{false, msg, "removeVariants", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified AvionicsSuite from the gorm
		//----------------------------------------------------------------------------
		return GetAvionicsSuite(avionicsSuiteId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more softwareLoadsIds as a SoftwareLoads to a AvionicsSuite
//----------------------------------------------------------------------------
func AddSoftwareLoadsToAvionicsSuite ( avionicsSuiteId uint64, softwareLoadsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AvionicsSuite with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAvionicsSuite(avionicsSuiteId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AvionicsSuite so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AvionicsSuite)

		// slice the ids on comma with no spaces
		ids := strings.Split( softwareLoadsIds, ",")

		for _, softwareLoadsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.SoftwareLoad

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a SoftwareLoad
			// with a matching softwareLoadsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , softwareLoadsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the SoftwareLoads using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("SoftwareLoads").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "SoftwareLoads", softwareLoadsId )
				return utils.RequestResult{false, msg, "unassignSoftwareLoads", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified AvionicsSuite from the gorm
		//----------------------------------------------------------------------------
		return GetAvionicsSuite(avionicsSuiteId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more softwareLoadsIds as a SoftwareLoads from a AvionicsSuite
//----------------------------------------------------------------------------
func RemoveSoftwareLoadsFromAvionicsSuite( avionicsSuiteId uint64, softwareLoadsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the AvionicsSuite with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAvionicsSuite(avionicsSuiteId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AvionicsSuite so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AvionicsSuite)

		// slice the ids on comma with no spaces
		ids := strings.Split( softwareLoadsIds, ",")

		for _, softwareLoadsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.SoftwareLoad

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a SoftwareLoad
			// with a matching softwareLoadsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , softwareLoadsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove SoftwareLoadObj from the SoftwareLoads array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("SoftwareLoads").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "SoftwareLoads", softwareLoadsId )
				return utils.RequestResult{false, msg, "removeSoftwareLoads", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified AvionicsSuite from the gorm
		//----------------------------------------------------------------------------
		return GetAvionicsSuite(avionicsSuiteId)

	} else {
		return parentRequestResult
	}
}

