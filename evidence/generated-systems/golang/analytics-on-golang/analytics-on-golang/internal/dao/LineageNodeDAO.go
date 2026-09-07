package dao

import (
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing LineageNodeDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateLineageNode - creates a new db entry
//----------------------------------------------------------------------------
func CreateLineageNode(obj model.LineageNode)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a LineageNode with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a LineageNode", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateLineageNode", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetLineageNode - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetLineageNode(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.LineageNode

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a LineageNode with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a LineageNode using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a LineageNode using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetLineageNode", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllLineageNode - returns all
//----------------------------------------------------------------------------
func GetAllLineageNode()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.LineageNode

	//----------------------------------------------------------------------------
	// Request the ORM to find all LineageNode
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all LineageNode" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all LineageNode", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllLineageNode", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateLineageNode - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateLineageNode(obj model.LineageNode)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a LineageNode using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a LineageNode using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateLineageNode", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteLineageNode - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteLineageNode(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the LineageNode with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetLineageNode(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LineageNode so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.LineageNode)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a LineageNode using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a LineageNode using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteLineageNode", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Workspace on a LineageNode
//----------------------------------------------------------------------------
func AssignWorkspaceToLineageNode( lineageNodeId uint64, workspaceId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the LineageNode with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLineageNode(lineageNodeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LineageNode so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LineageNode)

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
			// assign the Workspace	to the LineageNode
			//----------------------------------------------------------------------------
			parentObj.Workspace = &childObj

			//----------------------------------------------------------------------------
			// save the LineageNode
			//----------------------------------------------------------------------------
			return UpdateLineageNode(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Workspace", workspaceId )
			return utils.RequestResult{false, msg, "assignWorkspace", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Workspace on a LineageNode
//----------------------------------------------------------------------------
func UnassignWorkspaceFromLineageNode(lineageNodeId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the LineageNode with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLineageNode(lineageNodeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LineageNode so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LineageNode)

		//----------------------------------------------------------------------------
		// assign an empty AnalyticsWorkspace to the Workspace
		//----------------------------------------------------------------------------
		parentObj.Workspace = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Workspace
		//----------------------------------------------------------------------------
		parentObj.WorkspaceId = nil;

		//----------------------------------------------------------------------------
		// save the LineageNode
		//----------------------------------------------------------------------------
		return UpdateLineageNode(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more inputsIds as a Inputs to a LineageNode
//----------------------------------------------------------------------------
func AddInputsToLineageNode ( lineageNodeId uint64, inputsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the LineageNode with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLineageNode(lineageNodeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LineageNode so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LineageNode)

		// slice the ids on comma with no spaces
		ids := strings.Split( inputsIds, ",")

		for _, inputsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.LineageNode

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a LineageNode
			// with a matching inputsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , inputsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Inputs using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Inputs").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Inputs", inputsId )
				return utils.RequestResult{false, msg, "unassignInputs", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified LineageNode from the gorm
		//----------------------------------------------------------------------------
		return GetLineageNode(lineageNodeId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more inputsIds as a Inputs from a LineageNode
//----------------------------------------------------------------------------
func RemoveInputsFromLineageNode( lineageNodeId uint64, inputsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the LineageNode with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLineageNode(lineageNodeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LineageNode so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LineageNode)

		// slice the ids on comma with no spaces
		ids := strings.Split( inputsIds, ",")

		for _, inputsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.LineageNode

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a LineageNode
			// with a matching inputsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , inputsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove LineageNodeObj from the Inputs array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Inputs").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Inputs", inputsId )
				return utils.RequestResult{false, msg, "removeInputs", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified LineageNode from the gorm
		//----------------------------------------------------------------------------
		return GetLineageNode(lineageNodeId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more outputsIds as a Outputs to a LineageNode
//----------------------------------------------------------------------------
func AddOutputsToLineageNode ( lineageNodeId uint64, outputsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the LineageNode with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLineageNode(lineageNodeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LineageNode so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LineageNode)

		// slice the ids on comma with no spaces
		ids := strings.Split( outputsIds, ",")

		for _, outputsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.LineageNode

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a LineageNode
			// with a matching outputsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , outputsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Outputs using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Outputs").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Outputs", outputsId )
				return utils.RequestResult{false, msg, "unassignOutputs", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified LineageNode from the gorm
		//----------------------------------------------------------------------------
		return GetLineageNode(lineageNodeId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more outputsIds as a Outputs from a LineageNode
//----------------------------------------------------------------------------
func RemoveOutputsFromLineageNode( lineageNodeId uint64, outputsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the LineageNode with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLineageNode(lineageNodeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LineageNode so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LineageNode)

		// slice the ids on comma with no spaces
		ids := strings.Split( outputsIds, ",")

		for _, outputsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.LineageNode

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a LineageNode
			// with a matching outputsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , outputsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove LineageNodeObj from the Outputs array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Outputs").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Outputs", outputsId )
				return utils.RequestResult{false, msg, "removeOutputs", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified LineageNode from the gorm
		//----------------------------------------------------------------------------
		return GetLineageNode(lineageNodeId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more datasetsIds as a Datasets to a LineageNode
//----------------------------------------------------------------------------
func AddDatasetsToLineageNode ( lineageNodeId uint64, datasetsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the LineageNode with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLineageNode(lineageNodeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LineageNode so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LineageNode)

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
		// retrieve the modified LineageNode from the gorm
		//----------------------------------------------------------------------------
		return GetLineageNode(lineageNodeId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more datasetsIds as a Datasets from a LineageNode
//----------------------------------------------------------------------------
func RemoveDatasetsFromLineageNode( lineageNodeId uint64, datasetsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the LineageNode with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLineageNode(lineageNodeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LineageNode so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LineageNode)

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
		// retrieve the modified LineageNode from the gorm
		//----------------------------------------------------------------------------
		return GetLineageNode(lineageNodeId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more modelsIds as a Models to a LineageNode
//----------------------------------------------------------------------------
func AddModelsToLineageNode ( lineageNodeId uint64, modelsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the LineageNode with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLineageNode(lineageNodeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LineageNode so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LineageNode)

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
		// retrieve the modified LineageNode from the gorm
		//----------------------------------------------------------------------------
		return GetLineageNode(lineageNodeId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more modelsIds as a Models from a LineageNode
//----------------------------------------------------------------------------
func RemoveModelsFromLineageNode( lineageNodeId uint64, modelsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the LineageNode with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLineageNode(lineageNodeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LineageNode so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LineageNode)

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
		// retrieve the modified LineageNode from the gorm
		//----------------------------------------------------------------------------
		return GetLineageNode(lineageNodeId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more pipelinesIds as a Pipelines to a LineageNode
//----------------------------------------------------------------------------
func AddPipelinesToLineageNode ( lineageNodeId uint64, pipelinesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the LineageNode with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLineageNode(lineageNodeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LineageNode so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LineageNode)

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
		// retrieve the modified LineageNode from the gorm
		//----------------------------------------------------------------------------
		return GetLineageNode(lineageNodeId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more pipelinesIds as a Pipelines from a LineageNode
//----------------------------------------------------------------------------
func RemovePipelinesFromLineageNode( lineageNodeId uint64, pipelinesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the LineageNode with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLineageNode(lineageNodeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LineageNode so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LineageNode)

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
		// retrieve the modified LineageNode from the gorm
		//----------------------------------------------------------------------------
		return GetLineageNode(lineageNodeId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more dashboardsIds as a Dashboards to a LineageNode
//----------------------------------------------------------------------------
func AddDashboardsToLineageNode ( lineageNodeId uint64, dashboardsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the LineageNode with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLineageNode(lineageNodeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LineageNode so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LineageNode)

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
		// retrieve the modified LineageNode from the gorm
		//----------------------------------------------------------------------------
		return GetLineageNode(lineageNodeId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more dashboardsIds as a Dashboards from a LineageNode
//----------------------------------------------------------------------------
func RemoveDashboardsFromLineageNode( lineageNodeId uint64, dashboardsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the LineageNode with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLineageNode(lineageNodeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LineageNode so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LineageNode)

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
		// retrieve the modified LineageNode from the gorm
		//----------------------------------------------------------------------------
		return GetLineageNode(lineageNodeId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more reportsIds as a Reports to a LineageNode
//----------------------------------------------------------------------------
func AddReportsToLineageNode ( lineageNodeId uint64, reportsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the LineageNode with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLineageNode(lineageNodeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LineageNode so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LineageNode)

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
		// retrieve the modified LineageNode from the gorm
		//----------------------------------------------------------------------------
		return GetLineageNode(lineageNodeId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more reportsIds as a Reports from a LineageNode
//----------------------------------------------------------------------------
func RemoveReportsFromLineageNode( lineageNodeId uint64, reportsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the LineageNode with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetLineageNode(lineageNodeId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.LineageNode so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.LineageNode)

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
		// retrieve the modified LineageNode from the gorm
		//----------------------------------------------------------------------------
		return GetLineageNode(lineageNodeId)

	} else {
		return parentRequestResult
	}
}

