package dao

import (
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
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
// assigns a Finding on a CorrectiveAction
//----------------------------------------------------------------------------
func AssignFindingToCorrectiveAction( correctiveActionId uint64, findingId uint64 )(utils.RequestResult){

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
		var childObj model.AuditFinding

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a AuditFinding with a
		// matching findingId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, findingId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Finding	to the CorrectiveAction
			//----------------------------------------------------------------------------
			parentObj.Finding = &childObj

			//----------------------------------------------------------------------------
			// save the CorrectiveAction
			//----------------------------------------------------------------------------
			return UpdateCorrectiveAction(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Finding", findingId )
			return utils.RequestResult{false, msg, "assignFinding", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Finding on a CorrectiveAction
//----------------------------------------------------------------------------
func UnassignFindingFromCorrectiveAction(correctiveActionId uint64)(utils.RequestResult) {

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
		// assign an empty AuditFinding to the Finding
		//----------------------------------------------------------------------------
		parentObj.Finding = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Finding
		//----------------------------------------------------------------------------
		parentObj.FindingId = nil;

		//----------------------------------------------------------------------------
		// save the CorrectiveAction
		//----------------------------------------------------------------------------
		return UpdateCorrectiveAction(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Issue on a CorrectiveAction
//----------------------------------------------------------------------------
func AssignIssueToCorrectiveAction( correctiveActionId uint64, issueId uint64 )(utils.RequestResult){

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
		var childObj model.Issue

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Issue with a
		// matching issueId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, issueId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Issue	to the CorrectiveAction
			//----------------------------------------------------------------------------
			parentObj.Issue = &childObj

			//----------------------------------------------------------------------------
			// save the CorrectiveAction
			//----------------------------------------------------------------------------
			return UpdateCorrectiveAction(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Issue", issueId )
			return utils.RequestResult{false, msg, "assignIssue", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Issue on a CorrectiveAction
//----------------------------------------------------------------------------
func UnassignIssueFromCorrectiveAction(correctiveActionId uint64)(utils.RequestResult) {

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
		// assign an empty Issue to the Issue
		//----------------------------------------------------------------------------
		parentObj.Issue = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Issue
		//----------------------------------------------------------------------------
		parentObj.IssueId = nil;

		//----------------------------------------------------------------------------
		// save the CorrectiveAction
		//----------------------------------------------------------------------------
		return UpdateCorrectiveAction(parentObj)

	} else {
		return parentRequestResult
	}

}


