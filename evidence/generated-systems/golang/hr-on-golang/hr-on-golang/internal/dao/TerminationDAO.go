package dao

import (
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing TerminationDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateTermination - creates a new db entry
//----------------------------------------------------------------------------
func CreateTermination(obj model.Termination)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Termination with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Termination", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateTermination", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetTermination - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetTermination(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Termination

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Termination with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Termination using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Termination using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetTermination", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllTermination - returns all
//----------------------------------------------------------------------------
func GetAllTermination()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Termination

	//----------------------------------------------------------------------------
	// Request the ORM to find all Termination
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Termination" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Termination", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllTermination", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateTermination - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateTermination(obj model.Termination)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Termination using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Termination using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateTermination", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteTermination - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteTermination(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Termination with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetTermination(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Termination so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Termination)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Termination using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Termination using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteTermination", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Employee on a Termination
//----------------------------------------------------------------------------
func AssignEmployeeToTermination( terminationId uint64, employeeId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Termination with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTermination(terminationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Termination so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Termination)

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
			// assign the Employee	to the Termination
			//----------------------------------------------------------------------------
			parentObj.Employee = &childObj

			//----------------------------------------------------------------------------
			// save the Termination
			//----------------------------------------------------------------------------
			return UpdateTermination(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Employee", employeeId )
			return utils.RequestResult{false, msg, "assignEmployee", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Employee on a Termination
//----------------------------------------------------------------------------
func UnassignEmployeeFromTermination(terminationId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Termination with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTermination(terminationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Termination so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Termination)

		//----------------------------------------------------------------------------
		// assign an empty Employee to the Employee
		//----------------------------------------------------------------------------
		parentObj.Employee = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Employee
		//----------------------------------------------------------------------------
		parentObj.EmployeeId = nil;

		//----------------------------------------------------------------------------
		// save the Termination
		//----------------------------------------------------------------------------
		return UpdateTermination(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Assignment on a Termination
//----------------------------------------------------------------------------
func AssignAssignmentToTermination( terminationId uint64, assignmentId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Termination with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTermination(terminationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Termination so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Termination)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.EmploymentAssignment

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a EmploymentAssignment with a
		// matching assignmentId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, assignmentId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Assignment	to the Termination
			//----------------------------------------------------------------------------
			parentObj.Assignment = &childObj

			//----------------------------------------------------------------------------
			// save the Termination
			//----------------------------------------------------------------------------
			return UpdateTermination(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Assignment", assignmentId )
			return utils.RequestResult{false, msg, "assignAssignment", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Assignment on a Termination
//----------------------------------------------------------------------------
func UnassignAssignmentFromTermination(terminationId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Termination with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTermination(terminationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Termination so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Termination)

		//----------------------------------------------------------------------------
		// assign an empty EmploymentAssignment to the Assignment
		//----------------------------------------------------------------------------
		parentObj.Assignment = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Assignment
		//----------------------------------------------------------------------------
		parentObj.AssignmentId = nil;

		//----------------------------------------------------------------------------
		// save the Termination
		//----------------------------------------------------------------------------
		return UpdateTermination(parentObj)

	} else {
		return parentRequestResult
	}

}


