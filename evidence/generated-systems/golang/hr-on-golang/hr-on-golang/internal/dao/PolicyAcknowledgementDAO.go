package dao

import (
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing PolicyAcknowledgementDAO..." ) )
}

//----------------------------------------------------------------------------
// CreatePolicyAcknowledgement - creates a new db entry
//----------------------------------------------------------------------------
func CreatePolicyAcknowledgement(obj model.PolicyAcknowledgement)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a PolicyAcknowledgement with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a PolicyAcknowledgement", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreatePolicyAcknowledgement", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetPolicyAcknowledgement - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetPolicyAcknowledgement(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.PolicyAcknowledgement

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a PolicyAcknowledgement with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a PolicyAcknowledgement using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a PolicyAcknowledgement using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetPolicyAcknowledgement", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllPolicyAcknowledgement - returns all
//----------------------------------------------------------------------------
func GetAllPolicyAcknowledgement()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.PolicyAcknowledgement

	//----------------------------------------------------------------------------
	// Request the ORM to find all PolicyAcknowledgement
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all PolicyAcknowledgement" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all PolicyAcknowledgement", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllPolicyAcknowledgement", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdatePolicyAcknowledgement - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdatePolicyAcknowledgement(obj model.PolicyAcknowledgement)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a PolicyAcknowledgement using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a PolicyAcknowledgement using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdatePolicyAcknowledgement", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeletePolicyAcknowledgement - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeletePolicyAcknowledgement(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the PolicyAcknowledgement with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetPolicyAcknowledgement(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PolicyAcknowledgement so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.PolicyAcknowledgement)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a PolicyAcknowledgement using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a PolicyAcknowledgement using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeletePolicyAcknowledgement", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Policy on a PolicyAcknowledgement
//----------------------------------------------------------------------------
func AssignPolicyToPolicyAcknowledgement( policyAcknowledgementId uint64, policyId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the PolicyAcknowledgement with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPolicyAcknowledgement(policyAcknowledgementId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PolicyAcknowledgement so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PolicyAcknowledgement)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Policy

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Policy with a
		// matching policyId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, policyId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Policy	to the PolicyAcknowledgement
			//----------------------------------------------------------------------------
			parentObj.Policy = &childObj

			//----------------------------------------------------------------------------
			// save the PolicyAcknowledgement
			//----------------------------------------------------------------------------
			return UpdatePolicyAcknowledgement(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Policy", policyId )
			return utils.RequestResult{false, msg, "assignPolicy", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Policy on a PolicyAcknowledgement
//----------------------------------------------------------------------------
func UnassignPolicyFromPolicyAcknowledgement(policyAcknowledgementId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the PolicyAcknowledgement with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPolicyAcknowledgement(policyAcknowledgementId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PolicyAcknowledgement so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PolicyAcknowledgement)

		//----------------------------------------------------------------------------
		// assign an empty Policy to the Policy
		//----------------------------------------------------------------------------
		parentObj.Policy = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Policy
		//----------------------------------------------------------------------------
		parentObj.PolicyId = nil;

		//----------------------------------------------------------------------------
		// save the PolicyAcknowledgement
		//----------------------------------------------------------------------------
		return UpdatePolicyAcknowledgement(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Employee on a PolicyAcknowledgement
//----------------------------------------------------------------------------
func AssignEmployeeToPolicyAcknowledgement( policyAcknowledgementId uint64, employeeId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the PolicyAcknowledgement with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPolicyAcknowledgement(policyAcknowledgementId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PolicyAcknowledgement so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PolicyAcknowledgement)

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
			// assign the Employee	to the PolicyAcknowledgement
			//----------------------------------------------------------------------------
			parentObj.Employee = &childObj

			//----------------------------------------------------------------------------
			// save the PolicyAcknowledgement
			//----------------------------------------------------------------------------
			return UpdatePolicyAcknowledgement(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Employee", employeeId )
			return utils.RequestResult{false, msg, "assignEmployee", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Employee on a PolicyAcknowledgement
//----------------------------------------------------------------------------
func UnassignEmployeeFromPolicyAcknowledgement(policyAcknowledgementId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the PolicyAcknowledgement with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPolicyAcknowledgement(policyAcknowledgementId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.PolicyAcknowledgement so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.PolicyAcknowledgement)

		//----------------------------------------------------------------------------
		// assign an empty Employee to the Employee
		//----------------------------------------------------------------------------
		parentObj.Employee = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Employee
		//----------------------------------------------------------------------------
		parentObj.EmployeeId = nil;

		//----------------------------------------------------------------------------
		// save the PolicyAcknowledgement
		//----------------------------------------------------------------------------
		return UpdatePolicyAcknowledgement(parentObj)

	} else {
		return parentRequestResult
	}

}


