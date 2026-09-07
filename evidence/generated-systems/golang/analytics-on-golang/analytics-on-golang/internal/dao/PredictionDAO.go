package dao

import (
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing PredictionDAO..." ) )
}

//----------------------------------------------------------------------------
// CreatePrediction - creates a new db entry
//----------------------------------------------------------------------------
func CreatePrediction(obj model.Prediction)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Prediction with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Prediction", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreatePrediction", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetPrediction - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetPrediction(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Prediction

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Prediction with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Prediction using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Prediction using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetPrediction", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllPrediction - returns all
//----------------------------------------------------------------------------
func GetAllPrediction()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Prediction

	//----------------------------------------------------------------------------
	// Request the ORM to find all Prediction
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Prediction" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Prediction", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllPrediction", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdatePrediction - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdatePrediction(obj model.Prediction)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Prediction using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Prediction using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdatePrediction", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeletePrediction - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeletePrediction(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Prediction with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetPrediction(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Prediction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Prediction)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Prediction using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Prediction using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeletePrediction", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Endpoint on a Prediction
//----------------------------------------------------------------------------
func AssignEndpointToPrediction( predictionId uint64, endpointId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Prediction with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPrediction(predictionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Prediction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Prediction)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.InferenceEndpoint

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a InferenceEndpoint with a
		// matching endpointId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, endpointId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Endpoint	to the Prediction
			//----------------------------------------------------------------------------
			parentObj.Endpoint = &childObj

			//----------------------------------------------------------------------------
			// save the Prediction
			//----------------------------------------------------------------------------
			return UpdatePrediction(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Endpoint", endpointId )
			return utils.RequestResult{false, msg, "assignEndpoint", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Endpoint on a Prediction
//----------------------------------------------------------------------------
func UnassignEndpointFromPrediction(predictionId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Prediction with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPrediction(predictionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Prediction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Prediction)

		//----------------------------------------------------------------------------
		// assign an empty InferenceEndpoint to the Endpoint
		//----------------------------------------------------------------------------
		parentObj.Endpoint = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Endpoint
		//----------------------------------------------------------------------------
		parentObj.EndpointId = nil;

		//----------------------------------------------------------------------------
		// save the Prediction
		//----------------------------------------------------------------------------
		return UpdatePrediction(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a ModelVersion on a Prediction
//----------------------------------------------------------------------------
func AssignModelVersionToPrediction( predictionId uint64, modelVersionId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Prediction with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPrediction(predictionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Prediction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Prediction)

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
			// assign the ModelVersion	to the Prediction
			//----------------------------------------------------------------------------
			parentObj.ModelVersion = &childObj

			//----------------------------------------------------------------------------
			// save the Prediction
			//----------------------------------------------------------------------------
			return UpdatePrediction(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ModelVersion", modelVersionId )
			return utils.RequestResult{false, msg, "assignModelVersion", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a ModelVersion on a Prediction
//----------------------------------------------------------------------------
func UnassignModelVersionFromPrediction(predictionId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Prediction with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPrediction(predictionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Prediction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Prediction)

		//----------------------------------------------------------------------------
		// assign an empty ModelVersion to the ModelVersion
		//----------------------------------------------------------------------------
		parentObj.ModelVersion = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the ModelVersion
		//----------------------------------------------------------------------------
		parentObj.ModelVersionId = nil;

		//----------------------------------------------------------------------------
		// save the Prediction
		//----------------------------------------------------------------------------
		return UpdatePrediction(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Dataset on a Prediction
//----------------------------------------------------------------------------
func AssignDatasetToPrediction( predictionId uint64, datasetId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Prediction with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPrediction(predictionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Prediction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Prediction)

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
			// assign the Dataset	to the Prediction
			//----------------------------------------------------------------------------
			parentObj.Dataset = &childObj

			//----------------------------------------------------------------------------
			// save the Prediction
			//----------------------------------------------------------------------------
			return UpdatePrediction(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Dataset", datasetId )
			return utils.RequestResult{false, msg, "assignDataset", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Dataset on a Prediction
//----------------------------------------------------------------------------
func UnassignDatasetFromPrediction(predictionId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Prediction with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetPrediction(predictionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Prediction so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Prediction)

		//----------------------------------------------------------------------------
		// assign an empty DataSet to the Dataset
		//----------------------------------------------------------------------------
		parentObj.Dataset = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Dataset
		//----------------------------------------------------------------------------
		parentObj.DatasetId = nil;

		//----------------------------------------------------------------------------
		// save the Prediction
		//----------------------------------------------------------------------------
		return UpdatePrediction(parentObj)

	} else {
		return parentRequestResult
	}

}


