package dao

import (
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing LeaveRequestDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateLeaveRequest - creates a new db entry
//----------------------------------------------------------------------------
func CreateLeaveRequest(obj model.LeaveRequest)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a LeaveRequest with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a LeaveRequest", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateLeaveRequest", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetLeaveRequest - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetLeaveRequest(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.LeaveRequest

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a LeaveRequest with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a LeaveRequest using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a LeaveRequest using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetLeaveRequest", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllLeaveRequest - returns all
//----------------------------------------------------------------------------
func GetAllLeaveRequest()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.LeaveRequest

	//----------------------------------------------------------------------------
	// Request the ORM to find all LeaveRequest
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all LeaveRequest" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all LeaveRequest", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllLeaveRequest", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateLeaveRequest - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateLeaveRequest(obj model.LeaveRequest)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a LeaveRequest using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a LeaveRequest using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateLeaveRequest", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteLeaveRequest - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteLeaveRequest(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the LeaveRequest with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetLeaveRequest(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LeaveRequest so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.LeaveRequest)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a LeaveRequest using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a LeaveRequest using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteLeaveRequest", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Employee on a LeaveRequest
//----------------------------------------------------------------------------
func AssignEmployeeToLeaveRequest( leaveRequestId uint64, employeeId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the LeaveRequest with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLeaveRequest(leaveRequestId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LeaveRequest so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LeaveRequest)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Employee

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Employee with a
		// matching employeeId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, employeeId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Employee	to the LeaveRequest
			//----------------------------------------------------------------------------
			parentObj.Employee = &childObj

			//----------------------------------------------------------------------------
			// save the LeaveRequest
			//----------------------------------------------------------------------------
			return UpdateLeaveRequest(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Employee", employeeId )
			return utils.RequestResult{false, msg, "assignEmployee", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Employee on a LeaveRequest
//----------------------------------------------------------------------------
func UnassignEmployeeFromLeaveRequest(leaveRequestId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the LeaveRequest with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLeaveRequest(leaveRequestId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LeaveRequest so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LeaveRequest)

		//----------------------------------------------------------------------------
		// assign an empty Employee to the Employee
		//----------------------------------------------------------------------------
		parentObj.Employee = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Employee
		//----------------------------------------------------------------------------
		parentObj.EmployeeId = nil;

		//----------------------------------------------------------------------------
		// save the LeaveRequest
		//----------------------------------------------------------------------------
		return UpdateLeaveRequest(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a LeavePolicy on a LeaveRequest
//----------------------------------------------------------------------------
func AssignLeavePolicyToLeaveRequest( leaveRequestId uint64, leavePolicyId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the LeaveRequest with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLeaveRequest(leaveRequestId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LeaveRequest so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LeaveRequest)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.LeavePolicy

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a LeavePolicy with a
		// matching leavePolicyId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, leavePolicyId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the LeavePolicy	to the LeaveRequest
			//----------------------------------------------------------------------------
			parentObj.LeavePolicy = &childObj

			//----------------------------------------------------------------------------
			// save the LeaveRequest
			//----------------------------------------------------------------------------
			return UpdateLeaveRequest(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "LeavePolicy", leavePolicyId )
			return utils.RequestResult{false, msg, "assignLeavePolicy", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a LeavePolicy on a LeaveRequest
//----------------------------------------------------------------------------
func UnassignLeavePolicyFromLeaveRequest(leaveRequestId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the LeaveRequest with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLeaveRequest(leaveRequestId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LeaveRequest so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LeaveRequest)

		//----------------------------------------------------------------------------
		// assign an empty LeavePolicy to the LeavePolicy
		//----------------------------------------------------------------------------
		parentObj.LeavePolicy = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the LeavePolicy
		//----------------------------------------------------------------------------
		parentObj.LeavePolicyId = nil;

		//----------------------------------------------------------------------------
		// save the LeaveRequest
		//----------------------------------------------------------------------------
		return UpdateLeaveRequest(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more approvalsIds as a Approvals to a LeaveRequest
//----------------------------------------------------------------------------
func AddApprovalsToLeaveRequest ( leaveRequestId uint64, approvalsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the LeaveRequest with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLeaveRequest(leaveRequestId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LeaveRequest so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LeaveRequest)

		// slice the ids on comma with no spaces
		ids := strings.Split( approvalsIds, ",")

		for _, approvalsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Approval

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Approval
			// with a matching approvalsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , approvalsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Approvals using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Approvals").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Approvals", approvalsId )
				return utils.RequestResult{false, msg, "unassignApprovals", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified LeaveRequest from the gorm
		//----------------------------------------------------------------------------
		return GetLeaveRequest(leaveRequestId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more approvalsIds as a Approvals from a LeaveRequest
//----------------------------------------------------------------------------
func RemoveApprovalsFromLeaveRequest( leaveRequestId uint64, approvalsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the LeaveRequest with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLeaveRequest(leaveRequestId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LeaveRequest so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LeaveRequest)

		// slice the ids on comma with no spaces
		ids := strings.Split( approvalsIds, ",")

		for _, approvalsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Approval

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Approval
			// with a matching approvalsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , approvalsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ApprovalObj from the Approvals array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Approvals").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Approvals", approvalsId )
				return utils.RequestResult{false, msg, "removeApprovals", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified LeaveRequest from the gorm
		//----------------------------------------------------------------------------
		return GetLeaveRequest(leaveRequestId)

	} else {
		return parentRequestResult
	}
}

