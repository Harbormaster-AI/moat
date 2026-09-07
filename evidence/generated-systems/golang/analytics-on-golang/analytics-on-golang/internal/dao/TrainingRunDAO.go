package dao

import (
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing TrainingRunDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateTrainingRun - creates a new db entry
//----------------------------------------------------------------------------
func CreateTrainingRun(obj model.TrainingRun)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a TrainingRun with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a TrainingRun", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateTrainingRun", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetTrainingRun - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetTrainingRun(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.TrainingRun

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a TrainingRun with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a TrainingRun using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a TrainingRun using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetTrainingRun", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllTrainingRun - returns all
//----------------------------------------------------------------------------
func GetAllTrainingRun()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.TrainingRun

	//----------------------------------------------------------------------------
	// Request the ORM to find all TrainingRun
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all TrainingRun" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all TrainingRun", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllTrainingRun", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateTrainingRun - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateTrainingRun(obj model.TrainingRun)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a TrainingRun using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a TrainingRun using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateTrainingRun", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteTrainingRun - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteTrainingRun(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the TrainingRun with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetTrainingRun(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TrainingRun so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.TrainingRun)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a TrainingRun using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a TrainingRun using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteTrainingRun", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Experiment on a TrainingRun
//----------------------------------------------------------------------------
func AssignExperimentToTrainingRun( trainingRunId uint64, experimentId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the TrainingRun with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTrainingRun(trainingRunId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TrainingRun so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TrainingRun)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Experiment

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Experiment with a
		// matching experimentId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, experimentId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Experiment	to the TrainingRun
			//----------------------------------------------------------------------------
			parentObj.Experiment = &childObj

			//----------------------------------------------------------------------------
			// save the TrainingRun
			//----------------------------------------------------------------------------
			return UpdateTrainingRun(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Experiment", experimentId )
			return utils.RequestResult{false, msg, "assignExperiment", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Experiment on a TrainingRun
//----------------------------------------------------------------------------
func UnassignExperimentFromTrainingRun(trainingRunId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the TrainingRun with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTrainingRun(trainingRunId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TrainingRun so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TrainingRun)

		//----------------------------------------------------------------------------
		// assign an empty Experiment to the Experiment
		//----------------------------------------------------------------------------
		parentObj.Experiment = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Experiment
		//----------------------------------------------------------------------------
		parentObj.ExperimentId = nil;

		//----------------------------------------------------------------------------
		// save the TrainingRun
		//----------------------------------------------------------------------------
		return UpdateTrainingRun(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a ModelVersion on a TrainingRun
//----------------------------------------------------------------------------
func AssignModelVersionToTrainingRun( trainingRunId uint64, modelVersionId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the TrainingRun with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTrainingRun(trainingRunId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TrainingRun so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TrainingRun)

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
			// assign the ModelVersion	to the TrainingRun
			//----------------------------------------------------------------------------
			parentObj.ModelVersion = &childObj

			//----------------------------------------------------------------------------
			// save the TrainingRun
			//----------------------------------------------------------------------------
			return UpdateTrainingRun(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ModelVersion", modelVersionId )
			return utils.RequestResult{false, msg, "assignModelVersion", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a ModelVersion on a TrainingRun
//----------------------------------------------------------------------------
func UnassignModelVersionFromTrainingRun(trainingRunId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the TrainingRun with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTrainingRun(trainingRunId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TrainingRun so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TrainingRun)

		//----------------------------------------------------------------------------
		// assign an empty ModelVersion to the ModelVersion
		//----------------------------------------------------------------------------
		parentObj.ModelVersion = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the ModelVersion
		//----------------------------------------------------------------------------
		parentObj.ModelVersionId = nil;

		//----------------------------------------------------------------------------
		// save the TrainingRun
		//----------------------------------------------------------------------------
		return UpdateTrainingRun(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more inputDatasetsIds as a InputDatasets to a TrainingRun
//----------------------------------------------------------------------------
func AddInputDatasetsToTrainingRun ( trainingRunId uint64, inputDatasetsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the TrainingRun with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTrainingRun(trainingRunId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TrainingRun so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TrainingRun)

		// slice the ids on comma with no spaces
		ids := strings.Split( inputDatasetsIds, ",")

		for _, inputDatasetsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.DataSet

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a DataSet
			// with a matching inputDatasetsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , inputDatasetsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the InputDatasets using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("InputDatasets").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "InputDatasets", inputDatasetsId )
				return utils.RequestResult{false, msg, "unassignInputDatasets", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified TrainingRun from the gorm
		//----------------------------------------------------------------------------
		return GetTrainingRun(trainingRunId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more inputDatasetsIds as a InputDatasets from a TrainingRun
//----------------------------------------------------------------------------
func RemoveInputDatasetsFromTrainingRun( trainingRunId uint64, inputDatasetsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the TrainingRun with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTrainingRun(trainingRunId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TrainingRun so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TrainingRun)

		// slice the ids on comma with no spaces
		ids := strings.Split( inputDatasetsIds, ",")

		for _, inputDatasetsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.DataSet

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a DataSet
			// with a matching inputDatasetsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , inputDatasetsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove DataSetObj from the InputDatasets array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("InputDatasets").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "InputDatasets", inputDatasetsId )
				return utils.RequestResult{false, msg, "removeInputDatasets", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified TrainingRun from the gorm
		//----------------------------------------------------------------------------
		return GetTrainingRun(trainingRunId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more featuresIds as a Features to a TrainingRun
//----------------------------------------------------------------------------
func AddFeaturesToTrainingRun ( trainingRunId uint64, featuresIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the TrainingRun with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTrainingRun(trainingRunId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TrainingRun so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TrainingRun)

		// slice the ids on comma with no spaces
		ids := strings.Split( featuresIds, ",")

		for _, featuresId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Feature

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Feature
			// with a matching featuresId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , featuresId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Features using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Features").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Features", featuresId )
				return utils.RequestResult{false, msg, "unassignFeatures", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified TrainingRun from the gorm
		//----------------------------------------------------------------------------
		return GetTrainingRun(trainingRunId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more featuresIds as a Features from a TrainingRun
//----------------------------------------------------------------------------
func RemoveFeaturesFromTrainingRun( trainingRunId uint64, featuresIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the TrainingRun with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTrainingRun(trainingRunId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TrainingRun so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TrainingRun)

		// slice the ids on comma with no spaces
		ids := strings.Split( featuresIds, ",")

		for _, featuresId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Feature

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Feature
			// with a matching featuresId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , featuresId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove FeatureObj from the Features array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Features").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Features", featuresId )
				return utils.RequestResult{false, msg, "removeFeatures", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified TrainingRun from the gorm
		//----------------------------------------------------------------------------
		return GetTrainingRun(trainingRunId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more runMetricsIds as a RunMetrics to a TrainingRun
//----------------------------------------------------------------------------
func AddRunMetricsToTrainingRun ( trainingRunId uint64, runMetricsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the TrainingRun with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTrainingRun(trainingRunId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TrainingRun so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TrainingRun)

		// slice the ids on comma with no spaces
		ids := strings.Split( runMetricsIds, ",")

		for _, runMetricsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.RunMetric

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a RunMetric
			// with a matching runMetricsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , runMetricsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the RunMetrics using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("RunMetrics").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "RunMetrics", runMetricsId )
				return utils.RequestResult{false, msg, "unassignRunMetrics", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified TrainingRun from the gorm
		//----------------------------------------------------------------------------
		return GetTrainingRun(trainingRunId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more runMetricsIds as a RunMetrics from a TrainingRun
//----------------------------------------------------------------------------
func RemoveRunMetricsFromTrainingRun( trainingRunId uint64, runMetricsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the TrainingRun with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTrainingRun(trainingRunId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TrainingRun so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TrainingRun)

		// slice the ids on comma with no spaces
		ids := strings.Split( runMetricsIds, ",")

		for _, runMetricsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.RunMetric

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a RunMetric
			// with a matching runMetricsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , runMetricsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove RunMetricObj from the RunMetrics array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("RunMetrics").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "RunMetrics", runMetricsId )
				return utils.RequestResult{false, msg, "removeRunMetrics", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified TrainingRun from the gorm
		//----------------------------------------------------------------------------
		return GetTrainingRun(trainingRunId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more runParametersIds as a RunParameters to a TrainingRun
//----------------------------------------------------------------------------
func AddRunParametersToTrainingRun ( trainingRunId uint64, runParametersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the TrainingRun with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTrainingRun(trainingRunId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TrainingRun so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TrainingRun)

		// slice the ids on comma with no spaces
		ids := strings.Split( runParametersIds, ",")

		for _, runParametersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.RunParameter

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a RunParameter
			// with a matching runParametersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , runParametersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the RunParameters using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("RunParameters").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "RunParameters", runParametersId )
				return utils.RequestResult{false, msg, "unassignRunParameters", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified TrainingRun from the gorm
		//----------------------------------------------------------------------------
		return GetTrainingRun(trainingRunId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more runParametersIds as a RunParameters from a TrainingRun
//----------------------------------------------------------------------------
func RemoveRunParametersFromTrainingRun( trainingRunId uint64, runParametersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the TrainingRun with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTrainingRun(trainingRunId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TrainingRun so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TrainingRun)

		// slice the ids on comma with no spaces
		ids := strings.Split( runParametersIds, ",")

		for _, runParametersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.RunParameter

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a RunParameter
			// with a matching runParametersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , runParametersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove RunParameterObj from the RunParameters array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("RunParameters").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "RunParameters", runParametersId )
				return utils.RequestResult{false, msg, "removeRunParameters", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified TrainingRun from the gorm
		//----------------------------------------------------------------------------
		return GetTrainingRun(trainingRunId)

	} else {
		return parentRequestResult
	}
}

