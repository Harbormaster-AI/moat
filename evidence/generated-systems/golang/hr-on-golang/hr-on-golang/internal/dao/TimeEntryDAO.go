package dao

import (
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing TimeEntryDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateTimeEntry - creates a new db entry
//----------------------------------------------------------------------------
func CreateTimeEntry(obj model.TimeEntry)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a TimeEntry with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a TimeEntry", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateTimeEntry", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetTimeEntry - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetTimeEntry(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.TimeEntry

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a TimeEntry with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a TimeEntry using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a TimeEntry using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetTimeEntry", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllTimeEntry - returns all
//----------------------------------------------------------------------------
func GetAllTimeEntry()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.TimeEntry

	//----------------------------------------------------------------------------
	// Request the ORM to find all TimeEntry
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all TimeEntry" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all TimeEntry", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllTimeEntry", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateTimeEntry - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateTimeEntry(obj model.TimeEntry)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a TimeEntry using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a TimeEntry using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateTimeEntry", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteTimeEntry - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteTimeEntry(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the TimeEntry with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetTimeEntry(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TimeEntry so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.TimeEntry)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a TimeEntry using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a TimeEntry using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteTimeEntry", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Timesheet on a TimeEntry
//----------------------------------------------------------------------------
func AssignTimesheetToTimeEntry( timeEntryId uint64, timesheetId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the TimeEntry with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTimeEntry(timeEntryId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TimeEntry so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TimeEntry)

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
			// assign the Timesheet	to the TimeEntry
			//----------------------------------------------------------------------------
			parentObj.Timesheet = &childObj

			//----------------------------------------------------------------------------
			// save the TimeEntry
			//----------------------------------------------------------------------------
			return UpdateTimeEntry(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Timesheet", timesheetId )
			return utils.RequestResult{false, msg, "assignTimesheet", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Timesheet on a TimeEntry
//----------------------------------------------------------------------------
func UnassignTimesheetFromTimeEntry(timeEntryId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the TimeEntry with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTimeEntry(timeEntryId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TimeEntry so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TimeEntry)

		//----------------------------------------------------------------------------
		// assign an empty Timesheet to the Timesheet
		//----------------------------------------------------------------------------
		parentObj.Timesheet = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Timesheet
		//----------------------------------------------------------------------------
		parentObj.TimesheetId = nil;

		//----------------------------------------------------------------------------
		// save the TimeEntry
		//----------------------------------------------------------------------------
		return UpdateTimeEntry(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Employee on a TimeEntry
//----------------------------------------------------------------------------
func AssignEmployeeToTimeEntry( timeEntryId uint64, employeeId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the TimeEntry with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTimeEntry(timeEntryId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TimeEntry so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TimeEntry)

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
			// assign the Employee	to the TimeEntry
			//----------------------------------------------------------------------------
			parentObj.Employee = &childObj

			//----------------------------------------------------------------------------
			// save the TimeEntry
			//----------------------------------------------------------------------------
			return UpdateTimeEntry(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Employee", employeeId )
			return utils.RequestResult{false, msg, "assignEmployee", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Employee on a TimeEntry
//----------------------------------------------------------------------------
func UnassignEmployeeFromTimeEntry(timeEntryId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the TimeEntry with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTimeEntry(timeEntryId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TimeEntry so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TimeEntry)

		//----------------------------------------------------------------------------
		// assign an empty Employee to the Employee
		//----------------------------------------------------------------------------
		parentObj.Employee = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Employee
		//----------------------------------------------------------------------------
		parentObj.EmployeeId = nil;

		//----------------------------------------------------------------------------
		// save the TimeEntry
		//----------------------------------------------------------------------------
		return UpdateTimeEntry(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a CostCenter on a TimeEntry
//----------------------------------------------------------------------------
func AssignCostCenterToTimeEntry( timeEntryId uint64, costCenterId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the TimeEntry with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTimeEntry(timeEntryId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TimeEntry so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TimeEntry)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.CostCenter

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a CostCenter with a
		// matching costCenterId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, costCenterId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the CostCenter	to the TimeEntry
			//----------------------------------------------------------------------------
			parentObj.CostCenter = &childObj

			//----------------------------------------------------------------------------
			// save the TimeEntry
			//----------------------------------------------------------------------------
			return UpdateTimeEntry(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "CostCenter", costCenterId )
			return utils.RequestResult{false, msg, "assignCostCenter", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a CostCenter on a TimeEntry
//----------------------------------------------------------------------------
func UnassignCostCenterFromTimeEntry(timeEntryId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the TimeEntry with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTimeEntry(timeEntryId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TimeEntry so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TimeEntry)

		//----------------------------------------------------------------------------
		// assign an empty CostCenter to the CostCenter
		//----------------------------------------------------------------------------
		parentObj.CostCenter = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the CostCenter
		//----------------------------------------------------------------------------
		parentObj.CostCenterId = nil;

		//----------------------------------------------------------------------------
		// save the TimeEntry
		//----------------------------------------------------------------------------
		return UpdateTimeEntry(parentObj)

	} else {
		return parentRequestResult
	}

}


