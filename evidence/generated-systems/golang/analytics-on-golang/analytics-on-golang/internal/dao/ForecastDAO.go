package dao

import (
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing ForecastDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateForecast - creates a new db entry
//----------------------------------------------------------------------------
func CreateForecast(obj model.Forecast)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Forecast with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Forecast", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateForecast", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetForecast - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetForecast(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Forecast

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Forecast with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Forecast using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Forecast using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetForecast", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllForecast - returns all
//----------------------------------------------------------------------------
func GetAllForecast()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Forecast

	//----------------------------------------------------------------------------
	// Request the ORM to find all Forecast
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Forecast" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Forecast", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllForecast", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateForecast - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateForecast(obj model.Forecast)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Forecast using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Forecast using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateForecast", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteForecast - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteForecast(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Forecast with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetForecast(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Forecast so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Forecast)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Forecast using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Forecast using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteForecast", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a ModelVersion on a Forecast
//----------------------------------------------------------------------------
func AssignModelVersionToForecast( forecastId uint64, modelVersionId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Forecast with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetForecast(forecastId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Forecast so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Forecast)

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
			// assign the ModelVersion	to the Forecast
			//----------------------------------------------------------------------------
			parentObj.ModelVersion = &childObj

			//----------------------------------------------------------------------------
			// save the Forecast
			//----------------------------------------------------------------------------
			return UpdateForecast(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ModelVersion", modelVersionId )
			return utils.RequestResult{false, msg, "assignModelVersion", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a ModelVersion on a Forecast
//----------------------------------------------------------------------------
func UnassignModelVersionFromForecast(forecastId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Forecast with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetForecast(forecastId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Forecast so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Forecast)

		//----------------------------------------------------------------------------
		// assign an empty ModelVersion to the ModelVersion
		//----------------------------------------------------------------------------
		parentObj.ModelVersion = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the ModelVersion
		//----------------------------------------------------------------------------
		parentObj.ModelVersionId = nil;

		//----------------------------------------------------------------------------
		// save the Forecast
		//----------------------------------------------------------------------------
		return UpdateForecast(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a TimeSeries on a Forecast
//----------------------------------------------------------------------------
func AssignTimeSeriesToForecast( forecastId uint64, timeSeriesId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Forecast with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetForecast(forecastId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Forecast so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Forecast)

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
			// assign the TimeSeries	to the Forecast
			//----------------------------------------------------------------------------
			parentObj.TimeSeries = &childObj

			//----------------------------------------------------------------------------
			// save the Forecast
			//----------------------------------------------------------------------------
			return UpdateForecast(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "TimeSeries", timeSeriesId )
			return utils.RequestResult{false, msg, "assignTimeSeries", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a TimeSeries on a Forecast
//----------------------------------------------------------------------------
func UnassignTimeSeriesFromForecast(forecastId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Forecast with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetForecast(forecastId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Forecast so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Forecast)

		//----------------------------------------------------------------------------
		// assign an empty TimeSeries to the TimeSeries
		//----------------------------------------------------------------------------
		parentObj.TimeSeries = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the TimeSeries
		//----------------------------------------------------------------------------
		parentObj.TimeSeriesId = nil;

		//----------------------------------------------------------------------------
		// save the Forecast
		//----------------------------------------------------------------------------
		return UpdateForecast(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more datasetsIds as a Datasets to a Forecast
//----------------------------------------------------------------------------
func AddDatasetsToForecast ( forecastId uint64, datasetsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Forecast with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetForecast(forecastId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Forecast so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Forecast)

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
		// retrieve the modified Forecast from the gorm
		//----------------------------------------------------------------------------
		return GetForecast(forecastId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more datasetsIds as a Datasets from a Forecast
//----------------------------------------------------------------------------
func RemoveDatasetsFromForecast( forecastId uint64, datasetsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Forecast with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetForecast(forecastId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Forecast so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Forecast)

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
		// retrieve the modified Forecast from the gorm
		//----------------------------------------------------------------------------
		return GetForecast(forecastId)

	} else {
		return parentRequestResult
	}
}

