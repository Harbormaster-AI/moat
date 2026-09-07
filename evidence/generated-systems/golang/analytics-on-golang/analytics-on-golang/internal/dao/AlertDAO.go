package dao

import (
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing AlertDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateAlert - creates a new db entry
//----------------------------------------------------------------------------
func CreateAlert(obj model.Alert)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Alert with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Alert", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateAlert", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetAlert - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetAlert(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Alert

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Alert with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Alert using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Alert using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetAlert", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllAlert - returns all
//----------------------------------------------------------------------------
func GetAllAlert()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Alert

	//----------------------------------------------------------------------------
	// Request the ORM to find all Alert
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Alert" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Alert", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllAlert", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateAlert - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateAlert(obj model.Alert)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Alert using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Alert using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateAlert", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteAlert - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteAlert(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Alert with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetAlert(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Alert so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Alert)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Alert using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Alert using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteAlert", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Metric on a Alert
//----------------------------------------------------------------------------
func AssignMetricToAlert( alertId uint64, metricId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Alert with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAlert(alertId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Alert so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Alert)

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
			// assign the Metric	to the Alert
			//----------------------------------------------------------------------------
			parentObj.Metric = &childObj

			//----------------------------------------------------------------------------
			// save the Alert
			//----------------------------------------------------------------------------
			return UpdateAlert(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Metric", metricId )
			return utils.RequestResult{false, msg, "assignMetric", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Metric on a Alert
//----------------------------------------------------------------------------
func UnassignMetricFromAlert(alertId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Alert with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAlert(alertId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Alert so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Alert)

		//----------------------------------------------------------------------------
		// assign an empty Metric to the Metric
		//----------------------------------------------------------------------------
		parentObj.Metric = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Metric
		//----------------------------------------------------------------------------
		parentObj.MetricId = nil;

		//----------------------------------------------------------------------------
		// save the Alert
		//----------------------------------------------------------------------------
		return UpdateAlert(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Dashboard on a Alert
//----------------------------------------------------------------------------
func AssignDashboardToAlert( alertId uint64, dashboardId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Alert with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAlert(alertId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Alert so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Alert)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Dashboard

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Dashboard with a
		// matching dashboardId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, dashboardId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Dashboard	to the Alert
			//----------------------------------------------------------------------------
			parentObj.Dashboard = &childObj

			//----------------------------------------------------------------------------
			// save the Alert
			//----------------------------------------------------------------------------
			return UpdateAlert(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Dashboard", dashboardId )
			return utils.RequestResult{false, msg, "assignDashboard", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Dashboard on a Alert
//----------------------------------------------------------------------------
func UnassignDashboardFromAlert(alertId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Alert with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAlert(alertId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Alert so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Alert)

		//----------------------------------------------------------------------------
		// assign an empty Dashboard to the Dashboard
		//----------------------------------------------------------------------------
		parentObj.Dashboard = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Dashboard
		//----------------------------------------------------------------------------
		parentObj.DashboardId = nil;

		//----------------------------------------------------------------------------
		// save the Alert
		//----------------------------------------------------------------------------
		return UpdateAlert(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Dataset on a Alert
//----------------------------------------------------------------------------
func AssignDatasetToAlert( alertId uint64, datasetId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Alert with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAlert(alertId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Alert so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Alert)

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
			// assign the Dataset	to the Alert
			//----------------------------------------------------------------------------
			parentObj.Dataset = &childObj

			//----------------------------------------------------------------------------
			// save the Alert
			//----------------------------------------------------------------------------
			return UpdateAlert(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Dataset", datasetId )
			return utils.RequestResult{false, msg, "assignDataset", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Dataset on a Alert
//----------------------------------------------------------------------------
func UnassignDatasetFromAlert(alertId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Alert with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAlert(alertId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Alert so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Alert)

		//----------------------------------------------------------------------------
		// assign an empty DataSet to the Dataset
		//----------------------------------------------------------------------------
		parentObj.Dataset = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Dataset
		//----------------------------------------------------------------------------
		parentObj.DatasetId = nil;

		//----------------------------------------------------------------------------
		// save the Alert
		//----------------------------------------------------------------------------
		return UpdateAlert(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Rule on a Alert
//----------------------------------------------------------------------------
func AssignRuleToAlert( alertId uint64, ruleId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Alert with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAlert(alertId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Alert so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Alert)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.QualityRule

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a QualityRule with a
		// matching ruleId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, ruleId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Rule	to the Alert
			//----------------------------------------------------------------------------
			parentObj.Rule = &childObj

			//----------------------------------------------------------------------------
			// save the Alert
			//----------------------------------------------------------------------------
			return UpdateAlert(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Rule", ruleId )
			return utils.RequestResult{false, msg, "assignRule", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Rule on a Alert
//----------------------------------------------------------------------------
func UnassignRuleFromAlert(alertId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Alert with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAlert(alertId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Alert so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Alert)

		//----------------------------------------------------------------------------
		// assign an empty QualityRule to the Rule
		//----------------------------------------------------------------------------
		parentObj.Rule = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Rule
		//----------------------------------------------------------------------------
		parentObj.RuleId = nil;

		//----------------------------------------------------------------------------
		// save the Alert
		//----------------------------------------------------------------------------
		return UpdateAlert(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more anomaliesIds as a Anomalies to a Alert
//----------------------------------------------------------------------------
func AddAnomaliesToAlert ( alertId uint64, anomaliesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Alert with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAlert(alertId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Alert so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Alert)

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
		// retrieve the modified Alert from the gorm
		//----------------------------------------------------------------------------
		return GetAlert(alertId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more anomaliesIds as a Anomalies from a Alert
//----------------------------------------------------------------------------
func RemoveAnomaliesFromAlert( alertId uint64, anomaliesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Alert with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAlert(alertId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Alert so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Alert)

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
		// retrieve the modified Alert from the gorm
		//----------------------------------------------------------------------------
		return GetAlert(alertId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more subscribersIds as a Subscribers to a Alert
//----------------------------------------------------------------------------
func AddSubscribersToAlert ( alertId uint64, subscribersIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Alert with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAlert(alertId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Alert so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Alert)

		// slice the ids on comma with no spaces
		ids := strings.Split( subscribersIds, ",")

		for _, subscribersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Subscriber

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Subscriber
			// with a matching subscribersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , subscribersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Subscribers using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Subscribers").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Subscribers", subscribersId )
				return utils.RequestResult{false, msg, "unassignSubscribers", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Alert from the gorm
		//----------------------------------------------------------------------------
		return GetAlert(alertId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more subscribersIds as a Subscribers from a Alert
//----------------------------------------------------------------------------
func RemoveSubscribersFromAlert( alertId uint64, subscribersIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Alert with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAlert(alertId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Alert so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Alert)

		// slice the ids on comma with no spaces
		ids := strings.Split( subscribersIds, ",")

		for _, subscribersId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Subscriber

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Subscriber
			// with a matching subscribersId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , subscribersId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove SubscriberObj from the Subscribers array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Subscribers").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Subscribers", subscribersId )
				return utils.RequestResult{false, msg, "removeSubscribers", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Alert from the gorm
		//----------------------------------------------------------------------------
		return GetAlert(alertId)

	} else {
		return parentRequestResult
	}
}

