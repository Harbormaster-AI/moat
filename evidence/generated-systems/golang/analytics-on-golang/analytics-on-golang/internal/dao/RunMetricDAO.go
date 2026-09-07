package dao

import (
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing RunMetricDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateRunMetric - creates a new db entry
//----------------------------------------------------------------------------
func CreateRunMetric(obj model.RunMetric)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a RunMetric with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a RunMetric", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateRunMetric", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetRunMetric - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetRunMetric(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.RunMetric

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a RunMetric with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a RunMetric using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a RunMetric using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetRunMetric", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllRunMetric - returns all
//----------------------------------------------------------------------------
func GetAllRunMetric()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.RunMetric

	//----------------------------------------------------------------------------
	// Request the ORM to find all RunMetric
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all RunMetric" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all RunMetric", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllRunMetric", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateRunMetric - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateRunMetric(obj model.RunMetric)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a RunMetric using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a RunMetric using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateRunMetric", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteRunMetric - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteRunMetric(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the RunMetric with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetRunMetric(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.RunMetric so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.RunMetric)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a RunMetric using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a RunMetric using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteRunMetric", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a TrainingRun on a RunMetric
//----------------------------------------------------------------------------
func AssignTrainingRunToRunMetric( runMetricId uint64, trainingRunId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the RunMetric with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRunMetric(runMetricId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.RunMetric so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.RunMetric)

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
			// assign the TrainingRun	to the RunMetric
			//----------------------------------------------------------------------------
			parentObj.TrainingRun = &childObj

			//----------------------------------------------------------------------------
			// save the RunMetric
			//----------------------------------------------------------------------------
			return UpdateRunMetric(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "TrainingRun", trainingRunId )
			return utils.RequestResult{false, msg, "assignTrainingRun", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a TrainingRun on a RunMetric
//----------------------------------------------------------------------------
func UnassignTrainingRunFromRunMetric(runMetricId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the RunMetric with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRunMetric(runMetricId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.RunMetric so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.RunMetric)

		//----------------------------------------------------------------------------
		// assign an empty TrainingRun to the TrainingRun
		//----------------------------------------------------------------------------
		parentObj.TrainingRun = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the TrainingRun
		//----------------------------------------------------------------------------
		parentObj.TrainingRunId = nil;

		//----------------------------------------------------------------------------
		// save the RunMetric
		//----------------------------------------------------------------------------
		return UpdateRunMetric(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Metric on a RunMetric
//----------------------------------------------------------------------------
func AssignMetricToRunMetric( runMetricId uint64, metricId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the RunMetric with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRunMetric(runMetricId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.RunMetric so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.RunMetric)

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
			// assign the Metric	to the RunMetric
			//----------------------------------------------------------------------------
			parentObj.Metric = &childObj

			//----------------------------------------------------------------------------
			// save the RunMetric
			//----------------------------------------------------------------------------
			return UpdateRunMetric(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Metric", metricId )
			return utils.RequestResult{false, msg, "assignMetric", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Metric on a RunMetric
//----------------------------------------------------------------------------
func UnassignMetricFromRunMetric(runMetricId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the RunMetric with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRunMetric(runMetricId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.RunMetric so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.RunMetric)

		//----------------------------------------------------------------------------
		// assign an empty Metric to the Metric
		//----------------------------------------------------------------------------
		parentObj.Metric = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Metric
		//----------------------------------------------------------------------------
		parentObj.MetricId = nil;

		//----------------------------------------------------------------------------
		// save the RunMetric
		//----------------------------------------------------------------------------
		return UpdateRunMetric(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Dataset on a RunMetric
//----------------------------------------------------------------------------
func AssignDatasetToRunMetric( runMetricId uint64, datasetId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the RunMetric with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRunMetric(runMetricId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.RunMetric so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.RunMetric)

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
			// assign the Dataset	to the RunMetric
			//----------------------------------------------------------------------------
			parentObj.Dataset = &childObj

			//----------------------------------------------------------------------------
			// save the RunMetric
			//----------------------------------------------------------------------------
			return UpdateRunMetric(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Dataset", datasetId )
			return utils.RequestResult{false, msg, "assignDataset", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Dataset on a RunMetric
//----------------------------------------------------------------------------
func UnassignDatasetFromRunMetric(runMetricId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the RunMetric with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetRunMetric(runMetricId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.RunMetric so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.RunMetric)

		//----------------------------------------------------------------------------
		// assign an empty DataSet to the Dataset
		//----------------------------------------------------------------------------
		parentObj.Dataset = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Dataset
		//----------------------------------------------------------------------------
		parentObj.DatasetId = nil;

		//----------------------------------------------------------------------------
		// save the RunMetric
		//----------------------------------------------------------------------------
		return UpdateRunMetric(parentObj)

	} else {
		return parentRequestResult
	}

}


