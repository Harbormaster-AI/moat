package dao

import (
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing TimeSeriesDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateTimeSeries - creates a new db entry
//----------------------------------------------------------------------------
func CreateTimeSeries(obj model.TimeSeries)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a TimeSeries with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a TimeSeries", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateTimeSeries", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetTimeSeries - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetTimeSeries(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.TimeSeries

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a TimeSeries with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a TimeSeries using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a TimeSeries using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetTimeSeries", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllTimeSeries - returns all
//----------------------------------------------------------------------------
func GetAllTimeSeries()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.TimeSeries

	//----------------------------------------------------------------------------
	// Request the ORM to find all TimeSeries
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all TimeSeries" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all TimeSeries", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllTimeSeries", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateTimeSeries - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateTimeSeries(obj model.TimeSeries)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a TimeSeries using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a TimeSeries using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateTimeSeries", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteTimeSeries - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteTimeSeries(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the TimeSeries with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetTimeSeries(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TimeSeries so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.TimeSeries)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a TimeSeries using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a TimeSeries using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteTimeSeries", requestResult.Data}

	}

	return requestResult
}



//----------------------------------------------------------------------------
// adds one or more datasetsIds as a Datasets to a TimeSeries
//----------------------------------------------------------------------------
func AddDatasetsToTimeSeries ( timeSeriesId uint64, datasetsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the TimeSeries with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTimeSeries(timeSeriesId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TimeSeries so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TimeSeries)

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
		// retrieve the modified TimeSeries from the gorm
		//----------------------------------------------------------------------------
		return GetTimeSeries(timeSeriesId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more datasetsIds as a Datasets from a TimeSeries
//----------------------------------------------------------------------------
func RemoveDatasetsFromTimeSeries( timeSeriesId uint64, datasetsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the TimeSeries with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTimeSeries(timeSeriesId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TimeSeries so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TimeSeries)

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
		// retrieve the modified TimeSeries from the gorm
		//----------------------------------------------------------------------------
		return GetTimeSeries(timeSeriesId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more forecastsIds as a Forecasts to a TimeSeries
//----------------------------------------------------------------------------
func AddForecastsToTimeSeries ( timeSeriesId uint64, forecastsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the TimeSeries with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTimeSeries(timeSeriesId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TimeSeries so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TimeSeries)

		// slice the ids on comma with no spaces
		ids := strings.Split( forecastsIds, ",")

		for _, forecastsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Forecast

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Forecast
			// with a matching forecastsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , forecastsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Forecasts using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Forecasts").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Forecasts", forecastsId )
				return utils.RequestResult{false, msg, "unassignForecasts", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified TimeSeries from the gorm
		//----------------------------------------------------------------------------
		return GetTimeSeries(timeSeriesId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more forecastsIds as a Forecasts from a TimeSeries
//----------------------------------------------------------------------------
func RemoveForecastsFromTimeSeries( timeSeriesId uint64, forecastsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the TimeSeries with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTimeSeries(timeSeriesId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TimeSeries so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TimeSeries)

		// slice the ids on comma with no spaces
		ids := strings.Split( forecastsIds, ",")

		for _, forecastsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Forecast

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Forecast
			// with a matching forecastsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , forecastsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ForecastObj from the Forecasts array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Forecasts").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Forecasts", forecastsId )
				return utils.RequestResult{false, msg, "removeForecasts", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified TimeSeries from the gorm
		//----------------------------------------------------------------------------
		return GetTimeSeries(timeSeriesId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more anomaliesIds as a Anomalies to a TimeSeries
//----------------------------------------------------------------------------
func AddAnomaliesToTimeSeries ( timeSeriesId uint64, anomaliesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the TimeSeries with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTimeSeries(timeSeriesId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TimeSeries so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TimeSeries)

		// slice the ids on comma with no spaces
		ids := strings.Split( anomaliesIds, ",")

		for _, anomaliesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Anomaly

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Anomaly
			// with a matching anomaliesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , anomaliesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Anomalies using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Anomalies").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Anomalies", anomaliesId )
				return utils.RequestResult{false, msg, "unassignAnomalies", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified TimeSeries from the gorm
		//----------------------------------------------------------------------------
		return GetTimeSeries(timeSeriesId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more anomaliesIds as a Anomalies from a TimeSeries
//----------------------------------------------------------------------------
func RemoveAnomaliesFromTimeSeries( timeSeriesId uint64, anomaliesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the TimeSeries with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetTimeSeries(timeSeriesId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.TimeSeries so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.TimeSeries)

		// slice the ids on comma with no spaces
		ids := strings.Split( anomaliesIds, ",")

		for _, anomaliesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Anomaly

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Anomaly
			// with a matching anomaliesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , anomaliesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove AnomalyObj from the Anomalies array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Anomalies").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Anomalies", anomaliesId )
				return utils.RequestResult{false, msg, "removeAnomalies", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified TimeSeries from the gorm
		//----------------------------------------------------------------------------
		return GetTimeSeries(timeSeriesId)

	} else {
		return parentRequestResult
	}
}

