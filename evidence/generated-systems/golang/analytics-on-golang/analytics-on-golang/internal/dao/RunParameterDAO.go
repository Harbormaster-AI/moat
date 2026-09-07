package dao

import (
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing RunParameterDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateRunParameter - creates a new db entry
//----------------------------------------------------------------------------
func CreateRunParameter(obj model.RunParameter)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a RunParameter with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a RunParameter", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateRunParameter", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetRunParameter - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetRunParameter(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.RunParameter

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a RunParameter with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a RunParameter using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a RunParameter using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetRunParameter", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllRunParameter - returns all
//----------------------------------------------------------------------------
func GetAllRunParameter()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.RunParameter

	//----------------------------------------------------------------------------
	// Request the ORM to find all RunParameter
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all RunParameter" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all RunParameter", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllRunParameter", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateRunParameter - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateRunParameter(obj model.RunParameter)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a RunParameter using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a RunParameter using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateRunParameter", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteRunParameter - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteRunParameter(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the RunParameter with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetRunParameter(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.RunParameter so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.RunParameter)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a RunParameter using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a RunParameter using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteRunParameter", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a TrainingRun on a RunParameter
//----------------------------------------------------------------------------
func AssignTrainingRunToRunParameter( runParameterId uint64, trainingRunId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the RunParameter with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRunParameter(runParameterId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.RunParameter so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.RunParameter)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.TrainingRun

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a TrainingRun with a
		// matching trainingRunId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, trainingRunId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the TrainingRun	to the RunParameter
			//----------------------------------------------------------------------------
			parentObj.TrainingRun = &childObj

			//----------------------------------------------------------------------------
			// save the RunParameter
			//----------------------------------------------------------------------------
			return UpdateRunParameter(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "TrainingRun", trainingRunId )
			return utils.RequestResult{false, msg, "assignTrainingRun", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a TrainingRun on a RunParameter
//----------------------------------------------------------------------------
func UnassignTrainingRunFromRunParameter(runParameterId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the RunParameter with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRunParameter(runParameterId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.RunParameter so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.RunParameter)

		//----------------------------------------------------------------------------
		// assign an empty TrainingRun to the TrainingRun
		//----------------------------------------------------------------------------
		parentObj.TrainingRun = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the TrainingRun
		//----------------------------------------------------------------------------
		parentObj.TrainingRunId = nil;

		//----------------------------------------------------------------------------
		// save the RunParameter
		//----------------------------------------------------------------------------
		return UpdateRunParameter(parentObj)

	} else {
		return parentRequestResult
	}

}


