package dao

import (
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing TimesheetDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateTimesheet - creates a new db entry
//----------------------------------------------------------------------------
func CreateTimesheet(obj model.Timesheet)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Timesheet with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Timesheet", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateTimesheet", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetTimesheet - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetTimesheet(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Timesheet

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Timesheet with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Timesheet using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Timesheet using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetTimesheet", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllTimesheet - returns all
//----------------------------------------------------------------------------
func GetAllTimesheet()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Timesheet

	//----------------------------------------------------------------------------
	// Request the ORM to find all Timesheet
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Timesheet" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Timesheet", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllTimesheet", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateTimesheet - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateTimesheet(obj model.Timesheet)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Timesheet using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Timesheet using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateTimesheet", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteTimesheet - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteTimesheet(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Timesheet with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetTimesheet(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Timesheet so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Timesheet)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Timesheet using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Timesheet using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteTimesheet", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Employee on a Timesheet
//----------------------------------------------------------------------------
func AssignEmployeeToTimesheet( timesheetId uint64, employeeId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Timesheet with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTimesheet(timesheetId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Timesheet so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Timesheet)

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
			// assign the Employee	to the Timesheet
			//----------------------------------------------------------------------------
			parentObj.Employee = &childObj

			//----------------------------------------------------------------------------
			// save the Timesheet
			//----------------------------------------------------------------------------
			return UpdateTimesheet(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Employee", employeeId )
			return utils.RequestResult{false, msg, "assignEmployee", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Employee on a Timesheet
//----------------------------------------------------------------------------
func UnassignEmployeeFromTimesheet(timesheetId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Timesheet with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTimesheet(timesheetId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Timesheet so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Timesheet)

		//----------------------------------------------------------------------------
		// assign an empty Employee to the Employee
		//----------------------------------------------------------------------------
		parentObj.Employee = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Employee
		//----------------------------------------------------------------------------
		parentObj.EmployeeId = nil;

		//----------------------------------------------------------------------------
		// save the Timesheet
		//----------------------------------------------------------------------------
		return UpdateTimesheet(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more timeEntriesIds as a TimeEntries to a Timesheet
//----------------------------------------------------------------------------
func AddTimeEntriesToTimesheet ( timesheetId uint64, timeEntriesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Timesheet with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTimesheet(timesheetId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Timesheet so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Timesheet)

		// slice the ids on comma with no spaces
		ids := strings.Split( timeEntriesIds, ",")

		for _, timeEntriesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.TimeEntry

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a TimeEntry
			// with a matching timeEntriesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , timeEntriesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the TimeEntries using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("TimeEntries").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "TimeEntries", timeEntriesId )
				return utils.RequestResult{false, msg, "unassignTimeEntries", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Timesheet from the gorm
		//----------------------------------------------------------------------------
		return GetTimesheet(timesheetId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more timeEntriesIds as a TimeEntries from a Timesheet
//----------------------------------------------------------------------------
func RemoveTimeEntriesFromTimesheet( timesheetId uint64, timeEntriesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Timesheet with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTimesheet(timesheetId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Timesheet so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Timesheet)

		// slice the ids on comma with no spaces
		ids := strings.Split( timeEntriesIds, ",")

		for _, timeEntriesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.TimeEntry

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a TimeEntry
			// with a matching timeEntriesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , timeEntriesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove TimeEntryObj from the TimeEntries array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("TimeEntries").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "TimeEntries", timeEntriesId )
				return utils.RequestResult{false, msg, "removeTimeEntries", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Timesheet from the gorm
		//----------------------------------------------------------------------------
		return GetTimesheet(timesheetId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more approvalsIds as a Approvals to a Timesheet
//----------------------------------------------------------------------------
func AddApprovalsToTimesheet ( timesheetId uint64, approvalsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Timesheet with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTimesheet(timesheetId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Timesheet so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Timesheet)

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
		// retrieve the modified Timesheet from the gorm
		//----------------------------------------------------------------------------
		return GetTimesheet(timesheetId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more approvalsIds as a Approvals from a Timesheet
//----------------------------------------------------------------------------
func RemoveApprovalsFromTimesheet( timesheetId uint64, approvalsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Timesheet with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTimesheet(timesheetId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Timesheet so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Timesheet)

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
		// retrieve the modified Timesheet from the gorm
		//----------------------------------------------------------------------------
		return GetTimesheet(timesheetId)

	} else {
		return parentRequestResult
	}
}

