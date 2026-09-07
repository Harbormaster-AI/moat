package dao

import (
    "aerospace-on-golang/internal/model"
    "aerospace-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing AircraftDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateAircraft - creates a new db entry
//----------------------------------------------------------------------------
func CreateAircraft(obj model.Aircraft)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Aircraft with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Aircraft", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateAircraft", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetAircraft - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetAircraft(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Aircraft

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Aircraft with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Aircraft using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Aircraft using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetAircraft", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllAircraft - returns all
//----------------------------------------------------------------------------
func GetAllAircraft()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Aircraft

	//----------------------------------------------------------------------------
	// Request the ORM to find all Aircraft
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Aircraft" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Aircraft", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllAircraft", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateAircraft - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateAircraft(obj model.Aircraft)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Aircraft using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Aircraft using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateAircraft", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteAircraft - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteAircraft(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Aircraft with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetAircraft(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Aircraft so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Aircraft)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Aircraft using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Aircraft using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteAircraft", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Variant on a Aircraft
//----------------------------------------------------------------------------
func AssignVariantToAircraft( aircraftId uint64, variantId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Aircraft with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAircraft(aircraftId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Aircraft so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Aircraft)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.AircraftVariant

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a AircraftVariant with a
		// matching variantId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, variantId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Variant	to the Aircraft
			//----------------------------------------------------------------------------
			parentObj.Variant = &childObj

			//----------------------------------------------------------------------------
			// save the Aircraft
			//----------------------------------------------------------------------------
			return UpdateAircraft(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Variant", variantId )
			return utils.RequestResult{false, msg, "assignVariant", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Variant on a Aircraft
//----------------------------------------------------------------------------
func UnassignVariantFromAircraft(aircraftId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Aircraft with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAircraft(aircraftId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Aircraft so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Aircraft)

		//----------------------------------------------------------------------------
		// assign an empty AircraftVariant to the Variant
		//----------------------------------------------------------------------------
		parentObj.Variant = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Variant
		//----------------------------------------------------------------------------
		parentObj.VariantId = nil;

		//----------------------------------------------------------------------------
		// save the Aircraft
		//----------------------------------------------------------------------------
		return UpdateAircraft(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Operator on a Aircraft
//----------------------------------------------------------------------------
func AssignOperatorToAircraft( aircraftId uint64, operatorId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Aircraft with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAircraft(aircraftId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Aircraft so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Aircraft)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Operator

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Operator with a
		// matching operatorId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, operatorId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Operator	to the Aircraft
			//----------------------------------------------------------------------------
			parentObj.Operator = &childObj

			//----------------------------------------------------------------------------
			// save the Aircraft
			//----------------------------------------------------------------------------
			return UpdateAircraft(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Operator", operatorId )
			return utils.RequestResult{false, msg, "assignOperator", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Operator on a Aircraft
//----------------------------------------------------------------------------
func UnassignOperatorFromAircraft(aircraftId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Aircraft with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAircraft(aircraftId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Aircraft so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Aircraft)

		//----------------------------------------------------------------------------
		// assign an empty Operator to the Operator
		//----------------------------------------------------------------------------
		parentObj.Operator = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Operator
		//----------------------------------------------------------------------------
		parentObj.OperatorId = nil;

		//----------------------------------------------------------------------------
		// save the Aircraft
		//----------------------------------------------------------------------------
		return UpdateAircraft(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Registration on a Aircraft
//----------------------------------------------------------------------------
func AssignRegistrationToAircraft( aircraftId uint64, registrationId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Aircraft with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAircraft(aircraftId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Aircraft so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Aircraft)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Registration

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Registration with a
		// matching registrationId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, registrationId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Registration	to the Aircraft
			//----------------------------------------------------------------------------
			parentObj.Registration = &childObj

			//----------------------------------------------------------------------------
			// save the Aircraft
			//----------------------------------------------------------------------------
			return UpdateAircraft(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Registration", registrationId )
			return utils.RequestResult{false, msg, "assignRegistration", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Registration on a Aircraft
//----------------------------------------------------------------------------
func UnassignRegistrationFromAircraft(aircraftId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Aircraft with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAircraft(aircraftId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Aircraft so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Aircraft)

		//----------------------------------------------------------------------------
		// assign an empty Registration to the Registration
		//----------------------------------------------------------------------------
		parentObj.Registration = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Registration
		//----------------------------------------------------------------------------
		parentObj.RegistrationId = nil;

		//----------------------------------------------------------------------------
		// save the Aircraft
		//----------------------------------------------------------------------------
		return UpdateAircraft(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Warranty on a Aircraft
//----------------------------------------------------------------------------
func AssignWarrantyToAircraft( aircraftId uint64, warrantyId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Aircraft with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAircraft(aircraftId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Aircraft so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Aircraft)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Warranty

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Warranty with a
		// matching warrantyId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, warrantyId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Warranty	to the Aircraft
			//----------------------------------------------------------------------------
			parentObj.Warranty = &childObj

			//----------------------------------------------------------------------------
			// save the Aircraft
			//----------------------------------------------------------------------------
			return UpdateAircraft(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Warranty", warrantyId )
			return utils.RequestResult{false, msg, "assignWarranty", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Warranty on a Aircraft
//----------------------------------------------------------------------------
func UnassignWarrantyFromAircraft(aircraftId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Aircraft with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAircraft(aircraftId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Aircraft so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Aircraft)

		//----------------------------------------------------------------------------
		// assign an empty Warranty to the Warranty
		//----------------------------------------------------------------------------
		parentObj.Warranty = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Warranty
		//----------------------------------------------------------------------------
		parentObj.WarrantyId = nil;

		//----------------------------------------------------------------------------
		// save the Aircraft
		//----------------------------------------------------------------------------
		return UpdateAircraft(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a ConnectedAircraft on a Aircraft
//----------------------------------------------------------------------------
func AssignConnectedAircraftToAircraft( aircraftId uint64, connectedAircraftId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Aircraft with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAircraft(aircraftId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Aircraft so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Aircraft)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.ConnectedAircraft

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a ConnectedAircraft with a
		// matching connectedAircraftId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, connectedAircraftId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the ConnectedAircraft	to the Aircraft
			//----------------------------------------------------------------------------
			parentObj.ConnectedAircraft = &childObj

			//----------------------------------------------------------------------------
			// save the Aircraft
			//----------------------------------------------------------------------------
			return UpdateAircraft(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ConnectedAircraft", connectedAircraftId )
			return utils.RequestResult{false, msg, "assignConnectedAircraft", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a ConnectedAircraft on a Aircraft
//----------------------------------------------------------------------------
func UnassignConnectedAircraftFromAircraft(aircraftId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Aircraft with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAircraft(aircraftId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Aircraft so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Aircraft)

		//----------------------------------------------------------------------------
		// assign an empty ConnectedAircraft to the ConnectedAircraft
		//----------------------------------------------------------------------------
		parentObj.ConnectedAircraft = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the ConnectedAircraft
		//----------------------------------------------------------------------------
		parentObj.ConnectedAircraftId = nil;

		//----------------------------------------------------------------------------
		// save the Aircraft
		//----------------------------------------------------------------------------
		return UpdateAircraft(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a CabinLayout on a Aircraft
//----------------------------------------------------------------------------
func AssignCabinLayoutToAircraft( aircraftId uint64, cabinLayoutId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Aircraft with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAircraft(aircraftId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Aircraft so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Aircraft)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.CabinLayout

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a CabinLayout with a
		// matching cabinLayoutId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, cabinLayoutId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the CabinLayout	to the Aircraft
			//----------------------------------------------------------------------------
			parentObj.CabinLayout = &childObj

			//----------------------------------------------------------------------------
			// save the Aircraft
			//----------------------------------------------------------------------------
			return UpdateAircraft(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "CabinLayout", cabinLayoutId )
			return utils.RequestResult{false, msg, "assignCabinLayout", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a CabinLayout on a Aircraft
//----------------------------------------------------------------------------
func UnassignCabinLayoutFromAircraft(aircraftId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Aircraft with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAircraft(aircraftId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Aircraft so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Aircraft)

		//----------------------------------------------------------------------------
		// assign an empty CabinLayout to the CabinLayout
		//----------------------------------------------------------------------------
		parentObj.CabinLayout = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the CabinLayout
		//----------------------------------------------------------------------------
		parentObj.CabinLayoutId = nil;

		//----------------------------------------------------------------------------
		// save the Aircraft
		//----------------------------------------------------------------------------
		return UpdateAircraft(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more maintenanceRecordsIds as a MaintenanceRecords to a Aircraft
//----------------------------------------------------------------------------
func AddMaintenanceRecordsToAircraft ( aircraftId uint64, maintenanceRecordsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Aircraft with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAircraft(aircraftId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Aircraft so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Aircraft)

		// slice the ids on comma with no spaces
		ids := strings.Split( maintenanceRecordsIds, ",")

		for _, maintenanceRecordsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.MaintenanceWorkOrder

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a MaintenanceWorkOrder
			// with a matching maintenanceRecordsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , maintenanceRecordsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the MaintenanceRecords using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("MaintenanceRecords").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "MaintenanceRecords", maintenanceRecordsId )
				return utils.RequestResult{false, msg, "unassignMaintenanceRecords", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Aircraft from the gorm
		//----------------------------------------------------------------------------
		return GetAircraft(aircraftId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more maintenanceRecordsIds as a MaintenanceRecords from a Aircraft
//----------------------------------------------------------------------------
func RemoveMaintenanceRecordsFromAircraft( aircraftId uint64, maintenanceRecordsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Aircraft with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAircraft(aircraftId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Aircraft so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Aircraft)

		// slice the ids on comma with no spaces
		ids := strings.Split( maintenanceRecordsIds, ",")

		for _, maintenanceRecordsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.MaintenanceWorkOrder

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a MaintenanceWorkOrder
			// with a matching maintenanceRecordsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , maintenanceRecordsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove MaintenanceWorkOrderObj from the MaintenanceRecords array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("MaintenanceRecords").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "MaintenanceRecords", maintenanceRecordsId )
				return utils.RequestResult{false, msg, "removeMaintenanceRecords", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Aircraft from the gorm
		//----------------------------------------------------------------------------
		return GetAircraft(aircraftId)

	} else {
		return parentRequestResult
	}
}

