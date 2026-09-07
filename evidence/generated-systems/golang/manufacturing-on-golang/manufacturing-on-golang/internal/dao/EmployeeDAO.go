package dao

import (
    "manufacturing-on-golang/internal/model"
    "manufacturing-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing EmployeeDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateEmployee - creates a new db entry
//----------------------------------------------------------------------------
func CreateEmployee(obj model.Employee)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Employee with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Employee", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateEmployee", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetEmployee - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetEmployee(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Employee

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Employee with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Employee using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Employee using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetEmployee", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllEmployee - returns all
//----------------------------------------------------------------------------
func GetAllEmployee()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Employee

	//----------------------------------------------------------------------------
	// Request the ORM to find all Employee
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Employee" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Employee", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllEmployee", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateEmployee - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateEmployee(obj model.Employee)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Employee using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Employee using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateEmployee", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteEmployee - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteEmployee(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Employee with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetEmployee(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Employee so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Employee)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Employee using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Employee using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteEmployee", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a WorkCenter on a Employee
//----------------------------------------------------------------------------
func AssignWorkCenterToEmployee( employeeId uint64, workCenterId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Employee with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEmployee(employeeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Employee so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Employee)

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
			// assign the WorkCenter	to the Employee
			//----------------------------------------------------------------------------
			parentObj.WorkCenter = &childObj

			//----------------------------------------------------------------------------
			// save the Employee
			//----------------------------------------------------------------------------
			return UpdateEmployee(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "WorkCenter", workCenterId )
			return utils.RequestResult{false, msg, "assignWorkCenter", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a WorkCenter on a Employee
//----------------------------------------------------------------------------
func UnassignWorkCenterFromEmployee(employeeId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Employee with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEmployee(employeeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Employee so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Employee)

		//----------------------------------------------------------------------------
		// assign an empty WorkCenter to the WorkCenter
		//----------------------------------------------------------------------------
		parentObj.WorkCenter = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the WorkCenter
		//----------------------------------------------------------------------------
		parentObj.WorkCenterId = nil;

		//----------------------------------------------------------------------------
		// save the Employee
		//----------------------------------------------------------------------------
		return UpdateEmployee(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more shiftAssignmentsIds as a ShiftAssignments to a Employee
//----------------------------------------------------------------------------
func AddShiftAssignmentsToEmployee ( employeeId uint64, shiftAssignmentsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Employee with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEmployee(employeeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Employee so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Employee)

		// slice the ids on comma with no spaces
		ids := strings.Split( shiftAssignmentsIds, ",")

		for _, shiftAssignmentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ShiftAssignment

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ShiftAssignment
			// with a matching shiftAssignmentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , shiftAssignmentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the ShiftAssignments using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("ShiftAssignments").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ShiftAssignments", shiftAssignmentsId )
				return utils.RequestResult{false, msg, "unassignShiftAssignments", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Employee from the gorm
		//----------------------------------------------------------------------------
		return GetEmployee(employeeId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more shiftAssignmentsIds as a ShiftAssignments from a Employee
//----------------------------------------------------------------------------
func RemoveShiftAssignmentsFromEmployee( employeeId uint64, shiftAssignmentsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Employee with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEmployee(employeeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Employee so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Employee)

		// slice the ids on comma with no spaces
		ids := strings.Split( shiftAssignmentsIds, ",")

		for _, shiftAssignmentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.ShiftAssignment

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a ShiftAssignment
			// with a matching shiftAssignmentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , shiftAssignmentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ShiftAssignmentObj from the ShiftAssignments array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("ShiftAssignments").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ShiftAssignments", shiftAssignmentsId )
				return utils.RequestResult{false, msg, "removeShiftAssignments", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Employee from the gorm
		//----------------------------------------------------------------------------
		return GetEmployee(employeeId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more correctiveActionsIds as a CorrectiveActions to a Employee
//----------------------------------------------------------------------------
func AddCorrectiveActionsToEmployee ( employeeId uint64, correctiveActionsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Employee with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEmployee(employeeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Employee so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Employee)

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
		// retrieve the modified Employee from the gorm
		//----------------------------------------------------------------------------
		return GetEmployee(employeeId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more correctiveActionsIds as a CorrectiveActions from a Employee
//----------------------------------------------------------------------------
func RemoveCorrectiveActionsFromEmployee( employeeId uint64, correctiveActionsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Employee with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEmployee(employeeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Employee so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Employee)

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
		// retrieve the modified Employee from the gorm
		//----------------------------------------------------------------------------
		return GetEmployee(employeeId)

	} else {
		return parentRequestResult
	}
}

