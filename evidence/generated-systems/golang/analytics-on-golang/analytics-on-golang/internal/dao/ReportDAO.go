package dao

import (
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing ReportDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateReport - creates a new db entry
//----------------------------------------------------------------------------
func CreateReport(obj model.Report)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a Report with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a Report", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateReport", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetReport - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetReport(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.Report

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a Report with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a Report using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a Report using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetReport", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllReport - returns all
//----------------------------------------------------------------------------
func GetAllReport()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.Report

	//----------------------------------------------------------------------------
	// Request the ORM to find all Report
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all Report" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all Report", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllReport", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateReport - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateReport(obj model.Report)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a Report using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a Report using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateReport", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteReport - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteReport(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the Report with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetReport(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Report so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.Report)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a Report using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a Report using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteReport", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Workspace on a Report
//----------------------------------------------------------------------------
func AssignWorkspaceToReport( reportId uint64, workspaceId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the Report with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetReport(reportId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Report so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Report)

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
			// assign the Workspace	to the Report
			//----------------------------------------------------------------------------
			parentObj.Workspace = &childObj

			//----------------------------------------------------------------------------
			// save the Report
			//----------------------------------------------------------------------------
			return UpdateReport(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Workspace", workspaceId )
			return utils.RequestResult{false, msg, "assignWorkspace", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Workspace on a Report
//----------------------------------------------------------------------------
func UnassignWorkspaceFromReport(reportId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Report with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetReport(reportId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Report so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Report)

		//----------------------------------------------------------------------------
		// assign an empty AnalyticsWorkspace to the Workspace
		//----------------------------------------------------------------------------
		parentObj.Workspace = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Workspace
		//----------------------------------------------------------------------------
		parentObj.WorkspaceId = nil;

		//----------------------------------------------------------------------------
		// save the Report
		//----------------------------------------------------------------------------
		return UpdateReport(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more visualizationsIds as a Visualizations to a Report
//----------------------------------------------------------------------------
func AddVisualizationsToReport ( reportId uint64, visualizationsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Report with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetReport(reportId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Report so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Report)

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
		// retrieve the modified Report from the gorm
		//----------------------------------------------------------------------------
		return GetReport(reportId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more visualizationsIds as a Visualizations from a Report
//----------------------------------------------------------------------------
func RemoveVisualizationsFromReport( reportId uint64, visualizationsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Report with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetReport(reportId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Report so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Report)

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
		// retrieve the modified Report from the gorm
		//----------------------------------------------------------------------------
		return GetReport(reportId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more datasetsIds as a Datasets to a Report
//----------------------------------------------------------------------------
func AddDatasetsToReport ( reportId uint64, datasetsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Report with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetReport(reportId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Report so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Report)

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
		// retrieve the modified Report from the gorm
		//----------------------------------------------------------------------------
		return GetReport(reportId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more datasetsIds as a Datasets from a Report
//----------------------------------------------------------------------------
func RemoveDatasetsFromReport( reportId uint64, datasetsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Report with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetReport(reportId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Report so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Report)

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
		// retrieve the modified Report from the gorm
		//----------------------------------------------------------------------------
		return GetReport(reportId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more semanticModelsIds as a SemanticModels to a Report
//----------------------------------------------------------------------------
func AddSemanticModelsToReport ( reportId uint64, semanticModelsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Report with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetReport(reportId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Report so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Report)

		// slice the ids on comma with no spaces
		ids := strings.Split( semanticModelsIds, ",")

		for _, semanticModelsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.SemanticModel

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a SemanticModel
			// with a matching semanticModelsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , semanticModelsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the SemanticModels using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("SemanticModels").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "SemanticModels", semanticModelsId )
				return utils.RequestResult{false, msg, "unassignSemanticModels", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Report from the gorm
		//----------------------------------------------------------------------------
		return GetReport(reportId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more semanticModelsIds as a SemanticModels from a Report
//----------------------------------------------------------------------------
func RemoveSemanticModelsFromReport( reportId uint64, semanticModelsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Report with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetReport(reportId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Report so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Report)

		// slice the ids on comma with no spaces
		ids := strings.Split( semanticModelsIds, ",")

		for _, semanticModelsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.SemanticModel

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a SemanticModel
			// with a matching semanticModelsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , semanticModelsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove SemanticModelObj from the SemanticModels array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("SemanticModels").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "SemanticModels", semanticModelsId )
				return utils.RequestResult{false, msg, "removeSemanticModels", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified Report from the gorm
		//----------------------------------------------------------------------------
		return GetReport(reportId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more queriesIds as a Queries to a Report
//----------------------------------------------------------------------------
func AddQueriesToReport ( reportId uint64, queriesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Report with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetReport(reportId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Report so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Report)

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
		// retrieve the modified Report from the gorm
		//----------------------------------------------------------------------------
		return GetReport(reportId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more queriesIds as a Queries from a Report
//----------------------------------------------------------------------------
func RemoveQueriesFromReport( reportId uint64, queriesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Report with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetReport(reportId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Report so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Report)

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
		// retrieve the modified Report from the gorm
		//----------------------------------------------------------------------------
		return GetReport(reportId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more tagsIds as a Tags to a Report
//----------------------------------------------------------------------------
func AddTagsToReport ( reportId uint64, tagsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the Report with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetReport(reportId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Report so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Report)

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
		// retrieve the modified Report from the gorm
		//----------------------------------------------------------------------------
		return GetReport(reportId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more tagsIds as a Tags from a Report
//----------------------------------------------------------------------------
func RemoveTagsFromReport( reportId uint64, tagsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the Report with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetReport(reportId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.Report so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.Report)

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
		// retrieve the modified Report from the gorm
		//----------------------------------------------------------------------------
		return GetReport(reportId)

	} else {
		return parentRequestResult
	}
}

