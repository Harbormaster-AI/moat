package dao

import (
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing EvaluationMetricDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateEvaluationMetric - creates a new db entry
//----------------------------------------------------------------------------
func CreateEvaluationMetric(obj model.EvaluationMetric)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a EvaluationMetric with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a EvaluationMetric", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateEvaluationMetric", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetEvaluationMetric - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetEvaluationMetric(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.EvaluationMetric

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a EvaluationMetric with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a EvaluationMetric using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a EvaluationMetric using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetEvaluationMetric", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllEvaluationMetric - returns all
//----------------------------------------------------------------------------
func GetAllEvaluationMetric()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.EvaluationMetric

	//----------------------------------------------------------------------------
	// Request the ORM to find all EvaluationMetric
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all EvaluationMetric" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all EvaluationMetric", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllEvaluationMetric", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateEvaluationMetric - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateEvaluationMetric(obj model.EvaluationMetric)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a EvaluationMetric using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a EvaluationMetric using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateEvaluationMetric", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteEvaluationMetric - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteEvaluationMetric(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the EvaluationMetric with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetEvaluationMetric(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.EvaluationMetric so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.EvaluationMetric)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a EvaluationMetric using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a EvaluationMetric using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteEvaluationMetric", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a ModelVersion on a EvaluationMetric
//----------------------------------------------------------------------------
func AssignModelVersionToEvaluationMetric( evaluationMetricId uint64, modelVersionId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the EvaluationMetric with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEvaluationMetric(evaluationMetricId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.EvaluationMetric so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.EvaluationMetric)

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
			// assign the ModelVersion	to the EvaluationMetric
			//----------------------------------------------------------------------------
			parentObj.ModelVersion = &childObj

			//----------------------------------------------------------------------------
			// save the EvaluationMetric
			//----------------------------------------------------------------------------
			return UpdateEvaluationMetric(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ModelVersion", modelVersionId )
			return utils.RequestResult{false, msg, "assignModelVersion", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a ModelVersion on a EvaluationMetric
//----------------------------------------------------------------------------
func UnassignModelVersionFromEvaluationMetric(evaluationMetricId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the EvaluationMetric with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEvaluationMetric(evaluationMetricId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.EvaluationMetric so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.EvaluationMetric)

		//----------------------------------------------------------------------------
		// assign an empty ModelVersion to the ModelVersion
		//----------------------------------------------------------------------------
		parentObj.ModelVersion = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the ModelVersion
		//----------------------------------------------------------------------------
		parentObj.ModelVersionId = nil;

		//----------------------------------------------------------------------------
		// save the EvaluationMetric
		//----------------------------------------------------------------------------
		return UpdateEvaluationMetric(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Metric on a EvaluationMetric
//----------------------------------------------------------------------------
func AssignMetricToEvaluationMetric( evaluationMetricId uint64, metricId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the EvaluationMetric with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEvaluationMetric(evaluationMetricId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.EvaluationMetric so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.EvaluationMetric)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Metric

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Metric with a
		// matching metricId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, metricId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Metric	to the EvaluationMetric
			//----------------------------------------------------------------------------
			parentObj.Metric = &childObj

			//----------------------------------------------------------------------------
			// save the EvaluationMetric
			//----------------------------------------------------------------------------
			return UpdateEvaluationMetric(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Metric", metricId )
			return utils.RequestResult{false, msg, "assignMetric", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Metric on a EvaluationMetric
//----------------------------------------------------------------------------
func UnassignMetricFromEvaluationMetric(evaluationMetricId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the EvaluationMetric with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEvaluationMetric(evaluationMetricId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.EvaluationMetric so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.EvaluationMetric)

		//----------------------------------------------------------------------------
		// assign an empty Metric to the Metric
		//----------------------------------------------------------------------------
		parentObj.Metric = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Metric
		//----------------------------------------------------------------------------
		parentObj.MetricId = nil;

		//----------------------------------------------------------------------------
		// save the EvaluationMetric
		//----------------------------------------------------------------------------
		return UpdateEvaluationMetric(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Dataset on a EvaluationMetric
//----------------------------------------------------------------------------
func AssignDatasetToEvaluationMetric( evaluationMetricId uint64, datasetId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the EvaluationMetric with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEvaluationMetric(evaluationMetricId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.EvaluationMetric so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.EvaluationMetric)

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
			// assign the Dataset	to the EvaluationMetric
			//----------------------------------------------------------------------------
			parentObj.Dataset = &childObj

			//----------------------------------------------------------------------------
			// save the EvaluationMetric
			//----------------------------------------------------------------------------
			return UpdateEvaluationMetric(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Dataset", datasetId )
			return utils.RequestResult{false, msg, "assignDataset", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Dataset on a EvaluationMetric
//----------------------------------------------------------------------------
func UnassignDatasetFromEvaluationMetric(evaluationMetricId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the EvaluationMetric with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetEvaluationMetric(evaluationMetricId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.EvaluationMetric so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.EvaluationMetric)

		//----------------------------------------------------------------------------
		// assign an empty DataSet to the Dataset
		//----------------------------------------------------------------------------
		parentObj.Dataset = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Dataset
		//----------------------------------------------------------------------------
		parentObj.DatasetId = nil;

		//----------------------------------------------------------------------------
		// save the EvaluationMetric
		//----------------------------------------------------------------------------
		return UpdateEvaluationMetric(parentObj)

	} else {
		return parentRequestResult
	}

}


