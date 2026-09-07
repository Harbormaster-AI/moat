package dao

import (
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing LeavePolicyDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateLeavePolicy - creates a new db entry
//----------------------------------------------------------------------------
func CreateLeavePolicy(obj model.LeavePolicy)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a LeavePolicy with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a LeavePolicy", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateLeavePolicy", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetLeavePolicy - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetLeavePolicy(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.LeavePolicy

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a LeavePolicy with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a LeavePolicy using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a LeavePolicy using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetLeavePolicy", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllLeavePolicy - returns all
//----------------------------------------------------------------------------
func GetAllLeavePolicy()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.LeavePolicy

	//----------------------------------------------------------------------------
	// Request the ORM to find all LeavePolicy
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all LeavePolicy" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all LeavePolicy", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllLeavePolicy", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateLeavePolicy - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateLeavePolicy(obj model.LeavePolicy)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a LeavePolicy using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a LeavePolicy using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateLeavePolicy", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteLeavePolicy - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteLeavePolicy(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the LeavePolicy with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetLeavePolicy(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LeavePolicy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.LeavePolicy)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a LeavePolicy using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a LeavePolicy using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteLeavePolicy", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Organization on a LeavePolicy
//----------------------------------------------------------------------------
func AssignOrganizationToLeavePolicy( leavePolicyId uint64, organizationId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the LeavePolicy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLeavePolicy(leavePolicyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LeavePolicy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LeavePolicy)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Organization

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Organization with a
		// matching organizationId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, organizationId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Organization	to the LeavePolicy
			//----------------------------------------------------------------------------
			parentObj.Organization = &childObj

			//----------------------------------------------------------------------------
			// save the LeavePolicy
			//----------------------------------------------------------------------------
			return UpdateLeavePolicy(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Organization", organizationId )
			return utils.RequestResult{false, msg, "assignOrganization", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Organization on a LeavePolicy
//----------------------------------------------------------------------------
func UnassignOrganizationFromLeavePolicy(leavePolicyId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the LeavePolicy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLeavePolicy(leavePolicyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LeavePolicy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LeavePolicy)

		//----------------------------------------------------------------------------
		// assign an empty Organization to the Organization
		//----------------------------------------------------------------------------
		parentObj.Organization = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Organization
		//----------------------------------------------------------------------------
		parentObj.OrganizationId = nil;

		//----------------------------------------------------------------------------
		// save the LeavePolicy
		//----------------------------------------------------------------------------
		return UpdateLeavePolicy(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more leaveRequestsIds as a LeaveRequests to a LeavePolicy
//----------------------------------------------------------------------------
func AddLeaveRequestsToLeavePolicy ( leavePolicyId uint64, leaveRequestsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the LeavePolicy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLeavePolicy(leavePolicyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LeavePolicy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LeavePolicy)

		// slice the ids on comma with no spaces
		ids := strings.Split( leaveRequestsIds, ",")

		for _, leaveRequestsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.LeaveRequest

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a LeaveRequest
			// with a matching leaveRequestsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , leaveRequestsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the LeaveRequests using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("LeaveRequests").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "LeaveRequests", leaveRequestsId )
				return utils.RequestResult{false, msg, "unassignLeaveRequests", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified LeavePolicy from the gorm
		//----------------------------------------------------------------------------
		return GetLeavePolicy(leavePolicyId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more leaveRequestsIds as a LeaveRequests from a LeavePolicy
//----------------------------------------------------------------------------
func RemoveLeaveRequestsFromLeavePolicy( leavePolicyId uint64, leaveRequestsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the LeavePolicy with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLeavePolicy(leavePolicyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LeavePolicy so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LeavePolicy)

		// slice the ids on comma with no spaces
		ids := strings.Split( leaveRequestsIds, ",")

		for _, leaveRequestsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.LeaveRequest

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a LeaveRequest
			// with a matching leaveRequestsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , leaveRequestsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove LeaveRequestObj from the LeaveRequests array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("LeaveRequests").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "LeaveRequests", leaveRequestsId )
				return utils.RequestResult{false, msg, "removeLeaveRequests", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified LeavePolicy from the gorm
		//----------------------------------------------------------------------------
		return GetLeavePolicy(leavePolicyId)

	} else {
		return parentRequestResult
	}
}

