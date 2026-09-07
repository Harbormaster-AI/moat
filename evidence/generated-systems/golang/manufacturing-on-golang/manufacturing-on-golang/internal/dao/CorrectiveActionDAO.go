package dao

import (
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing CorrectiveActionDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateCorrectiveAction - creates a new db entry
//----------------------------------------------------------------------------
func CreateCorrectiveAction(obj model.CorrectiveAction)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a CorrectiveAction with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a CorrectiveAction", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateCorrectiveAction", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetCorrectiveAction - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetCorrectiveAction(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.CorrectiveAction

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a CorrectiveAction with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a CorrectiveAction using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a CorrectiveAction using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetCorrectiveAction", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllCorrectiveAction - returns all
//----------------------------------------------------------------------------
func GetAllCorrectiveAction()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.CorrectiveAction

	//----------------------------------------------------------------------------
	// Request the ORM to find all CorrectiveAction
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all CorrectiveAction" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all CorrectiveAction", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllCorrectiveAction", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateCorrectiveAction - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateCorrectiveAction(obj model.CorrectiveAction)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a CorrectiveAction using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a CorrectiveAction using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateCorrectiveAction", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteCorrectiveAction - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteCorrectiveAction(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the CorrectiveAction with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetCorrectiveAction(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CorrectiveAction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.CorrectiveAction)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a CorrectiveAction using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a CorrectiveAction using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteCorrectiveAction", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Nonconformance on a CorrectiveAction
//----------------------------------------------------------------------------
func AssignNonconformanceToCorrectiveAction( correctiveActionId uint64, nonconformanceId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the CorrectiveAction with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCorrectiveAction(correctiveActionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CorrectiveAction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CorrectiveAction)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Nonconformance

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Nonconformance with a
		// matching nonconformanceId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, nonconformanceId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Nonconformance	to the CorrectiveAction
			//----------------------------------------------------------------------------
			parentObj.Nonconformance = &childObj

			//----------------------------------------------------------------------------
			// save the CorrectiveAction
			//----------------------------------------------------------------------------
			return UpdateCorrectiveAction(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Nonconformance", nonconformanceId )
			return utils.RequestResult{false, msg, "assignNonconformance", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Nonconformance on a CorrectiveAction
//----------------------------------------------------------------------------
func UnassignNonconformanceFromCorrectiveAction(correctiveActionId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the CorrectiveAction with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCorrectiveAction(correctiveActionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CorrectiveAction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CorrectiveAction)

		//----------------------------------------------------------------------------
		// assign an empty Nonconformance to the Nonconformance
		//----------------------------------------------------------------------------
		parentObj.Nonconformance = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Nonconformance
		//----------------------------------------------------------------------------
		parentObj.NonconformanceId = nil;

		//----------------------------------------------------------------------------
		// save the CorrectiveAction
		//----------------------------------------------------------------------------
		return UpdateCorrectiveAction(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Owner on a CorrectiveAction
//----------------------------------------------------------------------------
func AssignOwnerToCorrectiveAction( correctiveActionId uint64, ownerId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the CorrectiveAction with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCorrectiveAction(correctiveActionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CorrectiveAction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CorrectiveAction)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Employee

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Employee with a
		// matching ownerId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, ownerId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Owner	to the CorrectiveAction
			//----------------------------------------------------------------------------
			parentObj.Owner = &childObj

			//----------------------------------------------------------------------------
			// save the CorrectiveAction
			//----------------------------------------------------------------------------
			return UpdateCorrectiveAction(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Owner", ownerId )
			return utils.RequestResult{false, msg, "assignOwner", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Owner on a CorrectiveAction
//----------------------------------------------------------------------------
func UnassignOwnerFromCorrectiveAction(correctiveActionId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the CorrectiveAction with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCorrectiveAction(correctiveActionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CorrectiveAction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CorrectiveAction)

		//----------------------------------------------------------------------------
		// assign an empty Employee to the Owner
		//----------------------------------------------------------------------------
		parentObj.Owner = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Owner
		//----------------------------------------------------------------------------
		parentObj.OwnerId = nil;

		//----------------------------------------------------------------------------
		// save the CorrectiveAction
		//----------------------------------------------------------------------------
		return UpdateCorrectiveAction(parentObj)

	} else {
		return parentRequestResult
	}

}


