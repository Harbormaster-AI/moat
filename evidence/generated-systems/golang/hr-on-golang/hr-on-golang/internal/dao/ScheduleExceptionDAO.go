package dao

import (
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing ScheduleExceptionDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateScheduleException - creates a new db entry
//----------------------------------------------------------------------------
func CreateScheduleException(obj model.ScheduleException)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a ScheduleException with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a ScheduleException", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateScheduleException", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetScheduleException - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetScheduleException(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.ScheduleException

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a ScheduleException with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a ScheduleException using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a ScheduleException using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetScheduleException", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllScheduleException - returns all
//----------------------------------------------------------------------------
func GetAllScheduleException()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.ScheduleException

	//----------------------------------------------------------------------------
	// Request the ORM to find all ScheduleException
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all ScheduleException" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all ScheduleException", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllScheduleException", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateScheduleException - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateScheduleException(obj model.ScheduleException)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a ScheduleException using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a ScheduleException using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateScheduleException", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteScheduleException - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteScheduleException(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the ScheduleException with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetScheduleException(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ScheduleException so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.ScheduleException)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a ScheduleException using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a ScheduleException using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteScheduleException", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a WorkSchedule on a ScheduleException
//----------------------------------------------------------------------------
func AssignWorkScheduleToScheduleException( scheduleExceptionId uint64, workScheduleId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the ScheduleException with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetScheduleException(scheduleExceptionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ScheduleException so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ScheduleException)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.WorkSchedule

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a WorkSchedule with a
		// matching workScheduleId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, workScheduleId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the WorkSchedule	to the ScheduleException
			//----------------------------------------------------------------------------
			parentObj.WorkSchedule = &childObj

			//----------------------------------------------------------------------------
			// save the ScheduleException
			//----------------------------------------------------------------------------
			return UpdateScheduleException(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "WorkSchedule", workScheduleId )
			return utils.RequestResult{false, msg, "assignWorkSchedule", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a WorkSchedule on a ScheduleException
//----------------------------------------------------------------------------
func UnassignWorkScheduleFromScheduleException(scheduleExceptionId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ScheduleException with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetScheduleException(scheduleExceptionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ScheduleException so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ScheduleException)

		//----------------------------------------------------------------------------
		// assign an empty WorkSchedule to the WorkSchedule
		//----------------------------------------------------------------------------
		parentObj.WorkSchedule = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the WorkSchedule
		//----------------------------------------------------------------------------
		parentObj.WorkScheduleId = nil;

		//----------------------------------------------------------------------------
		// save the ScheduleException
		//----------------------------------------------------------------------------
		return UpdateScheduleException(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Employee on a ScheduleException
//----------------------------------------------------------------------------
func AssignEmployeeToScheduleException( scheduleExceptionId uint64, employeeId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the ScheduleException with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetScheduleException(scheduleExceptionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ScheduleException so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ScheduleException)

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
			// assign the Employee	to the ScheduleException
			//----------------------------------------------------------------------------
			parentObj.Employee = &childObj

			//----------------------------------------------------------------------------
			// save the ScheduleException
			//----------------------------------------------------------------------------
			return UpdateScheduleException(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Employee", employeeId )
			return utils.RequestResult{false, msg, "assignEmployee", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Employee on a ScheduleException
//----------------------------------------------------------------------------
func UnassignEmployeeFromScheduleException(scheduleExceptionId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ScheduleException with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetScheduleException(scheduleExceptionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ScheduleException so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ScheduleException)

		//----------------------------------------------------------------------------
		// assign an empty Employee to the Employee
		//----------------------------------------------------------------------------
		parentObj.Employee = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Employee
		//----------------------------------------------------------------------------
		parentObj.EmployeeId = nil;

		//----------------------------------------------------------------------------
		// save the ScheduleException
		//----------------------------------------------------------------------------
		return UpdateScheduleException(parentObj)

	} else {
		return parentRequestResult
	}

}


