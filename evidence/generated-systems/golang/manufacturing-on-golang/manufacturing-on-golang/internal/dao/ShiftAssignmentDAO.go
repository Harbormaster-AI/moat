package dao

import (
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing ShiftAssignmentDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateShiftAssignment - creates a new db entry
//----------------------------------------------------------------------------
func CreateShiftAssignment(obj model.ShiftAssignment)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a ShiftAssignment with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a ShiftAssignment", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateShiftAssignment", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetShiftAssignment - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetShiftAssignment(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.ShiftAssignment

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a ShiftAssignment with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a ShiftAssignment using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a ShiftAssignment using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetShiftAssignment", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllShiftAssignment - returns all
//----------------------------------------------------------------------------
func GetAllShiftAssignment()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.ShiftAssignment

	//----------------------------------------------------------------------------
	// Request the ORM to find all ShiftAssignment
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all ShiftAssignment" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all ShiftAssignment", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllShiftAssignment", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateShiftAssignment - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateShiftAssignment(obj model.ShiftAssignment)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a ShiftAssignment using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a ShiftAssignment using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateShiftAssignment", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteShiftAssignment - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteShiftAssignment(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the ShiftAssignment with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetShiftAssignment(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ShiftAssignment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.ShiftAssignment)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a ShiftAssignment using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a ShiftAssignment using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteShiftAssignment", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Shift on a ShiftAssignment
//----------------------------------------------------------------------------
func AssignShiftToShiftAssignment( shiftAssignmentId uint64, shiftId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the ShiftAssignment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetShiftAssignment(shiftAssignmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ShiftAssignment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ShiftAssignment)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Shift

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Shift with a
		// matching shiftId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, shiftId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Shift	to the ShiftAssignment
			//----------------------------------------------------------------------------
			parentObj.Shift = &childObj

			//----------------------------------------------------------------------------
			// save the ShiftAssignment
			//----------------------------------------------------------------------------
			return UpdateShiftAssignment(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Shift", shiftId )
			return utils.RequestResult{false, msg, "assignShift", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Shift on a ShiftAssignment
//----------------------------------------------------------------------------
func UnassignShiftFromShiftAssignment(shiftAssignmentId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ShiftAssignment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetShiftAssignment(shiftAssignmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ShiftAssignment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ShiftAssignment)

		//----------------------------------------------------------------------------
		// assign an empty Shift to the Shift
		//----------------------------------------------------------------------------
		parentObj.Shift = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Shift
		//----------------------------------------------------------------------------
		parentObj.ShiftId = nil;

		//----------------------------------------------------------------------------
		// save the ShiftAssignment
		//----------------------------------------------------------------------------
		return UpdateShiftAssignment(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Employee on a ShiftAssignment
//----------------------------------------------------------------------------
func AssignEmployeeToShiftAssignment( shiftAssignmentId uint64, employeeId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the ShiftAssignment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetShiftAssignment(shiftAssignmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ShiftAssignment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ShiftAssignment)

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
			// assign the Employee	to the ShiftAssignment
			//----------------------------------------------------------------------------
			parentObj.Employee = &childObj

			//----------------------------------------------------------------------------
			// save the ShiftAssignment
			//----------------------------------------------------------------------------
			return UpdateShiftAssignment(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Employee", employeeId )
			return utils.RequestResult{false, msg, "assignEmployee", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Employee on a ShiftAssignment
//----------------------------------------------------------------------------
func UnassignEmployeeFromShiftAssignment(shiftAssignmentId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ShiftAssignment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetShiftAssignment(shiftAssignmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ShiftAssignment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ShiftAssignment)

		//----------------------------------------------------------------------------
		// assign an empty Employee to the Employee
		//----------------------------------------------------------------------------
		parentObj.Employee = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Employee
		//----------------------------------------------------------------------------
		parentObj.EmployeeId = nil;

		//----------------------------------------------------------------------------
		// save the ShiftAssignment
		//----------------------------------------------------------------------------
		return UpdateShiftAssignment(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a WorkCenter on a ShiftAssignment
//----------------------------------------------------------------------------
func AssignWorkCenterToShiftAssignment( shiftAssignmentId uint64, workCenterId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the ShiftAssignment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetShiftAssignment(shiftAssignmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ShiftAssignment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ShiftAssignment)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.WorkCenter

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a WorkCenter with a
		// matching workCenterId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, workCenterId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the WorkCenter	to the ShiftAssignment
			//----------------------------------------------------------------------------
			parentObj.WorkCenter = &childObj

			//----------------------------------------------------------------------------
			// save the ShiftAssignment
			//----------------------------------------------------------------------------
			return UpdateShiftAssignment(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "WorkCenter", workCenterId )
			return utils.RequestResult{false, msg, "assignWorkCenter", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a WorkCenter on a ShiftAssignment
//----------------------------------------------------------------------------
func UnassignWorkCenterFromShiftAssignment(shiftAssignmentId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ShiftAssignment with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetShiftAssignment(shiftAssignmentId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ShiftAssignment so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ShiftAssignment)

		//----------------------------------------------------------------------------
		// assign an empty WorkCenter to the WorkCenter
		//----------------------------------------------------------------------------
		parentObj.WorkCenter = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the WorkCenter
		//----------------------------------------------------------------------------
		parentObj.WorkCenterId = nil;

		//----------------------------------------------------------------------------
		// save the ShiftAssignment
		//----------------------------------------------------------------------------
		return UpdateShiftAssignment(parentObj)

	} else {
		return parentRequestResult
	}

}


