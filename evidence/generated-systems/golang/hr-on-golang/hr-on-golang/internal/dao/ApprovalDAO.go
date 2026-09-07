package dao

import (
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing ApprovalDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateApproval - creates a new db entry
//----------------------------------------------------------------------------
func CreateApproval(obj model.Approval)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Approval with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Approval", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateApproval", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetApproval - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetApproval(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Approval

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Approval with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Approval using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Approval using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetApproval", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllApproval - returns all
//----------------------------------------------------------------------------
func GetAllApproval()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Approval

	//----------------------------------------------------------------------------
	// Request the ORM to find all Approval
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Approval" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Approval", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllApproval", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateApproval - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateApproval(obj model.Approval)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Approval using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Approval using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateApproval", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteApproval - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteApproval(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Approval with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetApproval(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Approval so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Approval)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Approval using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Approval using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteApproval", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Approver on a Approval
//----------------------------------------------------------------------------
func AssignApproverToApproval( approvalId uint64, approverId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Approval with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetApproval(approvalId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Approval so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Approval)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Employee

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Employee with a
		// matching approverId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, approverId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Approver	to the Approval
			//----------------------------------------------------------------------------
			parentObj.Approver = &childObj

			//----------------------------------------------------------------------------
			// save the Approval
			//----------------------------------------------------------------------------
			return UpdateApproval(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Approver", approverId )
			return utils.RequestResult{false, msg, "assignApprover", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Approver on a Approval
//----------------------------------------------------------------------------
func UnassignApproverFromApproval(approvalId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Approval with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetApproval(approvalId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Approval so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Approval)

		//----------------------------------------------------------------------------
		// assign an empty Employee to the Approver
		//----------------------------------------------------------------------------
		parentObj.Approver = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Approver
		//----------------------------------------------------------------------------
		parentObj.ApproverId = nil;

		//----------------------------------------------------------------------------
		// save the Approval
		//----------------------------------------------------------------------------
		return UpdateApproval(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Timesheet on a Approval
//----------------------------------------------------------------------------
func AssignTimesheetToApproval( approvalId uint64, timesheetId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Approval with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetApproval(approvalId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Approval so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Approval)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Timesheet

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Timesheet with a
		// matching timesheetId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, timesheetId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Timesheet	to the Approval
			//----------------------------------------------------------------------------
			parentObj.Timesheet = &childObj

			//----------------------------------------------------------------------------
			// save the Approval
			//----------------------------------------------------------------------------
			return UpdateApproval(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Timesheet", timesheetId )
			return utils.RequestResult{false, msg, "assignTimesheet", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Timesheet on a Approval
//----------------------------------------------------------------------------
func UnassignTimesheetFromApproval(approvalId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Approval with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetApproval(approvalId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Approval so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Approval)

		//----------------------------------------------------------------------------
		// assign an empty Timesheet to the Timesheet
		//----------------------------------------------------------------------------
		parentObj.Timesheet = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Timesheet
		//----------------------------------------------------------------------------
		parentObj.TimesheetId = nil;

		//----------------------------------------------------------------------------
		// save the Approval
		//----------------------------------------------------------------------------
		return UpdateApproval(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a LeaveRequest on a Approval
//----------------------------------------------------------------------------
func AssignLeaveRequestToApproval( approvalId uint64, leaveRequestId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Approval with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetApproval(approvalId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Approval so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Approval)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.LeaveRequest

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a LeaveRequest with a
		// matching leaveRequestId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, leaveRequestId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the LeaveRequest	to the Approval
			//----------------------------------------------------------------------------
			parentObj.LeaveRequest = &childObj

			//----------------------------------------------------------------------------
			// save the Approval
			//----------------------------------------------------------------------------
			return UpdateApproval(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "LeaveRequest", leaveRequestId )
			return utils.RequestResult{false, msg, "assignLeaveRequest", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a LeaveRequest on a Approval
//----------------------------------------------------------------------------
func UnassignLeaveRequestFromApproval(approvalId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Approval with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetApproval(approvalId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Approval so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Approval)

		//----------------------------------------------------------------------------
		// assign an empty LeaveRequest to the LeaveRequest
		//----------------------------------------------------------------------------
		parentObj.LeaveRequest = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the LeaveRequest
		//----------------------------------------------------------------------------
		parentObj.LeaveRequestId = nil;

		//----------------------------------------------------------------------------
		// save the Approval
		//----------------------------------------------------------------------------
		return UpdateApproval(parentObj)

	} else {
		return parentRequestResult
	}

}


