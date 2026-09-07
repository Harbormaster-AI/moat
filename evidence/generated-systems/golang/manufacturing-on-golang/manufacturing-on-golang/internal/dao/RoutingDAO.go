package dao

import (
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing RoutingDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateRouting - creates a new db entry
//----------------------------------------------------------------------------
func CreateRouting(obj model.Routing)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Routing with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Routing", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateRouting", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetRouting - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetRouting(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Routing

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Routing with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Routing using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Routing using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetRouting", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllRouting - returns all
//----------------------------------------------------------------------------
func GetAllRouting()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Routing

	//----------------------------------------------------------------------------
	// Request the ORM to find all Routing
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Routing" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Routing", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllRouting", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateRouting - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateRouting(obj model.Routing)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Routing using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Routing using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateRouting", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteRouting - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteRouting(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Routing with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetRouting(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Routing so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Routing)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Routing using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Routing using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteRouting", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Item on a Routing
//----------------------------------------------------------------------------
func AssignItemToRouting( routingId uint64, itemId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Routing with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRouting(routingId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Routing so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Routing)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Item

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Item with a
		// matching itemId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, itemId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Item	to the Routing
			//----------------------------------------------------------------------------
			parentObj.Item = &childObj

			//----------------------------------------------------------------------------
			// save the Routing
			//----------------------------------------------------------------------------
			return UpdateRouting(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Item", itemId )
			return utils.RequestResult{false, msg, "assignItem", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Item on a Routing
//----------------------------------------------------------------------------
func UnassignItemFromRouting(routingId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Routing with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRouting(routingId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Routing so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Routing)

		//----------------------------------------------------------------------------
		// assign an empty Item to the Item
		//----------------------------------------------------------------------------
		parentObj.Item = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Item
		//----------------------------------------------------------------------------
		parentObj.ItemId = nil;

		//----------------------------------------------------------------------------
		// save the Routing
		//----------------------------------------------------------------------------
		return UpdateRouting(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more operationsIds as a Operations to a Routing
//----------------------------------------------------------------------------
func AddOperationsToRouting ( routingId uint64, operationsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Routing with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRouting(routingId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Routing so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Routing)

		// slice the ids on comma with no spaces
		ids := strings.Split( operationsIds, ",")

		for _, operationsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Operation

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Operation
			// with a matching operationsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , operationsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Operations using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Operations").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Operations", operationsId )
				return utils.RequestResult{false, msg, "unassignOperations", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Routing from the gorm
		//----------------------------------------------------------------------------
		return GetRouting(routingId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more operationsIds as a Operations from a Routing
//----------------------------------------------------------------------------
func RemoveOperationsFromRouting( routingId uint64, operationsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Routing with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRouting(routingId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Routing so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Routing)

		// slice the ids on comma with no spaces
		ids := strings.Split( operationsIds, ",")

		for _, operationsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Operation

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Operation
			// with a matching operationsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , operationsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove OperationObj from the Operations array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Operations").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Operations", operationsId )
				return utils.RequestResult{false, msg, "removeOperations", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Routing from the gorm
		//----------------------------------------------------------------------------
		return GetRouting(routingId)

	} else {
		return parentRequestResult
	}
}

