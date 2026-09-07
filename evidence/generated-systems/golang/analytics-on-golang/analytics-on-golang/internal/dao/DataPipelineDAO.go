package dao

import (
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing DataPipelineDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateDataPipeline - creates a new db entry
//----------------------------------------------------------------------------
func CreateDataPipeline(obj model.DataPipeline)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a DataPipeline with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a DataPipeline", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateDataPipeline", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetDataPipeline - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetDataPipeline(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.DataPipeline

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a DataPipeline with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a DataPipeline using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a DataPipeline using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetDataPipeline", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllDataPipeline - returns all
//----------------------------------------------------------------------------
func GetAllDataPipeline()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.DataPipeline

	//----------------------------------------------------------------------------
	// Request the ORM to find all DataPipeline
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all DataPipeline" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all DataPipeline", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllDataPipeline", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateDataPipeline - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateDataPipeline(obj model.DataPipeline)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a DataPipeline using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a DataPipeline using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateDataPipeline", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteDataPipeline - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteDataPipeline(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the DataPipeline with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetDataPipeline(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataPipeline so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.DataPipeline)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a DataPipeline using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a DataPipeline using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteDataPipeline", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Workspace on a DataPipeline
//----------------------------------------------------------------------------
func AssignWorkspaceToDataPipeline( dataPipelineId uint64, workspaceId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the DataPipeline with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataPipeline(dataPipelineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataPipeline so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataPipeline)

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
			// assign the Workspace	to the DataPipeline
			//----------------------------------------------------------------------------
			parentObj.Workspace = &childObj

			//----------------------------------------------------------------------------
			// save the DataPipeline
			//----------------------------------------------------------------------------
			return UpdateDataPipeline(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Workspace", workspaceId )
			return utils.RequestResult{false, msg, "assignWorkspace", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Workspace on a DataPipeline
//----------------------------------------------------------------------------
func UnassignWorkspaceFromDataPipeline(dataPipelineId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the DataPipeline with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataPipeline(dataPipelineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataPipeline so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataPipeline)

		//----------------------------------------------------------------------------
		// assign an empty AnalyticsWorkspace to the Workspace
		//----------------------------------------------------------------------------
		parentObj.Workspace = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Workspace
		//----------------------------------------------------------------------------
		parentObj.WorkspaceId = nil;

		//----------------------------------------------------------------------------
		// save the DataPipeline
		//----------------------------------------------------------------------------
		return UpdateDataPipeline(parentObj)

	} else {
		return parentRequestResult
	}

}

//----------------------------------------------------------------------------
// assigns a LineageNode on a DataPipeline
//----------------------------------------------------------------------------
func AssignLineageNodeToDataPipeline( dataPipelineId uint64, lineageNodeId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the DataPipeline with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataPipeline(dataPipelineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataPipeline so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataPipeline)

		//----------------------------------------------------------------------------
		// Pass the reference to the ORM to get
		//----------------------------------------------------------------------------
		var childObj model.LineageNode

		//----------------------------------------------------------------------------
		// Retrieve the 1st occurrence from the ORM of a LineageNode with a
		// matching lineageNodeId
		//----------------------------------------------------------------------------
		childRequestResult := utils.GetDB().First(&childObj, lineageNodeId).Error // find first using identifier

		if childRequestResult == nil {
			//----------------------------------------------------------------------------
			// assign the LineageNode	to the DataPipeline
			//----------------------------------------------------------------------------
			parentObj.LineageNode = &childObj

			//----------------------------------------------------------------------------
			// save the DataPipeline
			//----------------------------------------------------------------------------
			return UpdateDataPipeline(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "LineageNode", lineageNodeId )
			return utils.RequestResult{false, msg, "assignLineageNode", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a LineageNode on a DataPipeline
//----------------------------------------------------------------------------
func UnassignLineageNodeFromDataPipeline(dataPipelineId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the DataPipeline with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataPipeline(dataPipelineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataPipeline so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataPipeline)

		//----------------------------------------------------------------------------
		// assign an empty LineageNode to the LineageNode
		//----------------------------------------------------------------------------
		parentObj.LineageNode = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the LineageNode
		//----------------------------------------------------------------------------
		parentObj.LineageNodeId = nil;

		//----------------------------------------------------------------------------
		// save the DataPipeline
		//----------------------------------------------------------------------------
		return UpdateDataPipeline(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more tasksIds as a Tasks to a DataPipeline
//----------------------------------------------------------------------------
func AddTasksToDataPipeline ( dataPipelineId uint64, tasksIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the DataPipeline with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataPipeline(dataPipelineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataPipeline so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataPipeline)

		// slice the ids on comma with no spaces
		ids := strings.Split( tasksIds, ",")

		for _, tasksId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.DataTask

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a DataTask
			// with a matching tasksId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , tasksId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Tasks using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Tasks").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Tasks", tasksId )
				return utils.RequestResult{false, msg, "unassignTasks", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified DataPipeline from the gorm
		//----------------------------------------------------------------------------
		return GetDataPipeline(dataPipelineId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more tasksIds as a Tasks from a DataPipeline
//----------------------------------------------------------------------------
func RemoveTasksFromDataPipeline( dataPipelineId uint64, tasksIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the DataPipeline with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataPipeline(dataPipelineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataPipeline so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataPipeline)

		// slice the ids on comma with no spaces
		ids := strings.Split( tasksIds, ",")

		for _, tasksId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.DataTask

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a DataTask
			// with a matching tasksId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , tasksId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove DataTaskObj from the Tasks array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Tasks").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Tasks", tasksId )
				return utils.RequestResult{false, msg, "removeTasks", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified DataPipeline from the gorm
		//----------------------------------------------------------------------------
		return GetDataPipeline(dataPipelineId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more sourcesIds as a Sources to a DataPipeline
//----------------------------------------------------------------------------
func AddSourcesToDataPipeline ( dataPipelineId uint64, sourcesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the DataPipeline with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataPipeline(dataPipelineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataPipeline so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataPipeline)

		// slice the ids on comma with no spaces
		ids := strings.Split( sourcesIds, ",")

		for _, sourcesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.DataSource

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a DataSource
			// with a matching sourcesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , sourcesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the Sources using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Sources").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Sources", sourcesId )
				return utils.RequestResult{false, msg, "unassignSources", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified DataPipeline from the gorm
		//----------------------------------------------------------------------------
		return GetDataPipeline(dataPipelineId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more sourcesIds as a Sources from a DataPipeline
//----------------------------------------------------------------------------
func RemoveSourcesFromDataPipeline( dataPipelineId uint64, sourcesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the DataPipeline with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataPipeline(dataPipelineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataPipeline so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataPipeline)

		// slice the ids on comma with no spaces
		ids := strings.Split( sourcesIds, ",")

		for _, sourcesId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.DataSource

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a DataSource
			// with a matching sourcesId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , sourcesId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove DataSourceObj from the Sources array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Sources").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Sources", sourcesId )
				return utils.RequestResult{false, msg, "removeSources", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified DataPipeline from the gorm
		//----------------------------------------------------------------------------
		return GetDataPipeline(dataPipelineId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more outputsIds as a Outputs to a DataPipeline
//----------------------------------------------------------------------------
func AddOutputsToDataPipeline ( dataPipelineId uint64, outputsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the DataPipeline with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataPipeline(dataPipelineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataPipeline so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataPipeline)

		// slice the ids on comma with no spaces
		ids := strings.Split( outputsIds, ",")

		for _, outputsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.DataSet

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a DataSet
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
		// retrieve the modified DataPipeline from the gorm
		//----------------------------------------------------------------------------
		return GetDataPipeline(dataPipelineId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more outputsIds as a Outputs from a DataPipeline
//----------------------------------------------------------------------------
func RemoveOutputsFromDataPipeline( dataPipelineId uint64, outputsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the DataPipeline with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataPipeline(dataPipelineId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataPipeline so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataPipeline)

		// slice the ids on comma with no spaces
		ids := strings.Split( outputsIds, ",")

		for _, outputsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.DataSet

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a DataSet
			// with a matching outputsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , outputsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove DataSetObj from the Outputs array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("Outputs").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Outputs", outputsId )
				return utils.RequestResult{false, msg, "removeOutputs", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified DataPipeline from the gorm
		//----------------------------------------------------------------------------
		return GetDataPipeline(dataPipelineId)

	} else {
		return parentRequestResult
	}
}

