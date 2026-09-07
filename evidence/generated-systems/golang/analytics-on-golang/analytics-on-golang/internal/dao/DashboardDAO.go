package dao

import (
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing DashboardDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateDashboard - creates a new db entry
//----------------------------------------------------------------------------
func CreateDashboard(obj model.Dashboard)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Dashboard with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Dashboard", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateDashboard", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetDashboard - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetDashboard(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Dashboard

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Dashboard with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Dashboard using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Dashboard using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetDashboard", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllDashboard - returns all
//----------------------------------------------------------------------------
func GetAllDashboard()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Dashboard

	//----------------------------------------------------------------------------
	// Request the ORM to find all Dashboard
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Dashboard" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Dashboard", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllDashboard", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateDashboard - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateDashboard(obj model.Dashboard)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Dashboard using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Dashboard using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateDashboard", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteDashboard - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteDashboard(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Dashboard with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetDashboard(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Dashboard so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Dashboard)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Dashboard using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Dashboard using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteDashboard", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Workspace on a Dashboard
//----------------------------------------------------------------------------
func AssignWorkspaceToDashboard( dashboardId uint64, workspaceId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Dashboard with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDashboard(dashboardId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Dashboard so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Dashboard)

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
			// assign the Workspace	to the Dashboard
			//----------------------------------------------------------------------------
			parentObj.Workspace = &childObj

			//----------------------------------------------------------------------------
			// save the Dashboard
			//----------------------------------------------------------------------------
			return UpdateDashboard(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Workspace", workspaceId )
			return utils.RequestResult{false, msg, "assignWorkspace", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Workspace on a Dashboard
//----------------------------------------------------------------------------
func UnassignWorkspaceFromDashboard(dashboardId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Dashboard with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDashboard(dashboardId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Dashboard so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Dashboard)

		//----------------------------------------------------------------------------
		// assign an empty AnalyticsWorkspace to the Workspace
		//----------------------------------------------------------------------------
		parentObj.Workspace = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Workspace
		//----------------------------------------------------------------------------
		parentObj.WorkspaceId = nil;

		//----------------------------------------------------------------------------
		// save the Dashboard
		//----------------------------------------------------------------------------
		return UpdateDashboard(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more visualizationsIds as a Visualizations to a Dashboard
//----------------------------------------------------------------------------
func AddVisualizationsToDashboard ( dashboardId uint64, visualizationsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Dashboard with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDashboard(dashboardId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Dashboard so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Dashboard)

		// slice the ids on comma with no spaces
		ids := strings.Split( visualizationsIds, ",")

		for _, visualizationsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Visualization

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Visualization
			// with a matching visualizationsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , visualizationsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Visualizations using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Visualizations").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Visualizations", visualizationsId )
				return utils.RequestResult{false, msg, "unassignVisualizations", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Dashboard from the gorm
		//----------------------------------------------------------------------------
		return GetDashboard(dashboardId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more visualizationsIds as a Visualizations from a Dashboard
//----------------------------------------------------------------------------
func RemoveVisualizationsFromDashboard( dashboardId uint64, visualizationsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Dashboard with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDashboard(dashboardId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Dashboard so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Dashboard)

		// slice the ids on comma with no spaces
		ids := strings.Split( visualizationsIds, ",")

		for _, visualizationsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Visualization

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Visualization
			// with a matching visualizationsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , visualizationsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove VisualizationObj from the Visualizations array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Visualizations").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Visualizations", visualizationsId )
				return utils.RequestResult{false, msg, "removeVisualizations", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Dashboard from the gorm
		//----------------------------------------------------------------------------
		return GetDashboard(dashboardId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more reportsIds as a Reports to a Dashboard
//----------------------------------------------------------------------------
func AddReportsToDashboard ( dashboardId uint64, reportsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Dashboard with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDashboard(dashboardId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Dashboard so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Dashboard)

		// slice the ids on comma with no spaces
		ids := strings.Split( reportsIds, ",")

		for _, reportsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Report

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Report
			// with a matching reportsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , reportsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Reports using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Reports").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Reports", reportsId )
				return utils.RequestResult{false, msg, "unassignReports", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Dashboard from the gorm
		//----------------------------------------------------------------------------
		return GetDashboard(dashboardId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more reportsIds as a Reports from a Dashboard
//----------------------------------------------------------------------------
func RemoveReportsFromDashboard( dashboardId uint64, reportsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Dashboard with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDashboard(dashboardId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Dashboard so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Dashboard)

		// slice the ids on comma with no spaces
		ids := strings.Split( reportsIds, ",")

		for _, reportsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Report

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Report
			// with a matching reportsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , reportsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove ReportObj from the Reports array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Reports").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Reports", reportsId )
				return utils.RequestResult{false, msg, "removeReports", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Dashboard from the gorm
		//----------------------------------------------------------------------------
		return GetDashboard(dashboardId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more datasetsIds as a Datasets to a Dashboard
//----------------------------------------------------------------------------
func AddDatasetsToDashboard ( dashboardId uint64, datasetsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Dashboard with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDashboard(dashboardId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Dashboard so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Dashboard)

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
		// retrieve the modified Dashboard from the gorm
		//----------------------------------------------------------------------------
		return GetDashboard(dashboardId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more datasetsIds as a Datasets from a Dashboard
//----------------------------------------------------------------------------
func RemoveDatasetsFromDashboard( dashboardId uint64, datasetsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Dashboard with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDashboard(dashboardId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Dashboard so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Dashboard)

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
		// retrieve the modified Dashboard from the gorm
		//----------------------------------------------------------------------------
		return GetDashboard(dashboardId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more alertsIds as a Alerts to a Dashboard
//----------------------------------------------------------------------------
func AddAlertsToDashboard ( dashboardId uint64, alertsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Dashboard with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDashboard(dashboardId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Dashboard so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Dashboard)

		// slice the ids on comma with no spaces
		ids := strings.Split( alertsIds, ",")

		for _, alertsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Alert

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Alert
			// with a matching alertsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , alertsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Alerts using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Alerts").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Alerts", alertsId )
				return utils.RequestResult{false, msg, "unassignAlerts", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Dashboard from the gorm
		//----------------------------------------------------------------------------
		return GetDashboard(dashboardId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more alertsIds as a Alerts from a Dashboard
//----------------------------------------------------------------------------
func RemoveAlertsFromDashboard( dashboardId uint64, alertsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Dashboard with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDashboard(dashboardId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Dashboard so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Dashboard)

		// slice the ids on comma with no spaces
		ids := strings.Split( alertsIds, ",")

		for _, alertsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Alert

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Alert
			// with a matching alertsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , alertsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove AlertObj from the Alerts array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Alerts").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Alerts", alertsId )
				return utils.RequestResult{false, msg, "removeAlerts", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Dashboard from the gorm
		//----------------------------------------------------------------------------
		return GetDashboard(dashboardId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more queriesIds as a Queries to a Dashboard
//----------------------------------------------------------------------------
func AddQueriesToDashboard ( dashboardId uint64, queriesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Dashboard with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDashboard(dashboardId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Dashboard so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Dashboard)

		// slice the ids on comma with no spaces
		ids := strings.Split( queriesIds, ",")

		for _, queriesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.BIQuery

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a BIQuery
			// with a matching queriesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , queriesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Queries using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Queries").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Queries", queriesId )
				return utils.RequestResult{false, msg, "unassignQueries", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Dashboard from the gorm
		//----------------------------------------------------------------------------
		return GetDashboard(dashboardId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more queriesIds as a Queries from a Dashboard
//----------------------------------------------------------------------------
func RemoveQueriesFromDashboard( dashboardId uint64, queriesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Dashboard with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDashboard(dashboardId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Dashboard so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Dashboard)

		// slice the ids on comma with no spaces
		ids := strings.Split( queriesIds, ",")

		for _, queriesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.BIQuery

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a BIQuery
			// with a matching queriesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , queriesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove BIQueryObj from the Queries array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Queries").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Queries", queriesId )
				return utils.RequestResult{false, msg, "removeQueries", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Dashboard from the gorm
		//----------------------------------------------------------------------------
		return GetDashboard(dashboardId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more tagsIds as a Tags to a Dashboard
//----------------------------------------------------------------------------
func AddTagsToDashboard ( dashboardId uint64, tagsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Dashboard with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDashboard(dashboardId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Dashboard so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Dashboard)

		// slice the ids on comma with no spaces
		ids := strings.Split( tagsIds, ",")

		for _, tagsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Tag

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Tag
			// with a matching tagsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , tagsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Tags using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Tags").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Tags", tagsId )
				return utils.RequestResult{false, msg, "unassignTags", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Dashboard from the gorm
		//----------------------------------------------------------------------------
		return GetDashboard(dashboardId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more tagsIds as a Tags from a Dashboard
//----------------------------------------------------------------------------
func RemoveTagsFromDashboard( dashboardId uint64, tagsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Dashboard with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDashboard(dashboardId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Dashboard so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Dashboard)

		// slice the ids on comma with no spaces
		ids := strings.Split( tagsIds, ",")

		for _, tagsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Tag

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Tag
			// with a matching tagsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , tagsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove TagObj from the Tags array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Tags").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Tags", tagsId )
				return utils.RequestResult{false, msg, "removeTags", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Dashboard from the gorm
		//----------------------------------------------------------------------------
		return GetDashboard(dashboardId)

	} else {
		return parentRequestResult
	}
}

