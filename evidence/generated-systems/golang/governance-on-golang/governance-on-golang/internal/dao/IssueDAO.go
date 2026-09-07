package dao

import (
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing IssueDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateIssue - creates a new db entry
//----------------------------------------------------------------------------
func CreateIssue(obj model.Issue)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Issue with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Issue", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateIssue", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetIssue - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetIssue(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Issue

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Issue with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Issue using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Issue using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetIssue", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllIssue - returns all
//----------------------------------------------------------------------------
func GetAllIssue()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Issue

	//----------------------------------------------------------------------------
	// Request the ORM to find all Issue
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Issue" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Issue", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllIssue", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateIssue - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateIssue(obj model.Issue)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Issue using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Issue using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateIssue", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteIssue - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteIssue(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Issue with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetIssue(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Issue so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Issue)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Issue using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Issue using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteIssue", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Risk on a Issue
//----------------------------------------------------------------------------
func AssignRiskToIssue( issueId uint64, riskId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Issue with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetIssue(issueId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Issue so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Issue)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Risk

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Risk with a
		// matching riskId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, riskId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Risk	to the Issue
			//----------------------------------------------------------------------------
			parentObj.Risk = &childObj

			//----------------------------------------------------------------------------
			// save the Issue
			//----------------------------------------------------------------------------
			return UpdateIssue(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Risk", riskId )
			return utils.RequestResult{false, msg, "assignRisk", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Risk on a Issue
//----------------------------------------------------------------------------
func UnassignRiskFromIssue(issueId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Issue with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetIssue(issueId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Issue so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Issue)

		//----------------------------------------------------------------------------
		// assign an empty Risk to the Risk
		//----------------------------------------------------------------------------
		parentObj.Risk = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Risk
		//----------------------------------------------------------------------------
		parentObj.RiskId = nil;

		//----------------------------------------------------------------------------
		// save the Issue
		//----------------------------------------------------------------------------
		return UpdateIssue(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Finding on a Issue
//----------------------------------------------------------------------------
func AssignFindingToIssue( issueId uint64, findingId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Issue with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetIssue(issueId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Issue so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Issue)

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
			// assign the Finding	to the Issue
			//----------------------------------------------------------------------------
			parentObj.Finding = &childObj

			//----------------------------------------------------------------------------
			// save the Issue
			//----------------------------------------------------------------------------
			return UpdateIssue(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Finding", findingId )
			return utils.RequestResult{false, msg, "assignFinding", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Finding on a Issue
//----------------------------------------------------------------------------
func UnassignFindingFromIssue(issueId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Issue with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetIssue(issueId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Issue so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Issue)

		//----------------------------------------------------------------------------
		// assign an empty AuditFinding to the Finding
		//----------------------------------------------------------------------------
		parentObj.Finding = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Finding
		//----------------------------------------------------------------------------
		parentObj.FindingId = nil;

		//----------------------------------------------------------------------------
		// save the Issue
		//----------------------------------------------------------------------------
		return UpdateIssue(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Control on a Issue
//----------------------------------------------------------------------------
func AssignControlToIssue( issueId uint64, controlId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Issue with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetIssue(issueId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Issue so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Issue)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Control

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Control with a
		// matching controlId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, controlId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Control	to the Issue
			//----------------------------------------------------------------------------
			parentObj.Control = &childObj

			//----------------------------------------------------------------------------
			// save the Issue
			//----------------------------------------------------------------------------
			return UpdateIssue(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Control", controlId )
			return utils.RequestResult{false, msg, "assignControl", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Control on a Issue
//----------------------------------------------------------------------------
func UnassignControlFromIssue(issueId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Issue with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetIssue(issueId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Issue so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Issue)

		//----------------------------------------------------------------------------
		// assign an empty Control to the Control
		//----------------------------------------------------------------------------
		parentObj.Control = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Control
		//----------------------------------------------------------------------------
		parentObj.ControlId = nil;

		//----------------------------------------------------------------------------
		// save the Issue
		//----------------------------------------------------------------------------
		return UpdateIssue(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more correctiveActionsIds as a CorrectiveActions to a Issue
//----------------------------------------------------------------------------
func AddCorrectiveActionsToIssue ( issueId uint64, correctiveActionsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Issue with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetIssue(issueId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Issue so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Issue)

		// slice the ids on comma with no spaces
		ids := strings.Split( correctiveActionsIds, ",")

		for _, correctiveActionsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.CorrectiveAction

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a CorrectiveAction
			// with a matching correctiveActionsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , correctiveActionsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the CorrectiveActions using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("CorrectiveActions").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "CorrectiveActions", correctiveActionsId )
				return utils.RequestResult{false, msg, "unassignCorrectiveActions", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Issue from the gorm
		//----------------------------------------------------------------------------
		return GetIssue(issueId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more correctiveActionsIds as a CorrectiveActions from a Issue
//----------------------------------------------------------------------------
func RemoveCorrectiveActionsFromIssue( issueId uint64, correctiveActionsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Issue with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetIssue(issueId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Issue so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Issue)

		// slice the ids on comma with no spaces
		ids := strings.Split( correctiveActionsIds, ",")

		for _, correctiveActionsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.CorrectiveAction

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a CorrectiveAction
			// with a matching correctiveActionsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , correctiveActionsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove CorrectiveActionObj from the CorrectiveActions array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("CorrectiveActions").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "CorrectiveActions", correctiveActionsId )
				return utils.RequestResult{false, msg, "removeCorrectiveActions", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Issue from the gorm
		//----------------------------------------------------------------------------
		return GetIssue(issueId)

	} else {
		return parentRequestResult
	}
}

