package dao

import (
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing ModelVersionDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateModelVersion - creates a new db entry
//----------------------------------------------------------------------------
func CreateModelVersion(obj model.ModelVersion)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a ModelVersion with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a ModelVersion", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateModelVersion", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetModelVersion - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetModelVersion(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.ModelVersion

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a ModelVersion with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a ModelVersion using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a ModelVersion using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetModelVersion", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllModelVersion - returns all
//----------------------------------------------------------------------------
func GetAllModelVersion()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.ModelVersion

	//----------------------------------------------------------------------------
	// Request the ORM to find all ModelVersion
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all ModelVersion" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all ModelVersion", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllModelVersion", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateModelVersion - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateModelVersion(obj model.ModelVersion)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a ModelVersion using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a ModelVersion using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateModelVersion", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteModelVersion - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteModelVersion(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the ModelVersion with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetModelVersion(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ModelVersion so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.ModelVersion)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a ModelVersion using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a ModelVersion using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteModelVersion", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Model_ on a ModelVersion
//----------------------------------------------------------------------------
func AssignModel_ToModelVersion( modelVersionId uint64, model_Id uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the ModelVersion with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetModelVersion(modelVersionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ModelVersion so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ModelVersion)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Model_

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Model_ with a
		// matching model_Id
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, model_Id).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Model_	to the ModelVersion
			//----------------------------------------------------------------------------
			parentObj.Model_ = &childObj

			//----------------------------------------------------------------------------
			// save the ModelVersion
			//----------------------------------------------------------------------------
			return UpdateModelVersion(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Model_", model_Id )
			return utils.RequestResult{false, msg, "assignModel_", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Model_ on a ModelVersion
//----------------------------------------------------------------------------
func UnassignModel_FromModelVersion(modelVersionId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ModelVersion with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetModelVersion(modelVersionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ModelVersion so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ModelVersion)

		//----------------------------------------------------------------------------
		// assign an empty Model_ to the Model_
		//----------------------------------------------------------------------------
		parentObj.Model_ = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Model_
		//----------------------------------------------------------------------------
		parentObj.Model_Id = nil;

		//----------------------------------------------------------------------------
		// save the ModelVersion
		//----------------------------------------------------------------------------
		return UpdateModelVersion(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a TrainingRun on a ModelVersion
//----------------------------------------------------------------------------
func AssignTrainingRunToModelVersion( modelVersionId uint64, trainingRunId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the ModelVersion with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetModelVersion(modelVersionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ModelVersion so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ModelVersion)

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
			// assign the TrainingRun	to the ModelVersion
			//----------------------------------------------------------------------------
			parentObj.TrainingRun = &childObj

			//----------------------------------------------------------------------------
			// save the ModelVersion
			//----------------------------------------------------------------------------
			return UpdateModelVersion(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "TrainingRun", trainingRunId )
			return utils.RequestResult{false, msg, "assignTrainingRun", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a TrainingRun on a ModelVersion
//----------------------------------------------------------------------------
func UnassignTrainingRunFromModelVersion(modelVersionId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ModelVersion with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetModelVersion(modelVersionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ModelVersion so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ModelVersion)

		//----------------------------------------------------------------------------
		// assign an empty TrainingRun to the TrainingRun
		//----------------------------------------------------------------------------
		parentObj.TrainingRun = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the TrainingRun
		//----------------------------------------------------------------------------
		parentObj.TrainingRunId = nil;

		//----------------------------------------------------------------------------
		// save the ModelVersion
		//----------------------------------------------------------------------------
		return UpdateModelVersion(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more evaluationMetricsIds as a EvaluationMetrics to a ModelVersion
//----------------------------------------------------------------------------
func AddEvaluationMetricsToModelVersion ( modelVersionId uint64, evaluationMetricsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ModelVersion with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetModelVersion(modelVersionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ModelVersion so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ModelVersion)

		// slice the ids on comma with no spaces
		ids := strings.Split( evaluationMetricsIds, ",")

		for _, evaluationMetricsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.EvaluationMetric

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a EvaluationMetric
			// with a matching evaluationMetricsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , evaluationMetricsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the EvaluationMetrics using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("EvaluationMetrics").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "EvaluationMetrics", evaluationMetricsId )
				return utils.RequestResult{false, msg, "unassignEvaluationMetrics", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified ModelVersion from the gorm
		//----------------------------------------------------------------------------
		return GetModelVersion(modelVersionId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more evaluationMetricsIds as a EvaluationMetrics from a ModelVersion
//----------------------------------------------------------------------------
func RemoveEvaluationMetricsFromModelVersion( modelVersionId uint64, evaluationMetricsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the ModelVersion with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetModelVersion(modelVersionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ModelVersion so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ModelVersion)

		// slice the ids on comma with no spaces
		ids := strings.Split( evaluationMetricsIds, ",")

		for _, evaluationMetricsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.EvaluationMetric

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a EvaluationMetric
			// with a matching evaluationMetricsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , evaluationMetricsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove EvaluationMetricObj from the EvaluationMetrics array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("EvaluationMetrics").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "EvaluationMetrics", evaluationMetricsId )
				return utils.RequestResult{false, msg, "removeEvaluationMetrics", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified ModelVersion from the gorm
		//----------------------------------------------------------------------------
		return GetModelVersion(modelVersionId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more deploymentsIds as a Deployments to a ModelVersion
//----------------------------------------------------------------------------
func AddDeploymentsToModelVersion ( modelVersionId uint64, deploymentsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ModelVersion with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetModelVersion(modelVersionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ModelVersion so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ModelVersion)

		// slice the ids on comma with no spaces
		ids := strings.Split( deploymentsIds, ",")

		for _, deploymentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.InferenceEndpoint

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a InferenceEndpoint
			// with a matching deploymentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , deploymentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Deployments using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Deployments").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Deployments", deploymentsId )
				return utils.RequestResult{false, msg, "unassignDeployments", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified ModelVersion from the gorm
		//----------------------------------------------------------------------------
		return GetModelVersion(modelVersionId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more deploymentsIds as a Deployments from a ModelVersion
//----------------------------------------------------------------------------
func RemoveDeploymentsFromModelVersion( modelVersionId uint64, deploymentsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the ModelVersion with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetModelVersion(modelVersionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ModelVersion so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ModelVersion)

		// slice the ids on comma with no spaces
		ids := strings.Split( deploymentsIds, ",")

		for _, deploymentsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.InferenceEndpoint

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a InferenceEndpoint
			// with a matching deploymentsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , deploymentsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove InferenceEndpointObj from the Deployments array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Deployments").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Deployments", deploymentsId )
				return utils.RequestResult{false, msg, "removeDeployments", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified ModelVersion from the gorm
		//----------------------------------------------------------------------------
		return GetModelVersion(modelVersionId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more featureSetsIds as a FeatureSets to a ModelVersion
//----------------------------------------------------------------------------
func AddFeatureSetsToModelVersion ( modelVersionId uint64, featureSetsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ModelVersion with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetModelVersion(modelVersionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ModelVersion so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ModelVersion)

		// slice the ids on comma with no spaces
		ids := strings.Split( featureSetsIds, ",")

		for _, featureSetsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.FeatureSet

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a FeatureSet
			// with a matching featureSetsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , featureSetsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the FeatureSets using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("FeatureSets").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "FeatureSets", featureSetsId )
				return utils.RequestResult{false, msg, "unassignFeatureSets", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified ModelVersion from the gorm
		//----------------------------------------------------------------------------
		return GetModelVersion(modelVersionId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more featureSetsIds as a FeatureSets from a ModelVersion
//----------------------------------------------------------------------------
func RemoveFeatureSetsFromModelVersion( modelVersionId uint64, featureSetsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the ModelVersion with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetModelVersion(modelVersionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ModelVersion so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ModelVersion)

		// slice the ids on comma with no spaces
		ids := strings.Split( featureSetsIds, ",")

		for _, featureSetsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.FeatureSet

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a FeatureSet
			// with a matching featureSetsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , featureSetsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove FeatureSetObj from the FeatureSets array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("FeatureSets").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "FeatureSets", featureSetsId )
				return utils.RequestResult{false, msg, "removeFeatureSets", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified ModelVersion from the gorm
		//----------------------------------------------------------------------------
		return GetModelVersion(modelVersionId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more datasetsIds as a Datasets to a ModelVersion
//----------------------------------------------------------------------------
func AddDatasetsToModelVersion ( modelVersionId uint64, datasetsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the ModelVersion with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetModelVersion(modelVersionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ModelVersion so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ModelVersion)

		// slice the ids on comma with no spaces
		ids := strings.Split( datasetsIds, ",")

		for _, datasetsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.DataSet

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a DataSet
			// with a matching datasetsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , datasetsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Datasets using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Datasets").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Datasets", datasetsId )
				return utils.RequestResult{false, msg, "unassignDatasets", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified ModelVersion from the gorm
		//----------------------------------------------------------------------------
		return GetModelVersion(modelVersionId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more datasetsIds as a Datasets from a ModelVersion
//----------------------------------------------------------------------------
func RemoveDatasetsFromModelVersion( modelVersionId uint64, datasetsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the ModelVersion with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetModelVersion(modelVersionId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.ModelVersion so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.ModelVersion)

		// slice the ids on comma with no spaces
		ids := strings.Split( datasetsIds, ",")

		for _, datasetsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.DataSet

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a DataSet
			// with a matching datasetsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , datasetsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove DataSetObj from the Datasets array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Datasets").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Datasets", datasetsId )
				return utils.RequestResult{false, msg, "removeDatasets", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified ModelVersion from the gorm
		//----------------------------------------------------------------------------
		return GetModelVersion(modelVersionId)

	} else {
		return parentRequestResult
	}
}

