package dao

import (
    "governance-on-golang/internal/model"
    "governance-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing Exception_DAO..." ) )
}

//----------------------------------------------------------------------------
// CreateException_ - creates a new db entry
//----------------------------------------------------------------------------
func CreateException_(obj model.Exception_)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Exception_ with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Exception_", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateException_", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetException_ - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetException_(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Exception_

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Exception_ with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Exception_ using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Exception_ using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetException_", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllException_ - returns all
//----------------------------------------------------------------------------
func GetAllException_()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Exception_

	//----------------------------------------------------------------------------
	// Request the ORM to find all Exception_
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Exception_" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Exception_", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllException_", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateException_ - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateException_(obj model.Exception_)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Exception_ using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Exception_ using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateException_", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteException_ - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteException_(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Exception_ with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetException_(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Exception_ so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Exception_)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Exception_ using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Exception_ using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteException_", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a RetentionSchedule on a Exception_
//----------------------------------------------------------------------------
func AssignRetentionScheduleToException_( exception_Id uint64, retentionScheduleId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Exception_ with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetException_(exception_Id)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Exception_ so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Exception_)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.RetentionSchedule

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a RetentionSchedule with a
		// matching retentionScheduleId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, retentionScheduleId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the RetentionSchedule	to the Exception_
			//----------------------------------------------------------------------------
			parentObj.RetentionSchedule = &childObj

			//----------------------------------------------------------------------------
			// save the Exception_
			//----------------------------------------------------------------------------
			return UpdateException_(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "RetentionSchedule", retentionScheduleId )
			return utils.RequestResult{false, msg, "assignRetentionSchedule", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a RetentionSchedule on a Exception_
//----------------------------------------------------------------------------
func UnassignRetentionScheduleFromException_(exception_Id uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Exception_ with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetException_(exception_Id)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Exception_ so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Exception_)

		//----------------------------------------------------------------------------
		// assign an empty RetentionSchedule to the RetentionSchedule
		//----------------------------------------------------------------------------
		parentObj.RetentionSchedule = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the RetentionSchedule
		//----------------------------------------------------------------------------
		parentObj.RetentionScheduleId = nil;

		//----------------------------------------------------------------------------
		// save the Exception_
		//----------------------------------------------------------------------------
		return UpdateException_(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Policy on a Exception_
//----------------------------------------------------------------------------
func AssignPolicyToException_( exception_Id uint64, policyId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Exception_ with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetException_(exception_Id)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Exception_ so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Exception_)

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
			// assign the Policy	to the Exception_
			//----------------------------------------------------------------------------
			parentObj.Policy = &childObj

			//----------------------------------------------------------------------------
			// save the Exception_
			//----------------------------------------------------------------------------
			return UpdateException_(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Policy", policyId )
			return utils.RequestResult{false, msg, "assignPolicy", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Policy on a Exception_
//----------------------------------------------------------------------------
func UnassignPolicyFromException_(exception_Id uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Exception_ with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetException_(exception_Id)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Exception_ so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Exception_)

		//----------------------------------------------------------------------------
		// assign an empty Policy to the Policy
		//----------------------------------------------------------------------------
		parentObj.Policy = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Policy
		//----------------------------------------------------------------------------
		parentObj.PolicyId = nil;

		//----------------------------------------------------------------------------
		// save the Exception_
		//----------------------------------------------------------------------------
		return UpdateException_(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Control on a Exception_
//----------------------------------------------------------------------------
func AssignControlToException_( exception_Id uint64, controlId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Exception_ with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetException_(exception_Id)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Exception_ so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Exception_)

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
			// assign the Control	to the Exception_
			//----------------------------------------------------------------------------
			parentObj.Control = &childObj

			//----------------------------------------------------------------------------
			// save the Exception_
			//----------------------------------------------------------------------------
			return UpdateException_(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Control", controlId )
			return utils.RequestResult{false, msg, "assignControl", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Control on a Exception_
//----------------------------------------------------------------------------
func UnassignControlFromException_(exception_Id uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Exception_ with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetException_(exception_Id)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Exception_ so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Exception_)

		//----------------------------------------------------------------------------
		// assign an empty Control to the Control
		//----------------------------------------------------------------------------
		parentObj.Control = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Control
		//----------------------------------------------------------------------------
		parentObj.ControlId = nil;

		//----------------------------------------------------------------------------
		// save the Exception_
		//----------------------------------------------------------------------------
		return UpdateException_(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Risk on a Exception_
//----------------------------------------------------------------------------
func AssignRiskToException_( exception_Id uint64, riskId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Exception_ with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetException_(exception_Id)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Exception_ so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Exception_)

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
			// assign the Risk	to the Exception_
			//----------------------------------------------------------------------------
			parentObj.Risk = &childObj

			//----------------------------------------------------------------------------
			// save the Exception_
			//----------------------------------------------------------------------------
			return UpdateException_(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Risk", riskId )
			return utils.RequestResult{false, msg, "assignRisk", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Risk on a Exception_
//----------------------------------------------------------------------------
func UnassignRiskFromException_(exception_Id uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Exception_ with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetException_(exception_Id)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Exception_ so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Exception_)

		//----------------------------------------------------------------------------
		// assign an empty Risk to the Risk
		//----------------------------------------------------------------------------
		parentObj.Risk = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Risk
		//----------------------------------------------------------------------------
		parentObj.RiskId = nil;

		//----------------------------------------------------------------------------
		// save the Exception_
		//----------------------------------------------------------------------------
		return UpdateException_(parentObj)

	} else {
		return parentRequestResult
	}

}


