package dao

import (
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing NonconformanceDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateNonconformance - creates a new db entry
//----------------------------------------------------------------------------
func CreateNonconformance(obj model.Nonconformance)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Nonconformance with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Nonconformance", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateNonconformance", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetNonconformance - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetNonconformance(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Nonconformance

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Nonconformance with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Nonconformance using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Nonconformance using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetNonconformance", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllNonconformance - returns all
//----------------------------------------------------------------------------
func GetAllNonconformance()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Nonconformance

	//----------------------------------------------------------------------------
	// Request the ORM to find all Nonconformance
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Nonconformance" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Nonconformance", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllNonconformance", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateNonconformance - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateNonconformance(obj model.Nonconformance)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Nonconformance using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Nonconformance using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateNonconformance", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteNonconformance - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteNonconformance(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Nonconformance with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetNonconformance(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Nonconformance so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Nonconformance)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Nonconformance using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Nonconformance using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteNonconformance", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Item on a Nonconformance
//----------------------------------------------------------------------------
func AssignItemToNonconformance( nonconformanceId uint64, itemId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Nonconformance with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetNonconformance(nonconformanceId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Nonconformance so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Nonconformance)

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
			// assign the Item	to the Nonconformance
			//----------------------------------------------------------------------------
			parentObj.Item = &childObj

			//----------------------------------------------------------------------------
			// save the Nonconformance
			//----------------------------------------------------------------------------
			return UpdateNonconformance(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Item", itemId )
			return utils.RequestResult{false, msg, "assignItem", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Item on a Nonconformance
//----------------------------------------------------------------------------
func UnassignItemFromNonconformance(nonconformanceId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Nonconformance with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetNonconformance(nonconformanceId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Nonconformance so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Nonconformance)

		//----------------------------------------------------------------------------
		// assign an empty Item to the Item
		//----------------------------------------------------------------------------
		parentObj.Item = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Item
		//----------------------------------------------------------------------------
		parentObj.ItemId = nil;

		//----------------------------------------------------------------------------
		// save the Nonconformance
		//----------------------------------------------------------------------------
		return UpdateNonconformance(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a WorkOrder on a Nonconformance
//----------------------------------------------------------------------------
func AssignWorkOrderToNonconformance( nonconformanceId uint64, workOrderId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Nonconformance with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetNonconformance(nonconformanceId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Nonconformance so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Nonconformance)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.WorkOrder

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a WorkOrder with a
		// matching workOrderId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, workOrderId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the WorkOrder	to the Nonconformance
			//----------------------------------------------------------------------------
			parentObj.WorkOrder = &childObj

			//----------------------------------------------------------------------------
			// save the Nonconformance
			//----------------------------------------------------------------------------
			return UpdateNonconformance(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "WorkOrder", workOrderId )
			return utils.RequestResult{false, msg, "assignWorkOrder", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a WorkOrder on a Nonconformance
//----------------------------------------------------------------------------
func UnassignWorkOrderFromNonconformance(nonconformanceId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Nonconformance with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetNonconformance(nonconformanceId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Nonconformance so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Nonconformance)

		//----------------------------------------------------------------------------
		// assign an empty WorkOrder to the WorkOrder
		//----------------------------------------------------------------------------
		parentObj.WorkOrder = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the WorkOrder
		//----------------------------------------------------------------------------
		parentObj.WorkOrderId = nil;

		//----------------------------------------------------------------------------
		// save the Nonconformance
		//----------------------------------------------------------------------------
		return UpdateNonconformance(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a InspectionLot on a Nonconformance
//----------------------------------------------------------------------------
func AssignInspectionLotToNonconformance( nonconformanceId uint64, inspectionLotId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Nonconformance with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetNonconformance(nonconformanceId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Nonconformance so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Nonconformance)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.InspectionLot

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a InspectionLot with a
		// matching inspectionLotId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, inspectionLotId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the InspectionLot	to the Nonconformance
			//----------------------------------------------------------------------------
			parentObj.InspectionLot = &childObj

			//----------------------------------------------------------------------------
			// save the Nonconformance
			//----------------------------------------------------------------------------
			return UpdateNonconformance(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "InspectionLot", inspectionLotId )
			return utils.RequestResult{false, msg, "assignInspectionLot", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a InspectionLot on a Nonconformance
//----------------------------------------------------------------------------
func UnassignInspectionLotFromNonconformance(nonconformanceId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Nonconformance with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetNonconformance(nonconformanceId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Nonconformance so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Nonconformance)

		//----------------------------------------------------------------------------
		// assign an empty InspectionLot to the InspectionLot
		//----------------------------------------------------------------------------
		parentObj.InspectionLot = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the InspectionLot
		//----------------------------------------------------------------------------
		parentObj.InspectionLotId = nil;

		//----------------------------------------------------------------------------
		// save the Nonconformance
		//----------------------------------------------------------------------------
		return UpdateNonconformance(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a CorrectiveAction on a Nonconformance
//----------------------------------------------------------------------------
func AssignCorrectiveActionToNonconformance( nonconformanceId uint64, correctiveActionId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Nonconformance with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetNonconformance(nonconformanceId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Nonconformance so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Nonconformance)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.CorrectiveAction

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a CorrectiveAction with a
		// matching correctiveActionId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, correctiveActionId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the CorrectiveAction	to the Nonconformance
			//----------------------------------------------------------------------------
			parentObj.CorrectiveAction = &childObj

			//----------------------------------------------------------------------------
			// save the Nonconformance
			//----------------------------------------------------------------------------
			return UpdateNonconformance(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "CorrectiveAction", correctiveActionId )
			return utils.RequestResult{false, msg, "assignCorrectiveAction", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a CorrectiveAction on a Nonconformance
//----------------------------------------------------------------------------
func UnassignCorrectiveActionFromNonconformance(nonconformanceId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Nonconformance with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetNonconformance(nonconformanceId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Nonconformance so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Nonconformance)

		//----------------------------------------------------------------------------
		// assign an empty CorrectiveAction to the CorrectiveAction
		//----------------------------------------------------------------------------
		parentObj.CorrectiveAction = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the CorrectiveAction
		//----------------------------------------------------------------------------
		parentObj.CorrectiveActionId = nil;

		//----------------------------------------------------------------------------
		// save the Nonconformance
		//----------------------------------------------------------------------------
		return UpdateNonconformance(parentObj)

	} else {
		return parentRequestResult
	}

}


