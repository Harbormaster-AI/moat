package dao

import (
    "analytics-on-golang/internal/model"
    "analytics-on-golang/internal/utils"
    "fmt"
    "strings"
)


func init() {
	fmt.Println( strings.ToTitle( "Initializing DataSourceDAO..." ) )
}

//----------------------------------------------------------------------------
// CreateDataSource - creates a new db entry
//----------------------------------------------------------------------------
func CreateDataSource(obj model.DataSource)(utils.RequestResult){
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
	    createMsg = fmt.Sprintf( "Created a DataSource with ID=%v", obj.ID )
	    success = true
	} else {
		createMsg = fmt.Sprintf( "Failed trying to create a DataSource", result )
		success = false
	}

	requestResult = utils.RequestResult{success, createMsg, "CreateDataSource", obj}
	return requestResult
}


//----------------------------------------------------------------------------
// GetDataSource - returns the matching the provided identifier
//----------------------------------------------------------------------------
func GetDataSource(id uint64)(utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var requestResult utils.RequestResult
	var getMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Pass the reference to the ORM to create
	//----------------------------------------------------------------------------
	var obj model.DataSource

	//----------------------------------------------------------------------------
	// Retrieve the 1st occurrence from the ORM of a DataSource with a matching ID
	//----------------------------------------------------------------------------
	result := utils.GetDB().First(&obj, id).Error // find first using identifier

	if result == nil {
	    getMsg = fmt.Sprintf( "Retrieved a DataSource using ID=%v", id )
	    success = true
	} else {
		getMsg = fmt.Sprintf( "Failed trying to retrieve a DataSource using ID=%v", id )
		success = false
	}

	requestResult = utils.RequestResult{success, getMsg, "GetDataSource", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// GetAllDataSource - returns all
//----------------------------------------------------------------------------
func GetAllDataSource()(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var getAllMsg string
	var success bool
	var objs []model.DataSource

	//----------------------------------------------------------------------------
	// Request the ORM to find all DataSource
	//----------------------------------------------------------------------------
	result := utils.GetDB().Find(&objs).Error // find all

	if result == nil {
	    getAllMsg = fmt.Sprintf( "Retrieved all DataSource" )
	    success = true
	} else {
		getAllMsg = fmt.Sprintf( "Failed trying to retrieve all DataSource", result )
		success = false
	}

	requestResult = utils.RequestResult{success, getAllMsg, "GetAllDataSource", objs}
	return requestResult
}

//----------------------------------------------------------------------------
// UpdateDataSource - updates matching the provided identifier
//----------------------------------------------------------------------------
func UpdateDataSource(obj model.DataSource)(requestResult utils.RequestResult){
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
	    updateMsg = fmt.Sprintf( "Updated a DataSource using ID=%v", obj.ID )
	    success = true
	} else {
		updateMsg = fmt.Sprintf( "Failed trying to update a DataSource using ID=%v", obj.ID )
		success = false
	}

	requestResult = utils.RequestResult{success, updateMsg, "UpdateDataSource", obj}

	return requestResult
}

//----------------------------------------------------------------------------
// DeleteDataSource - deletes matching the provided identifier
//----------------------------------------------------------------------------
func DeleteDataSource(id uint64)(requestResult utils.RequestResult){
	//----------------------------------------------------------------------------
	// variable initialization
	//----------------------------------------------------------------------------
	var deleteMsg string
	var success bool

	//----------------------------------------------------------------------------
	// Obtain the DataSource with the matching identifier
	//----------------------------------------------------------------------------
	requestResult = GetDataSource(id)

	if requestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataSource so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		obj,_ := requestResult.Data. (model.DataSource)

		//----------------------------------------------------------------------------
		// Make call to the ORM to delete
		//----------------------------------------------------------------------------
		result := utils.GetDB().Delete(&obj).Error // pass pointer of data to Delete

		if result == nil {
		    deleteMsg = fmt.Sprintf( "Deleted a DataSource using ID=%v", id )
		    success = true
		} else {
			deleteMsg = fmt.Sprintf( "Failed trying to delete a DataSource using ID=%v", id )
			success = false
		}

		requestResult = utils.RequestResult{success, deleteMsg, "DeleteDataSource", requestResult.Data}

	}

	return requestResult
}


//----------------------------------------------------------------------------
// assigns a Workspace on a DataSource
//----------------------------------------------------------------------------
func AssignWorkspaceToDataSource( dataSourceId uint64, workspaceId uint64 )(utils.RequestResult){

	//----------------------------------------------------------------------------
	// Obtain the DataSource with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataSource(dataSourceId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataSource so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataSource)

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
			// assign the Workspace	to the DataSource
			//----------------------------------------------------------------------------
			parentObj.Workspace = &childObj

			//----------------------------------------------------------------------------
			// save the DataSource
			//----------------------------------------------------------------------------
			return UpdateDataSource(parentObj)
		} else {
			msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "Workspace", workspaceId )
			return utils.RequestResult{false, msg, "assignWorkspace", childObj}
		}
	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// unassigns a Workspace on a DataSource
//----------------------------------------------------------------------------
func UnassignWorkspaceFromDataSource(dataSourceId uint64)(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the DataSource with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataSource(dataSourceId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataSource so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataSource)

		//----------------------------------------------------------------------------
		// assign an empty AnalyticsWorkspace to the Workspace
		//----------------------------------------------------------------------------
		parentObj.Workspace = nil;

		//----------------------------------------------------------------------------
		// assign  nil to the Workspace
		//----------------------------------------------------------------------------
		parentObj.WorkspaceId = nil;

		//----------------------------------------------------------------------------
		// save the DataSource
		//----------------------------------------------------------------------------
		return UpdateDataSource(parentObj)

	} else {
		return parentRequestResult
	}

}


//----------------------------------------------------------------------------
// adds one or more producedDatasetsIds as a ProducedDatasets to a DataSource
//----------------------------------------------------------------------------
func AddProducedDatasetsToDataSource ( dataSourceId uint64, producedDatasetsIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the DataSource with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataSource(dataSourceId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataSource so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataSource)

		// slice the ids on comma with no spaces
		ids := strings.Split( producedDatasetsIds, ",")

		for _, producedDatasetsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.DataSet

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a DataSet
			// with a matching producedDatasetsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , producedDatasetsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// append to the ProducedDatasets using the gorm mechanism
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("ProducedDatasets").Append( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ProducedDatasets", producedDatasetsId )
				return utils.RequestResult{false, msg, "unassignProducedDatasets", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified DataSource from the gorm
		//----------------------------------------------------------------------------
		return GetDataSource(dataSourceId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more producedDatasetsIds as a ProducedDatasets from a DataSource
//----------------------------------------------------------------------------
func RemoveProducedDatasetsFromDataSource( dataSourceId uint64, producedDatasetsIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the DataSource with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataSource(dataSourceId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataSource so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataSource)

		// slice the ids on comma with no spaces
		ids := strings.Split( producedDatasetsIds, ",")

		for _, producedDatasetsId:= range ids {
			//----------------------------------------------------------------------------
			// Pass the reference to the ORM to get
			//----------------------------------------------------------------------------
			var childObj model.DataSet

			//----------------------------------------------------------------------------
			// Retrieve the 1st occurrence from the ORM of a DataSet
			// with a matching producedDatasetsId
			//----------------------------------------------------------------------------
			childRequestResult := utils.GetDB().First(&childObj , producedDatasetsId).Error // find first using identifier

			if childRequestResult == nil {
				//----------------------------------------------------------------------------
				// remove DataSetObj from the ProducedDatasets array, but wont delete it from db
				//----------------------------------------------------------------------------
				utils.GetDB().Model(&parentObj).Association("ProducedDatasets").Delete( &childObj )

			} else {
				msg := fmt.Sprintf( "Failed trying to read %s using ID=%v", "ProducedDatasets", producedDatasetsId )
				return utils.RequestResult{false, msg, "removeProducedDatasets", childObj}
			}
		}

		//----------------------------------------------------------------------------
		// retrieve the modified DataSource from the gorm
		//----------------------------------------------------------------------------
		return GetDataSource(dataSourceId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// adds one or more pipelinesIds as a Pipelines to a DataSource
//----------------------------------------------------------------------------
func AddPipelinesToDataSource ( dataSourceId uint64, pipelinesIds string )(utils.RequestResult) {

	//----------------------------------------------------------------------------
	// Obtain the DataSource with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataSource(dataSourceId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataSource so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataSource)

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
		// retrieve the modified DataSource from the gorm
		//----------------------------------------------------------------------------
		return GetDataSource(dataSourceId)

	} else {
		return parentRequestResult
	}
}

//----------------------------------------------------------------------------
// removes one or more pipelinesIds as a Pipelines from a DataSource
//----------------------------------------------------------------------------
func RemovePipelinesFromDataSource( dataSourceId uint64, pipelinesIds string )(utils.RequestResult) {
	//----------------------------------------------------------------------------
	// Obtain the DataSource with the matching identifier
	//----------------------------------------------------------------------------
	parentRequestResult := GetDataSource(dataSourceId)

	if parentRequestResult.Success == true {
		//----------------------------------------------------------------------------
		// Need to cast the interface to a model.DataSource so the ORM can figure
		// out which table to deal with
		//----------------------------------------------------------------------------
		parentObj,_ := parentRequestResult.Data. (model.DataSource)

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
		// retrieve the modified DataSource from the gorm
		//----------------------------------------------------------------------------
		return GetDataSource(dataSourceId)

	} else {
		return parentRequestResult
	}
}

