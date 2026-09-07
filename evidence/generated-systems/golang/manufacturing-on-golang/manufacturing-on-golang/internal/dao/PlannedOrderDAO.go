package dao

import (
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing PlannedOrderDAO..." ) )
}

//----------------------------------------------------------------------------
// CreatePlannedOrder - creates a new db entry
//----------------------------------------------------------------------------
func CreatePlannedOrder(obj model.PlannedOrder)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a PlannedOrder with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a PlannedOrder", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreatePlannedOrder", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetPlannedOrder - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetPlannedOrder(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.PlannedOrder

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a PlannedOrder with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a PlannedOrder using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a PlannedOrder using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetPlannedOrder", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllPlannedOrder - returns all
//----------------------------------------------------------------------------
func GetAllPlannedOrder()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.PlannedOrder

	//----------------------------------------------------------------------------
	// Request the ORM to find all PlannedOrder
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all PlannedOrder" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all PlannedOrder", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllPlannedOrder", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdatePlannedOrder - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdatePlannedOrder(obj model.PlannedOrder)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a PlannedOrder using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a PlannedOrder using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdatePlannedOrder", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeletePlannedOrder - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeletePlannedOrder(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the PlannedOrder with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetPlannedOrder(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PlannedOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.PlannedOrder)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a PlannedOrder using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a PlannedOrder using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeletePlannedOrder", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a MrpRun on a PlannedOrder
//----------------------------------------------------------------------------
func AssignMrpRunToPlannedOrder( plannedOrderId uint64, mrpRunId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the PlannedOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPlannedOrder(plannedOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PlannedOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PlannedOrder)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.MRPRun

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a MRPRun with a
		// matching mrpRunId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, mrpRunId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the MrpRun	to the PlannedOrder
			//----------------------------------------------------------------------------
			parentObj.MrpRun = &childObj

			//----------------------------------------------------------------------------
			// save the PlannedOrder
			//----------------------------------------------------------------------------
			return UpdatePlannedOrder(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "MrpRun", mrpRunId )
			return utils.RequestResult{false, msg, "assignMrpRun", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a MrpRun on a PlannedOrder
//----------------------------------------------------------------------------
func UnassignMrpRunFromPlannedOrder(plannedOrderId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the PlannedOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPlannedOrder(plannedOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PlannedOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PlannedOrder)

		//----------------------------------------------------------------------------
		// assign an empty MRPRun to the MrpRun
		//----------------------------------------------------------------------------
		parentObj.MrpRun = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the MrpRun
		//----------------------------------------------------------------------------
		parentObj.MrpRunId = nil;

		//----------------------------------------------------------------------------
		// save the PlannedOrder
		//----------------------------------------------------------------------------
		return UpdatePlannedOrder(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Item on a PlannedOrder
//----------------------------------------------------------------------------
func AssignItemToPlannedOrder( plannedOrderId uint64, itemId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the PlannedOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPlannedOrder(plannedOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PlannedOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PlannedOrder)

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
			// assign the Item	to the PlannedOrder
			//----------------------------------------------------------------------------
			parentObj.Item = &childObj

			//----------------------------------------------------------------------------
			// save the PlannedOrder
			//----------------------------------------------------------------------------
			return UpdatePlannedOrder(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Item", itemId )
			return utils.RequestResult{false, msg, "assignItem", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Item on a PlannedOrder
//----------------------------------------------------------------------------
func UnassignItemFromPlannedOrder(plannedOrderId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the PlannedOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPlannedOrder(plannedOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PlannedOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PlannedOrder)

		//----------------------------------------------------------------------------
		// assign an empty Item to the Item
		//----------------------------------------------------------------------------
		parentObj.Item = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Item
		//----------------------------------------------------------------------------
		parentObj.ItemId = nil;

		//----------------------------------------------------------------------------
		// save the PlannedOrder
		//----------------------------------------------------------------------------
		return UpdatePlannedOrder(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Plant on a PlannedOrder
//----------------------------------------------------------------------------
func AssignPlantToPlannedOrder( plannedOrderId uint64, plantId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the PlannedOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPlannedOrder(plannedOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PlannedOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PlannedOrder)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Plant

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Plant with a
		// matching plantId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, plantId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Plant	to the PlannedOrder
			//----------------------------------------------------------------------------
			parentObj.Plant = &childObj

			//----------------------------------------------------------------------------
			// save the PlannedOrder
			//----------------------------------------------------------------------------
			return UpdatePlannedOrder(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Plant", plantId )
			return utils.RequestResult{false, msg, "assignPlant", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Plant on a PlannedOrder
//----------------------------------------------------------------------------
func UnassignPlantFromPlannedOrder(plannedOrderId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the PlannedOrder with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPlannedOrder(plannedOrderId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PlannedOrder so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PlannedOrder)

		//----------------------------------------------------------------------------
		// assign an empty Plant to the Plant
		//----------------------------------------------------------------------------
		parentObj.Plant = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Plant
		//----------------------------------------------------------------------------
		parentObj.PlantId = nil;

		//----------------------------------------------------------------------------
		// save the PlannedOrder
		//----------------------------------------------------------------------------
		return UpdatePlannedOrder(parentObj)

	} else {
		return parentRequestResult
	}

}


