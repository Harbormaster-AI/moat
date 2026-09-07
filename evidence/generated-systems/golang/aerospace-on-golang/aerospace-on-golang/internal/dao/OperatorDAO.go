package dao

import (
    "aerospace-on-golang/internal/model"
    "aerospace-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing OperatorDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateOperator - creates a new db entry
//----------------------------------------------------------------------------
func CreateOperator(obj model.Operator)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Operator with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Operator", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateOperator", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetOperator - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetOperator(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Operator

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Operator with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Operator using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Operator using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetOperator", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllOperator - returns all
//----------------------------------------------------------------------------
func GetAllOperator()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Operator

	//----------------------------------------------------------------------------
	// Request the ORM to find all Operator
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Operator" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Operator", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllOperator", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateOperator - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateOperator(obj model.Operator)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Operator using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Operator using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateOperator", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteOperator - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteOperator(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Operator with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetOperator(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Operator so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Operator)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Operator using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Operator using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteOperator", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a SalesRegion on a Operator
//----------------------------------------------------------------------------
func AssignSalesRegionToOperator( operatorId uint64, salesRegionId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Operator with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOperator(operatorId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Operator so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Operator)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.SalesRegion

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a SalesRegion with a
		// matching salesRegionId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, salesRegionId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the SalesRegion	to the Operator
			//----------------------------------------------------------------------------
			parentObj.SalesRegion = &childObj

			//----------------------------------------------------------------------------
			// save the Operator
			//----------------------------------------------------------------------------
			return UpdateOperator(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "SalesRegion", salesRegionId )
			return utils.RequestResult{false, msg, "assignSalesRegion", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a SalesRegion on a Operator
//----------------------------------------------------------------------------
func UnassignSalesRegionFromOperator(operatorId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Operator with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOperator(operatorId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Operator so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Operator)

		//----------------------------------------------------------------------------
		// assign an empty SalesRegion to the SalesRegion
		//----------------------------------------------------------------------------
		parentObj.SalesRegion = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the SalesRegion
		//----------------------------------------------------------------------------
		parentObj.SalesRegionId = nil;

		//----------------------------------------------------------------------------
		// save the Operator
		//----------------------------------------------------------------------------
		return UpdateOperator(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more aircraftOrdersIds as a AircraftOrders to a Operator
//----------------------------------------------------------------------------
func AddAircraftOrdersToOperator ( operatorId uint64, aircraftOrdersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Operator with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOperator(operatorId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Operator so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Operator)

		// slice the ids on comma with no spaces
		ids := strings.Split( aircraftOrdersIds, ",")

		for _, aircraftOrdersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.AircraftOrder

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a AircraftOrder
			// with a matching aircraftOrdersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , aircraftOrdersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the AircraftOrders using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("AircraftOrders").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "AircraftOrders", aircraftOrdersId )
				return utils.RequestResult{false, msg, "unassignAircraftOrders", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Operator from the gorm
		//----------------------------------------------------------------------------
		return GetOperator(operatorId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more aircraftOrdersIds as a AircraftOrders from a Operator
//----------------------------------------------------------------------------
func RemoveAircraftOrdersFromOperator( operatorId uint64, aircraftOrdersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Operator with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOperator(operatorId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Operator so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Operator)

		// slice the ids on comma with no spaces
		ids := strings.Split( aircraftOrdersIds, ",")

		for _, aircraftOrdersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.AircraftOrder

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a AircraftOrder
			// with a matching aircraftOrdersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , aircraftOrdersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove AircraftOrderObj from the AircraftOrders array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("AircraftOrders").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "AircraftOrders", aircraftOrdersId )
				return utils.RequestResult{false, msg, "removeAircraftOrders", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Operator from the gorm
		//----------------------------------------------------------------------------
		return GetOperator(operatorId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more operatedAircraftIds as a OperatedAircraft to a Operator
//----------------------------------------------------------------------------
func AddOperatedAircraftToOperator ( operatorId uint64, operatedAircraftIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Operator with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOperator(operatorId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Operator so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Operator)

		// slice the ids on comma with no spaces
		ids := strings.Split( operatedAircraftIds, ",")

		for _, operatedAircraftId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Aircraft

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Aircraft
			// with a matching operatedAircraftId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , operatedAircraftId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the OperatedAircraft using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("OperatedAircraft").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "OperatedAircraft", operatedAircraftId )
				return utils.RequestResult{false, msg, "unassignOperatedAircraft", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Operator from the gorm
		//----------------------------------------------------------------------------
		return GetOperator(operatorId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more operatedAircraftIds as a OperatedAircraft from a Operator
//----------------------------------------------------------------------------
func RemoveOperatedAircraftFromOperator( operatorId uint64, operatedAircraftIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Operator with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetOperator(operatorId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Operator so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Operator)

		// slice the ids on comma with no spaces
		ids := strings.Split( operatedAircraftIds, ",")

		for _, operatedAircraftId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Aircraft

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Aircraft
			// with a matching operatedAircraftId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , operatedAircraftId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove AircraftObj from the OperatedAircraft array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("OperatedAircraft").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "OperatedAircraft", operatedAircraftId )
				return utils.RequestResult{false, msg, "removeOperatedAircraft", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Operator from the gorm
		//----------------------------------------------------------------------------
		return GetOperator(operatorId)

	} else {
		return parentRequestResult
	}
}

