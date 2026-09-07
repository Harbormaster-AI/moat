package dao

import (
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing InferenceEndpointDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateInferenceEndpoint - creates a new db entry
//----------------------------------------------------------------------------
func CreateInferenceEndpoint(obj model.InferenceEndpoint)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a InferenceEndpoint with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a InferenceEndpoint", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateInferenceEndpoint", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetInferenceEndpoint - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetInferenceEndpoint(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.InferenceEndpoint

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a InferenceEndpoint with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a InferenceEndpoint using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a InferenceEndpoint using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetInferenceEndpoint", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllInferenceEndpoint - returns all
//----------------------------------------------------------------------------
func GetAllInferenceEndpoint()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.InferenceEndpoint

	//----------------------------------------------------------------------------
	// Request the ORM to find all InferenceEndpoint
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all InferenceEndpoint" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all InferenceEndpoint", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllInferenceEndpoint", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateInferenceEndpoint - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateInferenceEndpoint(obj model.InferenceEndpoint)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a InferenceEndpoint using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a InferenceEndpoint using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateInferenceEndpoint", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteInferenceEndpoint - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteInferenceEndpoint(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the InferenceEndpoint with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetInferenceEndpoint(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InferenceEndpoint so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.InferenceEndpoint)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a InferenceEndpoint using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a InferenceEndpoint using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteInferenceEndpoint", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a ModelVersion on a InferenceEndpoint
//----------------------------------------------------------------------------
func AssignModelVersionToInferenceEndpoint( inferenceEndpointId uint64, modelVersionId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the InferenceEndpoint with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInferenceEndpoint(inferenceEndpointId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InferenceEndpoint so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InferenceEndpoint)

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
			// assign the ModelVersion	to the InferenceEndpoint
			//----------------------------------------------------------------------------
			parentObj.ModelVersion = &childObj

			//----------------------------------------------------------------------------
			// save the InferenceEndpoint
			//----------------------------------------------------------------------------
			return UpdateInferenceEndpoint(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ModelVersion", modelVersionId )
			return utils.RequestResult{false, msg, "assignModelVersion", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a ModelVersion on a InferenceEndpoint
//----------------------------------------------------------------------------
func UnassignModelVersionFromInferenceEndpoint(inferenceEndpointId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the InferenceEndpoint with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInferenceEndpoint(inferenceEndpointId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InferenceEndpoint so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InferenceEndpoint)

		//----------------------------------------------------------------------------
		// assign an empty ModelVersion to the ModelVersion
		//----------------------------------------------------------------------------
		parentObj.ModelVersion = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the ModelVersion
		//----------------------------------------------------------------------------
		parentObj.ModelVersionId = nil;

		//----------------------------------------------------------------------------
		// save the InferenceEndpoint
		//----------------------------------------------------------------------------
		return UpdateInferenceEndpoint(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Workspace on a InferenceEndpoint
//----------------------------------------------------------------------------
func AssignWorkspaceToInferenceEndpoint( inferenceEndpointId uint64, workspaceId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the InferenceEndpoint with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInferenceEndpoint(inferenceEndpointId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InferenceEndpoint so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InferenceEndpoint)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.AnalyticsWorkspace

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a AnalyticsWorkspace with a
		// matching workspaceId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, workspaceId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Workspace	to the InferenceEndpoint
			//----------------------------------------------------------------------------
			parentObj.Workspace = &childObj

			//----------------------------------------------------------------------------
			// save the InferenceEndpoint
			//----------------------------------------------------------------------------
			return UpdateInferenceEndpoint(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Workspace", workspaceId )
			return utils.RequestResult{false, msg, "assignWorkspace", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Workspace on a InferenceEndpoint
//----------------------------------------------------------------------------
func UnassignWorkspaceFromInferenceEndpoint(inferenceEndpointId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the InferenceEndpoint with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInferenceEndpoint(inferenceEndpointId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InferenceEndpoint so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InferenceEndpoint)

		//----------------------------------------------------------------------------
		// assign an empty AnalyticsWorkspace to the Workspace
		//----------------------------------------------------------------------------
		parentObj.Workspace = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Workspace
		//----------------------------------------------------------------------------
		parentObj.WorkspaceId = nil;

		//----------------------------------------------------------------------------
		// save the InferenceEndpoint
		//----------------------------------------------------------------------------
		return UpdateInferenceEndpoint(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more predictionsIds as a Predictions to a InferenceEndpoint
//----------------------------------------------------------------------------
func AddPredictionsToInferenceEndpoint ( inferenceEndpointId uint64, predictionsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the InferenceEndpoint with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInferenceEndpoint(inferenceEndpointId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InferenceEndpoint so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InferenceEndpoint)

		// slice the ids on comma with no spaces
		ids := strings.Split( predictionsIds, ",")

		for _, predictionsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Prediction

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Prediction
			// with a matching predictionsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , predictionsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Predictions using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Predictions").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Predictions", predictionsId )
				return utils.RequestResult{false, msg, "unassignPredictions", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified InferenceEndpoint from the gorm
		//----------------------------------------------------------------------------
		return GetInferenceEndpoint(inferenceEndpointId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more predictionsIds as a Predictions from a InferenceEndpoint
//----------------------------------------------------------------------------
func RemovePredictionsFromInferenceEndpoint( inferenceEndpointId uint64, predictionsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the InferenceEndpoint with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetInferenceEndpoint(inferenceEndpointId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.InferenceEndpoint so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.InferenceEndpoint)

		// slice the ids on comma with no spaces
		ids := strings.Split( predictionsIds, ",")

		for _, predictionsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Prediction

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Prediction
			// with a matching predictionsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , predictionsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove PredictionObj from the Predictions array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Predictions").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Predictions", predictionsId )
				return utils.RequestResult{false, msg, "removePredictions", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified InferenceEndpoint from the gorm
		//----------------------------------------------------------------------------
		return GetInferenceEndpoint(inferenceEndpointId)

	} else {
		return parentRequestResult
	}
}

