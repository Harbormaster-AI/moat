package dao

import (
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing BIQueryDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateBIQuery - creates a new db entry
//----------------------------------------------------------------------------
func CreateBIQuery(obj model.BIQuery)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a BIQuery with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a BIQuery", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateBIQuery", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetBIQuery - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetBIQuery(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.BIQuery

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a BIQuery with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a BIQuery using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a BIQuery using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetBIQuery", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllBIQuery - returns all
//----------------------------------------------------------------------------
func GetAllBIQuery()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.BIQuery

	//----------------------------------------------------------------------------
	// Request the ORM to find all BIQuery
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all BIQuery" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all BIQuery", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllBIQuery", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateBIQuery - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateBIQuery(obj model.BIQuery)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a BIQuery using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a BIQuery using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateBIQuery", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteBIQuery - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteBIQuery(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the BIQuery with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetBIQuery(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BIQuery so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.BIQuery)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a BIQuery using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a BIQuery using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteBIQuery", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Workspace on a BIQuery
//----------------------------------------------------------------------------
func AssignWorkspaceToBIQuery( bIQueryId uint64, workspaceId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the BIQuery with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBIQuery(bIQueryId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BIQuery so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.BIQuery)

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
			// assign the Workspace	to the BIQuery
			//----------------------------------------------------------------------------
			parentObj.Workspace = &childObj

			//----------------------------------------------------------------------------
			// save the BIQuery
			//----------------------------------------------------------------------------
			return UpdateBIQuery(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Workspace", workspaceId )
			return utils.RequestResult{false, msg, "assignWorkspace", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Workspace on a BIQuery
//----------------------------------------------------------------------------
func UnassignWorkspaceFromBIQuery(bIQueryId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the BIQuery with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBIQuery(bIQueryId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BIQuery so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.BIQuery)

		//----------------------------------------------------------------------------
		// assign an empty AnalyticsWorkspace to the Workspace
		//----------------------------------------------------------------------------
		parentObj.Workspace = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Workspace
		//----------------------------------------------------------------------------
		parentObj.WorkspaceId = nil;

		//----------------------------------------------------------------------------
		// save the BIQuery
		//----------------------------------------------------------------------------
		return UpdateBIQuery(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more datasetsIds as a Datasets to a BIQuery
//----------------------------------------------------------------------------
func AddDatasetsToBIQuery ( bIQueryId uint64, datasetsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the BIQuery with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBIQuery(bIQueryId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BIQuery so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.BIQuery)

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
		// retrieve the modified BIQuery from the gorm
		//----------------------------------------------------------------------------
		return GetBIQuery(bIQueryId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more datasetsIds as a Datasets from a BIQuery
//----------------------------------------------------------------------------
func RemoveDatasetsFromBIQuery( bIQueryId uint64, datasetsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the BIQuery with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBIQuery(bIQueryId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BIQuery so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.BIQuery)

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
		// retrieve the modified BIQuery from the gorm
		//----------------------------------------------------------------------------
		return GetBIQuery(bIQueryId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more reportsIds as a Reports to a BIQuery
//----------------------------------------------------------------------------
func AddReportsToBIQuery ( bIQueryId uint64, reportsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the BIQuery with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBIQuery(bIQueryId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BIQuery so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.BIQuery)

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
		// retrieve the modified BIQuery from the gorm
		//----------------------------------------------------------------------------
		return GetBIQuery(bIQueryId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more reportsIds as a Reports from a BIQuery
//----------------------------------------------------------------------------
func RemoveReportsFromBIQuery( bIQueryId uint64, reportsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the BIQuery with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBIQuery(bIQueryId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BIQuery so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.BIQuery)

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
		// retrieve the modified BIQuery from the gorm
		//----------------------------------------------------------------------------
		return GetBIQuery(bIQueryId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more dashboardsIds as a Dashboards to a BIQuery
//----------------------------------------------------------------------------
func AddDashboardsToBIQuery ( bIQueryId uint64, dashboardsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the BIQuery with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBIQuery(bIQueryId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BIQuery so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.BIQuery)

		// slice the ids on comma with no spaces
		ids := strings.Split( dashboardsIds, ",")

		for _, dashboardsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Dashboard

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Dashboard
			// with a matching dashboardsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , dashboardsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Dashboards using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Dashboards").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Dashboards", dashboardsId )
				return utils.RequestResult{false, msg, "unassignDashboards", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified BIQuery from the gorm
		//----------------------------------------------------------------------------
		return GetBIQuery(bIQueryId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more dashboardsIds as a Dashboards from a BIQuery
//----------------------------------------------------------------------------
func RemoveDashboardsFromBIQuery( bIQueryId uint64, dashboardsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the BIQuery with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBIQuery(bIQueryId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BIQuery so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.BIQuery)

		// slice the ids on comma with no spaces
		ids := strings.Split( dashboardsIds, ",")

		for _, dashboardsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Dashboard

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Dashboard
			// with a matching dashboardsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , dashboardsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove DashboardObj from the Dashboards array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Dashboards").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Dashboards", dashboardsId )
				return utils.RequestResult{false, msg, "removeDashboards", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified BIQuery from the gorm
		//----------------------------------------------------------------------------
		return GetBIQuery(bIQueryId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more notebooksIds as a Notebooks to a BIQuery
//----------------------------------------------------------------------------
func AddNotebooksToBIQuery ( bIQueryId uint64, notebooksIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the BIQuery with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBIQuery(bIQueryId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BIQuery so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.BIQuery)

		// slice the ids on comma with no spaces
		ids := strings.Split( notebooksIds, ",")

		for _, notebooksId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Notebook

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Notebook
			// with a matching notebooksId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , notebooksId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Notebooks using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Notebooks").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Notebooks", notebooksId )
				return utils.RequestResult{false, msg, "unassignNotebooks", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified BIQuery from the gorm
		//----------------------------------------------------------------------------
		return GetBIQuery(bIQueryId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more notebooksIds as a Notebooks from a BIQuery
//----------------------------------------------------------------------------
func RemoveNotebooksFromBIQuery( bIQueryId uint64, notebooksIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the BIQuery with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetBIQuery(bIQueryId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.BIQuery so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.BIQuery)

		// slice the ids on comma with no spaces
		ids := strings.Split( notebooksIds, ",")

		for _, notebooksId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Notebook

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Notebook
			// with a matching notebooksId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , notebooksId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove NotebookObj from the Notebooks array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Notebooks").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Notebooks", notebooksId )
				return utils.RequestResult{false, msg, "removeNotebooks", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified BIQuery from the gorm
		//----------------------------------------------------------------------------
		return GetBIQuery(bIQueryId)

	} else {
		return parentRequestResult
	}
}

