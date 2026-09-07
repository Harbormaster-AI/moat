package dao

import (
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing AnalyticsWorkspaceDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateAnalyticsWorkspace - creates a new db entry
//----------------------------------------------------------------------------
func CreateAnalyticsWorkspace(obj model.AnalyticsWorkspace)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a AnalyticsWorkspace with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a AnalyticsWorkspace", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateAnalyticsWorkspace", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetAnalyticsWorkspace - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetAnalyticsWorkspace(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.AnalyticsWorkspace

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a AnalyticsWorkspace with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a AnalyticsWorkspace using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a AnalyticsWorkspace using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetAnalyticsWorkspace", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllAnalyticsWorkspace - returns all
//----------------------------------------------------------------------------
func GetAllAnalyticsWorkspace()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.AnalyticsWorkspace

	//----------------------------------------------------------------------------
	// Request the ORM to find all AnalyticsWorkspace
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all AnalyticsWorkspace" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all AnalyticsWorkspace", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllAnalyticsWorkspace", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateAnalyticsWorkspace - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateAnalyticsWorkspace(obj model.AnalyticsWorkspace)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a AnalyticsWorkspace using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a AnalyticsWorkspace using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateAnalyticsWorkspace", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteAnalyticsWorkspace - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteAnalyticsWorkspace(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the AnalyticsWorkspace with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetAnalyticsWorkspace(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AnalyticsWorkspace so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.AnalyticsWorkspace)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a AnalyticsWorkspace using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a AnalyticsWorkspace using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteAnalyticsWorkspace", requestResult.Data}

	}

	return requestResult
}



//----------------------------------------------------------------------------
// adds one or more datasetsIds as a Datasets to a AnalyticsWorkspace
//----------------------------------------------------------------------------
func AddDatasetsToAnalyticsWorkspace ( analyticsWorkspaceId uint64, datasetsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AnalyticsWorkspace with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAnalyticsWorkspace(analyticsWorkspaceId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AnalyticsWorkspace so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AnalyticsWorkspace)

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
		// retrieve the modified AnalyticsWorkspace from the gorm
		//----------------------------------------------------------------------------
		return GetAnalyticsWorkspace(analyticsWorkspaceId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more datasetsIds as a Datasets from a AnalyticsWorkspace
//----------------------------------------------------------------------------
func RemoveDatasetsFromAnalyticsWorkspace( analyticsWorkspaceId uint64, datasetsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the AnalyticsWorkspace with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAnalyticsWorkspace(analyticsWorkspaceId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AnalyticsWorkspace so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AnalyticsWorkspace)

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
		// retrieve the modified AnalyticsWorkspace from the gorm
		//----------------------------------------------------------------------------
		return GetAnalyticsWorkspace(analyticsWorkspaceId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more dataSourcesIds as a DataSources to a AnalyticsWorkspace
//----------------------------------------------------------------------------
func AddDataSourcesToAnalyticsWorkspace ( analyticsWorkspaceId uint64, dataSourcesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AnalyticsWorkspace with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAnalyticsWorkspace(analyticsWorkspaceId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AnalyticsWorkspace so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AnalyticsWorkspace)

		// slice the ids on comma with no spaces
		ids := strings.Split( dataSourcesIds, ",")

		for _, dataSourcesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.DataSource

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a DataSource
			// with a matching dataSourcesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , dataSourcesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the DataSources using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("DataSources").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "DataSources", dataSourcesId )
				return utils.RequestResult{false, msg, "unassignDataSources", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified AnalyticsWorkspace from the gorm
		//----------------------------------------------------------------------------
		return GetAnalyticsWorkspace(analyticsWorkspaceId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more dataSourcesIds as a DataSources from a AnalyticsWorkspace
//----------------------------------------------------------------------------
func RemoveDataSourcesFromAnalyticsWorkspace( analyticsWorkspaceId uint64, dataSourcesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the AnalyticsWorkspace with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAnalyticsWorkspace(analyticsWorkspaceId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AnalyticsWorkspace so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AnalyticsWorkspace)

		// slice the ids on comma with no spaces
		ids := strings.Split( dataSourcesIds, ",")

		for _, dataSourcesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.DataSource

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a DataSource
			// with a matching dataSourcesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , dataSourcesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove DataSourceObj from the DataSources array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("DataSources").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "DataSources", dataSourcesId )
				return utils.RequestResult{false, msg, "removeDataSources", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified AnalyticsWorkspace from the gorm
		//----------------------------------------------------------------------------
		return GetAnalyticsWorkspace(analyticsWorkspaceId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more pipelinesIds as a Pipelines to a AnalyticsWorkspace
//----------------------------------------------------------------------------
func AddPipelinesToAnalyticsWorkspace ( analyticsWorkspaceId uint64, pipelinesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AnalyticsWorkspace with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAnalyticsWorkspace(analyticsWorkspaceId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AnalyticsWorkspace so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AnalyticsWorkspace)

		// slice the ids on comma with no spaces
		ids := strings.Split( pipelinesIds, ",")

		for _, pipelinesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.DataPipeline

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a DataPipeline
			// with a matching pipelinesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , pipelinesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Pipelines using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Pipelines").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Pipelines", pipelinesId )
				return utils.RequestResult{false, msg, "unassignPipelines", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified AnalyticsWorkspace from the gorm
		//----------------------------------------------------------------------------
		return GetAnalyticsWorkspace(analyticsWorkspaceId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more pipelinesIds as a Pipelines from a AnalyticsWorkspace
//----------------------------------------------------------------------------
func RemovePipelinesFromAnalyticsWorkspace( analyticsWorkspaceId uint64, pipelinesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the AnalyticsWorkspace with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAnalyticsWorkspace(analyticsWorkspaceId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AnalyticsWorkspace so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AnalyticsWorkspace)

		// slice the ids on comma with no spaces
		ids := strings.Split( pipelinesIds, ",")

		for _, pipelinesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.DataPipeline

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a DataPipeline
			// with a matching pipelinesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , pipelinesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove DataPipelineObj from the Pipelines array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Pipelines").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Pipelines", pipelinesId )
				return utils.RequestResult{false, msg, "removePipelines", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified AnalyticsWorkspace from the gorm
		//----------------------------------------------------------------------------
		return GetAnalyticsWorkspace(analyticsWorkspaceId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more dashboardsIds as a Dashboards to a AnalyticsWorkspace
//----------------------------------------------------------------------------
func AddDashboardsToAnalyticsWorkspace ( analyticsWorkspaceId uint64, dashboardsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AnalyticsWorkspace with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAnalyticsWorkspace(analyticsWorkspaceId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AnalyticsWorkspace so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AnalyticsWorkspace)

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
		// retrieve the modified AnalyticsWorkspace from the gorm
		//----------------------------------------------------------------------------
		return GetAnalyticsWorkspace(analyticsWorkspaceId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more dashboardsIds as a Dashboards from a AnalyticsWorkspace
//----------------------------------------------------------------------------
func RemoveDashboardsFromAnalyticsWorkspace( analyticsWorkspaceId uint64, dashboardsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the AnalyticsWorkspace with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAnalyticsWorkspace(analyticsWorkspaceId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AnalyticsWorkspace so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AnalyticsWorkspace)

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
		// retrieve the modified AnalyticsWorkspace from the gorm
		//----------------------------------------------------------------------------
		return GetAnalyticsWorkspace(analyticsWorkspaceId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more reportsIds as a Reports to a AnalyticsWorkspace
//----------------------------------------------------------------------------
func AddReportsToAnalyticsWorkspace ( analyticsWorkspaceId uint64, reportsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AnalyticsWorkspace with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAnalyticsWorkspace(analyticsWorkspaceId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AnalyticsWorkspace so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AnalyticsWorkspace)

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
		// retrieve the modified AnalyticsWorkspace from the gorm
		//----------------------------------------------------------------------------
		return GetAnalyticsWorkspace(analyticsWorkspaceId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more reportsIds as a Reports from a AnalyticsWorkspace
//----------------------------------------------------------------------------
func RemoveReportsFromAnalyticsWorkspace( analyticsWorkspaceId uint64, reportsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the AnalyticsWorkspace with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAnalyticsWorkspace(analyticsWorkspaceId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AnalyticsWorkspace so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AnalyticsWorkspace)

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
		// retrieve the modified AnalyticsWorkspace from the gorm
		//----------------------------------------------------------------------------
		return GetAnalyticsWorkspace(analyticsWorkspaceId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more notebooksIds as a Notebooks to a AnalyticsWorkspace
//----------------------------------------------------------------------------
func AddNotebooksToAnalyticsWorkspace ( analyticsWorkspaceId uint64, notebooksIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AnalyticsWorkspace with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAnalyticsWorkspace(analyticsWorkspaceId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AnalyticsWorkspace so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AnalyticsWorkspace)

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
		// retrieve the modified AnalyticsWorkspace from the gorm
		//----------------------------------------------------------------------------
		return GetAnalyticsWorkspace(analyticsWorkspaceId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more notebooksIds as a Notebooks from a AnalyticsWorkspace
//----------------------------------------------------------------------------
func RemoveNotebooksFromAnalyticsWorkspace( analyticsWorkspaceId uint64, notebooksIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the AnalyticsWorkspace with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAnalyticsWorkspace(analyticsWorkspaceId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AnalyticsWorkspace so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AnalyticsWorkspace)

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
		// retrieve the modified AnalyticsWorkspace from the gorm
		//----------------------------------------------------------------------------
		return GetAnalyticsWorkspace(analyticsWorkspaceId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more modelsIds as a Models to a AnalyticsWorkspace
//----------------------------------------------------------------------------
func AddModelsToAnalyticsWorkspace ( analyticsWorkspaceId uint64, modelsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AnalyticsWorkspace with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAnalyticsWorkspace(analyticsWorkspaceId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AnalyticsWorkspace so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AnalyticsWorkspace)

		// slice the ids on comma with no spaces
		ids := strings.Split( modelsIds, ",")

		for _, modelsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Model_

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Model_
			// with a matching modelsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , modelsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Models using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Models").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Models", modelsId )
				return utils.RequestResult{false, msg, "unassignModels", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified AnalyticsWorkspace from the gorm
		//----------------------------------------------------------------------------
		return GetAnalyticsWorkspace(analyticsWorkspaceId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more modelsIds as a Models from a AnalyticsWorkspace
//----------------------------------------------------------------------------
func RemoveModelsFromAnalyticsWorkspace( analyticsWorkspaceId uint64, modelsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the AnalyticsWorkspace with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAnalyticsWorkspace(analyticsWorkspaceId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AnalyticsWorkspace so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AnalyticsWorkspace)

		// slice the ids on comma with no spaces
		ids := strings.Split( modelsIds, ",")

		for _, modelsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.Model_

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a Model_
			// with a matching modelsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , modelsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove Model_Obj from the Models array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Models").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Models", modelsId )
				return utils.RequestResult{false, msg, "removeModels", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified AnalyticsWorkspace from the gorm
		//----------------------------------------------------------------------------
		return GetAnalyticsWorkspace(analyticsWorkspaceId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more featureSetsIds as a FeatureSets to a AnalyticsWorkspace
//----------------------------------------------------------------------------
func AddFeatureSetsToAnalyticsWorkspace ( analyticsWorkspaceId uint64, featureSetsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AnalyticsWorkspace with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAnalyticsWorkspace(analyticsWorkspaceId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AnalyticsWorkspace so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AnalyticsWorkspace)

		// slice the ids on comma with no spaces
		ids := strings.Split( featureSetsIds, ",")

		for _, featureSetsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.FeatureSet

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a FeatureSet
			// with a matching featureSetsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , featureSetsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the FeatureSets using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("FeatureSets").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "FeatureSets", featureSetsId )
				return utils.RequestResult{false, msg, "unassignFeatureSets", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified AnalyticsWorkspace from the gorm
		//----------------------------------------------------------------------------
		return GetAnalyticsWorkspace(analyticsWorkspaceId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more featureSetsIds as a FeatureSets from a AnalyticsWorkspace
//----------------------------------------------------------------------------
func RemoveFeatureSetsFromAnalyticsWorkspace( analyticsWorkspaceId uint64, featureSetsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the AnalyticsWorkspace with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAnalyticsWorkspace(analyticsWorkspaceId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AnalyticsWorkspace so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AnalyticsWorkspace)

		// slice the ids on comma with no spaces
		ids := strings.Split( featureSetsIds, ",")

		for _, featureSetsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.FeatureSet

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a FeatureSet
			// with a matching featureSetsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , featureSetsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove FeatureSetObj from the FeatureSets array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("FeatureSets").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "FeatureSets", featureSetsId )
				return utils.RequestResult{false, msg, "removeFeatureSets", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified AnalyticsWorkspace from the gorm
		//----------------------------------------------------------------------------
		return GetAnalyticsWorkspace(analyticsWorkspaceId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more policiesIds as a Policies to a AnalyticsWorkspace
//----------------------------------------------------------------------------
func AddPoliciesToAnalyticsWorkspace ( analyticsWorkspaceId uint64, policiesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AnalyticsWorkspace with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAnalyticsWorkspace(analyticsWorkspaceId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AnalyticsWorkspace so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AnalyticsWorkspace)

		// slice the ids on comma with no spaces
		ids := strings.Split( policiesIds, ",")

		for _, policiesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.AccessPolicy

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a AccessPolicy
			// with a matching policiesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , policiesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Policies using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Policies").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Policies", policiesId )
				return utils.RequestResult{false, msg, "unassignPolicies", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified AnalyticsWorkspace from the gorm
		//----------------------------------------------------------------------------
		return GetAnalyticsWorkspace(analyticsWorkspaceId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more policiesIds as a Policies from a AnalyticsWorkspace
//----------------------------------------------------------------------------
func RemovePoliciesFromAnalyticsWorkspace( analyticsWorkspaceId uint64, policiesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the AnalyticsWorkspace with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAnalyticsWorkspace(analyticsWorkspaceId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AnalyticsWorkspace so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AnalyticsWorkspace)

		// slice the ids on comma with no spaces
		ids := strings.Split( policiesIds, ",")

		for _, policiesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.AccessPolicy

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a AccessPolicy
			// with a matching policiesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , policiesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove AccessPolicyObj from the Policies array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Policies").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Policies", policiesId )
				return utils.RequestResult{false, msg, "removePolicies", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified AnalyticsWorkspace from the gorm
		//----------------------------------------------------------------------------
		return GetAnalyticsWorkspace(analyticsWorkspaceId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more lineageNodesIds as a LineageNodes to a AnalyticsWorkspace
//----------------------------------------------------------------------------
func AddLineageNodesToAnalyticsWorkspace ( analyticsWorkspaceId uint64, lineageNodesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the AnalyticsWorkspace with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAnalyticsWorkspace(analyticsWorkspaceId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AnalyticsWorkspace so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AnalyticsWorkspace)

		// slice the ids on comma with no spaces
		ids := strings.Split( lineageNodesIds, ",")

		for _, lineageNodesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.LineageNode

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a LineageNode
			// with a matching lineageNodesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , lineageNodesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the LineageNodes using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("LineageNodes").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "LineageNodes", lineageNodesId )
				return utils.RequestResult{false, msg, "unassignLineageNodes", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified AnalyticsWorkspace from the gorm
		//----------------------------------------------------------------------------
		return GetAnalyticsWorkspace(analyticsWorkspaceId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more lineageNodesIds as a LineageNodes from a AnalyticsWorkspace
//----------------------------------------------------------------------------
func RemoveLineageNodesFromAnalyticsWorkspace( analyticsWorkspaceId uint64, lineageNodesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the AnalyticsWorkspace with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetAnalyticsWorkspace(analyticsWorkspaceId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.AnalyticsWorkspace so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.AnalyticsWorkspace)

		// slice the ids on comma with no spaces
		ids := strings.Split( lineageNodesIds, ",")

		for _, lineageNodesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.LineageNode

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a LineageNode
			// with a matching lineageNodesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , lineageNodesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove LineageNodeObj from the LineageNodes array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("LineageNodes").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "LineageNodes", lineageNodesId )
				return utils.RequestResult{false, msg, "removeLineageNodes", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified AnalyticsWorkspace from the gorm
		//----------------------------------------------------------------------------
		return GetAnalyticsWorkspace(analyticsWorkspaceId)

	} else {
		return parentRequestResult
	}
}

