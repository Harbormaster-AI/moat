package dao

import (
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing EmploymentAssignmentDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateEmploymentAssignment - creates a new db entry
//----------------------------------------------------------------------------
func CreateEmploymentAssignment(obj model.EmploymentAssignment)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a EmploymentAssignment with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a EmploymentAssignment", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateEmploymentAssignment", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetEmploymentAssignment - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetEmploymentAssignment(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.EmploymentAssignment

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a EmploymentAssignment with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a EmploymentAssignment using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a EmploymentAssignment using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetEmploymentAssignment", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllEmploymentAssignment - returns all
//----------------------------------------------------------------------------
func GetAllEmploymentAssignment()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.EmploymentAssignment

	//----------------------------------------------------------------------------
	// Request the ORM to find all EmploymentAssignment
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all EmploymentAssignment" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all EmploymentAssignment", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllEmploymentAssignment", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateEmploymentAssignment - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateEmploymentAssignment(obj model.EmploymentAssignment)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a EmploymentAssignment using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a EmploymentAssignment using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateEmploymentAssignment", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteEmploymentAssignment - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteEmploymentAssignment(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the EmploymentAssignment with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetEmploymentAssignment(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.EmploymentAssignment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.EmploymentAssignment)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a EmploymentAssignment using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a EmploymentAssignment using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteEmploymentAssignment", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Employee on a EmploymentAssignment
//----------------------------------------------------------------------------
func AssignEmployeeToEmploymentAssignment( employmentAssignmentId uint64, employeeId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the EmploymentAssignment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEmploymentAssignment(employmentAssignmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.EmploymentAssignment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.EmploymentAssignment)

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
			// assign the Employee	to the EmploymentAssignment
			//----------------------------------------------------------------------------
			parentObj.Employee = &childObj

			//----------------------------------------------------------------------------
			// save the EmploymentAssignment
			//----------------------------------------------------------------------------
			return UpdateEmploymentAssignment(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Employee", employeeId )
			return utils.RequestResult{false, msg, "assignEmployee", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Employee on a EmploymentAssignment
//----------------------------------------------------------------------------
func UnassignEmployeeFromEmploymentAssignment(employmentAssignmentId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the EmploymentAssignment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEmploymentAssignment(employmentAssignmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.EmploymentAssignment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.EmploymentAssignment)

		//----------------------------------------------------------------------------
		// assign an empty Employee to the Employee
		//----------------------------------------------------------------------------
		parentObj.Employee = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Employee
		//----------------------------------------------------------------------------
		parentObj.EmployeeId = nil;

		//----------------------------------------------------------------------------
		// save the EmploymentAssignment
		//----------------------------------------------------------------------------
		return UpdateEmploymentAssignment(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Position on a EmploymentAssignment
//----------------------------------------------------------------------------
func AssignPositionToEmploymentAssignment( employmentAssignmentId uint64, positionId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the EmploymentAssignment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEmploymentAssignment(employmentAssignmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.EmploymentAssignment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.EmploymentAssignment)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Position

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Position with a
		// matching positionId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, positionId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Position	to the EmploymentAssignment
			//----------------------------------------------------------------------------
			parentObj.Position = &childObj

			//----------------------------------------------------------------------------
			// save the EmploymentAssignment
			//----------------------------------------------------------------------------
			return UpdateEmploymentAssignment(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Position", positionId )
			return utils.RequestResult{false, msg, "assignPosition", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Position on a EmploymentAssignment
//----------------------------------------------------------------------------
func UnassignPositionFromEmploymentAssignment(employmentAssignmentId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the EmploymentAssignment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEmploymentAssignment(employmentAssignmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.EmploymentAssignment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.EmploymentAssignment)

		//----------------------------------------------------------------------------
		// assign an empty Position to the Position
		//----------------------------------------------------------------------------
		parentObj.Position = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Position
		//----------------------------------------------------------------------------
		parentObj.PositionId = nil;

		//----------------------------------------------------------------------------
		// save the EmploymentAssignment
		//----------------------------------------------------------------------------
		return UpdateEmploymentAssignment(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Supervisor on a EmploymentAssignment
//----------------------------------------------------------------------------
func AssignSupervisorToEmploymentAssignment( employmentAssignmentId uint64, supervisorId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the EmploymentAssignment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEmploymentAssignment(employmentAssignmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.EmploymentAssignment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.EmploymentAssignment)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Employee

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Employee with a
		// matching supervisorId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, supervisorId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Supervisor	to the EmploymentAssignment
			//----------------------------------------------------------------------------
			parentObj.Supervisor = &childObj

			//----------------------------------------------------------------------------
			// save the EmploymentAssignment
			//----------------------------------------------------------------------------
			return UpdateEmploymentAssignment(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Supervisor", supervisorId )
			return utils.RequestResult{false, msg, "assignSupervisor", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Supervisor on a EmploymentAssignment
//----------------------------------------------------------------------------
func UnassignSupervisorFromEmploymentAssignment(employmentAssignmentId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the EmploymentAssignment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEmploymentAssignment(employmentAssignmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.EmploymentAssignment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.EmploymentAssignment)

		//----------------------------------------------------------------------------
		// assign an empty Employee to the Supervisor
		//----------------------------------------------------------------------------
		parentObj.Supervisor = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Supervisor
		//----------------------------------------------------------------------------
		parentObj.SupervisorId = nil;

		//----------------------------------------------------------------------------
		// save the EmploymentAssignment
		//----------------------------------------------------------------------------
		return UpdateEmploymentAssignment(parentObj)

	} else {
		return parentRequestResult
	}

}


