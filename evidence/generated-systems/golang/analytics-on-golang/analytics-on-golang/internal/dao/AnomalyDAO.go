package dao

import (
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing AnomalyDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateAnomaly - creates a new db entry
//----------------------------------------------------------------------------
func CreateAnomaly(obj model.Anomaly)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Anomaly with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Anomaly", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateAnomaly", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetAnomaly - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetAnomaly(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Anomaly

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Anomaly with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Anomaly using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Anomaly using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetAnomaly", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllAnomaly - returns all
//----------------------------------------------------------------------------
func GetAllAnomaly()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Anomaly

	//----------------------------------------------------------------------------
	// Request the ORM to find all Anomaly
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Anomaly" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Anomaly", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllAnomaly", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateAnomaly - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateAnomaly(obj model.Anomaly)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Anomaly using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Anomaly using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateAnomaly", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteAnomaly - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteAnomaly(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Anomaly with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetAnomaly(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Anomaly so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Anomaly)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Anomaly using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Anomaly using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteAnomaly", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a TimeSeries on a Anomaly
//----------------------------------------------------------------------------
func AssignTimeSeriesToAnomaly( anomalyId uint64, timeSeriesId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Anomaly with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAnomaly(anomalyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Anomaly so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Anomaly)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.TimeSeries

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a TimeSeries with a
		// matching timeSeriesId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, timeSeriesId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the TimeSeries	to the Anomaly
			//----------------------------------------------------------------------------
			parentObj.TimeSeries = &childObj

			//----------------------------------------------------------------------------
			// save the Anomaly
			//----------------------------------------------------------------------------
			return UpdateAnomaly(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "TimeSeries", timeSeriesId )
			return utils.RequestResult{false, msg, "assignTimeSeries", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a TimeSeries on a Anomaly
//----------------------------------------------------------------------------
func UnassignTimeSeriesFromAnomaly(anomalyId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Anomaly with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAnomaly(anomalyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Anomaly so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Anomaly)

		//----------------------------------------------------------------------------
		// assign an empty TimeSeries to the TimeSeries
		//----------------------------------------------------------------------------
		parentObj.TimeSeries = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the TimeSeries
		//----------------------------------------------------------------------------
		parentObj.TimeSeriesId = nil;

		//----------------------------------------------------------------------------
		// save the Anomaly
		//----------------------------------------------------------------------------
		return UpdateAnomaly(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Alert on a Anomaly
//----------------------------------------------------------------------------
func AssignAlertToAnomaly( anomalyId uint64, alertId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Anomaly with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAnomaly(anomalyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Anomaly so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Anomaly)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Alert

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Alert with a
		// matching alertId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, alertId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Alert	to the Anomaly
			//----------------------------------------------------------------------------
			parentObj.Alert = &childObj

			//----------------------------------------------------------------------------
			// save the Anomaly
			//----------------------------------------------------------------------------
			return UpdateAnomaly(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Alert", alertId )
			return utils.RequestResult{false, msg, "assignAlert", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Alert on a Anomaly
//----------------------------------------------------------------------------
func UnassignAlertFromAnomaly(anomalyId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Anomaly with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAnomaly(anomalyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Anomaly so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Anomaly)

		//----------------------------------------------------------------------------
		// assign an empty Alert to the Alert
		//----------------------------------------------------------------------------
		parentObj.Alert = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Alert
		//----------------------------------------------------------------------------
		parentObj.AlertId = nil;

		//----------------------------------------------------------------------------
		// save the Anomaly
		//----------------------------------------------------------------------------
		return UpdateAnomaly(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Dataset on a Anomaly
//----------------------------------------------------------------------------
func AssignDatasetToAnomaly( anomalyId uint64, datasetId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Anomaly with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAnomaly(anomalyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Anomaly so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Anomaly)

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
			// assign the Dataset	to the Anomaly
			//----------------------------------------------------------------------------
			parentObj.Dataset = &childObj

			//----------------------------------------------------------------------------
			// save the Anomaly
			//----------------------------------------------------------------------------
			return UpdateAnomaly(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Dataset", datasetId )
			return utils.RequestResult{false, msg, "assignDataset", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Dataset on a Anomaly
//----------------------------------------------------------------------------
func UnassignDatasetFromAnomaly(anomalyId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Anomaly with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAnomaly(anomalyId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Anomaly so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Anomaly)

		//----------------------------------------------------------------------------
		// assign an empty DataSet to the Dataset
		//----------------------------------------------------------------------------
		parentObj.Dataset = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Dataset
		//----------------------------------------------------------------------------
		parentObj.DatasetId = nil;

		//----------------------------------------------------------------------------
		// save the Anomaly
		//----------------------------------------------------------------------------
		return UpdateAnomaly(parentObj)

	} else {
		return parentRequestResult
	}

}


