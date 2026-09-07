package dao

import (
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing VisualizationDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateVisualization - creates a new db entry
//----------------------------------------------------------------------------
func CreateVisualization(obj model.Visualization)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Visualization with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Visualization", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateVisualization", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetVisualization - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetVisualization(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Visualization

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Visualization with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Visualization using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Visualization using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetVisualization", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllVisualization - returns all
//----------------------------------------------------------------------------
func GetAllVisualization()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Visualization

	//----------------------------------------------------------------------------
	// Request the ORM to find all Visualization
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Visualization" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Visualization", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllVisualization", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateVisualization - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateVisualization(obj model.Visualization)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Visualization using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Visualization using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateVisualization", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteVisualization - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteVisualization(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Visualization with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetVisualization(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Visualization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Visualization)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Visualization using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Visualization using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteVisualization", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Dashboard on a Visualization
//----------------------------------------------------------------------------
func AssignDashboardToVisualization( visualizationId uint64, dashboardId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Visualization with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetVisualization(visualizationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Visualization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Visualization)

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
			// assign the Dashboard	to the Visualization
			//----------------------------------------------------------------------------
			parentObj.Dashboard = &childObj

			//----------------------------------------------------------------------------
			// save the Visualization
			//----------------------------------------------------------------------------
			return UpdateVisualization(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Dashboard", dashboardId )
			return utils.RequestResult{false, msg, "assignDashboard", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Dashboard on a Visualization
//----------------------------------------------------------------------------
func UnassignDashboardFromVisualization(visualizationId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Visualization with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetVisualization(visualizationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Visualization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Visualization)

		//----------------------------------------------------------------------------
		// assign an empty Dashboard to the Dashboard
		//----------------------------------------------------------------------------
		parentObj.Dashboard = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Dashboard
		//----------------------------------------------------------------------------
		parentObj.DashboardId = nil;

		//----------------------------------------------------------------------------
		// save the Visualization
		//----------------------------------------------------------------------------
		return UpdateVisualization(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a Report on a Visualization
//----------------------------------------------------------------------------
func AssignReportToVisualization( visualizationId uint64, reportId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Visualization with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetVisualization(visualizationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Visualization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Visualization)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.Report

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a Report with a
		// matching reportId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, reportId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the Report	to the Visualization
			//----------------------------------------------------------------------------
			parentObj.Report = &childObj

			//----------------------------------------------------------------------------
			// save the Visualization
			//----------------------------------------------------------------------------
			return UpdateVisualization(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Report", reportId )
			return utils.RequestResult{false, msg, "assignReport", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Report on a Visualization
//----------------------------------------------------------------------------
func UnassignReportFromVisualization(visualizationId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Visualization with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetVisualization(visualizationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Visualization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Visualization)

		//----------------------------------------------------------------------------
		// assign an empty Report to the Report
		//----------------------------------------------------------------------------
		parentObj.Report = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Report
		//----------------------------------------------------------------------------
		parentObj.ReportId = nil;

		//----------------------------------------------------------------------------
		// save the Visualization
		//----------------------------------------------------------------------------
		return UpdateVisualization(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more metricsIds as a Metrics to a Visualization
//----------------------------------------------------------------------------
func AddMetricsToVisualization ( visualizationId uint64, metricsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Visualization with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetVisualization(visualizationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Visualization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Visualization)

		// slice the ids on comma with no spaces
		ids := strings.Split( metricsIds, ",")

		for _, metricsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Metric

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Metric
			// with a matching metricsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , metricsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Metrics using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Metrics").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Metrics", metricsId )
				return utils.RequestResult{false, msg, "unassignMetrics", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Visualization from the gorm
		//----------------------------------------------------------------------------
		return GetVisualization(visualizationId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more metricsIds as a Metrics from a Visualization
//----------------------------------------------------------------------------
func RemoveMetricsFromVisualization( visualizationId uint64, metricsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Visualization with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetVisualization(visualizationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Visualization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Visualization)

		// slice the ids on comma with no spaces
		ids := strings.Split( metricsIds, ",")

		for _, metricsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Metric

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Metric
			// with a matching metricsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , metricsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove MetricObj from the Metrics array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Metrics").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Metrics", metricsId )
				return utils.RequestResult{false, msg, "removeMetrics", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Visualization from the gorm
		//----------------------------------------------------------------------------
		return GetVisualization(visualizationId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more dimensionsIds as a Dimensions to a Visualization
//----------------------------------------------------------------------------
func AddDimensionsToVisualization ( visualizationId uint64, dimensionsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Visualization with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetVisualization(visualizationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Visualization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Visualization)

		// slice the ids on comma with no spaces
		ids := strings.Split( dimensionsIds, ",")

		for _, dimensionsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Dimension

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Dimension
			// with a matching dimensionsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , dimensionsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Dimensions using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Dimensions").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Dimensions", dimensionsId )
				return utils.RequestResult{false, msg, "unassignDimensions", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Visualization from the gorm
		//----------------------------------------------------------------------------
		return GetVisualization(visualizationId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more dimensionsIds as a Dimensions from a Visualization
//----------------------------------------------------------------------------
func RemoveDimensionsFromVisualization( visualizationId uint64, dimensionsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Visualization with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetVisualization(visualizationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Visualization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Visualization)

		// slice the ids on comma with no spaces
		ids := strings.Split( dimensionsIds, ",")

		for _, dimensionsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Dimension

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Dimension
			// with a matching dimensionsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , dimensionsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove DimensionObj from the Dimensions array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Dimensions").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Dimensions", dimensionsId )
				return utils.RequestResult{false, msg, "removeDimensions", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Visualization from the gorm
		//----------------------------------------------------------------------------
		return GetVisualization(visualizationId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more datasetsIds as a Datasets to a Visualization
//----------------------------------------------------------------------------
func AddDatasetsToVisualization ( visualizationId uint64, datasetsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Visualization with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetVisualization(visualizationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Visualization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Visualization)

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
		// retrieve the modified Visualization from the gorm
		//----------------------------------------------------------------------------
		return GetVisualization(visualizationId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more datasetsIds as a Datasets from a Visualization
//----------------------------------------------------------------------------
func RemoveDatasetsFromVisualization( visualizationId uint64, datasetsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Visualization with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetVisualization(visualizationId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Visualization so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Visualization)

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
		// retrieve the modified Visualization from the gorm
		//----------------------------------------------------------------------------
		return GetVisualization(visualizationId)

	} else {
		return parentRequestResult
	}
}

