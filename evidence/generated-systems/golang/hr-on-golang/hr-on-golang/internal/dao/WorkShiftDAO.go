package dao

import (
    "hr-on-golang/internal/model"
    "hr-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing WorkShiftDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateWorkShift - creates a new db entry
//----------------------------------------------------------------------------
func CreateWorkShift(obj model.WorkShift)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a WorkShift with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a WorkShift", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateWorkShift", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetWorkShift - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetWorkShift(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.WorkShift

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a WorkShift with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a WorkShift using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a WorkShift using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetWorkShift", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllWorkShift - returns all
//----------------------------------------------------------------------------
func GetAllWorkShift()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.WorkShift

	//----------------------------------------------------------------------------
	// Request the ORM to find all WorkShift
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all WorkShift" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all WorkShift", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllWorkShift", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateWorkShift - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateWorkShift(obj model.WorkShift)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a WorkShift using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a WorkShift using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateWorkShift", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteWorkShift - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteWorkShift(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the WorkShift with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetWorkShift(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.WorkShift so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.WorkShift)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a WorkShift using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a WorkShift using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteWorkShift", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a WorkSchedule on a WorkShift
//----------------------------------------------------------------------------
func AssignWorkScheduleToWorkShift( workShiftId uint64, workScheduleId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the WorkShift with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetWorkShift(workShiftId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.WorkShift so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.WorkShift)

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
			// assign the WorkSchedule	to the WorkShift
			//----------------------------------------------------------------------------
			parentObj.WorkSchedule = &childObj

			//----------------------------------------------------------------------------
			// save the WorkShift
			//----------------------------------------------------------------------------
			return UpdateWorkShift(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "WorkSchedule", workScheduleId )
			return utils.RequestResult{false, msg, "assignWorkSchedule", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a WorkSchedule on a WorkShift
//----------------------------------------------------------------------------
func UnassignWorkScheduleFromWorkShift(workShiftId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the WorkShift with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetWorkShift(workShiftId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.WorkShift so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.WorkShift)

		//----------------------------------------------------------------------------
		// assign an empty WorkSchedule to the WorkSchedule
		//----------------------------------------------------------------------------
		parentObj.WorkSchedule = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the WorkSchedule
		//----------------------------------------------------------------------------
		parentObj.WorkScheduleId = nil;

		//----------------------------------------------------------------------------
		// save the WorkShift
		//----------------------------------------------------------------------------
		return UpdateWorkShift(parentObj)

	} else {
		return parentRequestResult
	}

}


