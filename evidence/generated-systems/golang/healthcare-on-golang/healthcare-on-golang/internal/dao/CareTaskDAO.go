package dao

import (
    "healthcare-on-golang/internal/model"
    "healthcare-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing CareTaskDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateCareTask - creates a new db entry
//----------------------------------------------------------------------------
func CreateCareTask(obj model.CareTask)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a CareTask with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a CareTask", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateCareTask", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetCareTask - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetCareTask(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.CareTask

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a CareTask with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a CareTask using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a CareTask using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetCareTask", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllCareTask - returns all
//----------------------------------------------------------------------------
func GetAllCareTask()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.CareTask

	//----------------------------------------------------------------------------
	// Request the ORM to find all CareTask
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all CareTask" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all CareTask", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllCareTask", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateCareTask - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateCareTask(obj model.CareTask)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a CareTask using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a CareTask using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateCareTask", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteCareTask - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteCareTask(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the CareTask with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetCareTask(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CareTask so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.CareTask)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a CareTask using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a CareTask using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteCareTask", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a CarePlan on a CareTask
//----------------------------------------------------------------------------
func AssignCarePlanToCareTask( careTaskId uint64, carePlanId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the CareTask with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCareTask(careTaskId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CareTask so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CareTask)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.CarePlan

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a CarePlan with a
		// matching carePlanId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, carePlanId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the CarePlan	to the CareTask
			//----------------------------------------------------------------------------
			parentObj.CarePlan = &childObj

			//----------------------------------------------------------------------------
			// save the CareTask
			//----------------------------------------------------------------------------
			return UpdateCareTask(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "CarePlan", carePlanId )
			return utils.RequestResult{false, msg, "assignCarePlan", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a CarePlan on a CareTask
//----------------------------------------------------------------------------
func UnassignCarePlanFromCareTask(careTaskId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the CareTask with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCareTask(careTaskId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CareTask so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CareTask)

		//----------------------------------------------------------------------------
		// assign an empty CarePlan to the CarePlan
		//----------------------------------------------------------------------------
		parentObj.CarePlan = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the CarePlan
		//----------------------------------------------------------------------------
		parentObj.CarePlanId = nil;

		//----------------------------------------------------------------------------
		// save the CareTask
		//----------------------------------------------------------------------------
		return UpdateCareTask(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a AssignedTo on a CareTask
//----------------------------------------------------------------------------
func AssignAssignedToToCareTask( careTaskId uint64, assignedToId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the CareTask with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCareTask(careTaskId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CareTask so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CareTask)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Clinician

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Clinician with a
		// matching assignedToId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, assignedToId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the AssignedTo	to the CareTask
			//----------------------------------------------------------------------------
			parentObj.AssignedTo = &childObj

			//----------------------------------------------------------------------------
			// save the CareTask
			//----------------------------------------------------------------------------
			return UpdateCareTask(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "AssignedTo", assignedToId )
			return utils.RequestResult{false, msg, "assignAssignedTo", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a AssignedTo on a CareTask
//----------------------------------------------------------------------------
func UnassignAssignedToFromCareTask(careTaskId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the CareTask with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCareTask(careTaskId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CareTask so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CareTask)

		//----------------------------------------------------------------------------
		// assign an empty Clinician to the AssignedTo
		//----------------------------------------------------------------------------
		parentObj.AssignedTo = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the AssignedTo
		//----------------------------------------------------------------------------
		parentObj.AssignedToId = nil;

		//----------------------------------------------------------------------------
		// save the CareTask
		//----------------------------------------------------------------------------
		return UpdateCareTask(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Encounter on a CareTask
//----------------------------------------------------------------------------
func AssignEncounterToCareTask( careTaskId uint64, encounterId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the CareTask with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCareTask(careTaskId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CareTask so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CareTask)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Encounter

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Encounter with a
		// matching encounterId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, encounterId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Encounter	to the CareTask
			//----------------------------------------------------------------------------
			parentObj.Encounter = &childObj

			//----------------------------------------------------------------------------
			// save the CareTask
			//----------------------------------------------------------------------------
			return UpdateCareTask(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Encounter", encounterId )
			return utils.RequestResult{false, msg, "assignEncounter", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Encounter on a CareTask
//----------------------------------------------------------------------------
func UnassignEncounterFromCareTask(careTaskId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the CareTask with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetCareTask(careTaskId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.CareTask so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.CareTask)

		//----------------------------------------------------------------------------
		// assign an empty Encounter to the Encounter
		//----------------------------------------------------------------------------
		parentObj.Encounter = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Encounter
		//----------------------------------------------------------------------------
		parentObj.EncounterId = nil;

		//----------------------------------------------------------------------------
		// save the CareTask
		//----------------------------------------------------------------------------
		return UpdateCareTask(parentObj)

	} else {
		return parentRequestResult
	}

}


