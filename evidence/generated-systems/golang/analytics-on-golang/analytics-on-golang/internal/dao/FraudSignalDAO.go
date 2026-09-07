package dao

import (
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing FraudSignalDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateFraudSignal - creates a new db entry
//----------------------------------------------------------------------------
func CreateFraudSignal(obj model.FraudSignal)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a FraudSignal with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a FraudSignal", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateFraudSignal", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetFraudSignal - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetFraudSignal(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.FraudSignal

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a FraudSignal with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a FraudSignal using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a FraudSignal using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetFraudSignal", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllFraudSignal - returns all
//----------------------------------------------------------------------------
func GetAllFraudSignal()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.FraudSignal

	//----------------------------------------------------------------------------
	// Request the ORM to find all FraudSignal
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all FraudSignal" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all FraudSignal", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllFraudSignal", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateFraudSignal - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateFraudSignal(obj model.FraudSignal)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a FraudSignal using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a FraudSignal using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateFraudSignal", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteFraudSignal - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteFraudSignal(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the FraudSignal with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetFraudSignal(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FraudSignal so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.FraudSignal)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a FraudSignal using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a FraudSignal using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteFraudSignal", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Scenario on a FraudSignal
//----------------------------------------------------------------------------
func AssignScenarioToFraudSignal( fraudSignalId uint64, scenarioId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the FraudSignal with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFraudSignal(fraudSignalId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FraudSignal so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.FraudSignal)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.FraudScenario

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a FraudScenario with a
		// matching scenarioId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, scenarioId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Scenario	to the FraudSignal
			//----------------------------------------------------------------------------
			parentObj.Scenario = &childObj

			//----------------------------------------------------------------------------
			// save the FraudSignal
			//----------------------------------------------------------------------------
			return UpdateFraudSignal(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Scenario", scenarioId )
			return utils.RequestResult{false, msg, "assignScenario", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Scenario on a FraudSignal
//----------------------------------------------------------------------------
func UnassignScenarioFromFraudSignal(fraudSignalId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the FraudSignal with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFraudSignal(fraudSignalId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FraudSignal so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.FraudSignal)

		//----------------------------------------------------------------------------
		// assign an empty FraudScenario to the Scenario
		//----------------------------------------------------------------------------
		parentObj.Scenario = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Scenario
		//----------------------------------------------------------------------------
		parentObj.ScenarioId = nil;

		//----------------------------------------------------------------------------
		// save the FraudSignal
		//----------------------------------------------------------------------------
		return UpdateFraudSignal(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Dataset on a FraudSignal
//----------------------------------------------------------------------------
func AssignDatasetToFraudSignal( fraudSignalId uint64, datasetId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the FraudSignal with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFraudSignal(fraudSignalId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FraudSignal so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.FraudSignal)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.DataSet

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a DataSet with a
		// matching datasetId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, datasetId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Dataset	to the FraudSignal
			//----------------------------------------------------------------------------
			parentObj.Dataset = &childObj

			//----------------------------------------------------------------------------
			// save the FraudSignal
			//----------------------------------------------------------------------------
			return UpdateFraudSignal(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Dataset", datasetId )
			return utils.RequestResult{false, msg, "assignDataset", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Dataset on a FraudSignal
//----------------------------------------------------------------------------
func UnassignDatasetFromFraudSignal(fraudSignalId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the FraudSignal with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFraudSignal(fraudSignalId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FraudSignal so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.FraudSignal)

		//----------------------------------------------------------------------------
		// assign an empty DataSet to the Dataset
		//----------------------------------------------------------------------------
		parentObj.Dataset = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Dataset
		//----------------------------------------------------------------------------
		parentObj.DatasetId = nil;

		//----------------------------------------------------------------------------
		// save the FraudSignal
		//----------------------------------------------------------------------------
		return UpdateFraudSignal(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a ModelVersion on a FraudSignal
//----------------------------------------------------------------------------
func AssignModelVersionToFraudSignal( fraudSignalId uint64, modelVersionId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the FraudSignal with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFraudSignal(fraudSignalId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FraudSignal so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.FraudSignal)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.ModelVersion

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a ModelVersion with a
		// matching modelVersionId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, modelVersionId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the ModelVersion	to the FraudSignal
			//----------------------------------------------------------------------------
			parentObj.ModelVersion = &childObj

			//----------------------------------------------------------------------------
			// save the FraudSignal
			//----------------------------------------------------------------------------
			return UpdateFraudSignal(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ModelVersion", modelVersionId )
			return utils.RequestResult{false, msg, "assignModelVersion", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a ModelVersion on a FraudSignal
//----------------------------------------------------------------------------
func UnassignModelVersionFromFraudSignal(fraudSignalId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the FraudSignal with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetFraudSignal(fraudSignalId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.FraudSignal so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.FraudSignal)

		//----------------------------------------------------------------------------
		// assign an empty ModelVersion to the ModelVersion
		//----------------------------------------------------------------------------
		parentObj.ModelVersion = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the ModelVersion
		//----------------------------------------------------------------------------
		parentObj.ModelVersionId = nil;

		//----------------------------------------------------------------------------
		// save the FraudSignal
		//----------------------------------------------------------------------------
		return UpdateFraudSignal(parentObj)

	} else {
		return parentRequestResult
	}

}


