package dao

import (
    "healthcare-on-golang/internal/model"
    "healthcare-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing ProcedureDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateProcedure - creates a new db entry
//----------------------------------------------------------------------------
func CreateProcedure(obj model.Procedure)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Procedure with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Procedure", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateProcedure", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetProcedure - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetProcedure(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Procedure

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Procedure with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Procedure using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Procedure using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetProcedure", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllProcedure - returns all
//----------------------------------------------------------------------------
func GetAllProcedure()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Procedure

	//----------------------------------------------------------------------------
	// Request the ORM to find all Procedure
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Procedure" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Procedure", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllProcedure", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateProcedure - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateProcedure(obj model.Procedure)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Procedure using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Procedure using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateProcedure", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteProcedure - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteProcedure(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Procedure with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetProcedure(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Procedure so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Procedure)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Procedure using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Procedure using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteProcedure", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Encounter on a Procedure
//----------------------------------------------------------------------------
func AssignEncounterToProcedure( procedureId uint64, encounterId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Procedure with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetProcedure(procedureId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Procedure so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Procedure)

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
			// assign the Encounter	to the Procedure
			//----------------------------------------------------------------------------
			parentObj.Encounter = &childObj

			//----------------------------------------------------------------------------
			// save the Procedure
			//----------------------------------------------------------------------------
			return UpdateProcedure(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Encounter", encounterId )
			return utils.RequestResult{false, msg, "assignEncounter", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Encounter on a Procedure
//----------------------------------------------------------------------------
func UnassignEncounterFromProcedure(procedureId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Procedure with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetProcedure(procedureId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Procedure so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Procedure)

		//----------------------------------------------------------------------------
		// assign an empty Encounter to the Encounter
		//----------------------------------------------------------------------------
		parentObj.Encounter = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Encounter
		//----------------------------------------------------------------------------
		parentObj.EncounterId = nil;

		//----------------------------------------------------------------------------
		// save the Procedure
		//----------------------------------------------------------------------------
		return UpdateProcedure(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Performer on a Procedure
//----------------------------------------------------------------------------
func AssignPerformerToProcedure( procedureId uint64, performerId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Procedure with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetProcedure(procedureId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Procedure so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Procedure)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Clinician

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Clinician with a
		// matching performerId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, performerId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Performer	to the Procedure
			//----------------------------------------------------------------------------
			parentObj.Performer = &childObj

			//----------------------------------------------------------------------------
			// save the Procedure
			//----------------------------------------------------------------------------
			return UpdateProcedure(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Performer", performerId )
			return utils.RequestResult{false, msg, "assignPerformer", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Performer on a Procedure
//----------------------------------------------------------------------------
func UnassignPerformerFromProcedure(procedureId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Procedure with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetProcedure(procedureId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Procedure so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Procedure)

		//----------------------------------------------------------------------------
		// assign an empty Clinician to the Performer
		//----------------------------------------------------------------------------
		parentObj.Performer = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Performer
		//----------------------------------------------------------------------------
		parentObj.PerformerId = nil;

		//----------------------------------------------------------------------------
		// save the Procedure
		//----------------------------------------------------------------------------
		return UpdateProcedure(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a ProcedureOrder on a Procedure
//----------------------------------------------------------------------------
func AssignProcedureOrderToProcedure( procedureId uint64, procedureOrderId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Procedure with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetProcedure(procedureId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Procedure so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Procedure)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.ProcedureOrder

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a ProcedureOrder with a
		// matching procedureOrderId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, procedureOrderId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the ProcedureOrder	to the Procedure
			//----------------------------------------------------------------------------
			parentObj.ProcedureOrder = &childObj

			//----------------------------------------------------------------------------
			// save the Procedure
			//----------------------------------------------------------------------------
			return UpdateProcedure(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ProcedureOrder", procedureOrderId )
			return utils.RequestResult{false, msg, "assignProcedureOrder", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a ProcedureOrder on a Procedure
//----------------------------------------------------------------------------
func UnassignProcedureOrderFromProcedure(procedureId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Procedure with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetProcedure(procedureId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Procedure so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Procedure)

		//----------------------------------------------------------------------------
		// assign an empty ProcedureOrder to the ProcedureOrder
		//----------------------------------------------------------------------------
		parentObj.ProcedureOrder = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the ProcedureOrder
		//----------------------------------------------------------------------------
		parentObj.ProcedureOrderId = nil;

		//----------------------------------------------------------------------------
		// save the Procedure
		//----------------------------------------------------------------------------
		return UpdateProcedure(parentObj)

	} else {
		return parentRequestResult
	}

}


